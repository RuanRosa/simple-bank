package postgres

import (
	"embed"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

//go:embed migrations/*.sql
var fs embed.FS

func Migrate(databaseURL string) error {
	driver, err := iofs.New(fs, "migrations")
	if err != nil {
		return fmt.Errorf("could not get embedded migration files: %w", err)
	}

	migration, err := migrate.NewWithSourceInstance("iofs", driver, databaseURL)
	if err != nil {
		return fmt.Errorf("could not read migration files: %w", err)
	}

	err = migration.Up()

	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			logrus.Info("migrations successfully read, no changes")

			return nil
		}

		return err
	}

	logrus.Info("migrations successfully read and run")

	return nil
}
