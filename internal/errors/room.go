package errors

import "errors"

var (
	ErrRoomCreatedExists = errors.New("room already exists")
	ErrRoomFull          = errors.New("room already full")
	ErrRoomNonExistent   = errors.New("room non-existent")
	ErrNoRoomsFound      = errors.New("no rooms found")
	//Safety checks
	ErrNilRoomsNotAllowed     = errors.New("nil rooms not allowed")
	ErrEmptyRoomIDsNotAllowed = errors.New("empty IDs not allowed")
)
