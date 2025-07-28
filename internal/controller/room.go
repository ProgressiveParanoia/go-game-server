package controller

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/ProgressiveParanoia/go-game-server/internal/errors"
	"github.com/ProgressiveParanoia/go-game-server/internal/model/game"
	"github.com/ProgressiveParanoia/go-game-server/internal/repo"
	"github.com/coder/websocket"
	"github.com/google/uuid"
)

const (
	PlayerCountInRooms    = 2
	SubscriberBufferLimit = 16
)

type Room struct {
	//TODO: keep track of separate rooms. Maybe spawn a go routine for each room?
	// We are doing individual rooms for now

	repo repo.RoomRepository

	subscribersMux sync.Mutex
	subscribers    map[*RoomSubscriber]struct{}
}

type RoomSubscriber struct {
	msgs chan []byte
}

func NewRoomController(repo repo.RoomRepository) *Room {

	rm := &Room{
		repo:        repo,
		subscribers: make(map[*RoomSubscriber]struct{}),
	}

	//for testing
	randMessages := []string{"deez", "nuts", "hah", "got", "eeeem"}
	go rm.noiseMaker(randMessages)

	return rm
}

func (rm *Room) Create() (string, error) {
	id := uuid.New().String()

	gameRoom := &game.Room{
		ID:        id,
		Players:   make([]string, PlayerCountInRooms),
		CreatedAt: time.Now().UTC(),
	}

	err := rm.repo.Create(gameRoom)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (rm *Room) Join(roomID, userID string, w http.ResponseWriter, r *http.Request) (*game.Room, error) {

	if roomID == "" {
		return nil, errors.ErrEmptyRoomIDsNotAllowed
	}

	gameRoom, err := rm.repo.Join(roomID, userID)
	if err != nil {
		return nil, err
	}

	return gameRoom, nil
}

func (rm *Room) SubscribeToRoom(roomID, userID string, w http.ResponseWriter, r *http.Request) error {

	/*
		ongoingGameRoom, err := rm.repo.GetRoom(roomID)
		if err != nil {
			return err
		}

		userInRoom := false
		for _, id := range ongoingGameRoom.Players {
			if userID == id {
				userInRoom = true
				break
			}
		}

		if !userInRoom {
			return errors.ErrUserNotJoinedPriorSub
		}*/

	rms := &RoomSubscriber{
		msgs: make(chan []byte, SubscriberBufferLimit),
	}

	rm.addSubscriber(rms)
	defer rm.removeSubscriber(rms)

	//create socket for client to connect to
	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		return err
	}

	defer c.CloseNow()

	// Use close read to keep reading from connection to process control messages
	// and cancel the context if the connection drops.
	ctx := c.CloseRead(context.Background())

	for {
		select {
		case msg := <-rms.msgs:
			err := writeTimeout(ctx, time.Second*5, c, msg)
			if err != nil {
				return err
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

// Debugging purposes. Just inject random stuff into the messages channel of every subscriber
func (rm *Room) noiseMaker(randmsgs []string) {

	for {
		time.Sleep(5 * time.Second)
		randomIndex := rand.Intn(len(randmsgs))
		msg := randmsgs[randomIndex]
		byteMsg := []byte(msg)

		for sub, _ := range rm.subscribers {

			sub.msgs <- byteMsg
		}

		fmt.Printf("\n Sent message to subscribers %v in bytes %v \n sub size %v", msg, byteMsg, len(rm.subscribers))
	}
}

func (rm *Room) addSubscriber(s *RoomSubscriber) {
	rm.subscribersMux.Lock()
	rm.subscribers[s] = struct{}{}
	rm.subscribersMux.Unlock()
}

func (rm *Room) removeSubscriber(s *RoomSubscriber) {
	rm.subscribersMux.Lock()
	delete(rm.subscribers, s)
	rm.subscribersMux.Unlock()
}

func (rm *Room) Delete(id string) error {
	return rm.repo.Delete(id)
}

func (rm *Room) GetAll() ([]*game.Room, error) {
	return rm.repo.GetAll()
}

func writeTimeout(ctx context.Context, timeout time.Duration, c *websocket.Conn, msg []byte) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	return c.Write(ctx, websocket.MessageText, msg)
}
