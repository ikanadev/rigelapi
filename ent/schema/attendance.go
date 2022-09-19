package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Attendance holds the schema definition for the Attendance entity.
type Attendance struct {
	ent.Schema
}

// Fields of the Attendance.
func (Attendance) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.Enum("value").Values("Asistencia", "Falta", "Atraso", "Licencia"),
	}
}

// Edges of the Attendance.
func (Attendance) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("attendanceDay", AttendanceDay.Type).Ref("attendances").Unique(),
		edge.From("student", Student.Type).Ref("attendances").Unique(),
	}
}
