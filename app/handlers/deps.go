package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/ent/dpto"
)

func DepsHandler(db *ent.Client) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		deps, err := db.Dpto.Query().Order(ent.Asc(dpto.FieldName)).All(c.Context())
		if err != nil {
			return err
		}
		depsRes := make([]Dpto, len(deps))
		for i, dep := range deps {
			depsRes[i] = Dpto{dep.ID, dep.Name}
		}
		return c.JSON(depsRes)
	}
}
