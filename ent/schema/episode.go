package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Episode holds the schema definition for the Episode entity.
type Episode struct {
	ent.Schema
}

// Fields of the Episode.
func (Episode) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive().Unique().Immutable(),
		field.Int("episode_number").Positive(),
		field.String("title").NotEmpty(),
		field.String("overview").Optional(),
		field.Time("air_date").Optional(),
		field.Int("season_id").Positive().Unique().Optional(),
	}
}

// Edges of the Episode.
func (Episode) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("season", Season.Type).
			Ref("episodes").
			Field("season_id").
			Unique(),
		edge.To("cast_members", CastMember.Type),
	}
}
