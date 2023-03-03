package account

import (
	"context"

	"github.com/RuanRosa/simple-bank/pkg/domain/entities/account"
)

func (u *usecase) Deposit(ctx *context.Context, req *account.Entity) error {
	return u.repository.Deposit(ctx, req)
}
