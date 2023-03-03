package main

import (
	"log"

	"github.com/RuanRosa/simple-bank/pkg/common/configuration"
	"github.com/RuanRosa/simple-bank/pkg/domain/usecase/account"
	"github.com/RuanRosa/simple-bank/pkg/gateways/database/postgres"
	"github.com/RuanRosa/simple-bank/pkg/gateways/http"
	auth_middleware "github.com/RuanRosa/simple-bank/pkg/gateways/http/middleware/auth"

	account_repository "github.com/RuanRosa/simple-bank/pkg/gateways/database/postgres/entries/account"
	"github.com/RuanRosa/simple-bank/pkg/gateways/service/auth"
	"github.com/sirupsen/logrus"
)

func main() {
	// Load config.
	config := configuration.NewConfig()
	config.Load()

	// Load database and get connetion.
	logrus.Info("connecting in postgres...")

	db, err := postgres.GetConnection(config.PostgresDSN())
	if err != nil {
		log.Fatal(err)
	}

	logrus.Info("successfully connected !")

	// Load migrate.
	if err := postgres.Migrate(config.PostgresURL()); err != nil {
		logrus.Fatal(err)
	}

	// Repository load constructor
	accountRepository := account_repository.NewRepository(db)

	// Usecase load constructor
	accountUsecase := account.NewUsecase(accountRepository)

	// Service load constructor
	authService := auth.NewService(accountUsecase, config.ENV().AccessSecret)

	// Middleware load constructor
	middleware := auth_middleware.NewMiddleware(authService)

	// Create api and run.
	http.NewAPI(config.ENV().Port).Start(accountUsecase, authService, middleware)
}
