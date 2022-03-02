// Package entrypoints provides the server initialization and configuration
package entrypoints

import (
	"github.com/gorilla/mux"
	"tribalChallenge/repositories/customers"
	"tribalChallenge/services"
)

type Server struct {
	//credit   *services.CreditService
	costumer *services.CustomerService
	router   *mux.Router
	repo     *customers.Repository
}

func (s *Server) SetupRouter() {
	s.router.Methods("POST").Path("/eval").HandlerFunc(s.GetCredit)
}

func NewServer(customerService *services.CustomerService, router *mux.Router) *Server {
	return &Server{
		costumer: customerService,
		router:   router,
	}
}
