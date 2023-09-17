package sync

import (
	"context"
	"time"

	"github.com/vmkevv/rigelapi/app/common"
	"github.com/vmkevv/rigelapi/app/models"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/ent/activity"
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

func (ser SyncEntRepo) GetAttendances(
	teacherID string,
	yearID string,
) ([]models.AppAttendance, error) {
	entAtts, err := ser.ent.Attendance.Query().
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
		All(ser.ctx)
	if err != nil {
		return nil, err
	}
	atts := make([]models.AppAttendance, len(entAtts))
	for i, att := range entAtts {
		atts[i] = models.AppAttendance{
			ID:              att.ID,
			Value:           att.Value,
			StudentID:       att.Edges.Student.ID,
			AttendanceDayID: att.Edges.AttendanceDay.ID,
		}
	}
	return atts, nil
}

func (ser SyncEntRepo) SyncAttendances(attendanceTxs []common.AppAttendanceTx) error {
	tx, err := ser.ent.Tx(ser.ctx)
	if err != nil {
		return err
	}
	toAdd := []*ent.AttendanceCreate{}
	toUpdate := []models.AppAttendance{}
	for _, att := range attendanceTxs {
		switch att.Type {
		case common.Insert:
			{
				toAdd = append(
					toAdd,
					tx.Attendance.Create().
						SetID(att.Data.ID).
						SetValue(att.Data.Value).
						SetStudentID(att.Data.StudentID).
						SetAttendanceDayID(att.Data.AttendanceDayID),
				)
			}
		case common.Update:
			{
				toUpdate = append(toUpdate, att.Data)
			}
		}
	}
	// Add attendances
	err = tx.Attendance.CreateBulk(toAdd...).
		OnConflictColumns(attendance.FieldID).
		Ignore().
		Exec(ser.ctx)
	if err != nil {
		return common.RollbackTx(tx, err)
	}
	// Update attendances
	for _, att := range toUpdate {
		_, err := tx.Attendance.UpdateOneID(att.ID).SetValue(att.Value).Save(ser.ctx)
		if err != nil {
			return common.RollbackTx(tx, err)
		}
	}
	return tx.Commit()
}

func (ser SyncEntRepo) GetActivities(
	teacherID string,
	yearID string,
) ([]models.AppActivity, error) {
	entActivities, err := ser.ent.Activity.Query().
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
		All(ser.ctx)
	if err != nil {
		return nil, err
	}
	activities := make([]models.AppActivity, len(entActivities))
	for i, act := range entActivities {
		activities[i] = models.AppActivity{
			ID:            act.ID,
			Name:          act.Name,
			Date:          act.Date.UnixMilli(),
			AreaId:        act.Edges.Area.ID,
			ClassPeriodId: act.Edges.ClassPeriod.ID,
		}
	}
	return activities, nil
}

func (ser SyncEntRepo) SyncActivities(activityTxs []common.AppActivityTx) error {
	tx, err := ser.ent.Tx(ser.ctx)
	if err != nil {
		return err
	}
	toAdd := []*ent.ActivityCreate{}
	toUpdate := []models.AppActivity{}
	for _, act := range activityTxs {
		switch act.Type {
		case common.Insert:
			{
				toAdd = append(
					toAdd,
					tx.Activity.Create().
						SetID(act.Data.ID).
						SetName(act.Data.Name).
						SetDate(time.UnixMilli(act.Data.Date)).
						SetClassPeriodID(act.Data.ClassPeriodId).
						SetAreaID(act.Data.AreaId),
				)
			}
		case common.Update:
			{
				toUpdate = append(toUpdate, act.Data)
			}
		}
	}
	// Add activities
	err = tx.Activity.CreateBulk(toAdd...).OnConflictColumns(activity.FieldID).Ignore().Exec(ser.ctx)
	if err != nil {
		return common.RollbackTx(tx, err)
	}
	// Update activities
	for _, act := range toUpdate {
		_, err := tx.Activity.UpdateOneID(act.ID).SetName(act.Name).SetAreaID(act.AreaId).Save(ser.ctx)
		if err != nil {
			return common.RollbackTx(tx, err)
		}
	}
	return tx.Commit()
}

func (ser SyncEntRepo) GetScores(teacherID string, yearID string) ([]models.AppScore, error) {
	entScores, err := ser.ent.Score.Query().
		Where(
			score.HasActivityWith(
				activity.HasClassPeriodWith(
					classperiod.HasClassWith(
						class.HasTeacherWith(teacher.IDEQ(teacherID)),
						class.HasYearWith(year.IDEQ(yearID)),
					),
				),
			),
		).
		WithActivity().
		WithStudent().
		All(ser.ctx)
	if err != nil {
		return nil, err
	}
	scores := make([]models.AppScore, len(entScores))
	for i, score := range entScores {
		scores[i] = models.AppScore{
			ID:         score.ID,
			StudentId:  score.Edges.Student.ID,
			ActivityId: score.Edges.Activity.ID,
			Points:     score.Points,
		}
	}
	return scores, nil
}

func (ser SyncEntRepo) SyncScores(scoreTxs []common.AppScoreTx) error {
	tx, err := ser.ent.Tx(ser.ctx)
	if err != nil {
		return err
	}
	toAdd := []*ent.ScoreCreate{}
	toUpdate := []models.AppScore{}
	for _, score := range scoreTxs {
		switch score.Type {
		case common.Insert:
			{
				toAdd = append(
					toAdd,
					tx.Score.Create().
						SetID(score.Data.ID).
						SetPoints(score.Data.Points).
						SetStudentID(score.Data.StudentId).
						SetActivityID(score.Data.ActivityId),
				)
			}
		case common.Update:
			{
				toUpdate = append(toUpdate, score.Data)
			}
		}
	}
	// Add scores
	err = tx.Score.CreateBulk(toAdd...).OnConflictColumns(score.FieldID).Ignore().Exec(ser.ctx)
	if err != nil {
		return common.RollbackTx(tx, err)
	}
	// Update scores
	for _, score := range toUpdate {
		_, err := tx.Score.UpdateOneID(score.ID).SetPoints(score.Points).Save(ser.ctx)
		if err != nil {
			return common.RollbackTx(tx, err)
		}
	}
	return tx.Commit()
}
