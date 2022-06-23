package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Country holds the schema definition for the Country entity.
type Country struct {
	ent.Schema
}

// Fields of the Country.
func (Country) Fields() []ent.Field {
	return []ent.Field{
		// field.String("id").StorageKey("code")
		field.String("id").
			StorageKey("code").
			MinLen(3).
			MaxLen(3),
		field.String("name"),
		field.String("region"),
		field.Float("surface_area"),
		field.Int16("indep_year"),
		field.Int("population"),
		field.Float("life_expectancy"),
		field.Float("gnp"),
		field.Float("gnp_old"),
		field.String("local_name"),
		field.String("government_form"),
		field.String("head_of_state").Nillable().Optional(),
		// field.Int("capital"),
		field.String("code2").
			MinLen(2).
			MaxLen(2),
	}
}

// Edges of the Country.
func (Country) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("capital", City.Type).
			Unique(),
		edge.To("language", Country_language.Type).
			Unique(),
		edge.From("city", City.Type).
			Ref("country").
			Unique().
			Required(),
	}
}
