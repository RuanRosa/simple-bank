package account

import (
	"context"

	"github.com/RuanRosa/simple-bank/pkg/domain/entities/account"
	"github.com/jackc/pgx/v4"
)

func (r *repository) GetAccounts(ctx *context.Context) ([]account.Entity, error) {
	query := `
		SELECT 
			id,
			name,
			cpf,
			balance,
			created_at
		FROM account
	`
	accounts := []account.Entity{}

	rows, err := r.DB.Query(*ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		acc := account.Entity{}

		err := rows.Scan(
			&acc.ID,
			&acc.Name,
			&acc.CPF,
			&acc.Balance,
			&acc.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		accounts = append(accounts, acc)
	}

	if len(accounts) == 0 {
		return nil, pgx.ErrNoRows
	}

	return accounts, nil
}
