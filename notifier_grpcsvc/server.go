package grpcsvc

import (
	"context"
	"time"

	"github.com/dhij/go-notifier"
	"github.com/dhij/go-notifier/db/storer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	notifier.UnimplementedNotifierServer
	storer storer.Storer
}

func NewServer(storer storer.Storer) (*Server, error) {
	return &Server{
		storer: storer,
	}, nil
}

func (s *Server) EnqueueNotificationEvent(ctx context.Context, req *notifier.EnqueueNotificationEventReq) (*notifier.EnqueueNotificationEventRes, error) {
	followers, err := s.storer.ListFollowers(ctx, req.GetUserUuid())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "listing followers: %v", err)
	}

	for _, f := range followers {
		err = s.storer.EnqueueNotificationEvent(ctx, req.GetMessage(), req.GetUserUuid(), f.UUID)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "enqueuing notification event for user %v: %v", req.GetUserUuid(), err)
		}
	}

	return &notifier.EnqueueNotificationEventRes{}, nil
}

func (s *Server) StreamNotificationEvents(req *notifier.StreamNotificationEventsReq, srv notifier.Notifier_StreamNotificationEventsServer) error {
	ctx := srv.Context()

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for ctx.Err() == nil {
		events, err := s.storer.ListNotificationEvents(ctx)
		if err != nil {
			return status.Errorf(codes.Internal, "listing notification events: %v", err)
		}

		for _, ev := range events {
			res := &notifier.StreamNotificationEventsRes{
				Event: &notifier.NotificationEvent{
					Message:     ev.Message,
					FollowerUrl: ev.FollowerURL,
				},
			}
			err = srv.Send(res)
			if err != nil {
				return err
			}
		}
		select {
		case <-ticker.C:
		}
	}

	return status.Errorf(codes.Canceled, "context canceled")
}
