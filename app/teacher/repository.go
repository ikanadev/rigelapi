package teacher

import "github.com/vmkevv/rigelapi/app/models"

type TeacherRepository interface {
	GetProfile(teacherID string) (models.TeacherWithSubs, error)
}
