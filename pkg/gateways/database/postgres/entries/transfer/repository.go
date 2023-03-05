package transfer

import "github.com/jackc/pgx/v4/pgxpool"

type repository struct {
	DB *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *repository {
	return &repository{
		DB: db,
	}
}
