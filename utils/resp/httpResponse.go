package resp

import (
	"encoding/json"
	"net/http"
)

// HttpResponse is a helper function that logs a message with an http.Response
func HttpResponse(w http.ResponseWriter, statusCode int, payLoad interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(payLoad); err != nil {
		panic(err)
	}
}

// HttpResponseNoCache is a helper function that logs a message with an http.Response without cache
func HttpResponseNoCache(w http.ResponseWriter, statusCode int, payLoad interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Pragma", "no-cache")

	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(payLoad); err != nil {
		panic(err)
	}
}
