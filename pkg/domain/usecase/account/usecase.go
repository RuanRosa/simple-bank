package account

import (
	"context"

	"github.com/RuanRosa/simple-bank/pkg/domain/entities/account"
)

type usecase struct {
	repository account.IRepository
}

type IUsecase interface {
	All(ctx *context.Context) ([]account.Entity, error)
	Save(ctx *context.Context, req *account.Entity) error
	GetByCPF(ctx *context.Context, cpf string) (*account.Entity, error)
	GetByID(ctx *context.Context, AccountID int) (*account.Entity, error)
	Discount(ctx *context.Context, req *account.Entity) error
	Deposit(ctx *context.Context, req *account.Entity) error
}

func NewUsecase(repository account.IRepository) IUsecase {
	return &usecase{
		repository: repository,
	}
}
