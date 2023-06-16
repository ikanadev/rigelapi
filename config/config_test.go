package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvOrPanic(t *testing.T) {
	assert := assert.New(t)
	strVal := "RANDOM"
	os.Setenv(strVal, strVal)
	value := getEnvOrPanic(strVal)
	assert.Equal(strVal, value)
	assert.Panics(func() {
		getEnvOrPanic("INEXISTENT")
	})
}

// DB_NAME=dbtest DB_HOST=hosttest DB_PORT=porttest DB_USER=usertest DB_PASSWORD=passtest DB_SSL_MODE=disable APP_PORT=4000
func TestGetConfig(t *testing.T) {
	assert := assert.New(t)
	config := GetConfig()
	expected := Config{
		DB: dbData{
			Host:     "hosttest",
			Port:     "porttest",
			Name:     "dbtest",
			User:     "usertest",
			Password: "passtest",
			SslMode:  "disable",
		},
		App: appData{
			JWTKey: "ImsKLIZXipqsHJKo_e3z",
			Port:   "4000",
		},
	}
	assert.Equal(expected, config)
}
