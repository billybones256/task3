package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"task3/pkg/api"
	"task3/pkg/soup"
)

func main() {
	server := grpc.NewServer()
	microService := &soup.GRPCServer{}
	api.RegisterSoupServer(server, microService)
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	if err = server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
