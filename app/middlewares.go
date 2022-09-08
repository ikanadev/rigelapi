package app

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/vmkevv/rigelapi/app/handlers"
	"github.com/vmkevv/rigelapi/config"
)

func authMiddleware(config config.Config) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		tokenStr := c.Get("Authorization")
		if len(tokenStr) == 0 {
			return handlers.NewClientErr(fiber.StatusUnauthorized, "Not authorized")
		}
		token, err := jwt.ParseWithClaims(tokenStr, &handlers.AppClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(config.App.JWTKey), nil
		})
		if err != nil {
			return err
		}
		claims, ok := token.Claims.(*handlers.AppClaims)
		if !token.Valid || !ok {
			return errors.New("Could not parse token.")
		}
		c.Locals("id", claims.ID)
		return c.Next()
	}
}
