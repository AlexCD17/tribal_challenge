package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"tribalChallenge/entrypoints"
	"tribalChallenge/repositories/customers"
	"tribalChallenge/services"
)

func main() {

	repo := customers.NewRepository()
	customerService := services.NewCustomerService(repo)
	router := mux.NewRouter().StrictSlash(true)
	server := entrypoints.NewServer(customerService, router)
	server.SetupRouter()
	log.Fatal(http.ListenAndServe(":8080", router))

}
