package postgres

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

var dbConnection *pgxpool.Pool

func connectPool(uri string) error {
	config, err := pgxpool.ParseConfig(uri)
	if err != nil {
		return err
	}

	dbPool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return err
	}

	if err := dbPool.Ping(context.TODO()); err != nil {
		return err
	}

	dbConnection = dbPool
	return nil
}

func GetConnection(uri string) (*pgxpool.Pool, error) {
	if dbConnection == nil {
		if err := connectPool(uri); err != nil {
			return nil, err
		}
	}

	return dbConnection, nil
}
