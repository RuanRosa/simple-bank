package transfer

import "context"

type IRepository interface {
	Save(ctx *context.Context, transferInstance *Entity) (*Entity, error)
}
