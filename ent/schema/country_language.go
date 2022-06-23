package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Country_language holds the schema definition for the Country_language entity.
type Country_language struct {
	ent.Schema
}

// Fields of the Country_language.
func (Country_language) Fields() []ent.Field {
	return []ent.Field{
		field.String("country_code"),
		field.String("language"),
		field.Bool("is_official"),
		field.Float("percentage"),
	}
}

// Edges of the Country_language.
func (Country_language) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("country", Country.Type).
			Ref("language").
			Unique().
			Required(),
	}
}
