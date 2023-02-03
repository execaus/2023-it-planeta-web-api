package models

import (
	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Environment struct {
	PostgresUser     string `env:"POSTGRES_USER,required,notEmpty"`
	PostgresPassword string `env:"POSTGRES_USER,required,notEmpty"`
}

const envFilePath = ".env"

func LoadEnv() *Environment {
	environment := Environment{}
	if err := godotenv.Load(envFilePath); err != nil {
		logrus.Warning("load file not found, environment variables load from environment")
		if err := env.Parse(&environment); err != nil {
			logrus.Fatalf(`environment variables load from environment: %s`, err.Error())
		}
	}

	return &environment
}
