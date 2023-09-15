package class

import "github.com/vmkevv/rigelapi/app/models"

type NewClassData struct {
	YearID    string
	TeacherID string
	GradeID   string
	SubjectID string
	SchoolID  string
	Parallel  string
}

type ClassRepository interface {
	GetClassData(classID string) (models.ClassData, error)
	GetClassPeriodsData(classID string) ([]models.ClassPeriodData, error)
	GetStudentsData(
		classID string,
		classPeriodsData []models.ClassPeriodData,
	) ([]models.StudentData, error)
	GetTeacherClasses(teacherID string, yearID string) ([]models.ClassData, error)
	SaveClass(data NewClassData) error
}
