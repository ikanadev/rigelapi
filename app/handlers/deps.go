package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/ent"
)

// DepsHandler godoc
// @Summary List all deps
// @Produce json
// @Success 200 {object} handlers.DepsHandler.res
// @Router  /deps [get]
func DepsHandler(db *ent.Client) func(*fiber.Ctx) error {
	type res = []*ent.Dpto
	return func(c *fiber.Ctx) error {
		deps, err := db.Dpto.Query().All(c.Context())
		if err != nil {
			return err
		}
		return c.JSON(deps)
	}
}
