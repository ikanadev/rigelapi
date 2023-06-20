package extra

import (
	"context"
	"math"

	"github.com/vmkevv/rigelapi/app/models"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/ent/apperror"
	"github.com/vmkevv/rigelapi/ent/attendance"
	"github.com/vmkevv/rigelapi/ent/attendanceday"
	"github.com/vmkevv/rigelapi/ent/class"
	"github.com/vmkevv/rigelapi/ent/student"
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

func (eer ExtraEntRepo) GetClassData(classID string) (models.ClassData, error) {
	var classData models.ClassData
	entClass, err := eer.ent.Class.
		Query().
		Where(class.IDEQ(classID)).
		WithGrade().
		WithSubject().
		WithTeacher().
		WithYear(func(yq *ent.YearQuery) {
			yq.WithAreas()
		}).
		WithSchool(func(sq *ent.SchoolQuery) {
			sq.WithMunicipio(func(mq *ent.MunicipioQuery) {
				mq.WithProvincia(func(pq *ent.ProvinciaQuery) {
					pq.WithDepartamento()
				})
			})
		}).
		Only(eer.ctx)
	if err != nil {
		return classData, err
	}
	classData.ID = entClass.ID
	classData.Parallel = entClass.Parallel
	classData.Municipio = entClass.Edges.School.Edges.Municipio.Name
	classData.Provincia = entClass.Edges.
		School.Edges.
		Municipio.Edges.
		Provincia.Name
	classData.Departamento = entClass.Edges.
		School.Edges.
		Municipio.Edges.
		Provincia.Edges.
		Departamento.Name
	classData.School = models.School{
		ID:   entClass.Edges.School.ID,
		Name: entClass.Edges.School.Name,
		Lat:  entClass.Edges.School.Lat,
		Lon:  entClass.Edges.School.Lon,
	}
	classData.Teacher = models.Teacher{
		ID:       entClass.Edges.Teacher.ID,
		Name:     entClass.Edges.Teacher.Name,
		LastName: entClass.Edges.Teacher.LastName,
		Email:    entClass.Edges.Teacher.Email,
	}
	classData.Subject = models.Subject{
		ID:   entClass.Edges.Subject.ID,
		Name: entClass.Edges.Subject.Name,
	}
	classData.Grade = models.Grade{
		ID:   entClass.Edges.Grade.ID,
		Name: entClass.Edges.Grade.Name,
	}
	classData.Year = models.Year{
		ID:    entClass.Edges.Year.ID,
		Value: entClass.Edges.Year.Value,
	}
	return classData, nil
}

func (eer ExtraEntRepo) GetClassPeriodsData(classID string) (
	[]models.ClassPeriodData,
	error,
) {
	entClass, err := eer.ent.Class.Query().
		Where(class.IDEQ(classID)).
		WithYear(func(yq *ent.YearQuery) {
			yq.WithAreas()
		}).
		Only(eer.ctx)
	if err != nil {
		return []models.ClassPeriodData{}, nil
	}

	entClassPeriods, err := entClass.QueryClassPeriods().
		WithPeriod().
		WithAttendanceDays(func(adq *ent.AttendanceDayQuery) {
			adq.Order(ent.Asc(attendanceday.FieldDay))
		}).
		WithActivities(func(aq *ent.ActivityQuery) {
			aq.WithArea()
		}).
		All(eer.ctx)
	if err != nil {
		return []models.ClassPeriodData{}, nil
	}

	classPeriodsData := make([]models.ClassPeriodData, len(entClassPeriods))
	for i, cp := range entClassPeriods {
		areas := make(
			[]models.AreaWithActivities,
			len(entClass.Edges.Year.Edges.Areas),
		)
		for j, area := range entClass.Edges.Year.Edges.Areas {
			acts := make([]models.Activity, 0)
			for _, act := range cp.Edges.Activities {
				if act.Edges.Area.ID == area.ID {
					acts = append(acts, models.Activity{
						ID:   act.ID,
						Name: act.Name,
						Date: act.Date.UnixMilli(),
					})
				}
			}
			areas[j] = models.AreaWithActivities{
				Area: models.Area{
					ID:     area.ID,
					Name:   area.Name,
					Points: area.Points,
				},
				Activities: acts,
			}
		}

		attDays := make([]models.AttendanceDay, len(cp.Edges.AttendanceDays))
		for j, attDay := range cp.Edges.AttendanceDays {
			attDays[j] = models.AttendanceDay{
				ID:  attDay.ID,
				Day: attDay.Day.UnixMilli(),
			}
		}

		classPeriodsData[i] = models.ClassPeriodData{
			ID:       cp.ID,
			Start:    cp.Start.UnixMilli(),
			End:      cp.End.UnixMilli(),
			Finished: cp.Finished,
			Period: models.Period{
				ID:   cp.Edges.Period.ID,
				Name: cp.Edges.Period.Name,
			},
			Areas:          areas,
			AttendanceDays: attDays,
		}
	}

	return classPeriodsData, nil
}

func (eer ExtraEntRepo) GetStudentsData(
	classID string,
	classPeriodsData []models.ClassPeriodData,
) (
	[]models.StudentData,
	error,
) {
	entStudents, err := eer.ent.Student.Query().
		Where(student.HasClassWith(class.IDEQ(classID))).
		Order(ent.Asc(student.FieldLastName)).
		WithAttendances(func(aq *ent.AttendanceQuery) {
			aq.WithAttendanceDay(func(adq *ent.AttendanceDayQuery) {
				adq.WithClassPeriod()
			})
		}).
		WithScores(func(sq *ent.ScoreQuery) {
			sq.WithActivity()
		}).
		All(eer.ctx)
	if err != nil {
		return []models.StudentData{}, err
	}

	studentsData := make([]models.StudentData, len(entStudents))
	for i, student := range entStudents {
		studentData := models.StudentData{
			ID:       student.ID,
			Name:     student.Name,
			LastName: student.LastName,
			CI:       student.Ci,
		}
		// student attendances data
		attendancesMap := make(
			map[string]models.Attendance,
			len(student.Edges.Attendances),
		)
		classPeriodTotalAttendances := make(
			map[string]models.AttendanceTotals,
			len(classPeriodsData),
		)
		for _, cp := range classPeriodsData {
			classPeriodTotalAttendances[cp.ID] = models.AttendanceTotals{
				attendance.ValuePresente: 0,
				attendance.ValueFalta:    0,
				attendance.ValueAtraso:   0,
				attendance.ValueLicencia: 0,
			}
		}
		yearTotalAttendances := models.AttendanceTotals{
			attendance.ValuePresente: 0,
			attendance.ValueFalta:    0,
			attendance.ValueAtraso:   0,
			attendance.ValueLicencia: 0,
		}
		for _, att := range student.Edges.Attendances {
			classPeriodID := att.Edges.AttendanceDay.Edges.ClassPeriod.ID
			classPeriodTotalAttendances[classPeriodID][att.Value] += 1
			yearTotalAttendances[att.Value] += 1
			attendancesMap[att.Edges.AttendanceDay.ID] = models.Attendance{
				ID:    att.ID,
				Value: att.Value,
			}
		}
		studentData.AttendancesMap = attendancesMap
		studentData.ClassPeriodTotalAttendances = classPeriodTotalAttendances
		studentData.YearTotalAttendances = yearTotalAttendances

		// student scores data
		scoresMap := make(map[string]models.Score, len(student.Edges.Scores))
		for _, score := range student.Edges.Scores {
			scoresMap[score.Edges.Activity.ID] = models.Score{
				ID:     score.ID,
				Points: score.Points,
			}
		}
		studentData.ScoresMap = scoresMap

		// student period scores
		periodScores := make(map[string]models.PeriodScores, len(classPeriodsData))
		yearScoreSum := 0
		for _, cp := range classPeriodsData {
			periodScore := 0
			areaScores := make(map[string]int, len(cp.Areas))
			for _, area := range cp.Areas {
				sum := 0
				if len(area.Activities) == 0 {
					areaScores[area.ID] = 0
					continue
				}
				for _, act := range area.Activities {
					score, ok := studentData.ScoresMap[act.ID]
					if ok {
						sum += score.Points
					}
				}
				areaScore := int(
					math.Round(
						float64(sum*area.Points) / float64(len(area.Activities)*100),
					),
				)
				areaScores[area.ID] = areaScore
				periodScore += areaScore
			}
			yearScoreSum += periodScore
			periodScores[cp.ID] = models.PeriodScores{
				Score:      periodScore,
				AreaScores: areaScores,
			}
		}
		studentData.PeriodScores = periodScores
		studentData.YearScore = int(
			math.Round(
				float64(yearScoreSum) / float64(len(classPeriodsData)),
			),
		)
		studentsData[i] = studentData
	}

	return studentsData, nil
}
