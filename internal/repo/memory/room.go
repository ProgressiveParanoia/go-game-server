package memory

import (
	"fmt"

	"github.com/ProgressiveParanoia/go-game-server/internal/errors"
	"github.com/ProgressiveParanoia/go-game-server/internal/model/game"
)

type RoomMemoryRepository struct {
	rooms map[string]*game.Room

	//maybe have this as a config?
	maxPlayersPerRoom int
}

func NewRoomRepository(maxPlayersPerRoom int) *RoomMemoryRepository {
	return &RoomMemoryRepository{
		rooms:             make(map[string]*game.Room),
		maxPlayersPerRoom: maxPlayersPerRoom,
	}
}

func (rmr *RoomMemoryRepository) Create(room *game.Room) error {
	if room == nil {
		return fmt.Errorf("nil rooms not allowed")
	}

	roomID := room.ID
	if _, ok := rmr.rooms[roomID]; ok {
		return errors.ErrRoomCreatedExists
	}

	rmr.rooms[roomID] = room
	return nil
}

func (rmr *RoomMemoryRepository) Join(roomID, userID string) (*game.Room, error) {
	rm, err := rmr.GetRoom(roomID)
	if err != nil {
		return nil, err
	}

	userInRoom := false
	for _, id := range rm.Players {
		if userID == id {
			userInRoom = true
			break
		}
	}

	//NOTE: we'll only allow 1 session for now.
	// Users existing in a room could mean also mean a user was abruptly disconnected
	// so we'll add a reconnect option in the future
	if userInRoom {
		return nil, errors.ErrUserAlreadyJoinedRoom
	}

	if len(rm.Players)+1 > rmr.maxPlayersPerRoom {
		return nil, errors.ErrRoomFull
	}

	rm.Players = append(rm.Players, userID)

	return rm, nil
}

func (rmr *RoomMemoryRepository) GetRoom(roomID string) (*game.Room, error) {
	rm, ok := rmr.rooms[roomID]
	if !ok {
		return nil, errors.ErrRoomNonExistent
	}

	return rm, nil
}

func (rmr *RoomMemoryRepository) Delete(id string) error {
	if _, exists := rmr.rooms[id]; !exists {
		return errors.ErrRoomNonExistent
	}

	delete(rmr.rooms, id)
	return nil
}

func (rmr *RoomMemoryRepository) GetAll() ([]*game.Room, error) {
	roomLen := len(rmr.rooms)
	if roomLen == 0 {
		return nil, errors.ErrNoRoomsFound
	}

	rooms := make([]*game.Room, roomLen)
	for _, rm := range rmr.rooms {
		rooms = append(rooms, rm)
	}

	return rooms, nil
}
