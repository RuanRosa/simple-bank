package account

import (
	"context"

	"github.com/RuanRosa/simple-bank/pkg/domain/entities/account"
)

func (u *usecase) GetByID(ctx *context.Context, AccountID int) (*account.Entity, error) {
	return u.repository.GetByID(*ctx, AccountID)
}
