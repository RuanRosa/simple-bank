package account

import (
	"context"

	"github.com/RuanRosa/simple-bank/pkg/domain/entities/account"
)

func (u *usecase) Discount(ctx *context.Context, req *account.Entity) error {
	return u.repository.Discount(ctx, req)
}
