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
	"io/ioutil"
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
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		json.Unmarshal(body, &user)
		w.Header().Set("Content-Type", "application/json")
		encodedJSON, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(encodedJSON)
	}).Methods("POST")
	http.ListenAndServe(PORT, r)
}
