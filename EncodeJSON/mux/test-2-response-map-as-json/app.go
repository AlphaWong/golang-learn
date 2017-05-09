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
		// Warning
		// It will covert the map to byte code
		// An Error will be throw if covertion failed.
		//  http://stackoverflow.com/questions/21197239/decoding-json-in-golang-using-json-unmarshal-vs-json-newdecoder-decode
		encodedJSON, err := json.Marshal(userMap)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		// response will be fired via fmt
		fmt.Fprintf(w, string(encodedJSON))
		// Reference https://blog.golang.org/json-and-go
	})
	http.ListenAndServe(PORT, r)
}
