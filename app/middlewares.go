package app

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/vmkevv/rigelapi/app/common"
	"github.com/vmkevv/rigelapi/config"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/ent/teacher"
)

func authMiddleware(config config.Config) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		tokenStr := c.Get("Authorization")
		if len(tokenStr) == 0 {
			return common.NewClientErr(fiber.StatusUnauthorized, "Not authorized")
		}
		token, err := jwt.ParseWithClaims(tokenStr, &common.AppClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(config.App.JWTKey), nil
		})
		if err != nil {
			return err
		}
		claims, ok := token.Claims.(*common.AppClaims)
		if !token.Valid || !ok {
			return errors.New("Could not parse token.")
		}
		c.Locals("id", claims.ID)
		return c.Next()
	}
}

func adminMiddleware(db *ent.Client) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// id obtained from authMiddleware
		teacherID := c.Locals("id").(string)
		teacher, err := db.Teacher.Query().Where(teacher.IDEQ(teacherID)).First(c.Context())
		if err != nil {
			return err
		}
		if !teacher.IsAdmin {
			return common.NewClientErr(fiber.StatusUnauthorized, "Not authorized")
		}
		return c.Next()
	}
}

func logUserAgent() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		userAgent := c.Get("User-Agent")
		if len(userAgent) > 50 {
			userAgent = userAgent[0:45] + "..."
		}
		log.Printf("%s:\t%s\t%s\n", c.Method(), c.Path(), userAgent)
		return c.Next()
	}
}
