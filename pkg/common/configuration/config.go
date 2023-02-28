package configuration

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type IConfig interface {
	Load()
	ENV() variables
	PostgresURI() string
}

type config struct {
	loadedVariables variables
}

func NewConfig() IConfig {
	return &config{}
}

type variables struct {
	DbName     string `envconfig:"DATABASE_NAME"`
	DbUser     string `envconfig:"DATABASE_USER"`
	DbPassword string `envconfig:"DATABASE_PASSWORD"`
	DbHost     string `envconfig:"DATABASE_HOST"`
	DbPort     string `envconfig:"DATABASE_PORT"`
}

func (c *config) Load() {
	godotenv.Load()

	noPrefix := ""
	if err := envconfig.Process(noPrefix, &c.loadedVariables); err != nil {
		log.Fatal(err)
	}
}

func (c *config) ENV() variables {
	return c.loadedVariables
}

func (c *config) PostgresURI() string {
	env := c.ENV()

	return fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s",
		env.DbUser, env.DbPassword, env.DbHost, env.DbPort, env.DbName,
	)
}
