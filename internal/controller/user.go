package controller

import (
	"time"

	"github.com/ProgressiveParanoia/go-game-server/internal/model"
	"github.com/ProgressiveParanoia/go-game-server/internal/repo"
)

type User struct {
	repo repo.UserRepository
}

func NewUserController(repo repo.UserRepository) *User {
	return &User{repo: repo}
}

func (u *User) Create(name, deviceID, pf string) (string, error) {

	_, err := u.repo.GetByDeviceID(deviceID)
	if err == nil {
		return "", err
	}

	model := &model.User{
		Name:           name,
		DeviceID:       deviceID,
		ProfilePicture: pf,
		CreatedAt:      time.Now().Format(time.RFC3339),
		LastLogin:      time.Now().Format(time.RFC3339),
		Active:         true,
	}

	return u.repo.Create(model)
}

func (u *User) GetAll() ([]*model.User, error) {
	users, err := u.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *User) GetByDeviceID(ID string) (*model.User, error) {
	user, err := u.repo.GetByDeviceID(ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) GetByID(ID string) (*model.User, error) {
	user, err := u.repo.GetByID(ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) Update(user *model.User) error {
	existingUser, err := u.repo.GetByDeviceID(user.DeviceID)
	if err != nil {
		return err
	}

	existingUser.Name = user.Name
	existingUser.ProfilePicture = user.ProfilePicture
	existingUser.LastLogin = time.Now().Format(time.RFC3339)

	return u.repo.Update(existingUser)
}

func (u *User) Delete(ID string) error {
	_, err := u.repo.GetByDeviceID(ID)
	if err != nil {
		return err
	}
	return u.repo.Delete(ID)
}
