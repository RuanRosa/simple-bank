package auth

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func (s *Service) IsValid(credentials Credentials) (*int64, error) {
	jwt, err := s.VerifyToken(credentials.Secret)
	if err != nil {
		return nil, err
	}

	claims := map[string]interface{}{}

	claimsByte, err := json.Marshal(jwt.Claims)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(claimsByte, &claims); err != nil {
		return nil, err
	}

	for k, v := range claims {
		if k == "account_id" {
			str := fmt.Sprintf("%v", v)
			account_id, err := strconv.ParseInt(str, 10, 64)

			if err != nil {
				return nil, err
			}

			return &account_id, nil
		}
	}

	return nil, nil
}
