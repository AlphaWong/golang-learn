// It aims to parse json response (encoder based)

// testing simple request json
//{
//	"name":"Beta wong",
// 	"gender":"male",
// 	"registerYear":1234,
// 	"balance":19.89,
// 	"socialAccounts":[
// 		"facebook",
// 		"github"
// 	]
//}
package main

import (
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
)

// PORT of web server.
const PORT = ":8080"

func main() {
	r := mux.NewRouter()
	// r.Headers("Content-Type", "application/json")
	r.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		// Anonymous structs
		var user struct {
			Name           string   `json:"name"`
			Gender         string   `json:"gender"`
			RegisterYear   int      `json:"registerYear"`
			Balance        float32  `json:"balance"`
			SocialAccounts []string `json:"socialAccounts"`
		}

		// Decode must pass a pointer (&user)
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	})
	http.ListenAndServe(PORT, r)
}
