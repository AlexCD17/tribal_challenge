// Package services has all the business logic, such as the credit procedures based on the rules provided
package services

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strings"
	"time"
	"tribalChallenge/entitites"
)

type CreditService struct {
	customer *entitites.Customer
}

func (credit *CreditService) Filter() (int, string) {
	record := credit.customer

	tries := record.Tries
	lastTime := time.Now().UTC().Sub(record.RequestedDate)
	log.Println(fmt.Sprintf("last request time: %v", lastTime))

	if tries >= 3 && !record.Accepted {
		log.Println("fourth rule")
		return http.StatusTooManyRequests, "A sales agent will contact you"
	}

	if lastTime > 2*time.Minute && record.Accepted {
		record.Tries = 0
		lastTime = time.Now().UTC().Sub(record.RequestedDate)
		return http.StatusOK, fmt.Sprintf("credit line accepted for the following amount: $ %.2f", record.RequestedCredit)
	}

	if record.Accepted == false && lastTime < (30*time.Second) {
		log.Println("third rule")
		return http.StatusTooManyRequests, "try again in some time"
	}

	if record.Accepted && tries >= 3 && lastTime < (2*time.Minute) {
		log.Println("second rule")
		return http.StatusTooManyRequests, "Blocked for 2 min"
	}

	if record.Accepted {
		log.Println(" first rule")
		return http.StatusOK, fmt.Sprintf("A credit line was accepted for the following amount: $ %.2f", record.RequestedCredit)
	}

	return 0, ""

}

func (credit *CreditService) EvaluateCredit() bool {
	customer := credit.customer
	creditRequested := customer.RequestedCredit

	monthly := customer.MonthlyRevenue / 5

	if strings.ToLower(customer.FoundingType) == "startup" {

		bal := customer.CashBalance / 3
		log.Println(fmt.Sprintf("credit requested: %v\nBalance: %v\nmonthly: %v\n", creditRequested, bal, monthly))

		return creditRequested < math.Max(monthly, bal)
	}

	log.Println(fmt.Sprintf("credit requested: %v\nmonthly: %v\n", creditRequested, monthly))
	return creditRequested < monthly
}

func NewCreditService(customer *entitites.Customer) *CreditService {
	return &CreditService{
		customer: customer,
	}
}
