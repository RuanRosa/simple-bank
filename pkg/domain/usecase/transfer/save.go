package transfer

import (
	"context"

	"github.com/RuanRosa/simple-bank/pkg/domain/entities/transfer"
)

func (u *usecase) Save(ctx *context.Context, req *transfer.Entity) (*transfer.Entity, error) {
	return u.repository.Save(ctx, req)
}
