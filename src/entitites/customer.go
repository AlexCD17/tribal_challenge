package entitites

import "time"

type Customer struct {
	FoundingType    string    `json:"foundingType"`
	CashBalance     float64   `json:"cashBalance"`
	MonthlyRevenue  float64   `json:"monthlyRevenue"`
	RequestedCredit float64   `json:"requestedCreditLine"`
	RequestedDate   time.Time `json:"requestedDate"`
	Tries           uint
	Accepted        bool
}
