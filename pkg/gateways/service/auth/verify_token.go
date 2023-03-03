package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

var ErrInvalidToken error = errors.New("invalid token")

func (s *Service) VerifyToken(encryptedToken string) (*jwt.Token, error) {
	claims := jwt.MapClaims{}

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(s.AccessSecret), nil
	}

	jwtToken, err := jwt.ParseWithClaims(encryptedToken, claims, keyFunc)
	if err != nil {
		return nil, err
	}

	if !jwtToken.Valid {
		return nil, ErrInvalidToken
	}

	return jwtToken, nil
}
