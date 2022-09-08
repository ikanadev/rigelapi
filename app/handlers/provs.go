package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/ent/dpto"
	"github.com/vmkevv/rigelapi/ent/provincia"
)

func ProvsHandler(db *ent.Client) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		depID := c.Params("depid")
		provs, err := db.Provincia.Query().Where(provincia.HasDepartamentoWith(dpto.ID(depID))).All(c.Context())
		if err != nil {
			return err
		}
		return c.JSON(provs)
	}
}
