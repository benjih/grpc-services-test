package controllers

import (
	"strconv"

	"github.com/benjih/grpc-services-test/customer-contact-service/dao"
	"github.com/benjih/grpc-services-test/customer-contact-service/model"
)

func AddOrUpdateCustomerContact(record []string) error {
	id, err := strconv.Atoi(record[0])
	if err != nil {
		return err
	}

	return dao.AddOrUpdateCustomerContact(&model.CustomerContact{
		ID:     id,
		Name:   record[1],
		Email:  record[2],
		Number: record[3],
	})
}
