package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/ent/subject"
)

func StaticDataHandler(db *ent.Client) func(*fiber.Ctx) error {
	type Resp struct {
		Grades   []Grade   `json:"grades"`
		Subjects []Subject `json:"subjects"`
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

		gradesResp := make([]Grade, len(grades))
		subjectsResp := make([]Subject, len(subjects))
		for index, grade := range grades {
			gradesResp[index] = Grade{grade.ID, grade.Name}
		}
		for index, subject := range subjects {
			subjectsResp[index] = Subject{subject.ID, subject.Name}
		}

		return c.JSON(Resp{gradesResp, subjectsResp})
	}
}
