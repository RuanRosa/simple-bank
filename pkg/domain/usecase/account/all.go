package account

import (
	"context"

	"github.com/RuanRosa/simple-bank/pkg/domain/entities/account"
)

func (u *usecase) All(ctx *context.Context) ([]account.Entity, error) {
	account, err := u.repository.GetAccounts(*ctx)
	return account, err
}
