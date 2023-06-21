package class

import "github.com/vmkevv/rigelapi/app/models"

type ClassRepository interface {
	GetClassData(classID string) (models.ClassData, error)
	GetClassPeriodsData(classID string) ([]models.ClassPeriodData, error)
	GetStudentsData(
		classID string,
		classPeriodsData []models.ClassPeriodData,
	) ([]models.StudentData, error)
}
