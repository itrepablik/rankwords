package routers

import (
	"rankwords/cmd/contents"

	"github.com/gorilla/mux"
)

// RankWordsRouters sets up the routes for the API
func RankWordsRouters(r *mux.Router) {
	// Project Codes
	r.HandleFunc("/v1/contents", contents.AcceptContentsHandler).Methods("POST")
}
