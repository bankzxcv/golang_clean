package service

type CustomerResponse struct {
	CustomerID int    `json:"customer_id"`
	Name       string `json:"name"`
	Status     bool   `json:"status"`
}

type CustomerService interface {
	GetCustomers() ([]CustomerResponse, error)
	GetCustomerById(int) (*CustomerResponse, error)
}
