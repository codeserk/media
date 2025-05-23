package health

import (
	"media/internal/controller"
	"net/http"
)

type healthResponse struct {
	Ok bool `json:"ok"`
}

// @Tags         health
// @Summary      Health endpoint
// @Description  Endpoint to make sure the application is healthy
// @ID           health
// @Accept       json
// @Produce      json
// @Success      200  {object}  healthResponse
// @Failure      400  {object}  controller.HTTPError
// @Failure      500  {object}  controller.HTTPError
// @Router       /health [get]
func health() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		controller.SendJSON(w, healthResponse{Ok: true})
	})
}
