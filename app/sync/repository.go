package sync

import (
	"github.com/vmkevv/rigelapi/app/common"
	"github.com/vmkevv/rigelapi/app/models"
)

type SyncRepository interface {
	GetStudents(teacherID, yearID string) ([]models.Student, error)
	SyncStudents(studentTxs []common.StudentTx) error
	GetClassPeriods(teacherID, yearID string) ([]models.ClassPeriod, error)
	SyncClassPeriods(classPeriodTxs []common.ClassPeriodTx) error
}
