package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/ent/attendanceday"
	"github.com/vmkevv/rigelapi/ent/attendancedaysyncs"
	"github.com/vmkevv/rigelapi/ent/class"
	"github.com/vmkevv/rigelapi/ent/classperiod"
	"github.com/vmkevv/rigelapi/ent/teacher"
	"github.com/vmkevv/rigelapi/ent/year"
)

type AttendanceDay struct {
	ID            string `json:"id"`
	Day           int64  `json:"day"`
	ClassPeriodID string `json:"class_period_id"`
}

func AttendanceDaySyncStatus(db *ent.Client) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		teacherID := c.Locals("id").(string)
		resp := SyncIDResp{}
		attDaySync, err := db.
			AttendanceDaySyncs.
			Query().
			Where(attendancedaysyncs.HasTeacherWith(teacher.IDEQ(teacherID))).
			First(c.Context())
		if err != nil {
			if _, ok := err.(*ent.NotFoundError); ok {
				return c.JSON(resp)
			}
			return err
		}
		resp.LastSyncID = attDaySync.LastSyncID
		return c.JSON(resp)
	}
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
		attDays := make([]AttendanceDay, len(dbAttDays))
		for index, attDay := range dbAttDays {
			attDays[index] = AttendanceDay{
				ID:            attDay.ID,
				Day:           attDay.Day.UnixMilli(),
				ClassPeriodID: attDay.Edges.ClassPeriod.ID,
			}
		}
		return c.JSON(attDays)
	}
}

func SaveAttendanceDays(db *ent.Client, newID func() string) func(*fiber.Ctx) error {
	type AttendanceDayReq struct {
		SyncReqBase
		Data AttendanceDay `json:"data"`
	}
	return func(c *fiber.Ctx) error {
		teacherID := c.Locals("id").(string)
		attDays := []AttendanceDayReq{}
		if err := c.BodyParser(&attDays); err != nil {
			return err
		}

		tx, err := db.Tx(c.Context())
		if err != nil {
			return err
		}

		toAdd := []*ent.AttendanceDayCreate{}
		lastSyncID := ""

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
			lastSyncID = attDay.ID
		}

		_, err = tx.AttendanceDay.CreateBulk(toAdd...).Save(c.Context())
		if err != nil {
			return rollback(tx, err)
		}

		attDaySyncFound := true
		attDaySync, err := tx.AttendanceDaySyncs.
			Query().
			Where(attendancedaysyncs.HasTeacherWith(teacher.IDEQ(teacherID))).
			First(c.Context())
		if err != nil {
			if _, isNotFound := err.(*ent.NotFoundError); isNotFound {
				attDaySyncFound = false
			} else {
				return rollback(tx, err)
			}
		}

		if attDaySyncFound {
			_, err = attDaySync.Update().SetLastSyncID(lastSyncID).Save(c.Context())
		} else {
			_, err = tx.AttendanceDaySyncs.
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
