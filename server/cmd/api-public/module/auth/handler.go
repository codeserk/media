package auth

import (
	"media/module/user"

	"github.com/gorilla/mux"
)

func Handle(router *mux.Router, userAuth user.AuthService) {

	router.Handle("/auth/login", login(userAuth)).Methods("POST", "OPTIONS")
	router.Handle("/auth/register", register(userAuth)).Methods("POST", "OPTIONS")
	router.Handle("/auth/verify-email", verifyEmail(userAuth)).Methods("POST", "OPTIONS")
	router.Handle("/auth/send-verify-email", sendVerifyEmail(userAuth)).Methods("POST", "OPTIONS")
}

type handler struct {
	userAuth user.AuthService
}
