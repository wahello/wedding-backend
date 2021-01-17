// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"wedding/ent/backroomuser"
	"wedding/ent/predicate"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// BackroomUserUpdate is the builder for updating BackroomUser entities.
type BackroomUserUpdate struct {
	config
	hooks    []Hook
	mutation *BackroomUserMutation
}

// Where adds a new predicate for the BackroomUserUpdate builder.
func (buu *BackroomUserUpdate) Where(ps ...predicate.BackroomUser) *BackroomUserUpdate {
	buu.mutation.predicates = append(buu.mutation.predicates, ps...)
	return buu
}

// SetUsername sets the "username" field.
func (buu *BackroomUserUpdate) SetUsername(s string) *BackroomUserUpdate {
	buu.mutation.SetUsername(s)
	return buu
}

// SetPassword sets the "password" field.
func (buu *BackroomUserUpdate) SetPassword(s string) *BackroomUserUpdate {
	buu.mutation.SetPassword(s)
	return buu
}

// Mutation returns the BackroomUserMutation object of the builder.
func (buu *BackroomUserUpdate) Mutation() *BackroomUserMutation {
	return buu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (buu *BackroomUserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(buu.hooks) == 0 {
		if err = buu.check(); err != nil {
			return 0, err
		}
		affected, err = buu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BackroomUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = buu.check(); err != nil {
				return 0, err
			}
			buu.mutation = mutation
			affected, err = buu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(buu.hooks) - 1; i >= 0; i-- {
			mut = buu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, buu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (buu *BackroomUserUpdate) SaveX(ctx context.Context) int {
	affected, err := buu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (buu *BackroomUserUpdate) Exec(ctx context.Context) error {
	_, err := buu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buu *BackroomUserUpdate) ExecX(ctx context.Context) {
	if err := buu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (buu *BackroomUserUpdate) check() error {
	if v, ok := buu.mutation.Username(); ok {
		if err := backroomuser.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf("ent: validator failed for field \"username\": %w", err)}
		}
	}
	if v, ok := buu.mutation.Password(); ok {
		if err := backroomuser.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf("ent: validator failed for field \"password\": %w", err)}
		}
	}
	return nil
}

func (buu *BackroomUserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   backroomuser.Table,
			Columns: backroomuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: backroomuser.FieldID,
			},
		},
	}
	if ps := buu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buu.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: backroomuser.FieldUsername,
		})
	}
	if value, ok := buu.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: backroomuser.FieldPassword,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, buu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{backroomuser.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// BackroomUserUpdateOne is the builder for updating a single BackroomUser entity.
type BackroomUserUpdateOne struct {
	config
	hooks    []Hook
	mutation *BackroomUserMutation
}

// SetUsername sets the "username" field.
func (buuo *BackroomUserUpdateOne) SetUsername(s string) *BackroomUserUpdateOne {
	buuo.mutation.SetUsername(s)
	return buuo
}

// SetPassword sets the "password" field.
func (buuo *BackroomUserUpdateOne) SetPassword(s string) *BackroomUserUpdateOne {
	buuo.mutation.SetPassword(s)
	return buuo
}

// Mutation returns the BackroomUserMutation object of the builder.
func (buuo *BackroomUserUpdateOne) Mutation() *BackroomUserMutation {
	return buuo.mutation
}

// Save executes the query and returns the updated BackroomUser entity.
func (buuo *BackroomUserUpdateOne) Save(ctx context.Context) (*BackroomUser, error) {
	var (
		err  error
		node *BackroomUser
	)
	if len(buuo.hooks) == 0 {
		if err = buuo.check(); err != nil {
			return nil, err
		}
		node, err = buuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BackroomUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = buuo.check(); err != nil {
				return nil, err
			}
			buuo.mutation = mutation
			node, err = buuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(buuo.hooks) - 1; i >= 0; i-- {
			mut = buuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, buuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (buuo *BackroomUserUpdateOne) SaveX(ctx context.Context) *BackroomUser {
	node, err := buuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buuo *BackroomUserUpdateOne) Exec(ctx context.Context) error {
	_, err := buuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buuo *BackroomUserUpdateOne) ExecX(ctx context.Context) {
	if err := buuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (buuo *BackroomUserUpdateOne) check() error {
	if v, ok := buuo.mutation.Username(); ok {
		if err := backroomuser.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf("ent: validator failed for field \"username\": %w", err)}
		}
	}
	if v, ok := buuo.mutation.Password(); ok {
		if err := backroomuser.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf("ent: validator failed for field \"password\": %w", err)}
		}
	}
	return nil
}

func (buuo *BackroomUserUpdateOne) sqlSave(ctx context.Context) (_node *BackroomUser, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   backroomuser.Table,
			Columns: backroomuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: backroomuser.FieldID,
			},
		},
	}
	id, ok := buuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing BackroomUser.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := buuo.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: backroomuser.FieldUsername,
		})
	}
	if value, ok := buuo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: backroomuser.FieldPassword,
		})
	}
	_node = &BackroomUser{config: buuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{backroomuser.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
