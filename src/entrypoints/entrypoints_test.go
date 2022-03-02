package entrypoints

import (
	"bytes"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"tribalChallenge/repositories/customers"
	"tribalChallenge/services"
)

func TestServer_GetCredit(t *testing.T) {

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}

	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantHeader http.Header
		wantBody   string
	}{
		{
			name: "success getting credit line approved",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(
					http.MethodPost,
					"http://0.0.0.0/eval",
					strings.NewReader(`{
"foundingType": "STARTUP",
"cashBalance": 200.30,
"monthlyRevenue": 4435.45,
"requestedCreditLine": 200,
"requestedDate": "2021-07-19T16:32:59.860Z"
}`)),
			},
			wantStatus: http.StatusOK,
			wantHeader: http.Header{"Content-Type": {"application/json"}},
			wantBody:   "credit line accepted for the following amount: $ 200.00",
		},
		{
			name: "rejected credit line",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(
					http.MethodPost,
					"http://0.0.0.0/eval",
					strings.NewReader(`{
"foundingType": "SME",
"cashBalance": 200.30,
"monthlyRevenue": 4435.45,
"requestedCreditLine": 20000,
"requestedDate": "2021-07-19T16:32:59.860Z"
}`)),
			},
			wantStatus: http.StatusOK,
			wantHeader: http.Header{"Content-Type": {"application/json"}},
			wantBody:   "credit line declined",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := customers.NewRepository()
			customer := services.NewCustomerService(repo)
			router := mux.NewRouter()
			service := NewServer(customer, router)
			service.SetupRouter()

			service.GetCredit(tt.args.w, tt.args.r)

			rec := tt.args.w.(*httptest.ResponseRecorder)
			res := rec.Result()
			if !reflect.DeepEqual(res.StatusCode, tt.wantStatus) {
				t.Errorf("entrypoints.GetCredit() status code error,  got %v, wanted: %v", res.StatusCode, tt.wantStatus)
			}

			bodyBuffer := new(bytes.Buffer)
			_, _ = bodyBuffer.ReadFrom(res.Body)
			body := strings.TrimSpace(bodyBuffer.String())
			if body[1:len(body)-1] != tt.wantBody {
				t.Errorf("entrypoints.GetCredit() wrong response, got %v, wanted: %v", body, tt.wantBody)
			}

		})
	}

}
