package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNanoIDGenerator(t *testing.T) {
	assert := assert.New(t)
	gen := NanoIDGenerator()
	uuid := gen()
	assert.Equal(21, len(uuid))
}
