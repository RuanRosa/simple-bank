package account

import (
	"context"

	"github.com/RuanRosa/simple-bank/pkg/domain/entities/account"
)

func (r *repository) Save(ctx *context.Context, req *account.Entity) error {
	statement := `
		INSERT INTO account
			(name,
				cpf,
				secret,
				balance)
			VALUES ($1, $2, $3, $4)
		returning created_at, id`

	err := r.DB.QueryRow(*ctx, statement,
		req.Name,
		req.CPF,
		req.Secret,
		req.Balance,
	).Scan(&req.CreatedAt, &req.ID)

	return err
}
