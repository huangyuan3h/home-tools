package main

import (
	"os"

	"github.com/golobby/dotenv"
)

type Config struct {
	Database struct {
		dsn string `env:"DB_dsn"`
	}
}

func GetConfig() Config {

	config := Config{}
	file, err := os.Open(".env")

	if err != nil {
		panic(err)
	}

	err = dotenv.NewDecoder(file).Decode(&config)

	if err != nil {
		panic(err)
	}

	return config
}
