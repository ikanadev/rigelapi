package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/ent/class"
)

func StatsHandler(db *ent.Client) func(*fiber.Ctx) error {
	type Stats struct {
		Teachers   int `json:"teachers"`
		Classes    int `json:"classes"`
		Schools    int `json:"schools"`
		Activities int `json:"activities"`
	}
	type Resp struct {
		Stats Stats `json:"stats"`
	}
	return func(c *fiber.Ctx) error {
		teachers, err := db.Teacher.Query().Count(c.Context())
		if err != nil {
			return err
		}

		classes, err := db.Class.Query().Count(c.Context())
		if err != nil {
			return err
		}

		schoolNames, err := db.Class.Query().
			Unique(true).
			Select(class.SchoolColumn).
			Strings(c.Context())
		if err != nil {
			return err
		}

		activities, err := db.Activity.Query().Count(c.Context())
		if err != nil {
			return err
		}

		return c.JSON(Resp{
			Stats: Stats{
				Teachers:   teachers,
				Classes:    classes,
				Schools:    len(schoolNames),
				Activities: activities,
			},
		})
	}
}
