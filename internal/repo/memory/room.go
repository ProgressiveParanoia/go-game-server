package memory

import (
	"fmt"

	"github.com/ProgressiveParanoia/go-game-server/internal/errors"
	"github.com/ProgressiveParanoia/go-game-server/internal/model/game"
)

type RoomMemoryRepository struct {
	rooms map[string]*game.Room
}

func NewRoomRepository() *RoomMemoryRepository {
	return &RoomMemoryRepository{
		rooms: make(map[string]*game.Room),
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

func (rmr *RoomMemoryRepository) Join(id string) (string, error) {
	if _, ok := rmr.rooms[id]; !ok {
		return "", errors.ErrRoomNonExistent
	}

	return "", nil
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
