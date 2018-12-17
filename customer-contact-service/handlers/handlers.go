package handlers

import (
	"context"
	"encoding/csv"
	"log"
	"strings"

	"github.com/benjih/grpc-services-test/customer-contact-service/controllers"
	pb "github.com/benjih/grpc-services-test/grpc"
)

type Server struct{}

func (s *Server) AddOrUpdateCustomerContact(ctx context.Context, in *pb.CustomerContactRequest) (*pb.CustomerContactReply, error) {
	csvReader := csv.NewReader(strings.NewReader(in.Data))
	record, err := csvReader.Read()
	if err != nil {
		log.Print(err.Error())
		return &pb.CustomerContactReply{}, nil
	}

	if err := controllers.AddOrUpdateCustomerContact(record); err != nil {
		log.Print(err.Error())
		return &pb.CustomerContactReply{}, nil
	}

	log.Print("Customer contact added")

	return &pb.CustomerContactReply{}, nil
}
