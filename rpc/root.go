package rpc

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func Serve() {
	listener, err := net.Listen("tcp", ":5679")

	if err != nil {
		log.Panic(err)
	}

	server := grpc.NewServer()
	RegisterJobServer(server, &jobService{})

	if err = server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
