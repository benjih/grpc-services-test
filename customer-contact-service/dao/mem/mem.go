package mem

import (
	"fmt"

	"github.com/benjih/grpc-services-test/customer-contact-service/model"
)

type inMemoryDAO struct {
	customerContacts map[int]*model.CustomerContact
}

func NewInMemoryDAO() *inMemoryDAO {
	return &inMemoryDAO{
		customerContacts: map[int]*model.CustomerContact{},
	}
}

func (d *inMemoryDAO) AddOrUpdateCustomerContact(customerContact *model.CustomerContact) error {
	if customerContact == nil || customerContact.ID == 0 {
		return fmt.Errorf("customerContact not correctly set")
	}
	d.customerContacts[customerContact.ID] = customerContact
	return nil
}
