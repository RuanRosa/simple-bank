package transfer

import (
	"context"

	"github.com/RuanRosa/simple-bank/pkg/domain/entities/transfer"
)

func (u *usecase) Transfer(ctx *context.Context, req *transfer.Entity) (*transfer.Entity, error) {
	origin, err := u.accountUsecase.GetByID(ctx, req.AccountOriginID)
	if err != nil {
		return nil, err
	}

	if origin.Balance < req.Amount {
		return nil, transfer.ErrinsufficientFunds
	}

	destinantion, err := u.accountUsecase.GetByID(ctx, req.AccountDestinationID)
	if err != nil {
		return nil, err
	}

	if _, err := u.repository.Save(ctx, req); err != nil {
		return nil, err

	}

	origin.Balance = req.Amount

	if err := u.accountUsecase.Discount(ctx, origin); err != nil {
		return nil, err
	}

	destinantion.Balance = req.Amount

	if err := u.accountUsecase.Deposit(ctx, destinantion); err != nil {
		return nil, err
	}

	return nil, nil
}
