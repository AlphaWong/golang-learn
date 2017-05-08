// It aims to test json response (string based)
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// PORT of web server.
const PORT = ":8080"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w,
			`{
				"name":"Alpha String",
				"gender":"male",
				"registerYear": 2017,
				"balance":9.87,
				"socialAccounts": ["facebook","line"]
			}`,
		)
		// Reference https://elithrar.github.io/article/testing-http-handlers-go/
	})
	http.ListenAndServe(PORT, r)
}
