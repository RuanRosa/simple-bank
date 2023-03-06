package account

import (
	"context"

	"github.com/RuanRosa/simple-bank/pkg/domain/entities/account"
)

type UsecaseSPY struct {
}

func (u *UsecaseSPY) All(ctx *context.Context) ([]account.Entity, error) {
	return nil, nil
}

func (u *UsecaseSPY) Save(ctx *context.Context, req *account.Entity) error {
	return nil

}

func (u *UsecaseSPY) GetByCPF(ctx *context.Context, cpf string) (*account.Entity, error) {
	return nil, nil

}
func (u *UsecaseSPY) GetByID(ctx *context.Context, AccountID int) (*account.Entity, error) {
	return nil, nil

}
func (u *UsecaseSPY) Discount(ctx *context.Context, req *account.Entity) error {
	return nil
}
func (u *UsecaseSPY) Deposit(ctx *context.Context, req *account.Entity) error {
	return nil
}
