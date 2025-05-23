package auth

import (
	"media/internal/controller"
	"media/module/user"
	"net/http"
)

// @Tags         auth
// @Summary      Send verify email
// @Description  Re-sends the verification email
// @ID           sendVerifyEmail
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Failure      400  {object}  controller.HTTPError
// @Failure      500  {object}  controller.HTTPError
// @Router       /api/v1/auth/send-verify-email [post]
func sendVerifyEmail(userAuth user.AuthService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u := controller.RequireUser(w, r)
		if u == nil {
			controller.UnauthorizedError(w)
			return
		}

		err := userAuth.SendVerifyEmail(u)
		if err != nil {
			controller.InternalError(w, err)
			return
		}
	})
}
