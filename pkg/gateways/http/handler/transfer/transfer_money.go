package transfer

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/RuanRosa/simple-bank/pkg/util/response"
)

func (h *Handler) TransferMoney(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	response := response.Json{}

	req := requestBody{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	myAccountId, err := strconv.ParseInt(r.Header.Get("account_id"), 10, 64)
	if err != nil {
		response.WriteError(w, err, http.StatusBadRequest)
		return
	}

	req.AccountOriginID = int(myAccountId)

	entity := requestToEntity(req)

	if _, err := h.usecase.Transfer(&ctx, &entity); err != nil {
		response.WriteError(w, err, http.StatusBadRequest)
		return
	}

	response.WriteError(w, nil, http.StatusOK)
}
