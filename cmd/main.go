package main

import (
	"github.com/RuanRosa/simple-bank/pkg/common/configuration"
	"github.com/RuanRosa/simple-bank/pkg/gateways/database/postgres"
	"github.com/RuanRosa/simple-bank/pkg/gateways/http"
	"go.uber.org/zap"
)

func main() {
	// Create log
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	log := logger.Sugar()

	// Load config.
	config := configuration.NewConfig()
	config.Load()

	// Load database and get connetion.

	log.Info("postgres connection...")

	_, err := postgres.GetConnection(config.PostgresURI())
	if err != nil {
		log.Fatal(err)
	}

	log.Info("postgres successfully connected...")

	// Create api and run.
	http.NewAPI().Start()
}
