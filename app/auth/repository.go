package auth

import "github.com/vmkevv/rigelapi/app/models"

type AuthRepository interface {
	Register(name, lastName, email, password string) error
	GetTeacher(email, password string) (models.TeacherWithSubs, string, error)
	GetProfile(teacherID string) (models.TeacherWithSubs, error)
}
