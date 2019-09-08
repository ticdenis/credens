package user_interface

import (
	"credens/src/shared/user_interface/config"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Env   config.Env
	Debug config.Debug
}

func Bootstrap() Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Error loading .config file")
	}

	return Config{
		config.ParseEnv(os.Getenv("APP_ENV")),
		config.ParseDebug(os.Getenv("APP_DEBUG")),
	}
}
