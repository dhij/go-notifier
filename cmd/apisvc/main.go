package main

import (
	"log"
	"os"

	"github.com/dhij/go-notifier"
	apisvc "github.com/dhij/go-notifier/notifier_apisvc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var (
		apiAddr = "0.0.0.0:9090"
		svcAddr = os.Getenv("SVC_ADDR")
	)

	if svcAddr == "" {
		svcAddr = "0.0.0.0:9091"
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial(svcAddr, opts...)
	if err != nil {
		log.Fatal("connecting to grpc server:", err)
	}
	defer conn.Close()

	client := notifier.NewNotifierClient(conn)

	server := apisvc.NewServer(client)
	server.InitRouter()
	server.Start(apiAddr)
}
