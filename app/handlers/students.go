package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/ent/class"
	"github.com/vmkevv/rigelapi/ent/student"
	"github.com/vmkevv/rigelapi/ent/studentsync"
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

func StudentSyncStatus(db *ent.Client) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		teacherID := c.Locals("id").(string)
		resp := SyncIDResp{}
		studentSync, err := db.StudentSync.Query().Where(studentsync.HasTeacherWith(teacher.IDEQ(teacherID))).First(c.Context())
		if err != nil {
			if _, ok := err.(*ent.NotFoundError); ok {
				return c.JSON(resp)
			}
			return err
		}
		resp.LastSyncID = studentSync.LastSyncID
		return c.JSON(resp)
	}
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

func SaveStudent(db *ent.Client, newID func() string) func(*fiber.Ctx) error {
	type ReqStudent struct {
		SyncReqBase
		Data Student `json:"data"`
	}
	return func(c *fiber.Ctx) error {
		teacherID := c.Locals("id").(string)
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
		lastSyncId := ""
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
			}
			lastSyncId = st.ID
		}
		_, err = tx.Student.CreateBulk(toAdd...).Save(c.Context())
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

		// Check if sync was done before
		studentSync, err := tx.StudentSync.Query().Where(
			studentsync.HasTeacherWith(teacher.IDEQ(teacherID)),
		).First(c.Context())
		studentSyncFound := true
		if err != nil {
			if _, isNotFound := err.(*ent.NotFoundError); isNotFound {
				studentSyncFound = false
			} else {
				return err
			}
		}

		if studentSyncFound {
			_, err = studentSync.Update().SetLastSyncID(lastSyncId).Save(c.Context())
		} else {
			_, err = tx.StudentSync.Create().
				SetID(newID()).
				SetLastSyncID(lastSyncId).
				SetTeacherID(teacherID).
				Save(c.Context())
		}
		if err != nil {
			return rollback(tx, err)
		}

		err = tx.Commit()
		if err != nil {
			return err
		}
		return c.JSON(SyncIDResp{lastSyncId})
	}
}
