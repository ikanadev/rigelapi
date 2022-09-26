package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/ent/studentsync"
	"github.com/vmkevv/rigelapi/ent/teacher"
)

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

func SaveStudent(db *ent.Client, newID func() string) func(*fiber.Ctx) error {
	type studentToSave struct {
		id       string
		name     string
		lastName string
		ci       string
	}
	return func(c *fiber.Ctx) error {
		teacherID := c.Locals("id").(string)
		students := []SyncReq{}
		err := c.BodyParser(&students)
		if err != nil {
			return err
		}
    // Check if sync was done before
		studentSync, err := db.StudentSync.Query().Where(studentsync.HasTeacherWith(teacher.IDEQ(teacherID))).First(c.Context())
		studentSyncFound := true
		if err != nil {
			if _, isNotFound := err.(*ent.NotFoundError); isNotFound {
				studentSyncFound = false
			} else {
				return err
			}
		}

		tx, err := db.Tx(c.Context())
		if err != nil {
			return err
		}
		toAdd := []*ent.StudentCreate{}
		toUpdate := []studentToSave{}
		lastSyncId := ""
		for _, st := range students {
			switch st.Type {
			case Insert:
				{
					toAdd = append(
						toAdd,
						db.Student.Create().
							SetID(st.Data["id"].(string)).
							SetName(st.Data["name"].(string)).
							SetLastName(st.Data["last_name"].(string)).
							SetCi(st.Data["ci"].(string)).
							SetClassID(st.Data["class_id"].(string)),
					)
				}
			case Update:
				{
					toUpdate = append(toUpdate, studentToSave{
						st.Data["id"].(string),
						st.Data["name"].(string),
						st.Data["last_name"].(string),
						st.Data["ci"].(string),
					})
				}
			}
			lastSyncId = st.ID
		}
		_, err = tx.Student.CreateBulk(toAdd...).Save(c.Context())
		if err != nil {
			return rollback(tx, err)
		}
		for _, st := range toUpdate {
			_, err := tx.Student.UpdateOneID(st.id).SetName(st.name).SetLastName(st.lastName).SetCi(st.ci).Save(c.Context())
			if err != nil {
				return rollback(tx, err)
			}
		}
		if studentSyncFound {
			_, err := studentSync.Update().SetLastSyncID(lastSyncId).Save(c.Context())
			if err != nil {
				return err
			}
		} else {
			_, err = tx.StudentSync.Create().SetID(newID()).SetLastSyncID(lastSyncId).SetTeacherID(teacherID).Save(c.Context())
			if err != nil {
				return rollback(tx, err)
			}
		}
		err = tx.Commit()
		if err != nil {
			return err
		}
		return c.JSON(SyncIDResp{lastSyncId})
	}
}
