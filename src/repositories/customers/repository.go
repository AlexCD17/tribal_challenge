// Package customers provides the repository creation and interface for possible mockups down the road.
package customers

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
	"tribalChallenge/entitites"
)

type CustomerRepo interface {
	GetCustomer(customerID string) (entitites.Customer, error)
	SetCustomer(key string, customer entitites.Customer)
	UpdateCustomer(key string, customer entitites.Customer) error
	DeleteCustomer(key string)
}

type Repository struct {
	data *cache.Cache
}

func NewRepository() *Repository {

	c := cache.New(3*time.Hour, 10*time.Hour)

	return &Repository{
		data: c,
	}
}

func (r *Repository) GetCustomer(customerID string) (entitites.Customer, error) {

	record, found := r.data.Get(customerID)

	if found {

		customerData := record.(entitites.Customer)

		return customerData, nil
	}

	return entitites.Customer{}, fmt.Errorf("Failed to get user: %s\n", customerID)

}

func (r *Repository) SetCustomer(key string, customer entitites.Customer) bool {

	r.data.Set(key, customer, cache.NoExpiration)
	_, found := r.data.Get(key)
	return found

}

func (r *Repository) UpdateCustomer(key string, customer entitites.Customer) error {

	err := r.data.Replace(key, customer, cache.NoExpiration)
	if err != nil {
		return err
	}
	return nil

}

func (r *Repository) DeleteCustomer(key string) {

	r.data.Delete(key)

}
