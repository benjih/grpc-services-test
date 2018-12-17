package handlers

import (
	"context"
	"log"

	"github.com/benjih/grpc-services-test/customer-contact-service/controllers"
	pb "github.com/benjih/grpc-services-test/grpc"
)

type Server struct{}

func (s *Server) AddOrUpdateCustomerContact(ctx context.Context, in *pb.CustomerContactRequest) (*pb.CustomerContactReply, error) {
	if err := controllers.AddOrUpdateCustomerContact([]string{in.Id, in.Name, in.EmailAddress, in.TelephoneNumber}); err != nil {
		log.Print(err.Error())
		return &pb.CustomerContactReply{}, nil
	}

	log.Print("Customer contact added")

	return &pb.CustomerContactReply{}, nil
}
