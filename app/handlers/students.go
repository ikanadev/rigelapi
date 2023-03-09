package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/ent/attendance"
	"github.com/vmkevv/rigelapi/ent/class"
	"github.com/vmkevv/rigelapi/ent/score"
	"github.com/vmkevv/rigelapi/ent/student"
	"github.com/vmkevv/rigelapi/ent/teacher"
	"github.com/vmkevv/rigelapi/ent/year"
)

type Student struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	CI       string `json:"ci"`
	ClassID  string `json:"class_id"`
}

func GetStudents(db *ent.Client) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		teacherID := c.Locals("id").(string)
		yearID := c.Params("yearid")
		students, err := db.Student.
			Query().
			Where(
				student.HasClassWith(
					class.HasTeacherWith(teacher.IDEQ(teacherID)),
					class.HasYearWith(year.IDEQ(yearID)),
				),
			).
			WithClass().
			All(c.Context())
		if err != nil {
			return err
		}
		studentsResp := make([]Student, len(students))
		for index, student := range students {
			studentsResp[index] = Student{
				ID:       student.ID,
				Name:     student.Name,
				LastName: student.LastName,
				CI:       student.Ci,
				ClassID:  student.Edges.Class.ID,
			}
		}
		return c.JSON(studentsResp)
	}
}

func SaveStudent(db *ent.Client) func(*fiber.Ctx) error {
	type ReqStudent struct {
		SyncReqBase
		Data Student `json:"data"`
	}
	return func(c *fiber.Ctx) error {
		students := []ReqStudent{}
		err := c.BodyParser(&students)
		if err != nil {
			return err
		}

		tx, err := db.Tx(c.Context())
		if err != nil {
			return err
		}
		toAdd := []*ent.StudentCreate{}
		toUpdate := []Student{}
		toDelete := []string{}
		for _, st := range students {
			switch st.Type {
			case Insert:
				{
					toAdd = append(
						toAdd,
						tx.Student.Create().
							SetID(st.Data.ID).
							SetName(st.Data.Name).
							SetLastName(st.Data.LastName).
							SetCi(st.Data.CI).
							SetClassID(st.Data.ClassID),
					)
				}
			case Update:
				{
					toUpdate = append(toUpdate, st.Data)
				}
			case Delete:
				{
					toDelete = append(toDelete, st.Data.ID)
				}
			}
		}
		err = tx.Student.CreateBulk(toAdd...).OnConflict().Ignore().Exec(c.Context())
		if err != nil {
			return rollback(tx, err)
		}
		for _, st := range toUpdate {
			_, err := tx.Student.
				UpdateOneID(st.ID).
				SetName(st.Name).
				SetLastName(st.LastName).
				SetCi(st.CI).
				Save(c.Context())
			if err != nil {
				return rollback(tx, err)
			}
		}

		// Delete student attendances
		_, err = tx.Attendance.Delete().Where(attendance.HasStudentWith(student.IDIn(toDelete...))).Exec(c.Context())
		if err != nil {
			return rollback(tx, err)
		}
		// Delete student scores
		_, err = tx.Score.Delete().Where(score.HasStudentWith(student.IDIn(toDelete...))).Exec(c.Context())
		if err != nil {
			return rollback(tx, err)
		}
		// Delete student
		_, err = tx.Student.Delete().Where(student.IDIn(toDelete...)).Exec(c.Context())
		if err != nil {
			return rollback(tx, err)
		}

		err = tx.Commit()
		if err != nil {
			return err
		}
		return c.SendStatus(fiber.StatusNoContent)
	}
}
