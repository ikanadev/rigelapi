package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/ent/activity"
	"github.com/vmkevv/rigelapi/ent/class"
	"github.com/vmkevv/rigelapi/ent/classperiod"
	"github.com/vmkevv/rigelapi/ent/teacher"
	"github.com/vmkevv/rigelapi/ent/year"
)

type Activity struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	ClassPeriodId string `json:"class_period_id"`
	AreaId        string `json:"area_id"`
	Date          int64  `json:"date"`
}

func GetActivities(db *ent.Client) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		teacherID := c.Locals("id").(string)
		yearID := c.Params("yearid")
		serverActivities, err := db.Activity.
			Query().
			Where(
				activity.HasClassPeriodWith(
					classperiod.HasClassWith(
						class.HasTeacherWith(teacher.IDEQ(teacherID)),
						class.HasYearWith(year.IDEQ(yearID)),
					),
				),
			).
			WithArea().
			WithClassPeriod().
			All(c.Context())
		if err != nil {
			return err
		}
		activities := make([]Activity, len(serverActivities))
		for index, act := range serverActivities {
			activities[index] = Activity{
				ID:            act.ID,
				Name:          act.Name,
				Date:          act.Date.UnixMilli(),
				AreaId:        act.Edges.Area.ID,
				ClassPeriodId: act.Edges.ClassPeriod.ID,
			}
		}
		return c.JSON(activities)
	}
}

func SaveActivities(db *ent.Client) func(*fiber.Ctx) error {
	type ActivityReq struct {
		SyncReqBase
		Data Activity `json:"data"`
	}
	return func(c *fiber.Ctx) error {
		acts := []ActivityReq{}
		err := c.BodyParser(&acts)
		if err != nil {
			return err
		}
		tx, err := db.Tx(c.Context())
		if err != nil {
			return err
		}

		toAdd := []*ent.ActivityCreate{}
		toUpdate := []Activity{}
		for _, act := range acts {
			switch act.Type {
			case Insert:
				{
					toAdd = append(
						toAdd,
						tx.Activity.
							Create().
							SetID(act.Data.ID).
							SetName(act.Data.Name).
							SetDate(time.UnixMilli(act.Data.Date)).
							SetClassPeriodID(act.Data.ClassPeriodId).
							SetAreaID(act.Data.AreaId),
					)
				}
			case Update:
				{
					toUpdate = append(toUpdate, act.Data)
				}
			}
		}

		_, err = tx.Activity.CreateBulk(toAdd...).Save(c.Context())
		if err != nil {
			return rollback(tx, err)
		}
		for _, act := range toUpdate {
			_, err = tx.Activity.UpdateOneID(act.ID).SetName(act.Name).SetAreaID(act.AreaId).Save(c.Context())
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
