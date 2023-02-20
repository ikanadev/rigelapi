package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/ent/subscription"
	"github.com/vmkevv/rigelapi/ent/teacher"
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

func GetTeacher(db *ent.Client) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		teacherID := c.Params("id")
		teacher, err := db.Teacher.
			Query().
			Where(teacher.ID(teacherID)).
			WithSubscriptions(func(sq *ent.SubscriptionQuery) {
				sq.WithYear()
				sq.Order(ent.Asc(subscription.FieldDate))
			}).
			First(c.Context())
		if err != nil {
			return err
		}
    resp := buildTeacherProfile(teacher)
		return c.JSON(resp)
	}
}

// Get teacher profile and subs by token provided in headers
func GetProfile(db *ent.Client) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		teacherID := c.Locals("id").(string)
		teacher, err := db.Teacher.
			Query().
			Where(teacher.ID(teacherID)).
			WithSubscriptions(func(sq *ent.SubscriptionQuery) {
				sq.WithYear()
				sq.Order(ent.Asc(subscription.FieldDate))
			}).
			First(c.Context())
		if err != nil {
			return err
		}
    resp := buildTeacherProfile(teacher)
		return c.JSON(resp)
	}
}
