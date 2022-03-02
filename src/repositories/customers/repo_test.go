package customers

import (
	"reflect"
	"testing"
	"time"
	"tribalChallenge/entitites"
)

func TestRepository_SetCustomer(t *testing.T) {
	type fields struct {
		cache *Repository
	}
	type args struct {
		key      string
		customer entitites.Customer
	}

	now := time.Now().UTC()

	tests := []struct {
		name    string
		args    args
		want    entitites.Customer
		wantErr bool
	}{
		{
			name: "fail create customer",
			args: args{
				key: "15",
				customer: entitites.Customer{
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
		{
			name: "succesfully created customer",
			args: args{
				key: "13",
				customer: entitites.Customer{
					FoundingType:    "SME",
					CashBalance:     770,
					MonthlyRevenue:  2400,
					RequestedCredit: 100,
					RequestedDate:   now,
					Tries:           2,
					Accepted:        true,
				},
			},
			want: entitites.Customer{
				FoundingType:    "SME",
				CashBalance:     770,
				MonthlyRevenue:  2400,
				RequestedCredit: 100,
				RequestedDate:   now,
				Tries:           2,
				Accepted:        true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cachedb := NewRepository()

			c := &fields{
				cache: cachedb,
			}

			c.cache.SetCustomer(tt.args.key, tt.args.customer)
			rec, err := c.cache.GetCustomer(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("customerRepo.SetCustomer() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(rec, tt.want) {
				t.Errorf("SetCustomer() = %v, want %v", rec, tt.want)
			}

		})

	}

}

func TestRepository_UpdateCustomer(t *testing.T) {
	type fields struct {
		cache *Repository
	}
	type args struct {
		key      string
		customer entitites.Customer
	}

	now := time.Now().UTC()

	tests := []struct {
		name    string
		args    args
		want    entitites.Customer
		wantErr bool
	}{
		{
			name: "fail update customer",
			args: args{
				key: "19",
				customer: entitites.Customer{
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
		{
			name: "succesfully created customer",
			args: args{
				key: "13",
				customer: entitites.Customer{
					FoundingType:    "SME",
					CashBalance:     1300,
					MonthlyRevenue:  4400,
					RequestedCredit: 400,
					RequestedDate:   now,
					Tries:           3,
					Accepted:        true,
				},
			},
			want: entitites.Customer{
				FoundingType:    "SME",
				CashBalance:     1300,
				MonthlyRevenue:  4400,
				RequestedCredit: 400,
				RequestedDate:   now,
				Tries:           3,
				Accepted:        true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cachedb := NewRepository()

			c := &fields{
				cache: cachedb,
			}

			err := c.cache.UpdateCustomer(tt.args.key, tt.args.customer)
			if err != nil {
				t.Errorf("customerRepo.SetCustomer() error = %v, wantErr %v", err, tt.wantErr)
			}
			rec, err := c.cache.GetCustomer(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("customerRepo.SetCustomer() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(rec, tt.want) {
				t.Errorf("SetCustomer() = %v, want %v", rec, tt.want)
			}

		})

	}

}
