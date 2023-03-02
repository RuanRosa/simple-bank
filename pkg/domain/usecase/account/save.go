package account

import (
	"context"

	"github.com/RuanRosa/simple-bank/pkg/domain/entities/account"
	"github.com/jackc/pgx/v4"
)

func (u *usecase) Save(ctx *context.Context, req *account.Entity) error {
	acc, err := u.GetByCPF(ctx, req.CPF)
	if err != nil && err != pgx.ErrNoRows {
		return err
	}

	if acc != nil {
		return account.ErrCpfAlredyExists
	}

	return u.repository.Save(ctx, req)
}
