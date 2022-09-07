package config

import "os"

type dbData struct {
	Host     string
	Name     string
	Port     string
	User     string
	Password string
}
type appData struct {
	JWTKey string
	Port   string
}

type Config struct {
	DB  dbData
	App appData
}

func GetConfig() Config {
	return Config{
		DB: dbData{
			Host:     os.Getenv("DB_HOST"),
			Name:     os.Getenv("DB_NAME"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
		},
		App: appData{
			JWTKey: "ImsKLIZXipqsHJKo_e3z",
			Port:   "4000",
		},
	}
}
