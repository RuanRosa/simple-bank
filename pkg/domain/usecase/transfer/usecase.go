package transfer

import (
	"context"

	"github.com/RuanRosa/simple-bank/pkg/domain/entities/transfer"
	"github.com/RuanRosa/simple-bank/pkg/domain/usecase/account"
)

type usecase struct {
	repository     transfer.IRepository
	accountUsecase account.IUsecase
}

type IUsecase interface {
	Save(ctx *context.Context, req *transfer.Entity) (*transfer.Entity, error)
	Transfer(ctx *context.Context, req *transfer.Entity) (*transfer.Entity, error)
}

func NewUsecase(repository transfer.IRepository, accountUsecase account.IUsecase) IUsecase {
	return &usecase{
		repository:     repository,
		accountUsecase: accountUsecase,
	}
}
