package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/ent/attendanceday"
	"github.com/vmkevv/rigelapi/ent/class"
	"github.com/vmkevv/rigelapi/ent/classperiod"
	"github.com/vmkevv/rigelapi/ent/teacher"
	"github.com/vmkevv/rigelapi/ent/year"
)

type AttendanceDayRes struct {
	ID            string `json:"id"`
	Day           int64  `json:"day"`
	ClassPeriodID string `json:"class_period_id"`
}

func GetAttendanceDays(db *ent.Client) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		teacherID := c.Locals("id").(string)
		yearID := c.Params("yearid")
		dbAttDays, err := db.AttendanceDay.
			Query().
			Where(
				attendanceday.HasClassPeriodWith(
					classperiod.HasClassWith(
						class.HasYearWith(year.IDEQ(yearID)),
						class.HasTeacherWith(teacher.IDEQ(teacherID)),
					),
				),
			).
			WithClassPeriod().
			All(c.Context())
		if err != nil {
			return err
		}
		attDays := make([]AttendanceDayRes, len(dbAttDays))
		for index, attDay := range dbAttDays {
			attDays[index] = AttendanceDayRes{
				ID:            attDay.ID,
				Day:           attDay.Day.UnixMilli(),
				ClassPeriodID: attDay.Edges.ClassPeriod.ID,
			}
		}
		return c.JSON(attDays)
	}
}

func SaveAttendanceDays(db *ent.Client) func(*fiber.Ctx) error {
	type AttendanceDayReq struct {
		SyncReqBase
		Data AttendanceDayRes `json:"data"`
	}
	return func(c *fiber.Ctx) error {
		attDays := []AttendanceDayReq{}
		if err := c.BodyParser(&attDays); err != nil {
			return err
		}

		tx, err := db.Tx(c.Context())
		if err != nil {
			return err
		}

		toAdd := []*ent.AttendanceDayCreate{}

		for _, attDay := range attDays {
			switch attDay.Type {
			case Insert:
				{
					toAdd = append(
						toAdd,
						tx.AttendanceDay.
							Create().
							SetID(attDay.Data.ID).
							SetDay(time.UnixMilli(attDay.Data.Day)).
							SetClassPeriodID(attDay.Data.ClassPeriodID),
					)
				}
			}
		}

		_, err = tx.AttendanceDay.CreateBulk(toAdd...).Save(c.Context())
		if err != nil {
			return rollback(tx, err)
		}

		err = tx.Commit()
		if err != nil {
			return rollback(tx, err)
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}
