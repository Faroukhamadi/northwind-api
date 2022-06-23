package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// City holds the schema definition for the City entity.
type City struct {
	ent.Schema
}

// Fields of the City.
func (City) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		// field.String("country_code").MinLen(3).MaxLen(3),
		field.String("district"),
		field.Int("population"),
	}
}

// Edges of the City.
func (City) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("country_code", Country.Type).
			Ref("cities").
			Unique(),
		edge.From("country", Country.Type).
			Ref("capital").
			Unique().
			Required(),
		edge.To("country", Country.Type).
			Unique(),
	}
}
