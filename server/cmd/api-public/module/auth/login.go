package auth

import (
	uerapi "media/cmd/api-public/module/user"
	"media/internal/controller"
	"media/module/user"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// request

type loginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// @Tags         auth
// @Summary      Login
// @Description  Login using credentials
// @ID           login
// @Accept       json
// @Produce      json
// @Param        request  body      loginRequest  true  "Login request body"
// @Success      200      {object}  loginResponse
// @Failure      400      {object}  controller.HTTPError
// @Failure      500      {object}  controller.HTTPError
// @Router       /api/v1/auth/login [post]
func login(userAuth user.AuthService) http.Handler {
	validate := validator.New()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input loginRequest
		if controller.ValidateRequest(w, r, validate, &input) != nil {
			return
		}

		result, err := userAuth.Login(user.LoginParams(input))
		if err != nil {
			controller.InternalError(w, err)
			return
		}

		response := loginResponse{
			User:  uerapi.ToResponse(result.User),
			Token: result.Token,
		}

		controller.SendJSON(w, response)
	})
}
