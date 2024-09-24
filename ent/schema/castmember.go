package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CastMember holds the schema definition for the CastMember entity.
type CastMember struct {
	ent.Schema
}

// Fields of the CastMember.
func (CastMember) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive().Unique().Immutable(),
		field.String("name").NotEmpty(),
		field.String("character_name").Optional(),
		field.String("profile_path").Optional(),
		field.Int("season_id").Positive().Unique().Optional(),
	}
}

// Edges of the CastMember.
func (CastMember) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("season", Season.Type).
			Ref("cast_members").
			Field("season_id").
			Unique(),
	}
}
