package controller

import (
	"time"

	"github.com/ProgressiveParanoia/go-game-server/internal/errors"
	"github.com/ProgressiveParanoia/go-game-server/internal/model/game"
	"github.com/ProgressiveParanoia/go-game-server/internal/repo"
	"github.com/google/uuid"
)

const (
	PlayerCountInRooms = 2
)

type Room struct {
	//TODO: keep track of separate rooms. Maybe spawn a go routine for each room?
	// We are doing individual rooms for now

	repo repo.RoomRepository
}

func NewRoomController(repo repo.RoomRepository) *Room {
	return &Room{
		repo: repo,
	}
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

func (rm *Room) Join(id string) (*game.Room, error) {
	if id == "" {
		return nil, errors.ErrEmptyRoomIDsNotAllowed
	}

	gameRoom, err := rm.repo.Join(id)
	if err != nil {
		return nil, err
	}

	return gameRoom, nil
}

func (rm *Room) Delete(id string) error {
	return rm.repo.Delete(id)
}

func (rm *Room) GetAll() ([]*game.Room, error) {
	return rm.repo.GetAll()
}
