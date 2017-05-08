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
	Name           string   `json:"name"`
	Gender         string   `json:"gender"`
	RegisterYear   int      `json:"registerYear"`
	Balance        float32  `json:"balance"`
	SocialAccounts []string `json:"socialAccounts"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		dummyUser0 := User{
			Name:           "Alpha",
			Gender:         "male",
			RegisterYear:   2017,
			Balance:        31.146,
			SocialAccounts: []string{"facebook", "line"},
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
