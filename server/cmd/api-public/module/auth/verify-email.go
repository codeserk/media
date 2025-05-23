package auth

import (
	userapi "media/cmd/api-public/module/user"
	"media/internal/controller"
	"media/module/user"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// request

type verifyEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
	Token string `json:"token" validate:"required"`
}

// @Tags         auth
// @Summary      Verify email
// @Description  Verifies the user email
// @ID           verifyEmail
// @Accept       json
// @Produce      json
// @Param        request  body      verifyEmailRequest  true  "Verify email request body"
// @Success      200      {object}  user.Response
// @Failure      400      {object}  controller.HTTPError
// @Failure      500      {object}  controller.HTTPError
// @Router       /api/v1/auth/verify-email [post]
func verifyEmail(userAuth user.AuthService) http.Handler {
	validate := validator.New()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input verifyEmailRequest
		if controller.ValidateRequest(w, r, validate, &input) != nil {
			return
		}

		result, err := userAuth.VerifyEmail(input.Email, input.Token)
		if err != nil {
			controller.InternalError(w, err)
			return
		}

		response := userapi.ToResponse(result)

		controller.SendJSON(w, response)
	})
}
