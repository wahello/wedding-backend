// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"wedding/ent/invitee"
	"wedding/ent/predicate"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// InviteeDelete is the builder for deleting a Invitee entity.
type InviteeDelete struct {
	config
	hooks    []Hook
	mutation *InviteeMutation
}

// Where adds a new predicate to the InviteeDelete builder.
func (id *InviteeDelete) Where(ps ...predicate.Invitee) *InviteeDelete {
	id.mutation.predicates = append(id.mutation.predicates, ps...)
	return id
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (id *InviteeDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(id.hooks) == 0 {
		affected, err = id.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InviteeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			id.mutation = mutation
			affected, err = id.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(id.hooks) - 1; i >= 0; i-- {
			mut = id.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, id.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (id *InviteeDelete) ExecX(ctx context.Context) int {
	n, err := id.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (id *InviteeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: invitee.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: invitee.FieldID,
			},
		},
	}
	if ps := id.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, id.driver, _spec)
}

// InviteeDeleteOne is the builder for deleting a single Invitee entity.
type InviteeDeleteOne struct {
	id *InviteeDelete
}

// Exec executes the deletion query.
func (ido *InviteeDeleteOne) Exec(ctx context.Context) error {
	n, err := ido.id.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{invitee.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ido *InviteeDeleteOne) ExecX(ctx context.Context) {
	ido.id.ExecX(ctx)
}
