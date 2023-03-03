package auth

import "github.com/RuanRosa/simple-bank/pkg/gateways/service/auth"

type Middleware struct {
	authService auth.Service
}

func NewMiddleware(authService auth.Service) Middleware {
	return Middleware{
		authService: authService,
	}
}
