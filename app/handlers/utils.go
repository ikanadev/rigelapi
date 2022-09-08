package handlers

import "github.com/golang-jwt/jwt/v4"

type ClientErr struct {
	Status  int
	Message string
}

func (c ClientErr) Error() string {
	return c.Message
}
func NewClientErr(code int, msg string) ClientErr {
	return ClientErr{code, msg}
}

type AppClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}
