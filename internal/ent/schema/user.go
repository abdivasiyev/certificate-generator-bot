package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Unique().MinLen(5),
		field.Int64("telegram_id").Unique(),

		field.Int64("quiz_id").Default(0).NonNegative(),
		field.Int64("incorrect_quiz").Default(0).NonNegative(),
		field.Int64("correct_quiz").Default(0).NonNegative(),

		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional().UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
