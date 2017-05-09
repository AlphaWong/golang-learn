// It aims to test json response (string based)
package main

import (
	"net/http"

	"io"

	"github.com/gorilla/mux"
)

// PORT of web server.
const PORT = ":8080"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		// Reference https://elithrar.github.io/article/testing-http-handlers-go/
		w.Header().Set("Content-Type", "application/json")
		// response will be fired via io
		io.WriteString(w,
			`{
				"name":"Alpha String",
				"gender":"male",
				"registerYear": 2017,
				"balance":9.87,
				"socialAccounts": ["facebook","line"]
			}`,
		)
	})
	http.ListenAndServe(PORT, r)
}
