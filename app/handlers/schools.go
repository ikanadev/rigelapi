package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/ent/municipio"
	"github.com/vmkevv/rigelapi/ent/school"
)

// SchoolsHandler godoc
// @Summary List all schools from municipio with {munid}
// @Produce json
// @Param   munid path     int true "Municipio ID"
// @Success 200    {object} handlers.MunsHandler.res
// @Router  /schools/mun/{munid} [get]
func SchoolsHandler(db *ent.Client) func(*fiber.Ctx) error {
	type res = *ent.School
	return func(c *fiber.Ctx) error {
		munID := c.Params("munid")
		schools, err := db.School.Query().Where(school.HasMunicipioWith(municipio.ID(munID))).All(c.Context())
		if err != nil {
			return err
		}
		return c.JSON(schools)
	}
}
