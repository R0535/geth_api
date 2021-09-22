package main

import (
	"log"
	"net"
	"github.com/R0535/geth_api/proto"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", "8080")
	if err != nil {
		log.Fatal("unable to listen 8080 port %v", err)
	}

	srv := grpc.NewServer(listener)


}
