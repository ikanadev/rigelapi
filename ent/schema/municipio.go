package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Municipio holds the schema definition for the Municipio entity.
type Municipio struct {
	ent.Schema
}

// Fields of the Municipio.
func (Municipio) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("name"),
	}
}

// Edges of the Municipio.
func (Municipio) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("schools", School.Type),
		edge.From("provincia", Provincia.Type).Ref("municipios").Unique(),
	}
}
