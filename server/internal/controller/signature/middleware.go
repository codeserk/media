package signature

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"media/internal/controller"
	"net/http"
)

func (s *Service) Middleware(signatureApp SignatureApp) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			bodyBytes, err := io.ReadAll(r.Body)
			if err != nil {
				controller.InternalError(w, fmt.Errorf("failed to read body: %v", err))
				return
			}

			err = s.Check(r, string(bodyBytes), signatureApp)
			if err != nil {
				log.Printf("invalid signature: %v", err)
				controller.UnauthorizedError(w)
				return
			}

			r.Body = io.NopCloser(bytes.NewReader(bodyBytes))
			next.ServeHTTP(w, r)
		})
	}
}
