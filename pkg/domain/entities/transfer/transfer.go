package transfer

import (
	"errors"
	"time"
)

type Entity struct {
	ID                   int
	AccountOriginID      int
	AccountDestinationID int
	Amount               int
	CreatedAt            time.Time
}

var ErrinsufficientFunds error = errors.New("insufficient funds")
