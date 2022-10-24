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
		schools, err := db.School.Query().Where(school.HasMunicipioWith(municipio.ID(munID))).Order(ent.Asc(school.FieldName)).All(c.Context())
		if err != nil {
			return err
		}
		schoolsRes := make([]School, len(schools))
		for i, school := range schools {
			schoolsRes[i] = School{
				school.ID,
				school.Name,
				school.Lat,
				school.Lon,
			}
		}
		return c.JSON(schoolsRes)
	}
}
