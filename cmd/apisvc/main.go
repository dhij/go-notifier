package main

import (
	"log"

	"github.com/dhij/go-notifier"
	apisvc "github.com/dhij/go-notifier/notifier_apisvc"
	"github.com/ianschenck/envflag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var (
		apiAddr = envflag.String("LISTEN_ADDR", "0.0.0.0:9090", "address where the docrapi should be served")
		svcAddr = envflag.String("SVC_ADDR", "0.0.0.0:9091", "address where the docrsvc is listening")
	)

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.Dial(*svcAddr, opts...)
	if err != nil {
		log.Fatal("connecting to grpc server:", err)
	}
	defer conn.Close()

	client := notifier.NewNotifierClient(conn)

	server := apisvc.NewServer(client)
	server.InitRouter()
	server.Start(*apiAddr)
}
