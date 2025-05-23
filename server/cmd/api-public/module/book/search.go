package book

import (
	"media/internal/controller"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type analysisResponse struct {
	Genres    []string `json:"genres"`
	Themes    []string `json:"themes"`
	Mood      string   `json:"mood"`
	Settings  []string `json:"settings"`
	AgeGroups []string `json:"ageGroups"`
	Pacing    string   `json:"pacing"`
}

// request

type searchRequest struct {
	Query string `json:"query" validate:"required,min=3"`
}

// response

type searchResponse struct {
	Google    any `json:"google"`
	Open      any `json:"open"`
	Items     any `json:"items"`
	GoodReads any `json:"goodreads"`
}

// @Tags         books
// @Summary      Search
// @Description  Search books using a query
// @ID           searchBooks
// @Accept       json
// @Produce      json
// @Param        request  query     searchRequest  true  "Search request body"
// @Success      200      {object}  searchResponse
// @Failure      400      {object}  controller.HTTPError
// @Failure      500      {object}  controller.HTTPError
// @Router       /api/v1/books/search [get]
func (h *handler) search() http.Handler {
	validate := validator.New()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input searchRequest
		if controller.ValidateQuery(w, r, validate, &input) != nil {
			return
		}

		book, err := h.read.FromISBN(input.Query, true)
		if err != nil {
			controller.InternalError(w, err)
			return
		}

		// book, err := h.booksSources.FromISBN(input.Query)
		// if err != nil {
		// 	controller.InternalError(w, err)
		// 	return
		// }

		// imgs := book.Images()
		// if len(imgs) > 0 {
		// 	img := imgs[0]

		// 	path, _ := h.images.Download(img)
		// 	i, _ := h.images.Open(path)
		// 	pixelated, _ := h.images.Pixelate(i)
		// 	pixelatedBase64, _ := h.images.Base64(pixelated)
		// 	fmt.Printf("%v\n", pixelatedBase64)

		// 	small, _ := h.images.Resize(i, media.ResizeParams{Width: 32, Height: 0, KeepAspectRatio: true, Filter: media.Nearest})
		// 	eInk := h.images.EffectEInk(i)
		// 	h.images.Save(small, fmt.Sprintf("%s/small", book.Metadata().ISBN))
		// 	url, _ := h.images.Save(eInk, fmt.Sprintf("%s/e-ink", book.Metadata().ISBN))

		// 	fmt.Printf("url: %s\n", url)
		// }

		// 		chatRes, err := h.chat.Structured(fmt.Sprintf(`
		// "Analyze the following book details and extract the specified information:
		// Title: %s
		// Description: %s
		// Author: %s
		// Genres: %s

		// Response must be JSON, for each field only write one or multiple valid options (empty if none apply)
		// Based on the provided information, please identify and list the following:
		// - genres: Select the most relevant genres from the following list: Fantasy, Science Fiction, Mystery, Thriller, Romance, Horror, Historical Fiction, Adventure, Crime, Young Adult, Children's Literature, Biography, Self-Help, Poetry, Drama, Comedy, Non-Fiction, Dystopian, Paranormal, Literary Fiction.
		// - themes: Identify any themes present in the description from this list: Love, Betrayal, Survival, Identity, Power and Corruption, Good vs. Evil, Coming of Age, War and Peace, Freedom and Oppression, Journey and Quest.
		// - mood: Determine the mood or tone from the description using these options: Dark, Light-hearted, Suspenseful, Melancholic, Hopeful, Humorous, Inspirational, Tense, Whimsical, Nostalgic.
		// - settings: Infer the setting from the description using these options: Urban, Rural, Historical, Futuristic, Fantasy World.
		// - ageGroups: Suggest the most likely target age group from these options: Children, Young Adult, Adult, Senior.
		// - pacing: Determine the pacing or style from the description using these options: Fast-Paced, Slow-Burn, Epistolary, Stream of Consciousness, Non-Linear Narrative."
		// 			`, book.Metadata().Title, book.Metadata().Description, book.Metadata().Authors, book.Metadata().Genres,
		// 		), analysisResponse{})
		// 		if err != nil {
		// 			log.Printf("err: %v", err)
		// 			// log.Fatalf("Error getting chat completion: %v", err)
		// 		}
		// 		log.Printf("Chat completion: %+v\n", chatRes)

		// var response searchResponse
		controller.SendJSON(w, searchResponse{
			Google: book,
			// Open:   openItems,
			// Items:  libItems,
			// GoodReads: res,
		})
	})
}
