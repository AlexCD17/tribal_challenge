package services

import (
	"fmt"
	"net/http"
	"testing"
	"time"
	"tribalChallenge/entitites"
)

func TestCreditService_EvaluateCredit(t *testing.T) {
	now := time.Now().UTC()

	type args struct {
		input entitites.Customer
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Accepted Startup",
			args: args{
				input: entitites.Customer{
					FoundingType:    "STARTUP",
					CashBalance:     4000.89,
					MonthlyRevenue:  2330.9,
					RequestedCredit: 500,
					RequestedDate:   now,
					Tries:           1,
					Accepted:        false,
				},
			},
			want: true,
		},
		{
			name: "Rejected Startup",
			args: args{
				input: entitites.Customer{
					FoundingType:    "STARTUP",
					CashBalance:     400.89,
					MonthlyRevenue:  1330.9,
					RequestedCredit: 5000,
					RequestedDate:   now,
					Tries:           1,
					Accepted:        false,
				},
			},
			want: false,
		},
		{
			name: "Accepted SME",
			args: args{
				input: entitites.Customer{
					FoundingType:    "SME",
					CashBalance:     4000.89,
					MonthlyRevenue:  4330.9,
					RequestedCredit: 200,
					RequestedDate:   now,
					Tries:           1,
					Accepted:        false,
				},
			},
			want: true,
		},
		{
			name: "Rejected SME",
			args: args{
				input: entitites.Customer{
					FoundingType:    "SME",
					CashBalance:     400.89,
					MonthlyRevenue:  1030.9,
					RequestedCredit: 5000,
					RequestedDate:   now,
					Tries:           1,
					Accepted:        false,
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := NewCreditService(&tt.args.input)

			res := service.EvaluateCredit()
			if res != tt.want {
				t.Errorf("creditService.EvaluateCredit() Accepted: %v, wanted: %v", res, tt.want)
			}

		})
	}
}

func TestCreditService_Filter(t *testing.T) {

	now := time.Now().UTC()

	type args struct {
		input entitites.Customer
	}

	tests := []struct {
		name string
		args args
		code int
		str  string
	}{
		{
			name: "rule 1",
			args: args{
				input: entitites.Customer{
					FoundingType:    "STARTUP",
					CashBalance:     4000.89,
					MonthlyRevenue:  2330.9,
					RequestedCredit: 500,
					RequestedDate:   now,
					Tries:           2,
					Accepted:        true,
				},
			},
			code: http.StatusOK,
			str:  fmt.Sprintf("A credit line was accepted for the following amount: $ %.2f", 500.00),
		},
		{
			name: "rule 2",
			args: args{
				input: entitites.Customer{
					FoundingType:    "STARTUP",
					CashBalance:     400.89,
					MonthlyRevenue:  1330.9,
					RequestedCredit: 200,
					RequestedDate:   now,
					Tries:           3,
					Accepted:        true,
				},
			},
			code: http.StatusTooManyRequests,
			str:  "Blocked for 2 min",
		},
		{
			name: "rule 3",
			args: args{
				input: entitites.Customer{
					FoundingType:    "SME",
					CashBalance:     4000.89,
					MonthlyRevenue:  4330.9,
					RequestedCredit: 200,
					RequestedDate:   now,
					Tries:           2,
					Accepted:        false,
				},
			},
			code: http.StatusTooManyRequests,
			str:  "try again in some time",
		},
		{
			name: "rule 4",
			args: args{
				input: entitites.Customer{
					FoundingType:    "SME",
					CashBalance:     400.89,
					MonthlyRevenue:  1030.9,
					RequestedCredit: 5000,
					RequestedDate:   now,
					Tries:           4,
					Accepted:        false,
				},
			},
			code: http.StatusTooManyRequests,
			str:  "A sales agent will contact you",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := NewCreditService(&tt.args.input)

			code, msg := service.Filter()
			if code != tt.code {
				t.Errorf("creditService.Filter() http response failed code: %v, expected: %v", code, tt.code)
			}
			if msg != tt.str {
				t.Errorf("CreditService.Filter() message failed expected: %v, got: %v", msg, tt.str)
			}

		})
	}

}
