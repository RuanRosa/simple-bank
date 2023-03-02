package account

import (
	"errors"
	"net/http"

	"github.com/RuanRosa/simple-bank/pkg/util/response"
	"github.com/jackc/pgx/v4"
)

func (h *Handler) All(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	accounts, err := h.usecase.All(&ctx)

	respBody := []responseBody{}

	for _, account := range accounts {
		respBody = append(respBody, entiyToResponse(account))
	}

	response := response.Json{}
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			response.WriteError(w, nil, http.StatusNotFound)
			return
		}

		response.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	response.Write(w, respBody, http.StatusOK)
}
