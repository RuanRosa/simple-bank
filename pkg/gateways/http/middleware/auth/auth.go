package auth

import (
	"errors"
	"fmt"
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

		accountId, err := m.authService.IsValid(credentials)
		if err != nil {
			if errors.Is(err, auth.ErrInvalidToken) {
				response.WriteError(w, err, http.StatusUnauthorized)
				return
			}

			response.WriteError(w, err, http.StatusInternalServerError)
			return
		}

		r.Header.Add("account_id", fmt.Sprintf("%v", *accountId))
		next.ServeHTTP(w, r)
	})
}
