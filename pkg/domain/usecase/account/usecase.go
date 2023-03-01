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
}

func NewUsecase(repository account.IRepository) IUsecase {
	return &usecase{
		repository: repository,
	}
}
