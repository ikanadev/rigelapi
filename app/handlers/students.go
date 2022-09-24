package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/ent/class"
	"github.com/vmkevv/rigelapi/ent/studentsync"
)

func StudentSyncStatus(db *ent.Client) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		resp := SyncIDResp{}
		classID := c.Params("classid")
		studentSync, err := db.StudentSync.Query().Where(studentsync.HasClassWith(class.IDEQ(classID))).First(c.Context())
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

func SaveStudent(db *ent.Client) func(*fiber.Ctx) error {
	type studentToSave struct {
		id       string
		name     string
		lastName string
		ci       string
	}
	return func(c *fiber.Ctx) error {
		students := []SyncReq{}
    log.Println("Parsing req")
		err := c.BodyParser(&students)
		if err != nil {
			return err
		}

    log.Println("Initializing Transaction")
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
    err = tx.Commit()
		if err != nil {
			return err
		}
		return c.JSON(SyncIDResp{lastSyncId})
	}
}
