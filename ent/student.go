// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/vmkevv/rigelapi/ent/class"
	"github.com/vmkevv/rigelapi/ent/student"
)

// Student is the model entity for the Student schema.
type Student struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// LastName holds the value of the "last_name" field.
	LastName string `json:"last_name,omitempty"`
	// Ci holds the value of the "ci" field.
	Ci string `json:"ci,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StudentQuery when eager-loading is set.
	Edges          StudentEdges `json:"edges"`
	class_students *string
	selectValues   sql.SelectValues
}

// StudentEdges holds the relations/edges for other nodes in the graph.
type StudentEdges struct {
	// Attendances holds the value of the attendances edge.
	Attendances []*Attendance `json:"attendances,omitempty"`
	// Scores holds the value of the scores edge.
	Scores []*Score `json:"scores,omitempty"`
	// Class holds the value of the class edge.
	Class *Class `json:"class,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// AttendancesOrErr returns the Attendances value or an error if the edge
// was not loaded in eager-loading.
func (e StudentEdges) AttendancesOrErr() ([]*Attendance, error) {
	if e.loadedTypes[0] {
		return e.Attendances, nil
	}
	return nil, &NotLoadedError{edge: "attendances"}
}

// ScoresOrErr returns the Scores value or an error if the edge
// was not loaded in eager-loading.
func (e StudentEdges) ScoresOrErr() ([]*Score, error) {
	if e.loadedTypes[1] {
		return e.Scores, nil
	}
	return nil, &NotLoadedError{edge: "scores"}
}

// ClassOrErr returns the Class value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StudentEdges) ClassOrErr() (*Class, error) {
	if e.loadedTypes[2] {
		if e.Class == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: class.Label}
		}
		return e.Class, nil
	}
	return nil, &NotLoadedError{edge: "class"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Student) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case student.FieldID, student.FieldName, student.FieldLastName, student.FieldCi:
			values[i] = new(sql.NullString)
		case student.ForeignKeys[0]: // class_students
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Student fields.
func (s *Student) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case student.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				s.ID = value.String
			}
		case student.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case student.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				s.LastName = value.String
			}
		case student.FieldCi:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ci", values[i])
			} else if value.Valid {
				s.Ci = value.String
			}
		case student.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field class_students", values[i])
			} else if value.Valid {
				s.class_students = new(string)
				*s.class_students = value.String
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Student.
// This includes values selected through modifiers, order, etc.
func (s *Student) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryAttendances queries the "attendances" edge of the Student entity.
func (s *Student) QueryAttendances() *AttendanceQuery {
	return NewStudentClient(s.config).QueryAttendances(s)
}

// QueryScores queries the "scores" edge of the Student entity.
func (s *Student) QueryScores() *ScoreQuery {
	return NewStudentClient(s.config).QueryScores(s)
}

// QueryClass queries the "class" edge of the Student entity.
func (s *Student) QueryClass() *ClassQuery {
	return NewStudentClient(s.config).QueryClass(s)
}

// Update returns a builder for updating this Student.
// Note that you need to call Student.Unwrap() before calling this method if this Student
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Student) Update() *StudentUpdateOne {
	return NewStudentClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Student entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Student) Unwrap() *Student {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Student is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Student) String() string {
	var builder strings.Builder
	builder.WriteString("Student(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteString(", ")
	builder.WriteString("last_name=")
	builder.WriteString(s.LastName)
	builder.WriteString(", ")
	builder.WriteString("ci=")
	builder.WriteString(s.Ci)
	builder.WriteByte(')')
	return builder.String()
}

// Students is a parsable slice of Student.
type Students []*Student
