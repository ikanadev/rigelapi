package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/ent"
)

func DepsHandler(db *ent.Client) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		deps, err := db.Dpto.Query().All(c.Context())
		if err != nil {
			return err
		}
		return c.JSON(deps)
	}
}
