package main

import (
	"github.com/RuanRosa/simple-bank/pkg/gateways/http"
)

func main() {
	http.NewAPI().Start()
}
