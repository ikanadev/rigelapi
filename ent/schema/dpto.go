package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Dpto holds the schema definition for the Dpto entity.
type Dpto struct {
	ent.Schema
}

// Fields of the Dpto.
func (Dpto) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("name"),
	}
}

// Edges of the Dpto.
func (Dpto) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("provincias", Provincia.Type),
	}
}
