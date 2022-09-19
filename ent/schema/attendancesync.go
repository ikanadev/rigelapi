package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AttendanceSync holds the schema definition for the AttendanceSync entity.
type AttendanceSync struct {
	ent.Schema
}

// Fields of the AttendanceSync.
func (AttendanceSync) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("last_sync_id"),
	}
}

// Edges of the AttendanceSync.
func (AttendanceSync) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("attendanceDay", AttendanceDay.Type).Ref("attendanceSyncs").Unique(),
	}
}
