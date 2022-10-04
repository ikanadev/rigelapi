package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// AppError holds the schema definition for the AppError entity.
type AppError struct {
	ent.Schema
}

// Fields of the AppError.
func (AppError) Fields() []ent.Field {
	return []ent.Field{
    field.String("id"),
    field.String("user_id"),
    field.String("cause"),
    field.String("error_msg"),
    field.String("error_stack"),
  }
}

// Edges of the AppError.
func (AppError) Edges() []ent.Edge {
	return nil
}
