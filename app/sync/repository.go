package sync

import (
	"github.com/vmkevv/rigelapi/app/common"
	"github.com/vmkevv/rigelapi/app/models"
)

type SyncRepository interface {
	GetStudents(teacherID, yearID string) ([]models.AppStudent, error)
	SyncStudents(studentTxs []common.AppStudentTx) error
	GetClassPeriods(teacherID, yearID string) ([]models.AppClassPeriod, error)
	SyncClassPeriods(classPeriodTxs []common.AppClassPeriodTx) error
	GetAttendanceDays(teacherID, yearID string) ([]models.AppAttendanceDay, error)
	SyncAttendanceDays(attendanceDayTxs []common.AppAttendanceDayTx) error
	GetAttendances(teacherID, yearID string) ([]models.AppAttendance, error)
	SyncAttendances(attendanceTxs []common.AppAttendanceTx) error
	GetActivities(teacherID, yearID string) ([]models.AppActivity, error)
	SyncActivities(activityTx []common.AppActivityTx) error
}
