package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Provincia holds the schema definition for the Provincia entity.
type Provincia struct {
	ent.Schema
}

// Fields of the Provincia.
func (Provincia) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("name"),
	}
}

// Edges of the Provincia.
func (Provincia) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("municipios", Municipio.Type),
		edge.From("departamento", Dpto.Type).Ref("provincias").Unique(),
	}
}
