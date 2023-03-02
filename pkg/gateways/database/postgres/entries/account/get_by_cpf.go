package account

import (
	"context"

	"github.com/RuanRosa/simple-bank/pkg/domain/entities/account"
)

func (r repository) GetByCPF(ctx context.Context, CPF string) (*account.Entity, error) {
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
			cpf=$1`

	acc := account.Entity{}

	err := r.DB.QueryRow(ctx, query, CPF).Scan(
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
