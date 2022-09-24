package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// StudentSync holds the schema definition for the StudentSync entity.
type StudentSync struct {
	ent.Schema
}

// Fields of the StudentSync.
func (StudentSync) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("last_sync_id"),
	}
}

// Edges of the StudentSync.
func (StudentSync) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("teacher", Teacher.Type).Ref("studentSyncs").Unique(),
	}
}
