package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AttendanceDay holds the schema definition for the AttendanceDay entity.
type AttendanceDay struct {
	ent.Schema
}

// Fields of the AttendanceDay.
func (AttendanceDay) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.Time("day"),
	}
}

// Edges of the AttendanceDay.
func (AttendanceDay) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("attendances", Attendance.Type),
		edge.From("classPeriod", ClassPeriod.Type).Ref("attendanceDays").Unique(),
	}
}
