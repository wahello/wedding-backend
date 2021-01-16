// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"wedding/ent/inviteeparty"

	"github.com/facebook/ent/dialect/sql"
)

// InviteeParty is the model entity for the InviteeParty schema.
type InviteeParty struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the InviteePartyQuery when eager-loading is set.
	Edges InviteePartyEdges `json:"edges"`
}

// InviteePartyEdges holds the relations/edges for other nodes in the graph.
type InviteePartyEdges struct {
	// Invitees holds the value of the invitees edge.
	Invitees []*Invitee
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// InviteesOrErr returns the Invitees value or an error if the edge
// was not loaded in eager-loading.
func (e InviteePartyEdges) InviteesOrErr() ([]*Invitee, error) {
	if e.loadedTypes[0] {
		return e.Invitees, nil
	}
	return nil, &NotLoadedError{edge: "invitees"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*InviteeParty) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case inviteeparty.FieldID:
			values[i] = &sql.NullInt64{}
		case inviteeparty.FieldName, inviteeparty.FieldCode:
			values[i] = &sql.NullString{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type InviteeParty", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the InviteeParty fields.
func (ip *InviteeParty) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case inviteeparty.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ip.ID = int(value.Int64)
		case inviteeparty.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ip.Name = value.String
			}
		case inviteeparty.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				ip.Code = value.String
			}
		}
	}
	return nil
}

// QueryInvitees queries the "invitees" edge of the InviteeParty entity.
func (ip *InviteeParty) QueryInvitees() *InviteeQuery {
	return (&InviteePartyClient{config: ip.config}).QueryInvitees(ip)
}

// Update returns a builder for updating this InviteeParty.
// Note that you need to call InviteeParty.Unwrap() before calling this method if this InviteeParty
// was returned from a transaction, and the transaction was committed or rolled back.
func (ip *InviteeParty) Update() *InviteePartyUpdateOne {
	return (&InviteePartyClient{config: ip.config}).UpdateOne(ip)
}

// Unwrap unwraps the InviteeParty entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ip *InviteeParty) Unwrap() *InviteeParty {
	tx, ok := ip.config.driver.(*txDriver)
	if !ok {
		panic("ent: InviteeParty is not a transactional entity")
	}
	ip.config.driver = tx.drv
	return ip
}

// String implements the fmt.Stringer.
func (ip *InviteeParty) String() string {
	var builder strings.Builder
	builder.WriteString("InviteeParty(")
	builder.WriteString(fmt.Sprintf("id=%v", ip.ID))
	builder.WriteString(", name=")
	builder.WriteString(ip.Name)
	builder.WriteString(", code=")
	builder.WriteString(ip.Code)
	builder.WriteByte(')')
	return builder.String()
}

// InviteeParties is a parsable slice of InviteeParty.
type InviteeParties []*InviteeParty

func (ip InviteeParties) config(cfg config) {
	for _i := range ip {
		ip[_i].config = cfg
	}
}
