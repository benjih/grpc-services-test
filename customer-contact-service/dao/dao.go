package dao

import "github.com/benjih/grpc-services-test/customer-contact-service/model"

type DAO interface {
	AddOrUpdateCustomerContact(customerContact model.CustomerContact) error
}

var GlobalDAO DAO

func AddOrUpdateCustomerContact(customerContact model.CustomerContact) error {
	return GlobalDAO.AddOrUpdateCustomerContact(customerContact)
}
