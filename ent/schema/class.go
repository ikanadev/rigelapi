package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Class holds the schema definition for the Class entity.
type Class struct {
	ent.Schema
}

// Fields of the Class.
func (Class) Fields() []ent.Field {
	return []ent.Field{
    field.String("id"),
    field.String("parallel"),
  }
}

// Edges of the Class.
func (Class) Edges() []ent.Edge {
	return []ent.Edge{
    edge.To("students", Student.Type),
    edge.To("classPeriods", ClassPeriod.Type),
    edge.To("classPeriodSyncs", ClassPeriodSync.Type),
    edge.To("studentSyncs", StudentSync.Type),
    edge.From("school", School.Type).Ref("classes").Unique(),
    edge.From("teacher", Teacher.Type).Ref("classes").Unique(),
    edge.From("subject", Subject.Type).Ref("classes").Unique(),
    edge.From("grade", Grade.Type).Ref("classes").Unique(),
  }
}
