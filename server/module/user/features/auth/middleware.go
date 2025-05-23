package auth

import (
	"context"
	"media/module/user"
	"net/http"
)

func (s *service) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := r.Header.Get("Authorization")
		if t == "" {
			next.ServeHTTP(w, r)
			return
		}

		u, err := s.GetUserFromToken(t)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		newRequest := r.WithContext(context.WithValue(r.Context(), user.ContextKey, u))
		*r = *newRequest

		next.ServeHTTP(w, r)
	})
}
