package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/ent/class"
	"github.com/vmkevv/rigelapi/ent/classperiod"
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

func SaveClassPeriods(db *ent.Client) func(*fiber.Ctx) error {
	type ClassPeriodReq struct {
		SyncReqBase
		Data ClassPeriod `json:"data"`
	}
	return func(c *fiber.Ctx) error {
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

		err = tx.Commit()
		if err != nil {
			return rollback(tx, err)
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}
