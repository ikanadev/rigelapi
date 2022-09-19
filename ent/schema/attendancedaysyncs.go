package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AttendanceDaySyncs holds the schema definition for the AttendanceDaySyncs entity.
type AttendanceDaySyncs struct {
	ent.Schema
}

// Fields of the AttendanceDaySyncs.
func (AttendanceDaySyncs) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("last_sync_id"),
	}
}

// Edges of the AttendanceDaySyncs.
func (AttendanceDaySyncs) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("classPeriod", ClassPeriod.Type).Ref("attendanceDaySyncs").Unique(),
	}
}
