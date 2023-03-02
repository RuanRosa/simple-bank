package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/RuanRosa/simple-bank/pkg/gateways/service/auth"
	"github.com/RuanRosa/simple-bank/pkg/util/response"
	"github.com/jackc/pgx/v4"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	response := response.Json{}

	credentials := auth.Credentials{}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		response.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	token, err := h.service.Login(credentials)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			response.WriteError(w, nil, http.StatusNotFound)
			return
		}

		if errors.Is(err, auth.ErrInvalidSecret) {
			response.WriteError(w, err, http.StatusUnauthorized)
			return
		}

		response.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	responseBody := &responseBody{
		Token: *token,
	}

	response.Write(w, responseBody, http.StatusOK)
}
