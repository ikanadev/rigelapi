package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/ent/class"
	"github.com/vmkevv/rigelapi/ent/classperiod"
	"github.com/vmkevv/rigelapi/ent/classperiodsync"
	"github.com/vmkevv/rigelapi/ent/teacher"
	"github.com/vmkevv/rigelapi/ent/year"
)

type ClassPeriodPeriod struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type ClassPeriod struct {
	ID       string            `json:"id"`
	Finished bool              `json:"finished"`
	Start    int64             `json:"start"`
	End      int64             `json:"end"`
	Period   ClassPeriodPeriod `json:"period"`
	ClassID  string            `json:"class_id"`
}

func ClassPeriodSyncStatus(db *ent.Client) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		teacherID := c.Locals("id").(string)
		resp := SyncIDResp{}
		classPeriodSync, err := db.
			ClassPeriodSync.
			Query().
			Where(classperiodsync.HasTeacherWith(teacher.IDEQ(teacherID))).
			First(c.Context())
		if err != nil {
			if _, ok := err.(*ent.NotFoundError); ok {
				return c.JSON(resp)
			}
			return err
		}
		resp.LastSyncID = classPeriodSync.LastSyncID
		return c.JSON(resp)
	}
}

func GetClassPeriods(db *ent.Client) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		teacherID := c.Locals("id").(string)
		yearID := c.Params("yearid")
		dbClassPeriods, err := db.ClassPeriod.
			Query().
			Where(
				classperiod.HasClassWith(
					class.HasYearWith(year.IDEQ(yearID)),
					class.HasTeacherWith(teacher.IDEQ(teacherID)),
				),
			).
			WithClass().
			WithPeriod().
			All(c.Context())
		if err != nil {
			return err
		}
		classPeriods := make([]ClassPeriod, len(dbClassPeriods))
		for index, classP := range dbClassPeriods {
			classPeriods[index] = ClassPeriod{
				ID:       classP.ID,
				Start:    classP.Start.UnixMilli(),
				End:      classP.End.UnixMilli(),
				Finished: classP.Finished,
				ClassID:  classP.Edges.Class.ID,
				Period: ClassPeriodPeriod{
					ID:   classP.Edges.Period.ID,
					Name: classP.Edges.Period.Name,
				},
			}
		}
		return c.JSON(classPeriods)
	}
}

func SaveClassPeriods(db *ent.Client, newID func() string) func(*fiber.Ctx) error {
	type ClassPeriodReq struct {
		SyncReqBase
		Data ClassPeriod `json:"data"`
	}
	return func(c *fiber.Ctx) error {
		teacherID := c.Locals("id").(string)
		classPeriods := []ClassPeriodReq{}
		if err := c.BodyParser(&classPeriods); err != nil {
			return err
		}

		tx, err := db.Tx(c.Context())
		if err != nil {
			return err
		}
		toAdd := []*ent.ClassPeriodCreate{}
		toUpdate := []ClassPeriod{}
		lastSyncID := ""

		for _, classP := range classPeriods {
			switch classP.Type {
			case Insert:
				{
					toAdd = append(
						toAdd,
						tx.ClassPeriod.
							Create().
							SetID(classP.Data.ID).
							SetStart(time.UnixMilli(classP.Data.Start)).
							SetEnd(time.UnixMilli(classP.Data.End)).
							SetFinished(classP.Data.Finished).
							SetClassID(classP.Data.ClassID).
							SetPeriodID(classP.Data.Period.ID),
					)
				}
			case Update:
				{
					toUpdate = append(toUpdate, classP.Data)
				}
			}
			lastSyncID = classP.ID
		}

		_, err = tx.ClassPeriod.CreateBulk(toAdd...).Save(c.Context())
		if err != nil {
			return rollback(tx, err)
		}
		for _, classP := range toUpdate {
			_, err := tx.ClassPeriod.
				UpdateOneID(classP.ID).
				SetFinished(classP.Finished).
				SetEnd(time.UnixMilli(classP.End)).
				Save(c.Context())
			if err != nil {
				return rollback(tx, err)
			}
		}

		classPeriodsSyncFound := true
		classPeriodsSync, err := tx.ClassPeriodSync.
			Query().
			Where(classperiodsync.HasTeacherWith(teacher.IDEQ(teacherID))).
			First(c.Context())
		if err != nil {
			if _, isNotFound := err.(*ent.NotFoundError); isNotFound {
				classPeriodsSyncFound = false
			} else {
				return rollback(tx, err)
			}
		}

		if classPeriodsSyncFound {
			_, err = classPeriodsSync.Update().SetLastSyncID(lastSyncID).Save(c.Context())
		} else {
			_, err = tx.ClassPeriodSync.
				Create().
				SetID(newID()).
				SetLastSyncID(lastSyncID).
				SetTeacherID(teacherID).
				Save(c.Context())
		}
		if err != nil {
			return rollback(tx, err)
		}

		err = tx.Commit()
		if err != nil {
			return rollback(tx, err)
		}

		return c.JSON(SyncIDResp{lastSyncID})
	}
}
