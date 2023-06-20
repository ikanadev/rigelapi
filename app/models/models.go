package models

import "github.com/vmkevv/rigelapi/ent/attendance"

type AppError struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	Cause      string `json:"cause"`
	ErrorMsg   string `json:"error_msg"`
	ErrorStack string `json:"error_stack"`
}

type Teacher struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	IsAdmin  bool   `json:"is_admin"`
}
type SubWithYear struct {
	Subscription
	Year Year `json:"year"`
}
type TeacherWithSubs struct {
	Teacher
	Subscriptions []SubWithYear `json:"subscriptions"`
}

type School struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Lat  string `json:"lat"`
	Lon  string `json:"lon"`
}

type Dpto struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Provincia struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Municipio struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Subject struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Grade struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Period struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Area struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Points int    `json:"points"`
}

type Year struct {
	ID    string `json:"id"`
	Value int    `json:"value"`
}

type YearData struct {
	Year
	Periods []Period `json:"periods"`
	Areas   []Area   `json:"areas"`
}

type Class struct {
	ID       string `json:"id"`
	Parallel string `json:"parallel"`
}

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

type Subscription struct {
	ID     string `json:"id"`
	Method string `json:"method"`
	Qtty   int    `json:"qtty"`
	Date   int64  `json:"date"`
}

type Score struct {
	ID     string `json:"id"`
	Points int    `json:"points"`
}

type Attendance struct {
	ID    string           `json:"id"`
	Value attendance.Value `json:"value"`
}

type Activity struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Date int64  `json:"date"`
}

type AttendanceDay struct {
	ID  string `json:"id"`
	Day int64  `json:"day"`
}

type AttendanceTotals map[attendance.Value]int

type PeriodScores struct {
	Score int `json:"score"`
	// the map key is area_id
	AreaScores map[string]int `json:"area_scores"`
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
	PeriodScores map[string]PeriodScores `json:"scores"`
	// the map key is attendance_day_id
	AttendancesMap       map[string]Attendance `json:"attendances_map"`
	YearTotalAttendances AttendanceTotals      `json:"year_total_attendances"`
	// the map key is period_id
	ClassPeriodTotalAttendances map[string]AttendanceTotals `json:"period_total_attendances"`
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
