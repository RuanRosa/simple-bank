package main

import (
	"log"

	"github.com/RuanRosa/simple-bank/pkg/common/configuration"
	"github.com/RuanRosa/simple-bank/pkg/gateways/database/postgres"
	"github.com/RuanRosa/simple-bank/pkg/gateways/http"
	"github.com/sirupsen/logrus"
)

func main() {
	// Load config.
	config := configuration.NewConfig()
	config.Load()

	// Load database and get connetion.
	logrus.Info("connecting in postgres...")

	_, err := postgres.GetConnection(config.PostgresDSN())
	if err != nil {
		log.Fatal(err)
	}

	logrus.Info("successfully connected !")

	// Load migrate.
	if err := postgres.Migrate(config.PostgresURL()); err != nil {
		logrus.Fatal(err)
	}

	// Create api and run.
	http.NewAPI().Start()
}
