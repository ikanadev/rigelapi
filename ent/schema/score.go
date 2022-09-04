package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Score holds the schema definition for the Score entity.
type Score struct {
	ent.Schema
}

// Fields of the Score.
func (Score) Fields() []ent.Field {
	return []ent.Field{
    field.String("id"),
    field.Int("points"),
  }
}

// Edges of the Score.
func (Score) Edges() []ent.Edge {
	return []ent.Edge{
    edge.From("activity", Activity.Type).Ref("scores").Unique(),
    edge.From("student", Student.Type).Ref("scores").Unique(),
  }
}
