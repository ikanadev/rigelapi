package sync

import "github.com/vmkevv/rigelapi/app/models"

type SyncRepository interface {
	GetStudents(teacherID, yearID string) ([]models.Student, error)
}
