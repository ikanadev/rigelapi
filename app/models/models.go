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
