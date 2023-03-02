package account

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/RuanRosa/simple-bank/pkg/util/response"
	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v4"
)

func (h *Handler) GetBalance(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	response := response.Json{}

	id, err := strconv.ParseInt(chi.URLParam(r, "account_id"), 10, 64)
	if err != nil {
		response.WriteError(w, err, http.StatusBadRequest)
		return
	}

	account, err := h.usecase.GetByID(&ctx, int(id))

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			response.WriteError(w, nil, http.StatusNotFound)
			return
		}

		response.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	balance := map[string]interface{}{
		"balance": account.Balance,
	}

	response.Write(w, balance, http.StatusOK)
}
