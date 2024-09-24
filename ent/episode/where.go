// Code generated by ent, DO NOT EDIT.

package episode

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/nuomizi-fw/stargazer/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Episode {
	return predicate.Episode(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Episode {
	return predicate.Episode(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Episode {
	return predicate.Episode(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Episode {
	return predicate.Episode(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Episode {
	return predicate.Episode(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Episode {
	return predicate.Episode(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Episode {
	return predicate.Episode(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Episode {
	return predicate.Episode(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Episode {
	return predicate.Episode(sql.FieldLTE(FieldID, id))
}

// EpisodeNumber applies equality check predicate on the "episode_number" field. It's identical to EpisodeNumberEQ.
func EpisodeNumber(v int) predicate.Episode {
	return predicate.Episode(sql.FieldEQ(FieldEpisodeNumber, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Episode {
	return predicate.Episode(sql.FieldEQ(FieldTitle, v))
}

// Overview applies equality check predicate on the "overview" field. It's identical to OverviewEQ.
func Overview(v string) predicate.Episode {
	return predicate.Episode(sql.FieldEQ(FieldOverview, v))
}

// AirDate applies equality check predicate on the "air_date" field. It's identical to AirDateEQ.
func AirDate(v time.Time) predicate.Episode {
	return predicate.Episode(sql.FieldEQ(FieldAirDate, v))
}

// SeasonID applies equality check predicate on the "season_id" field. It's identical to SeasonIDEQ.
func SeasonID(v int) predicate.Episode {
	return predicate.Episode(sql.FieldEQ(FieldSeasonID, v))
}

// EpisodeNumberEQ applies the EQ predicate on the "episode_number" field.
func EpisodeNumberEQ(v int) predicate.Episode {
	return predicate.Episode(sql.FieldEQ(FieldEpisodeNumber, v))
}

// EpisodeNumberNEQ applies the NEQ predicate on the "episode_number" field.
func EpisodeNumberNEQ(v int) predicate.Episode {
	return predicate.Episode(sql.FieldNEQ(FieldEpisodeNumber, v))
}

// EpisodeNumberIn applies the In predicate on the "episode_number" field.
func EpisodeNumberIn(vs ...int) predicate.Episode {
	return predicate.Episode(sql.FieldIn(FieldEpisodeNumber, vs...))
}

// EpisodeNumberNotIn applies the NotIn predicate on the "episode_number" field.
func EpisodeNumberNotIn(vs ...int) predicate.Episode {
	return predicate.Episode(sql.FieldNotIn(FieldEpisodeNumber, vs...))
}

// EpisodeNumberGT applies the GT predicate on the "episode_number" field.
func EpisodeNumberGT(v int) predicate.Episode {
	return predicate.Episode(sql.FieldGT(FieldEpisodeNumber, v))
}

// EpisodeNumberGTE applies the GTE predicate on the "episode_number" field.
func EpisodeNumberGTE(v int) predicate.Episode {
	return predicate.Episode(sql.FieldGTE(FieldEpisodeNumber, v))
}

// EpisodeNumberLT applies the LT predicate on the "episode_number" field.
func EpisodeNumberLT(v int) predicate.Episode {
	return predicate.Episode(sql.FieldLT(FieldEpisodeNumber, v))
}

// EpisodeNumberLTE applies the LTE predicate on the "episode_number" field.
func EpisodeNumberLTE(v int) predicate.Episode {
	return predicate.Episode(sql.FieldLTE(FieldEpisodeNumber, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Episode {
	return predicate.Episode(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Episode {
	return predicate.Episode(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Episode {
	return predicate.Episode(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Episode {
	return predicate.Episode(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Episode {
	return predicate.Episode(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Episode {
	return predicate.Episode(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Episode {
	return predicate.Episode(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Episode {
	return predicate.Episode(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Episode {
	return predicate.Episode(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Episode {
	return predicate.Episode(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Episode {
	return predicate.Episode(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Episode {
	return predicate.Episode(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Episode {
	return predicate.Episode(sql.FieldContainsFold(FieldTitle, v))
}

// OverviewEQ applies the EQ predicate on the "overview" field.
func OverviewEQ(v string) predicate.Episode {
	return predicate.Episode(sql.FieldEQ(FieldOverview, v))
}

// OverviewNEQ applies the NEQ predicate on the "overview" field.
func OverviewNEQ(v string) predicate.Episode {
	return predicate.Episode(sql.FieldNEQ(FieldOverview, v))
}

// OverviewIn applies the In predicate on the "overview" field.
func OverviewIn(vs ...string) predicate.Episode {
	return predicate.Episode(sql.FieldIn(FieldOverview, vs...))
}

// OverviewNotIn applies the NotIn predicate on the "overview" field.
func OverviewNotIn(vs ...string) predicate.Episode {
	return predicate.Episode(sql.FieldNotIn(FieldOverview, vs...))
}

// OverviewGT applies the GT predicate on the "overview" field.
func OverviewGT(v string) predicate.Episode {
	return predicate.Episode(sql.FieldGT(FieldOverview, v))
}

// OverviewGTE applies the GTE predicate on the "overview" field.
func OverviewGTE(v string) predicate.Episode {
	return predicate.Episode(sql.FieldGTE(FieldOverview, v))
}

// OverviewLT applies the LT predicate on the "overview" field.
func OverviewLT(v string) predicate.Episode {
	return predicate.Episode(sql.FieldLT(FieldOverview, v))
}

// OverviewLTE applies the LTE predicate on the "overview" field.
func OverviewLTE(v string) predicate.Episode {
	return predicate.Episode(sql.FieldLTE(FieldOverview, v))
}

// OverviewContains applies the Contains predicate on the "overview" field.
func OverviewContains(v string) predicate.Episode {
	return predicate.Episode(sql.FieldContains(FieldOverview, v))
}

// OverviewHasPrefix applies the HasPrefix predicate on the "overview" field.
func OverviewHasPrefix(v string) predicate.Episode {
	return predicate.Episode(sql.FieldHasPrefix(FieldOverview, v))
}

// OverviewHasSuffix applies the HasSuffix predicate on the "overview" field.
func OverviewHasSuffix(v string) predicate.Episode {
	return predicate.Episode(sql.FieldHasSuffix(FieldOverview, v))
}

// OverviewIsNil applies the IsNil predicate on the "overview" field.
func OverviewIsNil() predicate.Episode {
	return predicate.Episode(sql.FieldIsNull(FieldOverview))
}

// OverviewNotNil applies the NotNil predicate on the "overview" field.
func OverviewNotNil() predicate.Episode {
	return predicate.Episode(sql.FieldNotNull(FieldOverview))
}

// OverviewEqualFold applies the EqualFold predicate on the "overview" field.
func OverviewEqualFold(v string) predicate.Episode {
	return predicate.Episode(sql.FieldEqualFold(FieldOverview, v))
}

// OverviewContainsFold applies the ContainsFold predicate on the "overview" field.
func OverviewContainsFold(v string) predicate.Episode {
	return predicate.Episode(sql.FieldContainsFold(FieldOverview, v))
}

// AirDateEQ applies the EQ predicate on the "air_date" field.
func AirDateEQ(v time.Time) predicate.Episode {
	return predicate.Episode(sql.FieldEQ(FieldAirDate, v))
}

// AirDateNEQ applies the NEQ predicate on the "air_date" field.
func AirDateNEQ(v time.Time) predicate.Episode {
	return predicate.Episode(sql.FieldNEQ(FieldAirDate, v))
}

// AirDateIn applies the In predicate on the "air_date" field.
func AirDateIn(vs ...time.Time) predicate.Episode {
	return predicate.Episode(sql.FieldIn(FieldAirDate, vs...))
}

// AirDateNotIn applies the NotIn predicate on the "air_date" field.
func AirDateNotIn(vs ...time.Time) predicate.Episode {
	return predicate.Episode(sql.FieldNotIn(FieldAirDate, vs...))
}

// AirDateGT applies the GT predicate on the "air_date" field.
func AirDateGT(v time.Time) predicate.Episode {
	return predicate.Episode(sql.FieldGT(FieldAirDate, v))
}

// AirDateGTE applies the GTE predicate on the "air_date" field.
func AirDateGTE(v time.Time) predicate.Episode {
	return predicate.Episode(sql.FieldGTE(FieldAirDate, v))
}

// AirDateLT applies the LT predicate on the "air_date" field.
func AirDateLT(v time.Time) predicate.Episode {
	return predicate.Episode(sql.FieldLT(FieldAirDate, v))
}

// AirDateLTE applies the LTE predicate on the "air_date" field.
func AirDateLTE(v time.Time) predicate.Episode {
	return predicate.Episode(sql.FieldLTE(FieldAirDate, v))
}

// AirDateIsNil applies the IsNil predicate on the "air_date" field.
func AirDateIsNil() predicate.Episode {
	return predicate.Episode(sql.FieldIsNull(FieldAirDate))
}

// AirDateNotNil applies the NotNil predicate on the "air_date" field.
func AirDateNotNil() predicate.Episode {
	return predicate.Episode(sql.FieldNotNull(FieldAirDate))
}

// SeasonIDEQ applies the EQ predicate on the "season_id" field.
func SeasonIDEQ(v int) predicate.Episode {
	return predicate.Episode(sql.FieldEQ(FieldSeasonID, v))
}

// SeasonIDNEQ applies the NEQ predicate on the "season_id" field.
func SeasonIDNEQ(v int) predicate.Episode {
	return predicate.Episode(sql.FieldNEQ(FieldSeasonID, v))
}

// SeasonIDIn applies the In predicate on the "season_id" field.
func SeasonIDIn(vs ...int) predicate.Episode {
	return predicate.Episode(sql.FieldIn(FieldSeasonID, vs...))
}

// SeasonIDNotIn applies the NotIn predicate on the "season_id" field.
func SeasonIDNotIn(vs ...int) predicate.Episode {
	return predicate.Episode(sql.FieldNotIn(FieldSeasonID, vs...))
}

// SeasonIDIsNil applies the IsNil predicate on the "season_id" field.
func SeasonIDIsNil() predicate.Episode {
	return predicate.Episode(sql.FieldIsNull(FieldSeasonID))
}

// SeasonIDNotNil applies the NotNil predicate on the "season_id" field.
func SeasonIDNotNil() predicate.Episode {
	return predicate.Episode(sql.FieldNotNull(FieldSeasonID))
}

// HasSeason applies the HasEdge predicate on the "season" edge.
func HasSeason() predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SeasonTable, SeasonColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSeasonWith applies the HasEdge predicate on the "season" edge with a given conditions (other predicates).
func HasSeasonWith(preds ...predicate.Season) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		step := newSeasonStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCastMembers applies the HasEdge predicate on the "cast_members" edge.
func HasCastMembers() predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CastMembersTable, CastMembersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCastMembersWith applies the HasEdge predicate on the "cast_members" edge with a given conditions (other predicates).
func HasCastMembersWith(preds ...predicate.CastMember) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		step := newCastMembersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Episode) predicate.Episode {
	return predicate.Episode(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Episode) predicate.Episode {
	return predicate.Episode(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Episode) predicate.Episode {
	return predicate.Episode(sql.NotPredicates(p))
}