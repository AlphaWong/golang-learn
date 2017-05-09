// It aims to test json response (string based)
package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// PORT of web server.
const PORT = ":8080"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		byteJSON := []byte(`
			{
				"name":"Alpha String",
				"gender":"male",
				"registerYear": 2017,
				"balance":9.87,
				"socialAccounts": ["facebook","line"]
			}`)
		// only w write can write byte
		w.Write(byteJSON)
		// Reference
		// 1. https://blog.golang.org/json-and-go
		// 2. https://eager.io/blog/go-and-json/
	})
	http.ListenAndServe(PORT, r)
}
