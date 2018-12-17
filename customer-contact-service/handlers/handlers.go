package handlers

import (
	"context"
	"encoding/csv"
	"log"
	"strconv"
	"strings"

	"github.com/benjih/grpc-services-test/customer-contact-service/dao"
	"github.com/benjih/grpc-services-test/customer-contact-service/model"
	pb "github.com/benjih/grpc-services-test/grpc"
)

type Server struct{}

func (s *Server) AddOrUpdateCustomerContact(ctx context.Context, in *pb.CustomerContactRequest) (*pb.CustomerContactReply, error) {
	csvReader := csv.NewReader(strings.NewReader(in.Data))
	record, err := csvReader.Read()
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	id, err := strconv.Atoi(record[0])
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	dao.AddOrUpdateCustomerContact(&model.CustomerContact{
		ID:     id,
		Name:   record[1],
		Email:  record[2],
		Number: record[3],
	})

	return &pb.CustomerContactReply{}, nil
}
