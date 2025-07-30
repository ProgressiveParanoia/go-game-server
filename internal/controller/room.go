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

	repo              repo.RoomRepository
	ongoingMatches    map[string]*RoomMatch
	ongoingMatchesMux sync.Mutex
}

type RoomSubscriber struct {
	msgs chan []byte
}

type RoomMatch struct {
	subscribersMux sync.Mutex
	//todo: add identifier for subscribers or have a separate reference point for IDs
	subscribers map[*RoomSubscriber]struct{}
	CreatedAt   time.Time

	isDone bool
}

func NewRoomController(repo repo.RoomRepository) *Room {
	rm := &Room{
		repo:           repo,
		ongoingMatches: make(map[string]*RoomMatch),
	}

	return rm
}

func (rm *Room) Create() (string, error) {
	id := uuid.New().String()

	m := &RoomMatch{
		subscribers: make(map[*RoomSubscriber]struct{}),
		CreatedAt:   time.Now().UTC(),
	}

	//for debugging purposes
	randStuff := []string{"deez", "nuts", "in", "ya", "mouth", "mah neeeega"}
	go m.noiseMaker(randStuff, id)

	rm.ongoingMatches[id] = m

	return id, nil
}

func (rm *Room) SubscribeToRoom(roomID, userID string, w http.ResponseWriter, r *http.Request) error {

	match, ok := rm.ongoingMatches[roomID]
	if !ok {
		fmt.Println("content of ongoing matches: ", rm.ongoingMatches, " id passed: ", roomID)
		return errors.ErrRoomNonExistent
	}

	rms := &RoomSubscriber{
		msgs: make(chan []byte, SubscriberBufferLimit),
	}

	match.addSubscriber(rms)
	defer match.removeSubscriber(rms)

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

func (r *Room) GetRooms() ([]game.Room, error) {
	room := make([]game.Room, 0)

	for k, om := range r.ongoingMatches {
		roomData := game.Room{
			ID:        k,
			Players:   []string{},
			CreatedAt: om.CreatedAt,
		}

		room = append(room, roomData)
	}

	return room, nil
}

func (rm *Room) CleanUpEmptyRoomAfterDisconnect(roomID string) error {

	m, ok := rm.ongoingMatches[roomID]

	if !ok {
		//return nil for now. Other subscribers may have cleaned up the room when the disconnect
		return nil
	}

	if len(m.subscribers) != 0 {
		return nil // Room not empty, so we return
	}

	rm.ongoingMatchesMux.Lock()
	m.isDone = true
	delete(rm.ongoingMatches, roomID)
	defer rm.ongoingMatchesMux.Unlock()

	return nil
}

// Debugging purposes. Just inject random stuff into the messages channel of every subscriber
func (m *RoomMatch) noiseMaker(randmsgs []string, id string) {
	for {
		time.Sleep(5 * time.Second)

		if m.isDone {
			fmt.Println("\n\n\n\nStop making noise!")
			return
		}

		randomIndex := rand.Intn(len(randmsgs))
		msg := randmsgs[randomIndex]
		byteMsg := []byte(msg)

		for sub, _ := range m.subscribers {
			sub.msgs <- byteMsg
		}

		fmt.Printf("\n [%v] Sent message to subscribers %v \n sub size %v", id, msg, len(m.subscribers))
	}
}

func (m *RoomMatch) addSubscriber(s *RoomSubscriber) {
	m.subscribersMux.Lock()
	m.subscribers[s] = struct{}{}
	m.subscribersMux.Unlock()
}

func (m *RoomMatch) removeSubscriber(s *RoomSubscriber) {
	m.subscribersMux.Lock()
	delete(m.subscribers, s)
	m.subscribersMux.Unlock()
}

func (r *Room) DeleteMatch(id string) error {
	if _, ok := r.ongoingMatches[id]; !ok {
		return errors.ErrRoomNonExistent
	}

	delete(r.ongoingMatches, id)
	return nil
}

func writeTimeout(ctx context.Context, timeout time.Duration, c *websocket.Conn, msg []byte) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	return c.Write(ctx, websocket.MessageText, msg)
}
