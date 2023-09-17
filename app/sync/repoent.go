package sync

import (
	"context"
	"time"

	"github.com/vmkevv/rigelapi/app/common"
	"github.com/vmkevv/rigelapi/app/models"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/ent/attendance"
	"github.com/vmkevv/rigelapi/ent/attendanceday"
	"github.com/vmkevv/rigelapi/ent/class"
	"github.com/vmkevv/rigelapi/ent/classperiod"
	"github.com/vmkevv/rigelapi/ent/score"
	"github.com/vmkevv/rigelapi/ent/student"
	"github.com/vmkevv/rigelapi/ent/teacher"
	"github.com/vmkevv/rigelapi/ent/year"
)

type SyncEntRepo struct {
	ent *ent.Client
	ctx context.Context
}

func NewSyncEntRepo(ent *ent.Client, ctx context.Context) SyncEntRepo {
	return SyncEntRepo{ent, ctx}
}

func (ser SyncEntRepo) GetStudents(teacherID, yearID string) ([]models.AppStudent, error) {
	entStudents, err := ser.ent.Student.
		Query().
		Where(
			student.HasClassWith(
				class.HasTeacherWith(teacher.IDEQ(teacherID)),
				class.HasYearWith(year.IDEQ(yearID)),
			),
		).
		WithClass().
		All(ser.ctx)
	if err != nil {
		return nil, err
	}
	students := make([]models.AppStudent, len(entStudents))
	for i, student := range entStudents {
		students[i] = models.AppStudent{
			ID:       student.ID,
			Name:     student.Name,
			LastName: student.LastName,
			CI:       student.Ci,
			ClassID:  student.Edges.Class.ID,
		}
	}
	return students, nil
}

func (ser SyncEntRepo) SyncStudents(studentTxs []common.AppStudentTx) error {
	tx, err := ser.ent.Tx(ser.ctx)
	if err != nil {
		return err
	}
	toAdd := []*ent.StudentCreate{}
	toUpdate := []models.AppStudent{}
	toDelete := []string{}
	for _, studentTx := range studentTxs {
		switch studentTx.Type {
		case common.Insert:
			{
				toAdd = append(
					toAdd,
					tx.Student.Create().
						SetID(studentTx.Data.ID).
						SetName(studentTx.Data.Name).
						SetLastName(studentTx.Data.LastName).
						SetCi(studentTx.Data.CI).
						SetClassID(studentTx.Data.ClassID),
				)
			}
		case common.Update:
			{
				toUpdate = append(toUpdate, studentTx.Data)
			}
		case common.Delete:
			{
				toDelete = append(toDelete, studentTx.Data.ID)
			}
		}
	}
	// Add students
	err = tx.Student.CreateBulk(toAdd...).OnConflictColumns(student.FieldID).Ignore().Exec(ser.ctx)
	if err != nil {
		return common.RollbackTx(tx, err)
	}
	// Update Students
	for _, st := range toUpdate {
		_, err := tx.Student.UpdateOneID(st.ID).
			SetName(st.Name).
			SetLastName(st.LastName).
			SetCi(st.CI).
			Save(ser.ctx)
		if err != nil {
			return common.RollbackTx(tx, err)
		}
	}
	// Delete student attendances
	_, err = tx.Attendance.Delete().Where(
		attendance.HasStudentWith(student.IDIn(toDelete...)),
	).Exec(ser.ctx)
	if err != nil {
		return common.RollbackTx(tx, err)
	}
	// Delete student scores
	_, err = tx.Score.Delete().Where(
		score.HasStudentWith(student.IDIn(toDelete...)),
	).Exec(ser.ctx)
	if err != nil {
		return common.RollbackTx(tx, err)
	}
	// Delete Students
	_, err = tx.Student.Delete().Where(student.IDIn(toDelete...)).Exec(ser.ctx)
	if err != nil {
		return common.RollbackTx(tx, err)
	}
	return tx.Commit()
}

func (ser SyncEntRepo) GetClassPeriods(
	teacherID string,
	yearID string,
) ([]models.AppClassPeriod, error) {
	entClassPeriods, err := ser.ent.ClassPeriod.
		Query().
		Where(
			classperiod.HasClassWith(
				class.HasYearWith(year.IDEQ(yearID)),
				class.HasTeacherWith(teacher.IDEQ(teacherID)),
			),
		).
		WithClass().
		WithPeriod().
		All(ser.ctx)
	if err != nil {
		return nil, err
	}
	classPeriods := make([]models.AppClassPeriod, len(entClassPeriods))
	for i, cp := range entClassPeriods {
		classPeriods[i] = models.AppClassPeriod{
			ID:       cp.ID,
			Finished: cp.Finished,
			Start:    cp.Start.UnixMilli(),
			End:      cp.End.UnixMilli(),
			ClassID:  cp.Edges.Class.ID,
			Period: models.ClassPeriodPeriod{
				ID:   cp.Edges.Period.ID,
				Name: cp.Edges.Period.Name,
			},
		}
	}
	return classPeriods, nil
}

func (ser SyncEntRepo) SyncClassPeriods(classPeriodTxs []common.AppClassPeriodTx) error {
	tx, err := ser.ent.Tx(ser.ctx)
	if err != nil {
		return err
	}
	toAdd := []*ent.ClassPeriodCreate{}
	toUpdate := []models.AppClassPeriod{}
	for _, cp := range classPeriodTxs {
		switch cp.Type {
		case common.Insert:
			{
				toAdd = append(
					toAdd,
					tx.ClassPeriod.
						Create().
						SetID(cp.Data.ID).
						SetStart(time.UnixMilli(cp.Data.Start)).
						SetEnd(time.UnixMilli(cp.Data.End)).
						SetFinished(cp.Data.Finished).
						SetClassID(cp.Data.ClassID).
						SetPeriodID(cp.Data.Period.ID),
				)
			}
		case common.Update:
			{
				toUpdate = append(toUpdate, cp.Data)
			}
		}
	}
	// Add class periods
	err = tx.ClassPeriod.CreateBulk(toAdd...).
		OnConflictColumns(classperiod.FieldID).
		Ignore().
		Exec(ser.ctx)
	if err != nil {
		return common.RollbackTx(tx, err)
	}
	// Update class periods
	for _, cp := range toUpdate {
		_, err := tx.ClassPeriod.UpdateOneID(cp.ID).
			SetFinished(cp.Finished).
			SetEnd(time.UnixMilli(cp.End)).
			Save(ser.ctx)
		if err != nil {
			return common.RollbackTx(tx, err)
		}
	}
	return tx.Commit()
}

func (ser SyncEntRepo) GetAttendanceDays(
	teacherID string,
	yearID string,
) ([]models.AppAttendanceDay, error) {
	entAttDays, err := ser.ent.AttendanceDay.Query().
		Where(
			attendanceday.HasClassPeriodWith(
				classperiod.HasClassWith(
					class.HasYearWith(year.IDEQ(yearID)),
					class.HasTeacherWith(teacher.IDEQ(teacherID)),
				),
			),
		).
		WithClassPeriod().
		All(ser.ctx)
	if err != nil {
		return nil, err
	}
	attDays := make([]models.AppAttendanceDay, len(entAttDays))
	for i, attDay := range entAttDays {
		attDays[i] = models.AppAttendanceDay{
			ID:            attDay.ID,
			Day:           attDay.Day.UnixMilli(),
			ClassPeriodID: attDay.Edges.ClassPeriod.ID,
		}
	}
	return attDays, nil
}

func (ser SyncEntRepo) SyncAttendanceDays(attendanceDayTxs []common.AppAttendanceDayTx) error {
	tx, err := ser.ent.Tx(ser.ctx)
	if err != nil {
		return err
	}
	toAdd := []*ent.AttendanceDayCreate{}
	for _, attDay := range attendanceDayTxs {
		switch attDay.Type {
		case common.Insert:
			{
				toAdd = append(
					toAdd,
					tx.AttendanceDay.Create().
						SetID(attDay.Data.ID).
						SetDay(time.UnixMilli(attDay.Data.Day)).
						SetClassPeriodID(attDay.Data.ClassPeriodID),
				)
			}
		}
	}
	// Add attendance day
	err = tx.AttendanceDay.CreateBulk(toAdd...).
		OnConflictColumns(attendanceday.FieldID).
		Ignore().
		Exec(ser.ctx)
	if err != nil {
		common.RollbackTx(tx, err)
	}
	return tx.Commit()
}
