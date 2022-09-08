package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// School holds the schema definition for the School entity.
type School struct {
	ent.Schema
}

// Fields of the School.
func (School) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("name"),
		field.String("lat"),
		field.String("lon"),
	}
}

// Edges of the School.
func (School) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("classes", Class.Type),
		edge.From("municipio", Municipio.Type).Ref("schools").Unique(),
	}
}
