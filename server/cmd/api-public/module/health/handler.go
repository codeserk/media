package health

import "github.com/gorilla/mux"

func Handle(router *mux.Router) {
	router.Handle("/health", health()).Methods("GET", "OPTIONS")
}
