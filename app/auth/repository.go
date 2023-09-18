package auth

import "github.com/vmkevv/rigelapi/app/models"

type AuthRepository interface {
	Register(name, lastName, email, password string) error
	GetTeacher(email, password string) (models.TeacherWithSubs, string, error)
	GetTeacherProfile(teacherID string) (models.TeacherWithSubs, error)
	GetTeachers() ([]models.Teacher, error)
}
