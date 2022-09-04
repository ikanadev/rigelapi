package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ClassPeriodSync holds the schema definition for the ClassPeriodSync entity.
type ClassPeriodSync struct {
	ent.Schema
}

// Fields of the ClassPeriodSync.
func (ClassPeriodSync) Fields() []ent.Field {
	return []ent.Field{
    field.String("id"),
    field.String("last_sync_id"),
  }
}

// Edges of the ClassPeriodSync.
func (ClassPeriodSync) Edges() []ent.Edge {
	return []ent.Edge{
    edge.From("class", Class.Type).Ref("classPeriodSyncs").Unique(),
  }
}
