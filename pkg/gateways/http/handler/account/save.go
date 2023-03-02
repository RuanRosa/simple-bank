package account

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/RuanRosa/simple-bank/pkg/domain/entities/account"
	"github.com/RuanRosa/simple-bank/pkg/util/response"
)

func (h *Handler) Save(w http.ResponseWriter, r *http.Request) {
	response := response.Json{}

	req := saveRequestBody{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	entity := requestToEntity(req)

	ctx := r.Context()

	if err := h.usecase.Save(&ctx, &entity); err != nil {
		if errors.Is(err, account.ErrCpfAlredyExists) {
			response.WriteError(w, err, http.StatusBadRequest)
			return
		}

		response.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	resp := entiyToResponse(entity)

	response.Write(w, resp, http.StatusCreated)
}
