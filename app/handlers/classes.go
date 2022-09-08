package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/ent/class"
	"github.com/vmkevv/rigelapi/ent/teacher"
)

func ClassListHandler(db *ent.Client) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		teacherID := c.Locals("id").(string)
		classes, err := db.Class.Query().Where(class.HasTeacherWith(teacher.ID(teacherID))).All(c.Context())
		if err != nil {
			return err
		}
		return c.JSON(classes)
	}
}

func NewClassHandler(db *ent.Client, newID func() string) func(*fiber.Ctx) error {
	type req struct {
		GradeID   string `json:"gradeId"`
		SubjectID string `json:"subjectId"`
		SchoolID  string `json:"schoolId"`
		YearID    string `json:"yearId"`
		Parallel  string `json:"parallel"`
	}
	return func(c *fiber.Ctx) error {
		teacherID := c.Locals("id").(string)
		var reqData req
		err := c.BodyParser(&reqData)
		if err != nil {
			return err
		}
		_, err = db.Class.
			Create().
			SetID(newID()).
			SetTeacherID(teacherID).
			SetGradeID(reqData.GradeID).
			SetSubjectID(reqData.SubjectID).
			SetSchoolID(reqData.SchoolID).
			SetParallel(reqData.Parallel).Save(c.Context())
		if err != nil {
			return err
		}
		return nil
	}
}
