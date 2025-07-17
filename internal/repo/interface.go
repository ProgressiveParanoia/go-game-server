package repo

import "github.com/ProgressiveParanoia/go-game-server/internal/model"

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
)
