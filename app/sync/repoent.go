package sync

import (
	"context"

	"github.com/vmkevv/rigelapi/app/common"
	"github.com/vmkevv/rigelapi/app/models"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/ent/attendance"
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

func (ser SyncEntRepo) GetStudents(teacherID, yearID string) ([]models.Student, error) {
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
	students := make([]models.Student, len(entStudents))
	for i, student := range entStudents {
		students[i] = models.Student{
			ID:       student.ID,
			Name:     student.Name,
			LastName: student.LastName,
			CI:       student.Ci,
			ClassID:  student.Edges.Class.ID,
		}
	}
	return students, nil
}

// SyncStudents implements SyncRepository
func (ser SyncEntRepo) SyncStudents(studentTxs []common.StudentTx) error {
	tx, err := ser.ent.Tx(ser.ctx)
	if err != nil {
		return err
	}
	toAdd := []*ent.StudentCreate{}
	toUpdate := []models.Student{}
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
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (ser SyncEntRepo) GetClassPeriods(
	teacherID string,
	yearID string,
) ([]models.ClassPeriod, error) {
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
	classPeriods := make([]models.ClassPeriod, len(entClassPeriods))
	for i, cp := range entClassPeriods {
		classPeriods[i] = models.ClassPeriod{
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
