package handler

import (
	"bank/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type customerHandler struct {
	customerService service.CustomerService
}

func NewCustomerHandler(customerService service.CustomerService) customerHandler {
	return customerHandler{customerService: customerService}
}

func (h customerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := h.customerService.GetCustomers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customers)
}

func (h customerHandler) GetCustomerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// convert string to int
	customerID, err := strconv.Atoi(vars["customer_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid Customer ID")
		return
	}
	customer, err := h.customerService.GetCustomerById(customerID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}
	if customer == nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Customer Not Found")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customer)
}
