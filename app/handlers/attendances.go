package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/ent/attendance"
	"github.com/vmkevv/rigelapi/ent/attendanceday"
	"github.com/vmkevv/rigelapi/ent/class"
	"github.com/vmkevv/rigelapi/ent/classperiod"
	"github.com/vmkevv/rigelapi/ent/teacher"
	"github.com/vmkevv/rigelapi/ent/year"
)

type AttendanceRes struct {
	ID              string           `json:"id"`
	Value           attendance.Value `json:"value"`
	StudentID       string           `json:"student_id"`
	AttendanceDayID string           `json:"attendance_day_id"`
}

func SaveAttendances(db *ent.Client) func(*fiber.Ctx) error {
	type AttendanceReq struct {
		SyncReqBase
		Data AttendanceRes `json:"data"`
	}
	return func(c *fiber.Ctx) error {
		atts := []AttendanceReq{}
		err := c.BodyParser(&atts)
		if err != nil {
			return nil
		}

		tx, err := db.Tx(c.Context())
		if err != nil {
			return nil
		}

		toAdd := []*ent.AttendanceCreate{}
		toUpdate := []AttendanceRes{}
		for _, attReq := range atts {
			switch attReq.Type {
			case Insert:
				{
					toAdd = append(
						toAdd,
						tx.Attendance.
							Create().
							SetID(attReq.Data.ID).
							SetValue(attReq.Data.Value).
							SetStudentID(attReq.Data.StudentID).
							SetAttendanceDayID(attReq.Data.AttendanceDayID),
					)
				}
			case Update:
				{
					toUpdate = append(toUpdate, attReq.Data)
				}
			}
		}

		err = tx.Attendance.CreateBulk(toAdd...).OnConflictColumns(attendance.FieldID).Ignore().Exec(c.Context())
		if err != nil {
			return rollback(tx, err)
		}
		for _, att := range toUpdate {
			_, err = tx.Attendance.UpdateOneID(att.ID).SetValue(att.Value).Save(c.Context())
			if err != nil {
				return rollback(tx, err)
			}
		}

		err = tx.Commit()
		if err != nil {
			return err
		}
		return c.SendStatus(fiber.StatusNoContent)
	}
}

func GetAttendances(db *ent.Client) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		teacherID := c.Locals("id").(string)
		yearID := c.Params("yearid")
		dbAtts, err := db.Attendance.
			Query().
			Where(
				attendance.HasAttendanceDayWith(
					attendanceday.HasClassPeriodWith(
						classperiod.HasClassWith(
							class.HasYearWith(year.IDEQ(yearID)),
							class.HasTeacherWith(teacher.IDEQ(teacherID)),
						),
					),
				),
			).
			WithStudent().
			WithAttendanceDay().
			All(c.Context())
		if err != nil {
			return err
		}

		atts := make([]AttendanceRes, len(dbAtts))
		for index, dbAtt := range dbAtts {
			atts[index] = AttendanceRes{
				ID:              dbAtt.ID,
				Value:           dbAtt.Value,
				StudentID:       dbAtt.Edges.Student.ID,
				AttendanceDayID: dbAtt.Edges.AttendanceDay.ID,
			}
		}

		return c.JSON(atts)
	}
}
