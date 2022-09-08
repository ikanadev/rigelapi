package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ScoreSync holds the schema definition for the ScoreSync entity.
type ScoreSync struct {
	ent.Schema
}

// Fields of the ScoreSync.
func (ScoreSync) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("last_sync_id"),
	}
}

// Edges of the ScoreSync.
func (ScoreSync) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("activity", Activity.Type).Ref("scoreSyncs").Unique(),
	}
}
