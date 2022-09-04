package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Student holds the schema definition for the Student entity.
type Student struct {
	ent.Schema
}

// Fields of the Student.
func (Student) Fields() []ent.Field {
	return []ent.Field{
    field.String("id"),
    field.String("name"),
    field.String("last_name"),
    field.String("ci"),
  }
}

// Edges of the Student.
func (Student) Edges() []ent.Edge {
	return []ent.Edge{
    edge.To("attendances", Attendance.Type),
    edge.To("scores", Score.Type),
    edge.From("class", Class.Type).Ref("students").Unique(),
  }
}
