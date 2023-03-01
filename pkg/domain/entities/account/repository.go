package account

import (
	"context"
)

type IRepository interface {
	GetAccounts(ctx context.Context) ([]Entity, error)
}
