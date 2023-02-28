package configuration

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type IEnv interface {
	Load()
	Get() variables
}

type env struct {
	loadedVariables variables
}

func NewEnv() IEnv {
	return &env{}
}

type variables struct {
	DbName     string `envconfig:"DATABASE_NAME"`
	DbUser     string `envconfig:"DATABASE_USER"`
	DbPassword string `envconfig:"DATABASE_PASSWORD"`
	DbHost     string `envconfig:"DATABASE_HOST"`
	DbPort     string `envconfig:"DATABASE_PORT"`
}

func (e *env) Load() {
	godotenv.Load()

	noPrefix := ""
	if err := envconfig.Process(noPrefix, &e.loadedVariables); err != nil {
		log.Fatal(err)
	}
}

func (e *env) Get() variables {
	return e.loadedVariables
}
