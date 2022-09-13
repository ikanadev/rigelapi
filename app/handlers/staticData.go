package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/ent/subject"
)

func StaticDataHandler(db *ent.Client) func(*fiber.Ctx) error {
	type resp struct {
		Grades   []*ent.Grade   `json:"grades"`
		Subjects []*ent.Subject `json:"subjects"`
	}
	return func(c *fiber.Ctx) error {
		grades, err := db.Grade.Query().All(c.Context())
		if err != nil {
			return err
		}
		subjects, err := db.Subject.Query().Order(ent.Asc(subject.FieldName)).All(c.Context())
		if err != nil {
			return err
		}
		return c.JSON(resp{grades, subjects})
	}
}
