package sync

import (
	"context"

	"github.com/vmkevv/rigelapi/app/models"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/ent/class"
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
