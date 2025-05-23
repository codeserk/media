package book

import (
	"media/internal/controller"
	"media/module/book"
	"net/http"

	"github.com/gorilla/mux"
)

// @Tags         books
// @Summary      Search
// @Description  Search books using a query
// @ID           searchBooks
// @Accept       json
// @Produce      json
// @Param        isbn  path      string  true  "ISBN to look for"
// @Success      200   {object}  Response
// @Failure      400   {object}  controller.HTTPError
// @Failure      500   {object}  controller.HTTPError
// @Router       /api/v1/books/isbn/{isbn} [get]
func (h *handler) getBookFromISBN() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isbn := book.ToISBN(mux.Vars(r)["isbn"])
		book, err := h.read.FromISBN(isbn, true)
		if err != nil {
			controller.InternalError(w, err)
			return
		}

		controller.SendJSON(w, ToResponse(book))
	})
}
