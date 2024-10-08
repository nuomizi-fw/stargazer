// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/nuomizi-fw/stargazer/ent/bangumi"
	"github.com/nuomizi-fw/stargazer/ent/season"
)

// Season is the model entity for the Season schema.
type Season struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// SeasonNumber holds the value of the "season_number" field.
	SeasonNumber int `json:"season_number,omitempty"`
	// AirDate holds the value of the "air_date" field.
	AirDate time.Time `json:"air_date,omitempty"`
	// EpisodeCount holds the value of the "episode_count" field.
	EpisodeCount int `json:"episode_count,omitempty"`
	// BangumiID holds the value of the "bangumi_id" field.
	BangumiID int `json:"bangumi_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SeasonQuery when eager-loading is set.
	Edges        SeasonEdges `json:"edges"`
	selectValues sql.SelectValues
}

// SeasonEdges holds the relations/edges for other nodes in the graph.
type SeasonEdges struct {
	// Bangumi holds the value of the bangumi edge.
	Bangumi *Bangumi `json:"bangumi,omitempty"`
	// Episodes holds the value of the episodes edge.
	Episodes []*Episode `json:"episodes,omitempty"`
	// CastMembers holds the value of the cast_members edge.
	CastMembers []*CastMember `json:"cast_members,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// BangumiOrErr returns the Bangumi value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SeasonEdges) BangumiOrErr() (*Bangumi, error) {
	if e.Bangumi != nil {
		return e.Bangumi, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: bangumi.Label}
	}
	return nil, &NotLoadedError{edge: "bangumi"}
}

// EpisodesOrErr returns the Episodes value or an error if the edge
// was not loaded in eager-loading.
func (e SeasonEdges) EpisodesOrErr() ([]*Episode, error) {
	if e.loadedTypes[1] {
		return e.Episodes, nil
	}
	return nil, &NotLoadedError{edge: "episodes"}
}

// CastMembersOrErr returns the CastMembers value or an error if the edge
// was not loaded in eager-loading.
func (e SeasonEdges) CastMembersOrErr() ([]*CastMember, error) {
	if e.loadedTypes[2] {
		return e.CastMembers, nil
	}
	return nil, &NotLoadedError{edge: "cast_members"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Season) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case season.FieldID, season.FieldSeasonNumber, season.FieldEpisodeCount, season.FieldBangumiID:
			values[i] = new(sql.NullInt64)
		case season.FieldAirDate:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Season fields.
func (s *Season) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case season.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case season.FieldSeasonNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field season_number", values[i])
			} else if value.Valid {
				s.SeasonNumber = int(value.Int64)
			}
		case season.FieldAirDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field air_date", values[i])
			} else if value.Valid {
				s.AirDate = value.Time
			}
		case season.FieldEpisodeCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field episode_count", values[i])
			} else if value.Valid {
				s.EpisodeCount = int(value.Int64)
			}
		case season.FieldBangumiID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field bangumi_id", values[i])
			} else if value.Valid {
				s.BangumiID = int(value.Int64)
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Season.
// This includes values selected through modifiers, order, etc.
func (s *Season) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryBangumi queries the "bangumi" edge of the Season entity.
func (s *Season) QueryBangumi() *BangumiQuery {
	return NewSeasonClient(s.config).QueryBangumi(s)
}

// QueryEpisodes queries the "episodes" edge of the Season entity.
func (s *Season) QueryEpisodes() *EpisodeQuery {
	return NewSeasonClient(s.config).QueryEpisodes(s)
}

// QueryCastMembers queries the "cast_members" edge of the Season entity.
func (s *Season) QueryCastMembers() *CastMemberQuery {
	return NewSeasonClient(s.config).QueryCastMembers(s)
}

// Update returns a builder for updating this Season.
// Note that you need to call Season.Unwrap() before calling this method if this Season
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Season) Update() *SeasonUpdateOne {
	return NewSeasonClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Season entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Season) Unwrap() *Season {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Season is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Season) String() string {
	var builder strings.Builder
	builder.WriteString("Season(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("season_number=")
	builder.WriteString(fmt.Sprintf("%v", s.SeasonNumber))
	builder.WriteString(", ")
	builder.WriteString("air_date=")
	builder.WriteString(s.AirDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("episode_count=")
	builder.WriteString(fmt.Sprintf("%v", s.EpisodeCount))
	builder.WriteString(", ")
	builder.WriteString("bangumi_id=")
	builder.WriteString(fmt.Sprintf("%v", s.BangumiID))
	builder.WriteByte(')')
	return builder.String()
}

// Seasons is a parsable slice of Season.
type Seasons []*Season
