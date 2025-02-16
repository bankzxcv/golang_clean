package service

import (
	"bank/repository"
	"log"
)

type customerService struct {
	customerRepository repository.CustomerRepository // reference to customer repository not reference to database directly
}

func NewCustomerService(customerRepository repository.CustomerRepository) CustomerService {
	return customerService{customerRepository: customerRepository}
}

// this is Business Logic
func (s customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := s.customerRepository.GetAll()
	if err != nil {
		log.Println("Error GetCustomers", err)
		return nil, err
	}

	// transform Customer (DB) to CustomerResponse (Service)
	customerResponses := []CustomerResponse{}
	for _, customer := range customers {
		customerResponse := CustomerResponse{
			CustomerID: customer.CustomerID,
			Name:       customer.Name,
			Status:     customer.Status,
		}
		customerResponses = append(customerResponses, customerResponse)
	}
	return customerResponses, nil
}

// this is Business Logic
func (s customerService) GetCustomerById(number int) (*CustomerResponse, error) {
	customer, err := s.customerRepository.GetById(number)
	if err != nil {
		log.Println("Error GetCustomerById", err)
		return nil, err
	}

	if customer == nil {
		return nil, nil
	}

	customerResponse := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}

	return &customerResponse, nil
}
