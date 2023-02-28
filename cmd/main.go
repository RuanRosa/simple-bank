package main

import (
	"github.com/RuanRosa/simple-bank/pkg/common/configuration"
	"github.com/RuanRosa/simple-bank/pkg/gateways/http"
)

func main() {
	env := configuration.NewEnv()
	env.Load()

	http.NewAPI().Start()
}
