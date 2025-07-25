package memory

import (
	"github.com/ProgressiveParanoia/go-game-server/internal/errors"

	//TODO: Fix Cyclic import
	"github.com/ProgressiveParanoia/go-game-server/internal/model"
	"github.com/google/uuid"
)

type UserMemoryRepository struct {
	deviceUserMap map[string]*model.User
	users         map[string]*model.User
}

func NewUserRepository() *UserMemoryRepository {
	return &UserMemoryRepository{
		users:         make(map[string]*model.User),
		deviceUserMap: make(map[string]*model.User),
	}
}

func (r *UserMemoryRepository) Create(user *model.User) (string, error) {
	if _, exists := r.deviceUserMap[user.DeviceID]; exists {
		return "", errors.ErrUserAlreadyExists
	}

	uniqueID := uuid.New().String()
	r.users[uniqueID] = user
	r.deviceUserMap[user.DeviceID] = user

	return uniqueID, nil
}

func (r *UserMemoryRepository) GetByDeviceID(deviceID string) (*model.User, error) {
	if user, exists := r.deviceUserMap[deviceID]; exists {
		return user, nil
	}
	return nil, errors.ErrDeviceIDNotFound
}

func (r *UserMemoryRepository) GetByID(id string) (*model.User, error) {
	if user, exists := r.users[id]; exists {
		return user, nil
	}
	return nil, errors.ErrUserNotFound
}

func (r *UserMemoryRepository) Update(user *model.User) error {
	if user, exists := r.users[user.ID]; exists {
		r.users[user.ID] = user
		return nil
	}

	return errors.ErrUserNotFound
}

func (r *UserMemoryRepository) Delete(id string) error {
	if _, exists := r.users[id]; exists {
		delete(r.users, id)
		for deviceID, user := range r.deviceUserMap {
			if user.ID == id {
				delete(r.deviceUserMap, deviceID)
				break
			}
		}
		return nil
	}

	return errors.ErrUserNotFound
}

func (r *UserMemoryRepository) GetAll() ([]*model.User, error) {

	usersLen := len(r.users)
	if usersLen == 0 {
		return nil, errors.ErrNoUsersFound
	}

	users := make([]*model.User, usersLen)
	for _, user := range r.users {
		users = append(users, user)
	}

	return users, nil
}
