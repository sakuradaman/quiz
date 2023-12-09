package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Quiz holds the schema definition for the User entity.
type Quiz struct {
	ent.Schema
}

// Fields of the Quiz.
func (Quiz) Fields() []ent.Field {
	return []ent.Field{
		field.String("problem").
			Default("unknown"),
		field.String("answer").
			Default("unknown"),
	}
}

// Edges of the Quiz.
func (Quiz) Edges() []ent.Edge {
	return nil
}
