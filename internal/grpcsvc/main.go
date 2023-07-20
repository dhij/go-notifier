package grpcsvc

import (
	"context"
	"database/sql"
	"log"
	"net"

	"github.com/dhij/go-notifier"
	"github.com/dhij/go-notifier/db/storer"
	grpcsvc "github.com/dhij/go-notifier/notifier_grpcsvc"
	"github.com/ianschenck/envflag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Main(ctx context.Context, db *sql.DB) error {
	var (
		svcAddr = envflag.String("SVC_ADDR", "0.0.0.0:9091", "address where the docrsvc is listening")
	)

	storer := storer.NewMySQL(db)
	srv, err := grpcsvc.NewServer(storer)
	if err != nil {
		log.Fatal("initializing new server:", err)
	}

	grpcSrv := grpc.NewServer()
	notifier.RegisterNotifierServer(grpcSrv, srv)
	reflection.Register(grpcSrv)

	listener, err := net.Listen("tcp", *svcAddr)
	if err != nil {
		log.Fatal("listener failed", err)
	}

	log.Printf("start gRPC server at %s", listener.Addr().String())
	err = grpcSrv.Serve(listener)
	if err != nil {
		log.Fatal("cannot start grpc server:", err)
	}
	return nil
}
