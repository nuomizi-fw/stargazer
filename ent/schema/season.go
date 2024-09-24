package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Season holds the schema definition for the Season entity.
type Season struct {
	ent.Schema
}

// Fields of the Season.
func (Season) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive().Unique().Immutable(),
		field.Int("season_number").Positive(),
		field.Time("air_date").Optional(),
		field.Int("episode_count").Positive(),
		field.Int("bangumi_id").Positive().Unique().Optional(),
	}
}

// Edges of the Season.
func (Season) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("bangumi", Bangumi.Type).
			Ref("seasons").
			Field("bangumi_id").
			Unique(),
		edge.To("episodes", Episode.Type),
		edge.To("cast_members", CastMember.Type),
	}
}
