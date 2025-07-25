package repo

import (
	"github.com/ProgressiveParanoia/go-game-server/internal/model"
	"github.com/ProgressiveParanoia/go-game-server/internal/model/game"
)

type (
	UserRepository interface {
		Create(user *model.User) (string, error)
		GetByDeviceID(id string) (*model.User, error)
		GetByID(id string) (*model.User, error)
		Update(user *model.User) error
		Delete(id string) error

		//FOR DEBUGGING
		GetAll() ([]*model.User, error)
	}

	RoomRepository interface {
		Create(room *game.Room) error
		Join(id string) (*game.Room, error)
		Delete(id string) error

		//FOR DEBUGGING
		GetAll() ([]*game.Room, error)
	}
)
