package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Bangumi holds the schema definition for the Bangumi entity.
type Bangumi struct {
	ent.Schema
}

// Fields of the Bangumi.
func (Bangumi) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive().Unique().Immutable(),
		field.String("title").NotEmpty(),
		field.String("english_title").Optional(),
		field.String("japanese_title").Optional(),
		field.String("type").NotEmpty(),   // e.g. "TV", "Movie", "OVA", etc.
		field.String("status").NotEmpty(), // e.g. "Airing", "Completed", "Canceled", etc.
		field.Float("score").Optional(),
		field.String("tags").Optional(),
		field.String("synopsis").Optional(),
		field.String("cover_image").Optional(),
		field.String("trailer_url").Optional(),
	}
}

// Edges of the Bangumi.
func (Bangumi) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("seasons", Season.Type),
		edge.To("cast_members", CastMember.Type),
	}
}
