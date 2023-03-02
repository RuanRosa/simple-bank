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
	PostgresDSN() string
	PostgresURL() string
}

type config struct {
	loadedVariables variables
	postgresDSN     string
	postgresURL     string
}

func NewConfig() IConfig {
	return &config{}
}

type variables struct {
	DbName       string `envconfig:"DATABASE_NAME"`
	DbUser       string `envconfig:"DATABASE_USER"`
	DbPassword   string `envconfig:"DATABASE_PASSWORD"`
	DbHost       string `envconfig:"DATABASE_HOST"`
	DbPort       string `envconfig:"DATABASE_PORT"`
	DbSSL        string `envconfig:"DATABASE_SSL"`
	Port         string `envconfig:"PORT"`
	AccessSecret string `envconfig:"ACCESS_SECRET"`
}

func (c *config) Load() {
	godotenv.Load()

	noPrefix := ""
	if err := envconfig.Process(noPrefix, &c.loadedVariables); err != nil {
		log.Fatal(err)
	}

	c.postgresDSN = fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s",
		c.loadedVariables.DbUser,
		c.loadedVariables.DbPassword,
		c.loadedVariables.DbHost,
		c.loadedVariables.DbPort,
		c.loadedVariables.DbName,
	)

	sslMode := new(string)

	if c.loadedVariables.DbSSL == "" {
		*sslMode = "disable"
	}

	c.postgresURL = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.loadedVariables.DbUser,
		c.loadedVariables.DbPassword,
		c.loadedVariables.DbHost,
		c.loadedVariables.DbPort,
		c.loadedVariables.DbName,
		*sslMode,
	)
}

func (c *config) ENV() variables {
	return c.loadedVariables
}

func (c *config) PostgresDSN() string {
	return c.postgresDSN
}

func (c *config) PostgresURL() string {
	return c.postgresURL
}
