package errors

import "errors"

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrDeviceIDNotFound  = errors.New("device ID not found")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrNoUsersFound      = errors.New("no users found")

	//new user post
	ErrNewUserDeviceIDEmpty   = errors.New("user device id empty")
	ErrNewUserNameEmpty       = errors.New("user name empty")
	ErrNewUserProfilePicEmpty = errors.New("user profile pic empty")
)
