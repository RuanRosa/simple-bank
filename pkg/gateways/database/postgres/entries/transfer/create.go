package transfer

import (
	"context"

	"github.com/RuanRosa/simple-bank/pkg/domain/entities/transfer"
)

func (r repository) Save(ctx *context.Context, transferInstance *transfer.Entity) (*transfer.Entity, error) {
	const query = `
		INSERT INTO
			transfer (
				account_origin_id,
				account_destination_id,
				amount
			)
		VALUES
			($1, $2, $3)
		RETURNING 
			id,
			account_origin_id,
			account_destination_id,
			amount
	`

	transferReturned := transfer.Entity{}

	err := r.DB.QueryRow(*ctx,
		query,
		transferInstance.AccountOriginID,
		transferInstance.AccountDestinationID,
		transferInstance.Amount,
	).Scan(
		&transferReturned.ID,
		&transferReturned.AccountOriginID,
		&transferReturned.AccountDestinationID,
		&transferReturned.Amount,
	)
	if err != nil {
		return nil, err
	}

	return &transferReturned, nil
}
