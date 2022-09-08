package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/ent"
)

func YearlyDataHandler(db *ent.Client) func(*fiber.Ctx) error {
  return func(c *fiber.Ctx) error {
    years, err := db.Year.Query().WithAreas().WithPeriods().All(c.Context())
    if err != nil {
      return err
    }
    return c.JSON(years)
  }
}
