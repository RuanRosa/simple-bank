package transfer

import "github.com/RuanRosa/simple-bank/pkg/domain/entities/transfer"

type requestBody struct {
	AccountOriginID      int
	AccountDestinationID int `json:"account_destination_id"`
	Amount               int `json:"amount"`
}

func requestToEntity(req requestBody) transfer.Entity {
	return transfer.Entity{
		AccountOriginID:      req.AccountOriginID,
		AccountDestinationID: req.AccountDestinationID,
		Amount:               req.Amount,
	}
}
