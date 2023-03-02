package account

import (
	"context"

	"github.com/RuanRosa/simple-bank/pkg/domain/entities/account"
)

func (u *usecase) GetByCPF(ctx *context.Context, cpf string) (*account.Entity, error) {
	return u.repository.GetByCPF(*ctx, cpf)
}
