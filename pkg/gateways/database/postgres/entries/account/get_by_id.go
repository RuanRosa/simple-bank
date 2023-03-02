package account

import (
	"context"

	"github.com/RuanRosa/simple-bank/pkg/domain/entities/account"
)

func (r repository) GetByID(ctx context.Context, AccountID int) (*account.Entity, error) {
	query := `
		SELECT
			id,
			name,
			cpf,
			balance,
			secret,
			created_at
		FROM
			account
		WHERE
			id=$1`

	acc := account.Entity{}

	err := r.DB.QueryRow(ctx, query, AccountID).Scan(
		&acc.ID,
		&acc.Name,
		&acc.CPF,
		&acc.Balance,
		&acc.Secret,
		&acc.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &acc, nil
}
