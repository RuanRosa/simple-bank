package account

import (
	"context"

	"github.com/RuanRosa/simple-bank/pkg/domain/entities/account"
)

func (r repository) Discount(ctx *context.Context, req *account.Entity) error {
	statement := `
		UPDATE account
		SET balance = (balance - $1)
		WHERE id = $2
		RETURNING
			id,
			name,
			cpf,
			secret,
			balance,
			created_at
	`

	return r.DB.QueryRow(*ctx, statement,
		req.Balance,
		req.ID,
	).Scan(
		&req.ID,
		&req.Name,
		&req.CPF,
		&req.Secret,
		&req.Balance,
		&req.CreatedAt,
	)
}
