package storer

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

type MySQLStorer struct {
	Querier
}

func NewMySQL(db Querier) *MySQLStorer {
	return &MySQLStorer{
		Querier: db,
	}
}

func (ms *MySQLStorer) ListFollowers(ctx context.Context, userUUID string) ([]*Follower, error) {
	var followers []*Follower

	q := squirrel.
		Select("uuid").
		From("followers").
		Where(squirrel.Eq{
			"user_uuid": userUUID,
		})

	rows, err := q.RunWith(ms.Querier).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("listing followers: %w", err)
	}

	for rows.Next() {
		follower := &Follower{}
		err := rows.Scan(
			&follower.UUID,
		)
		if err != nil {
			return nil, fmt.Errorf("scanning followers: %w", err)
		}

		followers = append(followers, follower)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterating rows: %w", err)
	}

	return followers, nil
}

func (ms *MySQLStorer) EnqueueNotificationEvent(ctx context.Context, message, userUUID, followerUUID string) error {
	return ms.inTx(ctx, func(ctx context.Context, tx *sql.Tx) error {
		s, err := insertNotificationState(ctx, tx)
		if err != nil {
			return fmt.Errorf("inserting notification state: %w", err)
		}

		_, err = insertNotificationEvent(ctx, tx, message, followerUUID, s.UUID)
		if err != nil {
			return fmt.Errorf("inserting notification event: %w", err)
		}

		return nil
	})
}

func insertNotificationEvent(ctx context.Context, tx *sql.Tx, message, followerUUID, stateUUID string) (*NotificationEvent, error) {
	var uuid = uuid.NewString()
	insertQ := squirrel.
		Insert("notification_events_queue").
		Columns("uuid", "message", "follower_uuid", "state_uuid", "attempts").
		Values(uuid, message, followerUUID, stateUUID, 0)

	_, err := insertQ.RunWith(tx).ExecContext(ctx)
	if err != nil {
		return nil, err
	}

	return &NotificationEvent{
		UUID:         uuid,
		Message:      message,
		FollowerUUID: followerUUID,
		StateUUID:    stateUUID,
		Attempts:     0,
	}, nil
}

func insertNotificationState(ctx context.Context, tx *sql.Tx) (*NotificationState, error) {
	var uuid = uuid.NewString()
	insertQ := squirrel.
		Insert("notification_states").
		Columns("uuid", "state").
		Values(uuid, Pending)

	_, err := insertQ.RunWith(tx).ExecContext(ctx)
	if err != nil {
		return nil, err
	}

	return &NotificationState{
		UUID: uuid,
	}, nil
}

func (ms *MySQLStorer) ListNotificationEvents(ctx context.Context) ([]*NotificationEvent, error) {
	var notificationEvents []*NotificationEvent

	q := squirrel.
		Select("e.message, u.url").
		From("notification_events_queue e").
		Join("followers f ON e.follower_uuid=f.uuid").
		Join("users u ON f.follower_uuid=u.uuid")

		// ORDER BY

	rows, err := q.RunWith(ms.Querier).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("listing followers: %w", err)
	}

	for rows.Next() {
		ev := &NotificationEvent{}
		err := rows.Scan(
			&ev.Message,
			&ev.FollowerURL,
		)
		if err != nil {
			return nil, fmt.Errorf("scanning followers: %w", err)
		}
		notificationEvents = append(notificationEvents, ev)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterating rows: %w", err)
	}

	return notificationEvents, nil
}

func (ms *MySQLStorer) inTx(ctx context.Context, fn func(ctx context.Context, tx *sql.Tx) error) error {
	tx, err := ms.Querier.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("starting transaction: %w", err)
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	return fn(ctx, tx)
}
