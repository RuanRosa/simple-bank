package auth

import (
	"context"
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type Credentials struct {
	CPF    string `json:"cpf"`
	Secret string `json:"secret"`
}

var ErrInvalidSecret error = errors.New("invalid secret")

func (s *Service) Login(credentials Credentials) (*string, error) {
	ctx := context.Background()

	acc, err := s.accountUsecase.GetByCPF(&ctx, credentials.CPF)
	if err != nil {
		return nil, err
	}

	if acc.Secret != credentials.Secret {
		return nil, ErrInvalidSecret
	}

	return s.createToken(&acc.ID)
}

func (s *Service) createToken(accountId *int) (*string, error) {
	claims := jwt.MapClaims{}
	claims["account_id"] = accountId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(s.AccessSecret))
	if err != nil {
		return nil, err
	}

	return &signedToken, nil
}
