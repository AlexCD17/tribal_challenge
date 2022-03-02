// Some test here, not covered 100% tho
package services

import (
	"testing"
	"time"
	"tribalChallenge/entitites"
	"tribalChallenge/repositories/customers"
)

func TestCustomerService_NewCustomer(t *testing.T) {

	now := time.Now().UTC()

	type args struct {
		key   string
		input entitites.Customer
	}

	tests := []struct {
		name    string
		args    args
		want    entitites.Customer
		wantErr bool
	}{
		{
			name: "customer creation",
			args: args{
				key: "15",
				input: entitites.Customer{
					FoundingType:    "STARTUP",
					CashBalance:     400.89,
					MonthlyRevenue:  2330.9,
					RequestedCredit: 5000,
					RequestedDate:   now,
					Tries:           1,
					Accepted:        false,
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := customers.NewRepository()
			service := NewCustomerService(repo)

			err := service.NewCustomer(tt.args.key, tt.args.input)
			if err != nil {
				t.Errorf("customerService.NewCustomer error creating new customer: %v", err)
			}
		})
	}
}

func TestCustomerService_GetCustomer(t *testing.T) {
	now := time.Now().UTC()

	type args struct {
		key   string
		input entitites.Customer
	}

	tests := []struct {
		name    string
		args    args
		want    entitites.Customer
		wantErr bool
	}{
		{
			name: "successful get",
			args: args{
				key: "234",
				input: entitites.Customer{
					FoundingType:    "STARTUP",
					CashBalance:     400.89,
					MonthlyRevenue:  2330.9,
					RequestedCredit: 5000,
					RequestedDate:   now,
					Tries:           1,
					Accepted:        false,
				},
			},
			want: entitites.Customer{
				FoundingType:    "STARTUP",
				CashBalance:     400.89,
				MonthlyRevenue:  2330.9,
				RequestedCredit: 5000,
				RequestedDate:   now,
				Tries:           1,
				Accepted:        false,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := customers.NewRepository()
			service := NewCustomerService(repo)

			err := service.AddCustomer(tt.args.key, tt.args.input)
			if err != nil {
				t.Errorf("error adding customer")
			}

			cust, err := service.GetCustomer(tt.args.key)
			if err != nil {
				t.Errorf("customerService.GetCustomer error getting customer: %v", err)
			}

			if *cust != tt.want {
				t.Errorf("customerService.GetCustomer error, got: %v, expected: %v", *cust, tt.want)

			}

		})
	}

}
