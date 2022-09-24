package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ClassPeriod holds the schema definition for the ClassPeriod entity.
type ClassPeriod struct {
	ent.Schema
}

// Fields of the ClassPeriod.
func (ClassPeriod) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.Time("start"),
		field.Time("end"),
		field.Bool("finished"),
	}
}

// Edges of the ClassPeriod.
func (ClassPeriod) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("attendanceDays", AttendanceDay.Type),
		edge.To("attendanceDaySyncs", AttendanceDaySyncs.Type),
		edge.To("activities", Activity.Type),
		edge.From("class", Class.Type).Ref("classPeriods").Unique(),
		edge.From("period", Period.Type).Ref("classPeriods").Unique(),
	}
}
