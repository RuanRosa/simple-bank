package auth

import "github.com/RuanRosa/simple-bank/pkg/domain/usecase/account"

type Service struct {
	accountUsecase account.IUsecase
	AccessSecret   string
}

func NewService(accountUsecase account.IUsecase, accessSecret string) Service {
	return Service{
		accountUsecase: accountUsecase,
		AccessSecret:   accessSecret,
	}
}
