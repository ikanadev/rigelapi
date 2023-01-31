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
		// the map key is period_id
		PeriodScores map[string]struct {
			Points int `json:"points"`
			// the map key is area_id
			AreaScores map[string]int `json:"area_scores"`
		} `json:"scores"`
		// the map key is attendance_day_id
		AttendancesMap  map[string]Attendance `json:"attendances_map"`
		YearAttendances AttendanceTotals      `json:"year_attendances"`
		// the map key is period_id
		PeriodAttendances map[string]AttendanceTotals `json:"period_attendances"`
	}
	type AreaWithActivities struct {
		Area
		Activities []Activity `json:"activities"`
	}
	type AttendanceData struct {
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
		Areas        []Area            `json:"areas"`
		Students     []StudentData     `json:"students"`
		ClassPeriods []ClassPeriodData `json:"class_periods"`
	}
	type RespX struct {
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
			WithActivities().
			All(c.Context())
		if err != nil {
			return err
		}

		students, err := class.QueryStudents().
			WithAttendances(func(aq *ent.AttendanceQuery) {
				aq.WithAttendanceDay()
			}).
			WithScores(func(sq *ent.ScoreQuery) {
				sq.WithActivity()
			}).
			All(c.Context())

		if err != nil {
			return err
		}

		return c.JSON(RespX{
			Class:        class,
			ClassPeriods: classPeriods,
			Students:     students,
		})
	}
}
