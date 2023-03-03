package auth

import (
	"errors"
	"net/http"

	"github.com/RuanRosa/simple-bank/pkg/gateways/service/auth"
	"github.com/RuanRosa/simple-bank/pkg/util/response"
)

func (m *Middleware) Check(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := response.Json{}
		secret := r.Header.Get("secret")

		credentials := auth.Credentials{
			Secret: secret,
		}

		_, err := m.authService.IsValid(credentials)
		if err != nil {
			if errors.Is(err, auth.ErrInvalidToken) {
				response.WriteError(w, err, http.StatusUnauthorized)
				return
			}

			response.WriteError(w, err, http.StatusInternalServerError)
			return
		}

		next.ServeHTTP(w, r)
	})
}
