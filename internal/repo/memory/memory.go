package memory

import (
	"github.com/ProgressiveParanoia/go-game-server/internal/controller" //TODO: Fix Cyclic import
	"github.com/ProgressiveParanoia/go-game-server/internal/model"
	"github.com/google/uuid"
)

type MemoryRepository struct {
	//TODO: Find a better way to handle device user maps. We have dupes for now
	deviceUserMap map[string]*model.User
	users         map[string]*model.User
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		users:         make(map[string]*model.User),
		deviceUserMap: make(map[string]*model.User),
	}
}

func (r *MemoryRepository) Create(user *model.User) (string, error) {
	if _, exists := r.deviceUserMap[user.DeviceID]; exists {
		return "", controller.ErrUserAlreadyExists
	}

	uniqueID := uuid.New().String()
	r.users[uniqueID] = user
	r.deviceUserMap[user.DeviceID] = user

	return uniqueID, nil
}

func (r *MemoryRepository) GetByDeviceID(deviceID string) (*model.User, error) {
	if user, exists := r.deviceUserMap[deviceID]; exists {
		return user, nil
	}
	return nil, controller.ErrDeviceIDNotFound
}

func (r *MemoryRepository) GetByID(id string) (*model.User, error) {
	if user, exists := r.users[id]; exists {
		return user, nil
	}
	return nil, controller.ErrUserNotFound
}

func (r *MemoryRepository) Update(user *model.User) error {
	if user, exists := r.users[user.ID]; exists {
		r.users[user.ID] = user
		return nil
	}

	return controller.ErrUserNotFound
}

func (r *MemoryRepository) Delete(id string) error {
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

	return controller.ErrUserNotFound
}

func (r *MemoryRepository) GetAll() ([]*model.User, error) {

	usersLen := len(r.users)
	if usersLen == 0 {
		return nil, controller.ErrNoUsersFound
	}

	users := make([]*model.User, usersLen)
	for _, user := range r.users {
		users = append(users, user)
	}

	return users, nil
}
