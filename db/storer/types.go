package storer

import "time"

type NotificationEventState string

const (
	Pending   NotificationEventState = "pending"
	Delivered NotificationEventState = "delivered"
	Failed    NotificationEventState = "failed"
)

type User struct {
	UUID string
	Name string
	Url  string
}

type Follower struct {
	UUID         string
	UserUUID     string
	FollowerUUID string
	Url          string
}

type FollowerURL struct {
	FollowerUUID string
	Url          string
}

type NotificationEvent struct {
	UUID         string
	Message      string
	FollowerUUID string
	StateUUID    string
	Attempts     int
	FollowerURL  string
	State        string
}

type NotificationState struct {
	UUID        string
	State       NotificationEventState
	Message     string
	RequestedAt time.Time
	CompletedAt time.Time
}
