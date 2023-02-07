package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/ent/attendance"
	"github.com/vmkevv/rigelapi/ent/class"
)

func ClassDetailsHandler(db *ent.Client) func(*fiber.Ctx) error {
	type AttendanceTotals map[attendance.Value]int
	type ClassData struct {
		ID           string  `json:"id"`
		Parallel     string  `json:"parallel"`
		Municipio    string  `json:"municipio"`
		Provincia    string  `json:"provincia"`
		Departamento string  `json:"departamento"`
		School       School  `json:"school"`
		Teacher      Teacher `json:"teacher"`
		Subject      Subject `json:"subject"`
		Grade        Grade   `json:"grade"`
		Year         Year    `json:"year"`
	}
	type StudentData struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		LastName  string `json:"last_name"`
		CI        string `json:"ci"`
		YearScore int    `json:"year_score"`
		// the map key is activity_id
		ScoresMap map[string]Score `json:"scores_map"`
		// the map key is class_period_id
		PeriodScores map[string]struct {
			Points int `json:"points"`
			// the map key is area_id
			AreaScores map[string]int `json:"area_scores"`
		} `json:"scores"`
		// the map key is attendance_day_id
		AttendancesMap       map[string]Attendance `json:"attendances_map"`
		YearTotalAttendances AttendanceTotals      `json:"year_total_attendances"`
		// the map key is period_id
		PeriodTotalAttendances map[string]AttendanceTotals `json:"period_total_attendances"`
	}
	type AreaWithActivities struct {
		Area
		Activities []Activity `json:"activities"`
	}
	type ClassPeriodData struct {
		ID             string               `json:"id"`
		Start          int64                `json:"start"`
		End            int64                `json:"end"`
		Finished       bool                 `json:"finished"`
		Period         Period               `json:"period"`
		Areas          []AreaWithActivities `json:"areas"`
		AttendanceDays []AttendanceDay      `json:"attendance_days"`
	}
	type Resp struct {
		ClassData    ClassData         `json:"class_data"`
		Areas        []Area            `json:"areas"` // TODO: check if this is needed
		Students     []StudentData     `json:"students"`
		ClassPeriods []ClassPeriodData `json:"class_periods"`
	}
	type RespX struct {
		Resp         Resp               `json:"resp"`
		Class        *ent.Class         `json:"class"`
		ClassPeriods []*ent.ClassPeriod `json:"class_periods"`
		Students     []*ent.Student     `json:"students"`
	}
	return func(c *fiber.Ctx) error {
		classID := c.Params("classid")
		class, err := db.Class.
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
			Only(c.Context())
		if err != nil {
			return c.SendStatus(fiber.StatusNotFound)
		}

		classPeriods, err := class.QueryClassPeriods().
			WithPeriod().
			WithAttendanceDays().
			WithActivities(func(aq *ent.ActivityQuery) {
				aq.WithArea()
			}).
			All(c.Context())
		if err != nil {
			return err
		}

		students, err := class.QueryStudents().
			WithAttendances(func(aq *ent.AttendanceQuery) {
				aq.WithAttendanceDay(func(adq *ent.AttendanceDayQuery) {
					adq.WithClassPeriod()
				})
			}).
			WithScores(func(sq *ent.ScoreQuery) {
				sq.WithActivity(func(aq *ent.ActivityQuery) {
					aq.WithArea()
				})
			}).
			All(c.Context())
		if err != nil {
			return err
		}

		resp := Resp{
			ClassData: ClassData{
				ID:           class.ID,
				Parallel:     class.Parallel,
				Municipio:    class.Edges.School.Edges.Municipio.Name,
				Provincia:    class.Edges.School.Edges.Municipio.Edges.Provincia.Name,
				Departamento: class.Edges.School.Edges.Municipio.Edges.Provincia.Edges.Departamento.Name,
				School: School{
					ID:   class.Edges.School.ID,
					Name: class.Edges.School.Name,
					Lat:  class.Edges.School.Lat,
					Lon:  class.Edges.School.Lon,
				},
				Teacher: Teacher{
					ID:       class.Edges.Teacher.ID,
					Name:     class.Edges.Teacher.Name,
					LastName: class.Edges.Teacher.LastName,
					Email:    class.Edges.Teacher.Email,
				},
				Subject: Subject{
					ID:   class.Edges.Subject.ID,
					Name: class.Edges.Subject.Name,
				},
				Grade: Grade{
					ID:   class.Edges.Grade.ID,
					Name: class.Edges.Grade.Name,
				},
				Year: Year{
					ID:    class.Edges.Year.ID,
					Value: class.Edges.Year.Value,
				},
			},
		}

		// class periods data
		classPeriodsData := make([]ClassPeriodData, len(classPeriods))
		for i, cp := range classPeriods {
			areas := make([]AreaWithActivities, len(class.Edges.Year.Edges.Areas))
			for j, area := range class.Edges.Year.Edges.Areas {
				acts := make([]Activity, 0)
				for _, act := range cp.Edges.Activities {
					if act.Edges.Area.ID == area.ID {
						acts = append(acts, Activity{
							ID:   act.ID,
							Name: act.Name,
							Date: act.Date.UnixMilli(),
						})
					}
				}
				areas[j] = AreaWithActivities{
					Area: Area{
						ID:     area.ID,
						Name:   area.Name,
						Points: area.Points,
					},
					Activities: acts,
				}
			}

			classPeriodsData[i] = ClassPeriodData{
				ID:       cp.ID,
				Start:    cp.Start.UnixMilli(),
				End:      cp.End.UnixMilli(),
				Finished: cp.Finished,
				Period: Period{
					ID:   cp.Edges.Period.ID,
					Name: cp.Edges.Period.Name,
				},
				Areas: areas,
			}
		}
		resp.ClassPeriods = classPeriodsData

		// Students data
		studentsData := make([]StudentData, len(students))
		for i, student := range students {
			// single student info
			studentData := StudentData{
				ID:       student.ID,
				Name:     student.Name,
				LastName: student.LastName,
				CI:       student.Ci,
			}

			// student attendances data
			attendancesMap := make(map[string]Attendance, len(student.Edges.Attendances))
			periodTotalAttendances := make(map[string]AttendanceTotals, len(classPeriods))
			for _, cp := range classPeriods {
				periodTotalAttendances[cp.ID] = AttendanceTotals{
					attendance.ValuePresente: 0,
					attendance.ValueFalta:    0,
					attendance.ValueAtraso:   0,
					attendance.ValueLicencia: 0,
				}
			}
			yearTotalAttendances := AttendanceTotals{
				attendance.ValuePresente: 0,
				attendance.ValueFalta:    0,
				attendance.ValueAtraso:   0,
				attendance.ValueLicencia: 0,
			}
			for _, att := range student.Edges.Attendances {
				periodTotalAttendances[att.Edges.AttendanceDay.Edges.ClassPeriod.ID][att.Value] += 1
				yearTotalAttendances[att.Value] += 1
				attendancesMap[att.Edges.AttendanceDay.ID] = Attendance{
					ID:    att.ID,
					Value: att.Value,
				}
			}
			studentData.AttendancesMap = attendancesMap
			studentData.PeriodTotalAttendances = periodTotalAttendances
			studentData.YearTotalAttendances = yearTotalAttendances

			studentsData[i] = studentData
		}
		resp.Students = studentsData

		return c.JSON(RespX{
			Resp:         resp,
			Class:        class,
			ClassPeriods: classPeriods,
			Students:     students,
		})
	}
}
