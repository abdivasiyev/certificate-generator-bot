package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Quiz holds the schema definition for the Quiz entity.
type Quiz struct {
	ent.Schema
}

// Fields of the Quiz.
func (Quiz) Fields() []ent.Field {
	return []ent.Field{
		field.Text("question").NotEmpty(),
		field.Int64("next_quiz_id").NonNegative(),
		field.Text("option_a").NotEmpty(),
		field.Text("option_b").NotEmpty(),
		field.Text("option_c").NotEmpty(),
		field.Text("option_d").NotEmpty(),

		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional().UpdateDefault(time.Now),
	}
}

// Edges of the Quiz.
func (Quiz) Edges() []ent.Edge {
	return nil
}
