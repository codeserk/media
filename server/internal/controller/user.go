package controller

import (
	"media/module/user"
	"net/http"
)

func User(r *http.Request) *user.Entity {
	if u, ok := r.Context().Value(user.ContextKey).(*user.Entity); ok {
		return u
	}

	return nil
}

func RequireUser(w http.ResponseWriter, r *http.Request) *user.Entity {
	u := User(r)
	if u == nil {
		UnauthorizedError(w)
		return nil
	}

	return u
}
