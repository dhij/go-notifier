package storer

import (
	"context"
	"errors"
)

var (
	ErrAlreadyExists = errors.New("name already exists")
)

type Storer interface {
	ListFollowers(ctx context.Context, userUUID string) ([]*Follower, error)
	EnqueueNotificationEvent(ctx context.Context, message, userUUID, followerUUID string) error
	ListNotificationEvents(ctx context.Context) ([]*NotificationEvent, error)
}
