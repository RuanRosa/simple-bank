package account

import (
	"time"

	"github.com/RuanRosa/simple-bank/pkg/domain/entities/account"
)

type responseBody struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CPF       string    `json:"cpf"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

type saveRequestBody struct {
	Name    string `json:"name"`
	CPF     string `json:"cpf"`
	Balance int    `json:"balance"`
	Secret  string `json:"secret"`
}

type getBalanceRequest struct {
	ID int `json:"id"`
}

func entityToResponse(entity account.Entity) responseBody {
	return responseBody{
		ID:        entity.ID,
		Name:      entity.Name,
		CPF:       entity.CPF,
		Balance:   entity.Balance,
		CreatedAt: entity.CreatedAt,
	}
}

func requestToEntity(req saveRequestBody) account.Entity {
	return account.Entity{
		Name:    req.Name,
		CPF:     req.CPF,
		Secret:  req.Secret,
		Balance: req.Balance,
	}
}
