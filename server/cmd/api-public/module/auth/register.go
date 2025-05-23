package auth

import (
	userapi "media/cmd/api-public/module/user"
	"media/internal/controller"
	"media/module/user"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type registerRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// @Tags         auth
// @Summary      Register
// @Description  Register using email and credentials
// @ID           register
// @Accept       json
// @Produce      json
// @Param        request  body      registerRequest  true  "Register request body"
// @Success      200      {object}  loginResponse
// @Failure      400      {object}  controller.HTTPError
// @Failure      500      {object}  controller.HTTPError
// @Router       /api/v1/auth/register [post]
func register(userAuth user.AuthService) http.Handler {
	validate := validator.New()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input registerRequest
		if controller.ValidateRequest(w, r, validate, &input) != nil {
			return
		}

		result, err := userAuth.Register(user.RegisterParams(input))
		if err != nil {
			controller.InternalError(w, err)
			return
		}

		response := loginResponse(loginResponse{
			User: userapi.Response{
				Id:        result.User.Id,
				Name:      result.User.Name,
				Email:     result.User.Email,
				CreatedAt: result.User.CreatedAt,
				UpdatedAt: result.User.UpdatedAt,
			},
			Token: result.Token,
		})

		controller.SendJSON(w, response)
	})
}
