// It aims to test json response (struct based)
package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// PORT of web server.
const PORT = ":8080"

// User for testing
type User struct {
	Name            string      `json:"name"`
	Gender          string      `json:"gender"`
	RegisterYear    *int        `json:"registerYear,omitempty"`
	MembershipLevel *int        `json:"membershipLevel,omitempty"`
	Ranking         int         `json:"ranking,omitempty"`
	Balance         float32     `json:"balance"`
	CompanyName     string      `json:"companyName,omitempty"`
	ReferralUserID  string      `json:"referralUserID"`
	SocialAccounts  []string    `json:"socialAccounts"`
	CreditCard      *CreditCard `json:"creditCard"`
	Phone           Phone       `json:"phone"`
}

// CreditCard for testing
type CreditCard struct {
	Number string `json:"number"`
	Issuer string `json:"issuer"`
}

// Phone for testing
type Phone struct {
	Number string `json:"number"`
	OS     string `json:"os"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user/{userName}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		membershipLevel := 0
		dummyUser0 := User{
			Name:            vars["userName"],
			Gender:          "male",
			RegisterYear:    nil,
			MembershipLevel: &membershipLevel,
			Ranking:         0,
			Balance:         31.146,
			SocialAccounts:  []string{"facebook", "line"},
			CreditCard: &CreditCard{
				Number: "12345-6789",
				Issuer: "VISA",
			},
			Phone: Phone{
				Number: "98765432",
				OS:     "window phone",
			},
		}
		// Do not set the header
		// https://github.com/golang/go/issues/17083
		// w.WriteHeader(http.StatusOK)

		// Follow statement is must.
		w.Header().Set("Content-Type", "application/json")
		// response will be fired via json
		json.NewEncoder(w).Encode(dummyUser0)
		// Reference https://kev.inburke.com/kevin/golang-json-http/
	})
	http.ListenAndServe(PORT, r)
}
