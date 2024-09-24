// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/nuomizi-fw/stargazer/ent/castmember"
	"github.com/nuomizi-fw/stargazer/ent/predicate"
)

// CastMemberDelete is the builder for deleting a CastMember entity.
type CastMemberDelete struct {
	config
	hooks    []Hook
	mutation *CastMemberMutation
}

// Where appends a list predicates to the CastMemberDelete builder.
func (cmd *CastMemberDelete) Where(ps ...predicate.CastMember) *CastMemberDelete {
	cmd.mutation.Where(ps...)
	return cmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cmd *CastMemberDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, cmd.sqlExec, cmd.mutation, cmd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cmd *CastMemberDelete) ExecX(ctx context.Context) int {
	n, err := cmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cmd *CastMemberDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(castmember.Table, sqlgraph.NewFieldSpec(castmember.FieldID, field.TypeInt))
	if ps := cmd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cmd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cmd.mutation.done = true
	return affected, err
}

// CastMemberDeleteOne is the builder for deleting a single CastMember entity.
type CastMemberDeleteOne struct {
	cmd *CastMemberDelete
}

// Where appends a list predicates to the CastMemberDelete builder.
func (cmdo *CastMemberDeleteOne) Where(ps ...predicate.CastMember) *CastMemberDeleteOne {
	cmdo.cmd.mutation.Where(ps...)
	return cmdo
}

// Exec executes the deletion query.
func (cmdo *CastMemberDeleteOne) Exec(ctx context.Context) error {
	n, err := cmdo.cmd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{castmember.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cmdo *CastMemberDeleteOne) ExecX(ctx context.Context) {
	if err := cmdo.Exec(ctx); err != nil {
		panic(err)
	}
}
