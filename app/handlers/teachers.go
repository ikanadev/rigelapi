package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/ent"
)

func GetTeachers(db *ent.Client) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		teachers, err := db.Teacher.Query().All(c.Context())
		if err != nil {
			return err
		}
		teachersResp := make([]Teacher, len(teachers))
		for i, teacher := range teachers {
			teachersResp[i] = Teacher{
				ID:       teacher.ID,
				Name:     teacher.Name,
				Email:    teacher.Email,
				LastName: teacher.LastName,
				IsAdmin:  teacher.IsAdmin,
			}
		}
		return c.JSON(teachersResp)
	}
}
