package main

import (
	"log"
	"net"

	"github.com/benjih/grpc-services-test/customer-contact-service/handlers"
	pb "github.com/benjih/grpc-services-test/grpc"

	"google.golang.org/grpc"
)

const (
	port = ":3001"
)

func main() {
	log.Print(" starting customer-contact-service on " + port)
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen to port: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterCustomerContactServiceServer(s, &handlers.Server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
