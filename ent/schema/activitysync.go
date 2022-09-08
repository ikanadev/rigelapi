package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ActivitySync holds the schema definition for the ActivitySync entity.
type ActivitySync struct {
	ent.Schema
}

// Fields of the ActivitySync.
func (ActivitySync) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("last_sync_id"),
	}
}

// Edges of the ActivitySync.
func (ActivitySync) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("classPeriod", ClassPeriod.Type).Ref("activitySyncs").Unique(),
	}
}
