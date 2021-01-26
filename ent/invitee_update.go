// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"wedding/ent/invitee"
	"wedding/ent/inviteeparty"
	"wedding/ent/predicate"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// InviteeUpdate is the builder for updating Invitee entities.
type InviteeUpdate struct {
	config
	hooks    []Hook
	mutation *InviteeMutation
}

// Where adds a new predicate for the InviteeUpdate builder.
func (iu *InviteeUpdate) Where(ps ...predicate.Invitee) *InviteeUpdate {
	iu.mutation.predicates = append(iu.mutation.predicates, ps...)
	return iu
}

// SetName sets the "name" field.
func (iu *InviteeUpdate) SetName(s string) *InviteeUpdate {
	iu.mutation.SetName(s)
	return iu
}

// SetIsChild sets the "is_child" field.
func (iu *InviteeUpdate) SetIsChild(b bool) *InviteeUpdate {
	iu.mutation.SetIsChild(b)
	return iu
}

// SetNillableIsChild sets the "is_child" field if the given value is not nil.
func (iu *InviteeUpdate) SetNillableIsChild(b *bool) *InviteeUpdate {
	if b != nil {
		iu.SetIsChild(*b)
	}
	return iu
}

// ClearIsChild clears the value of the "is_child" field.
func (iu *InviteeUpdate) ClearIsChild() *InviteeUpdate {
	iu.mutation.ClearIsChild()
	return iu
}

// SetHasPlusOne sets the "has_plus_one" field.
func (iu *InviteeUpdate) SetHasPlusOne(b bool) *InviteeUpdate {
	iu.mutation.SetHasPlusOne(b)
	return iu
}

// SetNillableHasPlusOne sets the "has_plus_one" field if the given value is not nil.
func (iu *InviteeUpdate) SetNillableHasPlusOne(b *bool) *InviteeUpdate {
	if b != nil {
		iu.SetHasPlusOne(*b)
	}
	return iu
}

// SetPlusOneName sets the "plus_one_name" field.
func (iu *InviteeUpdate) SetPlusOneName(s string) *InviteeUpdate {
	iu.mutation.SetPlusOneName(s)
	return iu
}

// SetNillablePlusOneName sets the "plus_one_name" field if the given value is not nil.
func (iu *InviteeUpdate) SetNillablePlusOneName(s *string) *InviteeUpdate {
	if s != nil {
		iu.SetPlusOneName(*s)
	}
	return iu
}

// ClearPlusOneName clears the value of the "plus_one_name" field.
func (iu *InviteeUpdate) ClearPlusOneName() *InviteeUpdate {
	iu.mutation.ClearPlusOneName()
	return iu
}

// SetPhone sets the "phone" field.
func (iu *InviteeUpdate) SetPhone(s string) *InviteeUpdate {
	iu.mutation.SetPhone(s)
	return iu
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (iu *InviteeUpdate) SetNillablePhone(s *string) *InviteeUpdate {
	if s != nil {
		iu.SetPhone(*s)
	}
	return iu
}

// ClearPhone clears the value of the "phone" field.
func (iu *InviteeUpdate) ClearPhone() *InviteeUpdate {
	iu.mutation.ClearPhone()
	return iu
}

// SetEmail sets the "email" field.
func (iu *InviteeUpdate) SetEmail(s string) *InviteeUpdate {
	iu.mutation.SetEmail(s)
	return iu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (iu *InviteeUpdate) SetNillableEmail(s *string) *InviteeUpdate {
	if s != nil {
		iu.SetEmail(*s)
	}
	return iu
}

// ClearEmail clears the value of the "email" field.
func (iu *InviteeUpdate) ClearEmail() *InviteeUpdate {
	iu.mutation.ClearEmail()
	return iu
}

// SetAddressLine1 sets the "address_line_1" field.
func (iu *InviteeUpdate) SetAddressLine1(s string) *InviteeUpdate {
	iu.mutation.SetAddressLine1(s)
	return iu
}

// SetNillableAddressLine1 sets the "address_line_1" field if the given value is not nil.
func (iu *InviteeUpdate) SetNillableAddressLine1(s *string) *InviteeUpdate {
	if s != nil {
		iu.SetAddressLine1(*s)
	}
	return iu
}

// ClearAddressLine1 clears the value of the "address_line_1" field.
func (iu *InviteeUpdate) ClearAddressLine1() *InviteeUpdate {
	iu.mutation.ClearAddressLine1()
	return iu
}

// SetAddressLine2 sets the "address_line_2" field.
func (iu *InviteeUpdate) SetAddressLine2(s string) *InviteeUpdate {
	iu.mutation.SetAddressLine2(s)
	return iu
}

// SetNillableAddressLine2 sets the "address_line_2" field if the given value is not nil.
func (iu *InviteeUpdate) SetNillableAddressLine2(s *string) *InviteeUpdate {
	if s != nil {
		iu.SetAddressLine2(*s)
	}
	return iu
}

// ClearAddressLine2 clears the value of the "address_line_2" field.
func (iu *InviteeUpdate) ClearAddressLine2() *InviteeUpdate {
	iu.mutation.ClearAddressLine2()
	return iu
}

// SetAddressCity sets the "address_city" field.
func (iu *InviteeUpdate) SetAddressCity(s string) *InviteeUpdate {
	iu.mutation.SetAddressCity(s)
	return iu
}

// SetNillableAddressCity sets the "address_city" field if the given value is not nil.
func (iu *InviteeUpdate) SetNillableAddressCity(s *string) *InviteeUpdate {
	if s != nil {
		iu.SetAddressCity(*s)
	}
	return iu
}

// ClearAddressCity clears the value of the "address_city" field.
func (iu *InviteeUpdate) ClearAddressCity() *InviteeUpdate {
	iu.mutation.ClearAddressCity()
	return iu
}

// SetAddressState sets the "address_state" field.
func (iu *InviteeUpdate) SetAddressState(s string) *InviteeUpdate {
	iu.mutation.SetAddressState(s)
	return iu
}

// SetNillableAddressState sets the "address_state" field if the given value is not nil.
func (iu *InviteeUpdate) SetNillableAddressState(s *string) *InviteeUpdate {
	if s != nil {
		iu.SetAddressState(*s)
	}
	return iu
}

// ClearAddressState clears the value of the "address_state" field.
func (iu *InviteeUpdate) ClearAddressState() *InviteeUpdate {
	iu.mutation.ClearAddressState()
	return iu
}

// SetAddressPostalCode sets the "address_postal_code" field.
func (iu *InviteeUpdate) SetAddressPostalCode(s string) *InviteeUpdate {
	iu.mutation.SetAddressPostalCode(s)
	return iu
}

// SetNillableAddressPostalCode sets the "address_postal_code" field if the given value is not nil.
func (iu *InviteeUpdate) SetNillableAddressPostalCode(s *string) *InviteeUpdate {
	if s != nil {
		iu.SetAddressPostalCode(*s)
	}
	return iu
}

// ClearAddressPostalCode clears the value of the "address_postal_code" field.
func (iu *InviteeUpdate) ClearAddressPostalCode() *InviteeUpdate {
	iu.mutation.ClearAddressPostalCode()
	return iu
}

// SetAddressCountry sets the "address_country" field.
func (iu *InviteeUpdate) SetAddressCountry(s string) *InviteeUpdate {
	iu.mutation.SetAddressCountry(s)
	return iu
}

// SetNillableAddressCountry sets the "address_country" field if the given value is not nil.
func (iu *InviteeUpdate) SetNillableAddressCountry(s *string) *InviteeUpdate {
	if s != nil {
		iu.SetAddressCountry(*s)
	}
	return iu
}

// ClearAddressCountry clears the value of the "address_country" field.
func (iu *InviteeUpdate) ClearAddressCountry() *InviteeUpdate {
	iu.mutation.ClearAddressCountry()
	return iu
}

// SetPartyID sets the "party" edge to the InviteeParty entity by ID.
func (iu *InviteeUpdate) SetPartyID(id int) *InviteeUpdate {
	iu.mutation.SetPartyID(id)
	return iu
}

// SetNillablePartyID sets the "party" edge to the InviteeParty entity by ID if the given value is not nil.
func (iu *InviteeUpdate) SetNillablePartyID(id *int) *InviteeUpdate {
	if id != nil {
		iu = iu.SetPartyID(*id)
	}
	return iu
}

// SetParty sets the "party" edge to the InviteeParty entity.
func (iu *InviteeUpdate) SetParty(i *InviteeParty) *InviteeUpdate {
	return iu.SetPartyID(i.ID)
}

// Mutation returns the InviteeMutation object of the builder.
func (iu *InviteeUpdate) Mutation() *InviteeMutation {
	return iu.mutation
}

// ClearParty clears the "party" edge to the InviteeParty entity.
func (iu *InviteeUpdate) ClearParty() *InviteeUpdate {
	iu.mutation.ClearParty()
	return iu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *InviteeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(iu.hooks) == 0 {
		if err = iu.check(); err != nil {
			return 0, err
		}
		affected, err = iu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InviteeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = iu.check(); err != nil {
				return 0, err
			}
			iu.mutation = mutation
			affected, err = iu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(iu.hooks) - 1; i >= 0; i-- {
			mut = iu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (iu *InviteeUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *InviteeUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *InviteeUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iu *InviteeUpdate) check() error {
	if v, ok := iu.mutation.Name(); ok {
		if err := invitee.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	return nil
}

func (iu *InviteeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   invitee.Table,
			Columns: invitee.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: invitee.FieldID,
			},
		},
	}
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: invitee.FieldName,
		})
	}
	if value, ok := iu.mutation.IsChild(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: invitee.FieldIsChild,
		})
	}
	if iu.mutation.IsChildCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: invitee.FieldIsChild,
		})
	}
	if value, ok := iu.mutation.HasPlusOne(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: invitee.FieldHasPlusOne,
		})
	}
	if value, ok := iu.mutation.PlusOneName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: invitee.FieldPlusOneName,
		})
	}
	if iu.mutation.PlusOneNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: invitee.FieldPlusOneName,
		})
	}
	if value, ok := iu.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: invitee.FieldPhone,
		})
	}
	if iu.mutation.PhoneCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: invitee.FieldPhone,
		})
	}
	if value, ok := iu.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: invitee.FieldEmail,
		})
	}
	if iu.mutation.EmailCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: invitee.FieldEmail,
		})
	}
	if value, ok := iu.mutation.AddressLine1(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: invitee.FieldAddressLine1,
		})
	}
	if iu.mutation.AddressLine1Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: invitee.FieldAddressLine1,
		})
	}
	if value, ok := iu.mutation.AddressLine2(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: invitee.FieldAddressLine2,
		})
	}
	if iu.mutation.AddressLine2Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: invitee.FieldAddressLine2,
		})
	}
	if value, ok := iu.mutation.AddressCity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: invitee.FieldAddressCity,
		})
	}
	if iu.mutation.AddressCityCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: invitee.FieldAddressCity,
		})
	}
	if value, ok := iu.mutation.AddressState(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: invitee.FieldAddressState,
		})
	}
	if iu.mutation.AddressStateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: invitee.FieldAddressState,
		})
	}
	if value, ok := iu.mutation.AddressPostalCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: invitee.FieldAddressPostalCode,
		})
	}
	if iu.mutation.AddressPostalCodeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: invitee.FieldAddressPostalCode,
		})
	}
	if value, ok := iu.mutation.AddressCountry(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: invitee.FieldAddressCountry,
		})
	}
	if iu.mutation.AddressCountryCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: invitee.FieldAddressCountry,
		})
	}
	if iu.mutation.PartyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   invitee.PartyTable,
			Columns: []string{invitee.PartyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: inviteeparty.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.PartyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   invitee.PartyTable,
			Columns: []string{invitee.PartyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: inviteeparty.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{invitee.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// InviteeUpdateOne is the builder for updating a single Invitee entity.
type InviteeUpdateOne struct {
	config
	hooks    []Hook
	mutation *InviteeMutation
}

// SetName sets the "name" field.
func (iuo *InviteeUpdateOne) SetName(s string) *InviteeUpdateOne {
	iuo.mutation.SetName(s)
	return iuo
}

// SetIsChild sets the "is_child" field.
func (iuo *InviteeUpdateOne) SetIsChild(b bool) *InviteeUpdateOne {
	iuo.mutation.SetIsChild(b)
	return iuo
}

// SetNillableIsChild sets the "is_child" field if the given value is not nil.
func (iuo *InviteeUpdateOne) SetNillableIsChild(b *bool) *InviteeUpdateOne {
	if b != nil {
		iuo.SetIsChild(*b)
	}
	return iuo
}

// ClearIsChild clears the value of the "is_child" field.
func (iuo *InviteeUpdateOne) ClearIsChild() *InviteeUpdateOne {
	iuo.mutation.ClearIsChild()
	return iuo
}

// SetHasPlusOne sets the "has_plus_one" field.
func (iuo *InviteeUpdateOne) SetHasPlusOne(b bool) *InviteeUpdateOne {
	iuo.mutation.SetHasPlusOne(b)
	return iuo
}

// SetNillableHasPlusOne sets the "has_plus_one" field if the given value is not nil.
func (iuo *InviteeUpdateOne) SetNillableHasPlusOne(b *bool) *InviteeUpdateOne {
	if b != nil {
		iuo.SetHasPlusOne(*b)
	}
	return iuo
}

// SetPlusOneName sets the "plus_one_name" field.
func (iuo *InviteeUpdateOne) SetPlusOneName(s string) *InviteeUpdateOne {
	iuo.mutation.SetPlusOneName(s)
	return iuo
}

// SetNillablePlusOneName sets the "plus_one_name" field if the given value is not nil.
func (iuo *InviteeUpdateOne) SetNillablePlusOneName(s *string) *InviteeUpdateOne {
	if s != nil {
		iuo.SetPlusOneName(*s)
	}
	return iuo
}

// ClearPlusOneName clears the value of the "plus_one_name" field.
func (iuo *InviteeUpdateOne) ClearPlusOneName() *InviteeUpdateOne {
	iuo.mutation.ClearPlusOneName()
	return iuo
}

// SetPhone sets the "phone" field.
func (iuo *InviteeUpdateOne) SetPhone(s string) *InviteeUpdateOne {
	iuo.mutation.SetPhone(s)
	return iuo
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (iuo *InviteeUpdateOne) SetNillablePhone(s *string) *InviteeUpdateOne {
	if s != nil {
		iuo.SetPhone(*s)
	}
	return iuo
}

// ClearPhone clears the value of the "phone" field.
func (iuo *InviteeUpdateOne) ClearPhone() *InviteeUpdateOne {
	iuo.mutation.ClearPhone()
	return iuo
}

// SetEmail sets the "email" field.
func (iuo *InviteeUpdateOne) SetEmail(s string) *InviteeUpdateOne {
	iuo.mutation.SetEmail(s)
	return iuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (iuo *InviteeUpdateOne) SetNillableEmail(s *string) *InviteeUpdateOne {
	if s != nil {
		iuo.SetEmail(*s)
	}
	return iuo
}

// ClearEmail clears the value of the "email" field.
func (iuo *InviteeUpdateOne) ClearEmail() *InviteeUpdateOne {
	iuo.mutation.ClearEmail()
	return iuo
}

// SetAddressLine1 sets the "address_line_1" field.
func (iuo *InviteeUpdateOne) SetAddressLine1(s string) *InviteeUpdateOne {
	iuo.mutation.SetAddressLine1(s)
	return iuo
}

// SetNillableAddressLine1 sets the "address_line_1" field if the given value is not nil.
func (iuo *InviteeUpdateOne) SetNillableAddressLine1(s *string) *InviteeUpdateOne {
	if s != nil {
		iuo.SetAddressLine1(*s)
	}
	return iuo
}

// ClearAddressLine1 clears the value of the "address_line_1" field.
func (iuo *InviteeUpdateOne) ClearAddressLine1() *InviteeUpdateOne {
	iuo.mutation.ClearAddressLine1()
	return iuo
}

// SetAddressLine2 sets the "address_line_2" field.
func (iuo *InviteeUpdateOne) SetAddressLine2(s string) *InviteeUpdateOne {
	iuo.mutation.SetAddressLine2(s)
	return iuo
}

// SetNillableAddressLine2 sets the "address_line_2" field if the given value is not nil.
func (iuo *InviteeUpdateOne) SetNillableAddressLine2(s *string) *InviteeUpdateOne {
	if s != nil {
		iuo.SetAddressLine2(*s)
	}
	return iuo
}

// ClearAddressLine2 clears the value of the "address_line_2" field.
func (iuo *InviteeUpdateOne) ClearAddressLine2() *InviteeUpdateOne {
	iuo.mutation.ClearAddressLine2()
	return iuo
}

// SetAddressCity sets the "address_city" field.
func (iuo *InviteeUpdateOne) SetAddressCity(s string) *InviteeUpdateOne {
	iuo.mutation.SetAddressCity(s)
	return iuo
}

// SetNillableAddressCity sets the "address_city" field if the given value is not nil.
func (iuo *InviteeUpdateOne) SetNillableAddressCity(s *string) *InviteeUpdateOne {
	if s != nil {
		iuo.SetAddressCity(*s)
	}
	return iuo
}

// ClearAddressCity clears the value of the "address_city" field.
func (iuo *InviteeUpdateOne) ClearAddressCity() *InviteeUpdateOne {
	iuo.mutation.ClearAddressCity()
	return iuo
}

// SetAddressState sets the "address_state" field.
func (iuo *InviteeUpdateOne) SetAddressState(s string) *InviteeUpdateOne {
	iuo.mutation.SetAddressState(s)
	return iuo
}

// SetNillableAddressState sets the "address_state" field if the given value is not nil.
func (iuo *InviteeUpdateOne) SetNillableAddressState(s *string) *InviteeUpdateOne {
	if s != nil {
		iuo.SetAddressState(*s)
	}
	return iuo
}

// ClearAddressState clears the value of the "address_state" field.
func (iuo *InviteeUpdateOne) ClearAddressState() *InviteeUpdateOne {
	iuo.mutation.ClearAddressState()
	return iuo
}

// SetAddressPostalCode sets the "address_postal_code" field.
func (iuo *InviteeUpdateOne) SetAddressPostalCode(s string) *InviteeUpdateOne {
	iuo.mutation.SetAddressPostalCode(s)
	return iuo
}

// SetNillableAddressPostalCode sets the "address_postal_code" field if the given value is not nil.
func (iuo *InviteeUpdateOne) SetNillableAddressPostalCode(s *string) *InviteeUpdateOne {
	if s != nil {
		iuo.SetAddressPostalCode(*s)
	}
	return iuo
}

// ClearAddressPostalCode clears the value of the "address_postal_code" field.
func (iuo *InviteeUpdateOne) ClearAddressPostalCode() *InviteeUpdateOne {
	iuo.mutation.ClearAddressPostalCode()
	return iuo
}

// SetAddressCountry sets the "address_country" field.
func (iuo *InviteeUpdateOne) SetAddressCountry(s string) *InviteeUpdateOne {
	iuo.mutation.SetAddressCountry(s)
	return iuo
}

// SetNillableAddressCountry sets the "address_country" field if the given value is not nil.
func (iuo *InviteeUpdateOne) SetNillableAddressCountry(s *string) *InviteeUpdateOne {
	if s != nil {
		iuo.SetAddressCountry(*s)
	}
	return iuo
}

// ClearAddressCountry clears the value of the "address_country" field.
func (iuo *InviteeUpdateOne) ClearAddressCountry() *InviteeUpdateOne {
	iuo.mutation.ClearAddressCountry()
	return iuo
}

// SetPartyID sets the "party" edge to the InviteeParty entity by ID.
func (iuo *InviteeUpdateOne) SetPartyID(id int) *InviteeUpdateOne {
	iuo.mutation.SetPartyID(id)
	return iuo
}

// SetNillablePartyID sets the "party" edge to the InviteeParty entity by ID if the given value is not nil.
func (iuo *InviteeUpdateOne) SetNillablePartyID(id *int) *InviteeUpdateOne {
	if id != nil {
		iuo = iuo.SetPartyID(*id)
	}
	return iuo
}

// SetParty sets the "party" edge to the InviteeParty entity.
func (iuo *InviteeUpdateOne) SetParty(i *InviteeParty) *InviteeUpdateOne {
	return iuo.SetPartyID(i.ID)
}

// Mutation returns the InviteeMutation object of the builder.
func (iuo *InviteeUpdateOne) Mutation() *InviteeMutation {
	return iuo.mutation
}

// ClearParty clears the "party" edge to the InviteeParty entity.
func (iuo *InviteeUpdateOne) ClearParty() *InviteeUpdateOne {
	iuo.mutation.ClearParty()
	return iuo
}

// Save executes the query and returns the updated Invitee entity.
func (iuo *InviteeUpdateOne) Save(ctx context.Context) (*Invitee, error) {
	var (
		err  error
		node *Invitee
	)
	if len(iuo.hooks) == 0 {
		if err = iuo.check(); err != nil {
			return nil, err
		}
		node, err = iuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InviteeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = iuo.check(); err != nil {
				return nil, err
			}
			iuo.mutation = mutation
			node, err = iuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(iuo.hooks) - 1; i >= 0; i-- {
			mut = iuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *InviteeUpdateOne) SaveX(ctx context.Context) *Invitee {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *InviteeUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *InviteeUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iuo *InviteeUpdateOne) check() error {
	if v, ok := iuo.mutation.Name(); ok {
		if err := invitee.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	return nil
}

func (iuo *InviteeUpdateOne) sqlSave(ctx context.Context) (_node *Invitee, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   invitee.Table,
			Columns: invitee.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: invitee.FieldID,
			},
		},
	}
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Invitee.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := iuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: invitee.FieldName,
		})
	}
	if value, ok := iuo.mutation.IsChild(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: invitee.FieldIsChild,
		})
	}
	if iuo.mutation.IsChildCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: invitee.FieldIsChild,
		})
	}
	if value, ok := iuo.mutation.HasPlusOne(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: invitee.FieldHasPlusOne,
		})
	}
	if value, ok := iuo.mutation.PlusOneName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: invitee.FieldPlusOneName,
		})
	}
	if iuo.mutation.PlusOneNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: invitee.FieldPlusOneName,
		})
	}
	if value, ok := iuo.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: invitee.FieldPhone,
		})
	}
	if iuo.mutation.PhoneCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: invitee.FieldPhone,
		})
	}
	if value, ok := iuo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: invitee.FieldEmail,
		})
	}
	if iuo.mutation.EmailCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: invitee.FieldEmail,
		})
	}
	if value, ok := iuo.mutation.AddressLine1(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: invitee.FieldAddressLine1,
		})
	}
	if iuo.mutation.AddressLine1Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: invitee.FieldAddressLine1,
		})
	}
	if value, ok := iuo.mutation.AddressLine2(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: invitee.FieldAddressLine2,
		})
	}
	if iuo.mutation.AddressLine2Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: invitee.FieldAddressLine2,
		})
	}
	if value, ok := iuo.mutation.AddressCity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: invitee.FieldAddressCity,
		})
	}
	if iuo.mutation.AddressCityCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: invitee.FieldAddressCity,
		})
	}
	if value, ok := iuo.mutation.AddressState(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: invitee.FieldAddressState,
		})
	}
	if iuo.mutation.AddressStateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: invitee.FieldAddressState,
		})
	}
	if value, ok := iuo.mutation.AddressPostalCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: invitee.FieldAddressPostalCode,
		})
	}
	if iuo.mutation.AddressPostalCodeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: invitee.FieldAddressPostalCode,
		})
	}
	if value, ok := iuo.mutation.AddressCountry(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: invitee.FieldAddressCountry,
		})
	}
	if iuo.mutation.AddressCountryCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: invitee.FieldAddressCountry,
		})
	}
	if iuo.mutation.PartyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   invitee.PartyTable,
			Columns: []string{invitee.PartyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: inviteeparty.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.PartyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   invitee.PartyTable,
			Columns: []string{invitee.PartyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: inviteeparty.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Invitee{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{invitee.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
