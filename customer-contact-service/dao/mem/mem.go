package mem

import (
	"github.com/benjih/grpc-services-test/customer-contact-service/model"
)

type inMemoryDAO struct {
	customerContacts map[int]model.CustomerContact
}

func NewInMemoryDAO() *inMemoryDAO {
	return &inMemoryDAO{
		customerContacts: map[int]model.CustomerContact{},
	}
}

func (d *inMemoryDAO) AddOrUpdateCustomerContact(customerContact model.CustomerContact) error {
	return nil
}
