package utils

import "github.com/jaevor/go-nanoid"

func NanoIDGenerator() func() string {
	gen, _ := nanoid.Standard(21)
	return gen
}
