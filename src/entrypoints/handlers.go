// Package entrypoints also provides the main entrypoint here, GetCredit, where all the bussines logic is invoked using DI to services/usecases
package entrypoints

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"tribalChallenge/entitites"
	"tribalChallenge/services"
)

func (s Server) GetKey(ip string) string {

	h := sha1.New()

	h.Write([]byte(ip))

	return string(h.Sum(nil))

}

func (s Server) GetCredit(w http.ResponseWriter, r *http.Request) {
	var body entitites.Customer

	log.Println("called")

	_ = json.NewDecoder(r.Body).Decode(&body)

	tp := strings.ToLower(body.FoundingType)
	if tp != "startup" && tp != "sme" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("wrong founding type")
		return
	}

	log.Println(fmt.Sprintf("body: %v", body))

	ip := strings.Split(r.RemoteAddr, ":")[0]

	key := s.GetKey(ip)
	get := s.costumer.GetCustomer
	add := s.costumer.NewCustomer
	update := s.costumer.UpdateTry

	_, err := get(key)
	if err != nil {
		log.Println("creating new user")
		creditService := services.NewCreditService(&body)
		err := add(key, body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		customer, err := get(key)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if creditService.EvaluateCredit() {
			customer.Accepted = true
			err := update(key, *customer)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode(fmt.Sprintf("credit line accepted for the following amount: $ %.2f", customer.RequestedCredit))
			return
		}

		er := update(key, *customer)
		if er != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode("credit line declined")

		return

	}

	customer, _ := get(key)
	creditService := services.NewCreditService(customer)

	code, message := creditService.Filter()
	if code != 0 {
		log.Println(fmt.Sprintf("csutomer: %v\n", customer))
		err := update(key, *customer)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(code)
		_ = json.NewEncoder(w).Encode(message)
		return
	}

	if creditService.EvaluateCredit() {
		customer.Accepted = true
		err := update(key, *customer)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(fmt.Sprintf("credit line accepted for the following amount: $ %.2f", customer.RequestedCredit))

		return
	}

	er := update(key, *customer)
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode("credit line declined")
	return

}
