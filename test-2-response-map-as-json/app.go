// It aims to test json response (map based)
package main

import (
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
)

// PORT of web server.
const PORT = ":8080"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		userMap := map[string]interface{}{
			"name":           "Alpha map",
			"gender":         "male",
			"registerYear":   2017,
			"balance":        9.87,
			"socialAccounts": []string{"facebook", "line"},
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		encodedJSON, _ := json.Marshal(userMap)
		fmt.Fprintf(w, string(encodedJSON))
	})
	http.ListenAndServe(PORT, r)
}
