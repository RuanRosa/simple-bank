package account

import (
	"net/http"

	"github.com/RuanRosa/simple-bank/pkg/util/response"
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
		response.WriteError(w, err, http.StatusInternalServerError)
	}

	response.Write(w, respBody, http.StatusOK)
}
