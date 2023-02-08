package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AdminAction holds the schema definition for the AdminAction entity.
type AdminAction struct {
	ent.Schema
}

// Fields of the AdminAction.
func (AdminAction) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("action"),
		field.String("info"),
	}
}

// Edges of the AdminAction.
func (AdminAction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("teacher", Teacher.Type).Ref("actions").Unique(),
	}
}
