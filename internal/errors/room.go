package errors

import (
	"context"
	"errors"
)

var (
	ErrRoomCreatedExists      = errors.New("room already exists")
	ErrRoomFull               = errors.New("room already full")
	ErrRoomNonExistent        = errors.New("room non-existent")
	ErrNoRoomsFound           = errors.New("no rooms found")
	ErrSubscriberDisconnected = errors.New("subscriber disconnected")

	//Player safety checks
	ErrUserAlreadyJoinedRoom = errors.New("user already joined room ")
	ErrUserNotJoinedPriorSub = errors.New("user did not join room prior to subscription")
	//Safety checks
	ErrNilRoomsNotAllowed     = errors.New("nil rooms not allowed")
	ErrEmptyRoomIDsNotAllowed = errors.New("empty IDs not allowed")
)

func HandleRoomContextError(ctx context.Context) error {
	ctxErr := ctx.Err()
	if errors.Is(ctxErr, context.Canceled) {
		return ErrSubscriberDisconnected
	}

	return ctxErr
}
