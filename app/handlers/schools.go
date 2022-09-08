package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/ent/municipio"
	"github.com/vmkevv/rigelapi/ent/school"
)

func SchoolsHandler(db *ent.Client) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		munID := c.Params("munid")
		schools, err := db.School.Query().Where(school.HasMunicipioWith(municipio.ID(munID))).All(c.Context())
		if err != nil {
			return err
		}
		return c.JSON(schools)
	}
}
