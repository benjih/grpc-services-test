package controllers

import (
	"fmt"
	"strconv"

	"github.com/benjih/grpc-services-test/customer-contact-service/dao"
	"github.com/benjih/grpc-services-test/customer-contact-service/model"
)

func AddOrUpdateCustomerContact(record []string) error {
	id, err := strconv.Atoi(record[0])
	if err != nil {
		return err
	}

	if !model.IsValidEmailAddress(record[2]) {
		return fmt.Errorf("Invalid email address")
	}

	if !model.IsValidTelephoneNumber(record[3]) {
		return fmt.Errorf("Invalid telephone number")
	}

	return dao.AddOrUpdateCustomerContact(&model.CustomerContact{
		ID:              id,
		Name:            record[1],
		EmailAddress:    record[2],
		TelephoneNumber: record[3],
	})
}
