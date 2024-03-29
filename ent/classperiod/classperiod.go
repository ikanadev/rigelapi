// Code generated by ent, DO NOT EDIT.

package classperiod

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the classperiod type in the database.
	Label = "class_period"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStart holds the string denoting the start field in the database.
	FieldStart = "start"
	// FieldEnd holds the string denoting the end field in the database.
	FieldEnd = "end"
	// FieldFinished holds the string denoting the finished field in the database.
	FieldFinished = "finished"
	// EdgeAttendanceDays holds the string denoting the attendancedays edge name in mutations.
	EdgeAttendanceDays = "attendanceDays"
	// EdgeActivities holds the string denoting the activities edge name in mutations.
	EdgeActivities = "activities"
	// EdgeClass holds the string denoting the class edge name in mutations.
	EdgeClass = "class"
	// EdgePeriod holds the string denoting the period edge name in mutations.
	EdgePeriod = "period"
	// Table holds the table name of the classperiod in the database.
	Table = "class_periods"
	// AttendanceDaysTable is the table that holds the attendanceDays relation/edge.
	AttendanceDaysTable = "attendance_days"
	// AttendanceDaysInverseTable is the table name for the AttendanceDay entity.
	// It exists in this package in order to avoid circular dependency with the "attendanceday" package.
	AttendanceDaysInverseTable = "attendance_days"
	// AttendanceDaysColumn is the table column denoting the attendanceDays relation/edge.
	AttendanceDaysColumn = "class_period_attendance_days"
	// ActivitiesTable is the table that holds the activities relation/edge.
	ActivitiesTable = "activities"
	// ActivitiesInverseTable is the table name for the Activity entity.
	// It exists in this package in order to avoid circular dependency with the "activity" package.
	ActivitiesInverseTable = "activities"
	// ActivitiesColumn is the table column denoting the activities relation/edge.
	ActivitiesColumn = "class_period_activities"
	// ClassTable is the table that holds the class relation/edge.
	ClassTable = "class_periods"
	// ClassInverseTable is the table name for the Class entity.
	// It exists in this package in order to avoid circular dependency with the "class" package.
	ClassInverseTable = "classes"
	// ClassColumn is the table column denoting the class relation/edge.
	ClassColumn = "class_class_periods"
	// PeriodTable is the table that holds the period relation/edge.
	PeriodTable = "class_periods"
	// PeriodInverseTable is the table name for the Period entity.
	// It exists in this package in order to avoid circular dependency with the "period" package.
	PeriodInverseTable = "periods"
	// PeriodColumn is the table column denoting the period relation/edge.
	PeriodColumn = "period_class_periods"
)

// Columns holds all SQL columns for classperiod fields.
var Columns = []string{
	FieldID,
	FieldStart,
	FieldEnd,
	FieldFinished,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "class_periods"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"class_class_periods",
	"period_class_periods",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the ClassPeriod queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByStart orders the results by the start field.
func ByStart(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStart, opts...).ToFunc()
}

// ByEnd orders the results by the end field.
func ByEnd(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEnd, opts...).ToFunc()
}

// ByFinished orders the results by the finished field.
func ByFinished(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFinished, opts...).ToFunc()
}

// ByAttendanceDaysCount orders the results by attendanceDays count.
func ByAttendanceDaysCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAttendanceDaysStep(), opts...)
	}
}

// ByAttendanceDays orders the results by attendanceDays terms.
func ByAttendanceDays(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAttendanceDaysStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByActivitiesCount orders the results by activities count.
func ByActivitiesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newActivitiesStep(), opts...)
	}
}

// ByActivities orders the results by activities terms.
func ByActivities(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newActivitiesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByClassField orders the results by class field.
func ByClassField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newClassStep(), sql.OrderByField(field, opts...))
	}
}

// ByPeriodField orders the results by period field.
func ByPeriodField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPeriodStep(), sql.OrderByField(field, opts...))
	}
}
func newAttendanceDaysStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AttendanceDaysInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AttendanceDaysTable, AttendanceDaysColumn),
	)
}
func newActivitiesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ActivitiesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ActivitiesTable, ActivitiesColumn),
	)
}
func newClassStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ClassInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ClassTable, ClassColumn),
	)
}
func newPeriodStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PeriodInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, PeriodTable, PeriodColumn),
	)
}
