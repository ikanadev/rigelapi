package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Year holds the schema definition for the Year entity.
type Year struct {
	ent.Schema
}

// Fields of the Year.
func (Year) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.Int("value"),
	}
}

// Edges of the Year.
func (Year) Edges() []ent.Edge {
	return []ent.Edge{
    edge.To("classes", Class.Type),
    edge.To("periods", Period.Type),
    edge.To("areas", Area.Type),
  }
}
