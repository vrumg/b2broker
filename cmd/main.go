package main

import (
	"fmt"
	"log"
	"net"

	"b2broker/internal/api"
	"b2broker/internal/client"
	"b2broker/internal/client/clientdb"
	"b2broker/internal/group"
	"b2broker/internal/group/groupdb"
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

	clientRepo := clientdb.New()
	clientService := client.New(clientRepo)

	groupRepo := groupdb.New()
	groupService := group.New(groupRepo)

	impl := api.NewAPI(clientService)
	s := grpc.NewServer()
	pb.RegisterMessageServiceServer(s, impl)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
