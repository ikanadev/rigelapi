package extra

import (
	"context"

	"github.com/vmkevv/rigelapi/app/models"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/ent/apperror"
	"github.com/vmkevv/rigelapi/ent/class"
	"github.com/vmkevv/rigelapi/ent/subject"
)

type ExtraEntRepo struct {
	ent *ent.Client
	ctx context.Context
}

func NewExtraEntRepo(ent *ent.Client, ctx context.Context) ExtraEntRepo {
	return ExtraEntRepo{ent, ctx}
}

func (eer ExtraEntRepo) GetYearsData() ([]models.YearData, error) {
	entYears, err := eer.ent.Year.Query().WithAreas().WithPeriods().All(eer.ctx)
	if err != nil {
		return nil, err
	}
	years := make([]models.YearData, len(entYears))
	for i, year := range entYears {
		periods := make([]models.Period, len(year.Edges.Periods))
		for j, period := range year.Edges.Periods {
			periods[j] = models.Period{ID: period.ID, Name: period.Name}
		}
		areas := make([]models.Area, len(year.Edges.Areas))
		for j, area := range year.Edges.Areas {
			areas[j] = models.Area{
				ID:     area.ID,
				Name:   area.Name,
				Points: area.Points,
			}
		}

		years[i] = models.YearData{
			Year: models.Year{
				ID:    year.ID,
				Value: year.Value,
			},
			Periods: periods,
			Areas:   areas,
		}
	}

	return years, nil
}

func (eer ExtraEntRepo) GetGrades() ([]models.Grade, error) {
	entGrades, err := eer.ent.Grade.Query().All(eer.ctx)
	if err != nil {
		return nil, err
	}
	grades := make([]models.Grade, len(entGrades))
	for i, grade := range entGrades {
		grades[i] = models.Grade{
			ID:   grade.ID,
			Name: grade.Name,
		}
	}
	return grades, nil
}

func (eer ExtraEntRepo) GetSubjects() ([]models.Subject, error) {
	entSubjects, err := eer.ent.Subject.Query().
		Order(ent.Asc(subject.FieldName)).
		All(eer.ctx)
	if err != nil {
		return nil, err
	}
	subjects := make([]models.Subject, len(entSubjects))
	for i, subject := range entSubjects {
		subjects[i] = models.Subject{ID: subject.ID, Name: subject.Name}
	}
	return subjects, nil
}

func (eer ExtraEntRepo) SaveAppErrors(appErrors []models.AppError) error {
	toAdd := make([]*ent.AppErrorCreate, len(appErrors))
	for i, appError := range appErrors {
		toAdd[i] = eer.ent.AppError.Create().
			SetID(appError.ID).
			SetUserID(appError.UserID).
			SetCause(appError.Cause).
			SetErrorMsg(appError.ErrorMsg).
			SetErrorStack(appError.ErrorStack)
	}
	err := eer.ent.AppError.CreateBulk(toAdd...).
		OnConflictColumns(apperror.FieldID).
		Ignore().
		Exec(eer.ctx)
	return err
}

func (eer ExtraEntRepo) GetTeachersCount() (int, error) {
	teachers, err := eer.ent.Teacher.Query().Count(eer.ctx)
	return teachers, err
}

func (eer ExtraEntRepo) GetClassesCount() (int, error) {
	classes, err := eer.ent.Class.Query().Count(eer.ctx)
	return classes, err
}

func (eer ExtraEntRepo) GetSchoolsCount() (int, error) {
	schools, err := eer.ent.Class.Query().
		Unique(true).
		Select(class.SchoolColumn).
		Strings(eer.ctx)
	return len(schools), err
}

func (eer ExtraEntRepo) GetActivitiesCount() (int, error) {
	acts, err := eer.ent.Activity.Query().Count(eer.ctx)
	return acts, err
}
