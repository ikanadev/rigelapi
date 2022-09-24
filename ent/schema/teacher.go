package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Teacher holds the schema definition for the Teacher entity.
type Teacher struct {
	ent.Schema
}

// Fields of the Teacher.
func (Teacher) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("name"),
		field.String("last_name"),
		field.String("email"),
		field.String("password"),
	}
}

// Edges of the Teacher.
func (Teacher) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("classes", Class.Type),
		edge.To("scoreSyncs", ScoreSync.Type),
		edge.To("studentSyncs", StudentSync.Type),
		edge.To("activitySyncs", ActivitySync.Type),
		edge.To("attendanceSyncs", AttendanceSync.Type),
		edge.To("classPeriodSyncs", ClassPeriodSync.Type),
	}
}
