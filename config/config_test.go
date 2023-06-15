package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvOrPanic(t *testing.T) {
	os.Setenv("RANDOM", "RANDOM")
	value := getEnvOrPanic("RANDOM")
	assert.Equal(t, value, "RANDOM")
}
