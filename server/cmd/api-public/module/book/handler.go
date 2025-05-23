package book

import (
	"media/module/book"

	"github.com/gorilla/mux"
)

func Handle(router *mux.Router, read book.ReadService) {
	h := handler{read}

	router.Handle("/books/search", h.search()).Methods("GET", "OPTIONS")

}

type handler struct {
	read book.ReadService
}
