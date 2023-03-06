package account

import (
	"context"
	"time"

	"github.com/bxcodec/faker/v3"
)

type RepositorySPY struct {
	GetAccountsEntity []Entity
	GetAccountsError  error
	GetByCpfEntity    *Entity
	GetByCpfError     error
	GetByIdEntity     *Entity
	GetByIdError      error
	SaveError         error
	DiscountError     error
	DepositError      error
}

func (r *RepositorySPY) GetAccounts(ctx *context.Context) ([]Entity, error) {
	if r.GetAccountsError == nil {
		faker.FakeData(&r.GetAccountsEntity)
	}

	return r.GetAccountsEntity, r.GetAccountsError
}

func (r *RepositorySPY) GetByCPF(ctx context.Context, CPF string) (*Entity, error) {
	return r.GetByCpfEntity, r.GetByCpfError
}

func (r *RepositorySPY) GetByID(ctx context.Context, AccountID int) (*Entity, error) {
	return r.GetByIdEntity, r.GetByIdError
}

func (r *RepositorySPY) Save(ctx *context.Context, req *Entity) error {
	if r.SaveError == nil {
		req.ID = 1
		req.CreatedAt = time.Now()
	}

	return r.SaveError
}

func (r *RepositorySPY) Discount(ctx *context.Context, req *Entity) error {
	return r.DiscountError
}

func (r *RepositorySPY) Deposit(ctx *context.Context, req *Entity) error {
	return r.DepositError
}
