package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/ent/class"
	"github.com/vmkevv/rigelapi/ent/teacher"
	"github.com/vmkevv/rigelapi/ent/year"
)

type ClassResp struct {
	Class
	Subject Subject `json:"subject"`
	Grade   Grade   `json:"grade"`
	Year    Year    `json:"year"`
}

func entClassToRespClass(entClasses []*ent.Class) []ClassResp {
	classesResp := make([]ClassResp, len(entClasses))
	for i, class := range entClasses {
		classesResp[i] = ClassResp{
			Class: Class{
				ID:       class.ID,
				Parallel: class.Parallel,
			},
			Subject: Subject{
				ID:   class.Edges.Subject.ID,
				Name: class.Edges.Subject.Name,
			},
			Grade: Grade{
				ID:   class.Edges.Grade.ID,
				Name: class.Edges.Grade.Name,
			},
			Year: Year{
				ID:    class.Edges.Year.ID,
				Value: class.Edges.Year.Value,
			},
		}
	}
	return classesResp
}

func ClassListHandler(db *ent.Client) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		teacherID := c.Locals("id").(string)
		yearID := c.Params("yearid")
		classes, err := db.Class.
			Query().
			Where(
				class.HasTeacherWith(teacher.ID(teacherID)),
				class.HasYearWith(year.ID(yearID)),
			).
			WithGrade().
			WithSubject().
			WithYear().
			All(c.Context())
		if err != nil {
			return err
		}
		classesResp := entClassToRespClass(classes)
		return c.JSON(classesResp)
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
			SetYearID(reqData.YearID).
			SetParallel(reqData.Parallel).
			Save(c.Context())
		if err != nil {
			return err
		}
		classes, err := db.Class.
			Query().
			Where(
				class.HasTeacherWith(teacher.ID(teacherID)),
				class.HasYearWith(year.ID(reqData.YearID)),
			).
			WithGrade().
			WithSubject().
			WithYear().
			All(c.Context())
		classesResp := entClassToRespClass(classes)
		return c.JSON(classesResp)
	}
}
