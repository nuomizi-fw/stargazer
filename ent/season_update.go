// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/nuomizi-fw/stargazer/ent/bangumi"
	"github.com/nuomizi-fw/stargazer/ent/castmember"
	"github.com/nuomizi-fw/stargazer/ent/episode"
	"github.com/nuomizi-fw/stargazer/ent/predicate"
	"github.com/nuomizi-fw/stargazer/ent/season"
)

// SeasonUpdate is the builder for updating Season entities.
type SeasonUpdate struct {
	config
	hooks    []Hook
	mutation *SeasonMutation
}

// Where appends a list predicates to the SeasonUpdate builder.
func (su *SeasonUpdate) Where(ps ...predicate.Season) *SeasonUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetSeasonNumber sets the "season_number" field.
func (su *SeasonUpdate) SetSeasonNumber(i int) *SeasonUpdate {
	su.mutation.ResetSeasonNumber()
	su.mutation.SetSeasonNumber(i)
	return su
}

// SetNillableSeasonNumber sets the "season_number" field if the given value is not nil.
func (su *SeasonUpdate) SetNillableSeasonNumber(i *int) *SeasonUpdate {
	if i != nil {
		su.SetSeasonNumber(*i)
	}
	return su
}

// AddSeasonNumber adds i to the "season_number" field.
func (su *SeasonUpdate) AddSeasonNumber(i int) *SeasonUpdate {
	su.mutation.AddSeasonNumber(i)
	return su
}

// SetAirDate sets the "air_date" field.
func (su *SeasonUpdate) SetAirDate(t time.Time) *SeasonUpdate {
	su.mutation.SetAirDate(t)
	return su
}

// SetNillableAirDate sets the "air_date" field if the given value is not nil.
func (su *SeasonUpdate) SetNillableAirDate(t *time.Time) *SeasonUpdate {
	if t != nil {
		su.SetAirDate(*t)
	}
	return su
}

// ClearAirDate clears the value of the "air_date" field.
func (su *SeasonUpdate) ClearAirDate() *SeasonUpdate {
	su.mutation.ClearAirDate()
	return su
}

// SetEpisodeCount sets the "episode_count" field.
func (su *SeasonUpdate) SetEpisodeCount(i int) *SeasonUpdate {
	su.mutation.ResetEpisodeCount()
	su.mutation.SetEpisodeCount(i)
	return su
}

// SetNillableEpisodeCount sets the "episode_count" field if the given value is not nil.
func (su *SeasonUpdate) SetNillableEpisodeCount(i *int) *SeasonUpdate {
	if i != nil {
		su.SetEpisodeCount(*i)
	}
	return su
}

// AddEpisodeCount adds i to the "episode_count" field.
func (su *SeasonUpdate) AddEpisodeCount(i int) *SeasonUpdate {
	su.mutation.AddEpisodeCount(i)
	return su
}

// SetBangumiID sets the "bangumi_id" field.
func (su *SeasonUpdate) SetBangumiID(i int) *SeasonUpdate {
	su.mutation.SetBangumiID(i)
	return su
}

// SetNillableBangumiID sets the "bangumi_id" field if the given value is not nil.
func (su *SeasonUpdate) SetNillableBangumiID(i *int) *SeasonUpdate {
	if i != nil {
		su.SetBangumiID(*i)
	}
	return su
}

// ClearBangumiID clears the value of the "bangumi_id" field.
func (su *SeasonUpdate) ClearBangumiID() *SeasonUpdate {
	su.mutation.ClearBangumiID()
	return su
}

// SetBangumi sets the "bangumi" edge to the Bangumi entity.
func (su *SeasonUpdate) SetBangumi(b *Bangumi) *SeasonUpdate {
	return su.SetBangumiID(b.ID)
}

// AddEpisodeIDs adds the "episodes" edge to the Episode entity by IDs.
func (su *SeasonUpdate) AddEpisodeIDs(ids ...int) *SeasonUpdate {
	su.mutation.AddEpisodeIDs(ids...)
	return su
}

// AddEpisodes adds the "episodes" edges to the Episode entity.
func (su *SeasonUpdate) AddEpisodes(e ...*Episode) *SeasonUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return su.AddEpisodeIDs(ids...)
}

// AddCastMemberIDs adds the "cast_members" edge to the CastMember entity by IDs.
func (su *SeasonUpdate) AddCastMemberIDs(ids ...int) *SeasonUpdate {
	su.mutation.AddCastMemberIDs(ids...)
	return su
}

// AddCastMembers adds the "cast_members" edges to the CastMember entity.
func (su *SeasonUpdate) AddCastMembers(c ...*CastMember) *SeasonUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return su.AddCastMemberIDs(ids...)
}

// Mutation returns the SeasonMutation object of the builder.
func (su *SeasonUpdate) Mutation() *SeasonMutation {
	return su.mutation
}

// ClearBangumi clears the "bangumi" edge to the Bangumi entity.
func (su *SeasonUpdate) ClearBangumi() *SeasonUpdate {
	su.mutation.ClearBangumi()
	return su
}

// ClearEpisodes clears all "episodes" edges to the Episode entity.
func (su *SeasonUpdate) ClearEpisodes() *SeasonUpdate {
	su.mutation.ClearEpisodes()
	return su
}

// RemoveEpisodeIDs removes the "episodes" edge to Episode entities by IDs.
func (su *SeasonUpdate) RemoveEpisodeIDs(ids ...int) *SeasonUpdate {
	su.mutation.RemoveEpisodeIDs(ids...)
	return su
}

// RemoveEpisodes removes "episodes" edges to Episode entities.
func (su *SeasonUpdate) RemoveEpisodes(e ...*Episode) *SeasonUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return su.RemoveEpisodeIDs(ids...)
}

// ClearCastMembers clears all "cast_members" edges to the CastMember entity.
func (su *SeasonUpdate) ClearCastMembers() *SeasonUpdate {
	su.mutation.ClearCastMembers()
	return su
}

// RemoveCastMemberIDs removes the "cast_members" edge to CastMember entities by IDs.
func (su *SeasonUpdate) RemoveCastMemberIDs(ids ...int) *SeasonUpdate {
	su.mutation.RemoveCastMemberIDs(ids...)
	return su
}

// RemoveCastMembers removes "cast_members" edges to CastMember entities.
func (su *SeasonUpdate) RemoveCastMembers(c ...*CastMember) *SeasonUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return su.RemoveCastMemberIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SeasonUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SeasonUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SeasonUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SeasonUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *SeasonUpdate) check() error {
	if v, ok := su.mutation.SeasonNumber(); ok {
		if err := season.SeasonNumberValidator(v); err != nil {
			return &ValidationError{Name: "season_number", err: fmt.Errorf(`ent: validator failed for field "Season.season_number": %w`, err)}
		}
	}
	if v, ok := su.mutation.EpisodeCount(); ok {
		if err := season.EpisodeCountValidator(v); err != nil {
			return &ValidationError{Name: "episode_count", err: fmt.Errorf(`ent: validator failed for field "Season.episode_count": %w`, err)}
		}
	}
	if v, ok := su.mutation.BangumiID(); ok {
		if err := season.BangumiIDValidator(v); err != nil {
			return &ValidationError{Name: "bangumi_id", err: fmt.Errorf(`ent: validator failed for field "Season.bangumi_id": %w`, err)}
		}
	}
	return nil
}

func (su *SeasonUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := su.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(season.Table, season.Columns, sqlgraph.NewFieldSpec(season.FieldID, field.TypeInt))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.SeasonNumber(); ok {
		_spec.SetField(season.FieldSeasonNumber, field.TypeInt, value)
	}
	if value, ok := su.mutation.AddedSeasonNumber(); ok {
		_spec.AddField(season.FieldSeasonNumber, field.TypeInt, value)
	}
	if value, ok := su.mutation.AirDate(); ok {
		_spec.SetField(season.FieldAirDate, field.TypeTime, value)
	}
	if su.mutation.AirDateCleared() {
		_spec.ClearField(season.FieldAirDate, field.TypeTime)
	}
	if value, ok := su.mutation.EpisodeCount(); ok {
		_spec.SetField(season.FieldEpisodeCount, field.TypeInt, value)
	}
	if value, ok := su.mutation.AddedEpisodeCount(); ok {
		_spec.AddField(season.FieldEpisodeCount, field.TypeInt, value)
	}
	if su.mutation.BangumiCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   season.BangumiTable,
			Columns: []string{season.BangumiColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bangumi.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.BangumiIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   season.BangumiTable,
			Columns: []string{season.BangumiColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bangumi.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.EpisodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   season.EpisodesTable,
			Columns: []string{season.EpisodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(episode.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedEpisodesIDs(); len(nodes) > 0 && !su.mutation.EpisodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   season.EpisodesTable,
			Columns: []string{season.EpisodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(episode.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.EpisodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   season.EpisodesTable,
			Columns: []string{season.EpisodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(episode.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.CastMembersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   season.CastMembersTable,
			Columns: []string{season.CastMembersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(castmember.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedCastMembersIDs(); len(nodes) > 0 && !su.mutation.CastMembersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   season.CastMembersTable,
			Columns: []string{season.CastMembersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(castmember.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.CastMembersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   season.CastMembersTable,
			Columns: []string{season.CastMembersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(castmember.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{season.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SeasonUpdateOne is the builder for updating a single Season entity.
type SeasonUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SeasonMutation
}

// SetSeasonNumber sets the "season_number" field.
func (suo *SeasonUpdateOne) SetSeasonNumber(i int) *SeasonUpdateOne {
	suo.mutation.ResetSeasonNumber()
	suo.mutation.SetSeasonNumber(i)
	return suo
}

// SetNillableSeasonNumber sets the "season_number" field if the given value is not nil.
func (suo *SeasonUpdateOne) SetNillableSeasonNumber(i *int) *SeasonUpdateOne {
	if i != nil {
		suo.SetSeasonNumber(*i)
	}
	return suo
}

// AddSeasonNumber adds i to the "season_number" field.
func (suo *SeasonUpdateOne) AddSeasonNumber(i int) *SeasonUpdateOne {
	suo.mutation.AddSeasonNumber(i)
	return suo
}

// SetAirDate sets the "air_date" field.
func (suo *SeasonUpdateOne) SetAirDate(t time.Time) *SeasonUpdateOne {
	suo.mutation.SetAirDate(t)
	return suo
}

// SetNillableAirDate sets the "air_date" field if the given value is not nil.
func (suo *SeasonUpdateOne) SetNillableAirDate(t *time.Time) *SeasonUpdateOne {
	if t != nil {
		suo.SetAirDate(*t)
	}
	return suo
}

// ClearAirDate clears the value of the "air_date" field.
func (suo *SeasonUpdateOne) ClearAirDate() *SeasonUpdateOne {
	suo.mutation.ClearAirDate()
	return suo
}

// SetEpisodeCount sets the "episode_count" field.
func (suo *SeasonUpdateOne) SetEpisodeCount(i int) *SeasonUpdateOne {
	suo.mutation.ResetEpisodeCount()
	suo.mutation.SetEpisodeCount(i)
	return suo
}

// SetNillableEpisodeCount sets the "episode_count" field if the given value is not nil.
func (suo *SeasonUpdateOne) SetNillableEpisodeCount(i *int) *SeasonUpdateOne {
	if i != nil {
		suo.SetEpisodeCount(*i)
	}
	return suo
}

// AddEpisodeCount adds i to the "episode_count" field.
func (suo *SeasonUpdateOne) AddEpisodeCount(i int) *SeasonUpdateOne {
	suo.mutation.AddEpisodeCount(i)
	return suo
}

// SetBangumiID sets the "bangumi_id" field.
func (suo *SeasonUpdateOne) SetBangumiID(i int) *SeasonUpdateOne {
	suo.mutation.SetBangumiID(i)
	return suo
}

// SetNillableBangumiID sets the "bangumi_id" field if the given value is not nil.
func (suo *SeasonUpdateOne) SetNillableBangumiID(i *int) *SeasonUpdateOne {
	if i != nil {
		suo.SetBangumiID(*i)
	}
	return suo
}

// ClearBangumiID clears the value of the "bangumi_id" field.
func (suo *SeasonUpdateOne) ClearBangumiID() *SeasonUpdateOne {
	suo.mutation.ClearBangumiID()
	return suo
}

// SetBangumi sets the "bangumi" edge to the Bangumi entity.
func (suo *SeasonUpdateOne) SetBangumi(b *Bangumi) *SeasonUpdateOne {
	return suo.SetBangumiID(b.ID)
}

// AddEpisodeIDs adds the "episodes" edge to the Episode entity by IDs.
func (suo *SeasonUpdateOne) AddEpisodeIDs(ids ...int) *SeasonUpdateOne {
	suo.mutation.AddEpisodeIDs(ids...)
	return suo
}

// AddEpisodes adds the "episodes" edges to the Episode entity.
func (suo *SeasonUpdateOne) AddEpisodes(e ...*Episode) *SeasonUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return suo.AddEpisodeIDs(ids...)
}

// AddCastMemberIDs adds the "cast_members" edge to the CastMember entity by IDs.
func (suo *SeasonUpdateOne) AddCastMemberIDs(ids ...int) *SeasonUpdateOne {
	suo.mutation.AddCastMemberIDs(ids...)
	return suo
}

// AddCastMembers adds the "cast_members" edges to the CastMember entity.
func (suo *SeasonUpdateOne) AddCastMembers(c ...*CastMember) *SeasonUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return suo.AddCastMemberIDs(ids...)
}

// Mutation returns the SeasonMutation object of the builder.
func (suo *SeasonUpdateOne) Mutation() *SeasonMutation {
	return suo.mutation
}

// ClearBangumi clears the "bangumi" edge to the Bangumi entity.
func (suo *SeasonUpdateOne) ClearBangumi() *SeasonUpdateOne {
	suo.mutation.ClearBangumi()
	return suo
}

// ClearEpisodes clears all "episodes" edges to the Episode entity.
func (suo *SeasonUpdateOne) ClearEpisodes() *SeasonUpdateOne {
	suo.mutation.ClearEpisodes()
	return suo
}

// RemoveEpisodeIDs removes the "episodes" edge to Episode entities by IDs.
func (suo *SeasonUpdateOne) RemoveEpisodeIDs(ids ...int) *SeasonUpdateOne {
	suo.mutation.RemoveEpisodeIDs(ids...)
	return suo
}

// RemoveEpisodes removes "episodes" edges to Episode entities.
func (suo *SeasonUpdateOne) RemoveEpisodes(e ...*Episode) *SeasonUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return suo.RemoveEpisodeIDs(ids...)
}

// ClearCastMembers clears all "cast_members" edges to the CastMember entity.
func (suo *SeasonUpdateOne) ClearCastMembers() *SeasonUpdateOne {
	suo.mutation.ClearCastMembers()
	return suo
}

// RemoveCastMemberIDs removes the "cast_members" edge to CastMember entities by IDs.
func (suo *SeasonUpdateOne) RemoveCastMemberIDs(ids ...int) *SeasonUpdateOne {
	suo.mutation.RemoveCastMemberIDs(ids...)
	return suo
}

// RemoveCastMembers removes "cast_members" edges to CastMember entities.
func (suo *SeasonUpdateOne) RemoveCastMembers(c ...*CastMember) *SeasonUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return suo.RemoveCastMemberIDs(ids...)
}

// Where appends a list predicates to the SeasonUpdate builder.
func (suo *SeasonUpdateOne) Where(ps ...predicate.Season) *SeasonUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SeasonUpdateOne) Select(field string, fields ...string) *SeasonUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Season entity.
func (suo *SeasonUpdateOne) Save(ctx context.Context) (*Season, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SeasonUpdateOne) SaveX(ctx context.Context) *Season {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SeasonUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SeasonUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *SeasonUpdateOne) check() error {
	if v, ok := suo.mutation.SeasonNumber(); ok {
		if err := season.SeasonNumberValidator(v); err != nil {
			return &ValidationError{Name: "season_number", err: fmt.Errorf(`ent: validator failed for field "Season.season_number": %w`, err)}
		}
	}
	if v, ok := suo.mutation.EpisodeCount(); ok {
		if err := season.EpisodeCountValidator(v); err != nil {
			return &ValidationError{Name: "episode_count", err: fmt.Errorf(`ent: validator failed for field "Season.episode_count": %w`, err)}
		}
	}
	if v, ok := suo.mutation.BangumiID(); ok {
		if err := season.BangumiIDValidator(v); err != nil {
			return &ValidationError{Name: "bangumi_id", err: fmt.Errorf(`ent: validator failed for field "Season.bangumi_id": %w`, err)}
		}
	}
	return nil
}

func (suo *SeasonUpdateOne) sqlSave(ctx context.Context) (_node *Season, err error) {
	if err := suo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(season.Table, season.Columns, sqlgraph.NewFieldSpec(season.FieldID, field.TypeInt))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Season.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, season.FieldID)
		for _, f := range fields {
			if !season.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != season.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.SeasonNumber(); ok {
		_spec.SetField(season.FieldSeasonNumber, field.TypeInt, value)
	}
	if value, ok := suo.mutation.AddedSeasonNumber(); ok {
		_spec.AddField(season.FieldSeasonNumber, field.TypeInt, value)
	}
	if value, ok := suo.mutation.AirDate(); ok {
		_spec.SetField(season.FieldAirDate, field.TypeTime, value)
	}
	if suo.mutation.AirDateCleared() {
		_spec.ClearField(season.FieldAirDate, field.TypeTime)
	}
	if value, ok := suo.mutation.EpisodeCount(); ok {
		_spec.SetField(season.FieldEpisodeCount, field.TypeInt, value)
	}
	if value, ok := suo.mutation.AddedEpisodeCount(); ok {
		_spec.AddField(season.FieldEpisodeCount, field.TypeInt, value)
	}
	if suo.mutation.BangumiCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   season.BangumiTable,
			Columns: []string{season.BangumiColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bangumi.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.BangumiIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   season.BangumiTable,
			Columns: []string{season.BangumiColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bangumi.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.EpisodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   season.EpisodesTable,
			Columns: []string{season.EpisodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(episode.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedEpisodesIDs(); len(nodes) > 0 && !suo.mutation.EpisodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   season.EpisodesTable,
			Columns: []string{season.EpisodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(episode.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.EpisodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   season.EpisodesTable,
			Columns: []string{season.EpisodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(episode.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.CastMembersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   season.CastMembersTable,
			Columns: []string{season.CastMembersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(castmember.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedCastMembersIDs(); len(nodes) > 0 && !suo.mutation.CastMembersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   season.CastMembersTable,
			Columns: []string{season.CastMembersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(castmember.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.CastMembersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   season.CastMembersTable,
			Columns: []string{season.CastMembersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(castmember.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Season{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{season.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
