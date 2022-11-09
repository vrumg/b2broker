package main

import (
	"fmt"
	"log"
	"net"

	"b2broker/internal/api"
	"b2broker/internal/client"
	"b2broker/internal/client/clientdb"
	pb "b2broker/pkg/b2brokerpb"

	"google.golang.org/grpc"
)

const (
	port = 13111
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	clientRepo := clientdb.NewDatabase()
	clientService := client.NewService(clientRepo)

	impl := api.NewAPI(clientService)
	s := grpc.NewServer()
	pb.RegisterMessageServiceServer(s, impl)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
