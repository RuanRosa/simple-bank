package account

import (
	"errors"
	"time"
)

type Entity struct {
	ID        int
	Name      string
	CPF       string
	Secret    string
	Balance   int
	CreatedAt time.Time
}

var ErrCpfAlredyExists error = errors.New("cpf already exists")
