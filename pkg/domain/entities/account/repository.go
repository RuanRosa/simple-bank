package account

import (
	"context"
)

type IRepository interface {
	GetAccounts(ctx *context.Context) ([]Entity, error)
	Save(ctx *context.Context, req *Entity) error
	GetByCPF(ctx context.Context, CPF string) (*Entity, error)
	GetByID(ctx context.Context, AccountID int) (*Entity, error)
}
