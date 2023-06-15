package config

import (
	"fmt"
	"os"
)

type dbData struct {
	Host     string
	Name     string
	Port     string
	User     string
	Password string
	SslMode  string
}
type appData struct {
	JWTKey string
	Port   string
}

type Config struct {
	DB  dbData
	App appData
}

func getEnvOrPanic(name string) string {
	value, exists := os.LookupEnv(name)
	if !exists {
		panic(fmt.Sprintf("env variable: %s is not set", name))
	}
	return value
}

func GetConfig() Config {
	return Config{
		DB: dbData{
			Host:     getEnvOrPanic("DB_HOST"),
			Name:     getEnvOrPanic("DB_NAME"),
			Port:     getEnvOrPanic("DB_PORT"),
			User:     getEnvOrPanic("DB_USER"),
			Password: getEnvOrPanic("DB_PASSWORD"),
			SslMode:  getEnvOrPanic("DB_SSL_MODE"),
		},
		App: appData{
			JWTKey: "ImsKLIZXipqsHJKo_e3z",
			Port:   getEnvOrPanic("APP_PORT"),
		},
	}
}
