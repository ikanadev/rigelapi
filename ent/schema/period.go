package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Period holds the schema definition for the Period entity.
type Period struct {
	ent.Schema
}

// Fields of the Period.
func (Period) Fields() []ent.Field {
	return []ent.Field{
    field.String("id"),
    field.String("name"),
  }
}

// Edges of the Period.
func (Period) Edges() []ent.Edge {
	return []ent.Edge{
    edge.To("classPeriods", ClassPeriod.Type),
    edge.From("year", Year.Type).Ref("periods").Unique(),
  }
}
