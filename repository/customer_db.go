package repository

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// create database adapter class
type customerRepositoryDB struct {
	db *sqlx.DB
}

func NewCustomerRepositoryDB(db *sqlx.DB) CustomerRepository {
	return customerRepositoryDB{db: db}
}

// method of adapter rely on interface
func (r customerRepositoryDB) GetAll() ([]Customer, error) {
	customers := []Customer{}
	query := `SELECT * from customers`
	err := r.db.Select(&customers, query)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

// method of adapter rely on interface
func (r customerRepositoryDB) GetById(id int) (*Customer, error) {
	customer := Customer{}
	fmt.Println("GetById", id)
	query := `SELECT * from customers WHERE customer_id = $1`
	err := r.db.Get(&customer, query, id)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {

		return nil, err
	}
	return &customer, nil
}
