package class

import (
	"context"
	"math"

	"github.com/vmkevv/rigelapi/app/models"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/ent/attendance"
	"github.com/vmkevv/rigelapi/ent/attendanceday"
	"github.com/vmkevv/rigelapi/ent/class"
	"github.com/vmkevv/rigelapi/ent/classperiod"
	"github.com/vmkevv/rigelapi/ent/student"
	"github.com/vmkevv/rigelapi/ent/teacher"
	"github.com/vmkevv/rigelapi/ent/year"
)

type ClassEntRepo struct {
	ent *ent.Client
	ctx context.Context
}

func NewClassEntRepo(ent *ent.Client, ctx context.Context) ClassEntRepo {
	return ClassEntRepo{ent, ctx}
}

func (cer ClassEntRepo) GetTeacherClasses(
	teacherID string,
	yearID string,
) ([]models.ClassData, error) {
	entClasses, err := cer.ent.Class.Query().
		Where(
			class.HasTeacherWith(teacher.ID(teacherID)),
			class.HasYearWith(year.ID(yearID)),
		).
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
		All(cer.ctx)
	if err != nil {
		return nil, err
	}
	classesData := make([]models.ClassData, len(entClasses))
	for i, class := range entClasses {
		classesData[i] = entClassToClassData(class)
	}
	return classesData, nil
}

func (cer ClassEntRepo) GetClassData(classID string) (models.ClassData, error) {
	var classData models.ClassData
	entClass, err := cer.ent.Class.
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
		Only(cer.ctx)
	if err != nil {
		return classData, err
	}
	classData = entClassToClassData(entClass)
	return classData, nil
}

func (cer ClassEntRepo) GetClassPeriodsData(classID string) (
	[]models.ClassPeriodData,
	error,
) {
	entClass, err := cer.ent.Class.Query().
		Where(class.IDEQ(classID)).
		WithYear(func(yq *ent.YearQuery) {
			yq.WithAreas()
		}).
		Only(cer.ctx)
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
		Order(ent.Asc(classperiod.FieldStart)).
		All(cer.ctx)
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

func (cer ClassEntRepo) GetStudentsData(
	classID string,
	classPeriodsData []models.ClassPeriodData,
) (
	[]models.StudentData,
	error,
) {
	entStudents, err := cer.ent.Student.Query().
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
		All(cer.ctx)
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

func entClassToClassData(entClass *ent.Class) models.ClassData {
	return models.ClassData{
		ID:        entClass.ID,
		Parallel:  entClass.Parallel,
		Municipio: entClass.Edges.School.Edges.Municipio.Name,
		Provincia: entClass.Edges.
			School.Edges.
			Municipio.Edges.
			Provincia.Name,
		Departamento: entClass.Edges.
			School.Edges.
			Municipio.Edges.
			Provincia.Edges.
			Departamento.Name,
		School: models.School{
			ID:   entClass.Edges.School.ID,
			Name: entClass.Edges.School.Name,
			Lat:  entClass.Edges.School.Lat,
			Lon:  entClass.Edges.School.Lon,
		},
		Teacher: models.Teacher{
			ID:       entClass.Edges.Teacher.ID,
			Name:     entClass.Edges.Teacher.Name,
			LastName: entClass.Edges.Teacher.LastName,
			Email:    entClass.Edges.Teacher.Email,
		},
		Subject: models.Subject{
			ID:   entClass.Edges.Subject.ID,
			Name: entClass.Edges.Subject.Name,
		},
		Grade: models.Grade{
			ID:   entClass.Edges.Grade.ID,
			Name: entClass.Edges.Grade.Name,
		},
		Year: models.Year{
			ID:    entClass.Edges.Year.ID,
			Value: entClass.Edges.Year.Value,
		},
	}
}
