package main

import (
	"bank/handler"
	"bank/repository"
	"bank/service"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Open("pgx", "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		fmt.Println("Connect to database error")
		panic(err)
	}
	// 1 start create repository with associated database
	// 2 creating service and push repository to service layer
	// transform repository layer to service layer
	customerRepository := repository.NewCustomerRepositoryDB(db)
	customerService := service.NewCustomerService(customerRepository)
	customerHandler := handler.NewCustomerHandler(customerService)
	router := mux.NewRouter()

	router.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", customerHandler.GetCustomerById).Methods(http.MethodGet)

	http.ListenAndServe(":8000", router)
}
