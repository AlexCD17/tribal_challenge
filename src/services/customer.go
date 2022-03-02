package services

import (
	"fmt"
	"time"
	"tribalChallenge/entitites"
	"tribalChallenge/repositories/customers"
)

type CustomerService struct {
	repo *customers.Repository
}

func (s *CustomerService) AddCustomer(key string, customer entitites.Customer) error {
	res := s.repo.SetCustomer(key, customer)
	if res == false {
		return fmt.Errorf("error adding customer")
	}
	return nil
}

func (s *CustomerService) NewCustomer(key string, customer entitites.Customer) error {
	customer.Accepted = false
	customer.Tries = 0
	customer.RequestedDate = time.Now().UTC()
	res := s.repo.SetCustomer(key, customer)
	if res == false {
		return fmt.Errorf("unable to create new customer")
	}

	return nil

}

func (s *CustomerService) GetCustomer(key string) (*entitites.Customer, error) {
	customer, err := s.repo.GetCustomer(key)
	if err != nil {
		return nil, err
	}

	return &customer, nil

}

func (s *CustomerService) UpdateCustomer(key string, customer entitites.Customer) error {
	err := s.repo.UpdateCustomer(key, customer)
	if err != nil {
		return err
	}

	return nil
}

func (s *CustomerService) UpdateTry(key string, customer entitites.Customer) error {

	err := s.repo.UpdateCustomer(key, entitites.Customer{
		FoundingType:    customer.FoundingType,
		CashBalance:     customer.CashBalance,
		MonthlyRevenue:  customer.MonthlyRevenue,
		RequestedCredit: customer.RequestedCredit,
		RequestedDate:   time.Now().UTC(),
		Tries:           customer.Tries + 1,
		Accepted:        customer.Accepted,
	})
	if err != nil {
		return err
	}

	return nil

}

func (s *CustomerService) DeleteCustomer(key string) {

	s.repo.DeleteCustomer(key)

}

func NewCustomerService(repo *customers.Repository) *CustomerService {
	return &CustomerService{repo: repo}
}
