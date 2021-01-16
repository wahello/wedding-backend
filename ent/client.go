// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"wedding/ent/migrate"

	"wedding/ent/invitee"
	"wedding/ent/inviteeparty"

	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Invitee is the client for interacting with the Invitee builders.
	Invitee *InviteeClient
	// InviteeParty is the client for interacting with the InviteeParty builders.
	InviteeParty *InviteePartyClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Invitee = NewInviteeClient(c.config)
	c.InviteeParty = NewInviteePartyClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		Invitee:      NewInviteeClient(cfg),
		InviteeParty: NewInviteePartyClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config:       cfg,
		Invitee:      NewInviteeClient(cfg),
		InviteeParty: NewInviteePartyClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Invitee.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Invitee.Use(hooks...)
	c.InviteeParty.Use(hooks...)
}

// InviteeClient is a client for the Invitee schema.
type InviteeClient struct {
	config
}

// NewInviteeClient returns a client for the Invitee from the given config.
func NewInviteeClient(c config) *InviteeClient {
	return &InviteeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `invitee.Hooks(f(g(h())))`.
func (c *InviteeClient) Use(hooks ...Hook) {
	c.hooks.Invitee = append(c.hooks.Invitee, hooks...)
}

// Create returns a create builder for Invitee.
func (c *InviteeClient) Create() *InviteeCreate {
	mutation := newInviteeMutation(c.config, OpCreate)
	return &InviteeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Invitee entities.
func (c *InviteeClient) CreateBulk(builders ...*InviteeCreate) *InviteeCreateBulk {
	return &InviteeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Invitee.
func (c *InviteeClient) Update() *InviteeUpdate {
	mutation := newInviteeMutation(c.config, OpUpdate)
	return &InviteeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *InviteeClient) UpdateOne(i *Invitee) *InviteeUpdateOne {
	mutation := newInviteeMutation(c.config, OpUpdateOne, withInvitee(i))
	return &InviteeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *InviteeClient) UpdateOneID(id int) *InviteeUpdateOne {
	mutation := newInviteeMutation(c.config, OpUpdateOne, withInviteeID(id))
	return &InviteeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Invitee.
func (c *InviteeClient) Delete() *InviteeDelete {
	mutation := newInviteeMutation(c.config, OpDelete)
	return &InviteeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *InviteeClient) DeleteOne(i *Invitee) *InviteeDeleteOne {
	return c.DeleteOneID(i.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *InviteeClient) DeleteOneID(id int) *InviteeDeleteOne {
	builder := c.Delete().Where(invitee.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &InviteeDeleteOne{builder}
}

// Query returns a query builder for Invitee.
func (c *InviteeClient) Query() *InviteeQuery {
	return &InviteeQuery{config: c.config}
}

// Get returns a Invitee entity by its id.
func (c *InviteeClient) Get(ctx context.Context, id int) (*Invitee, error) {
	return c.Query().Where(invitee.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *InviteeClient) GetX(ctx context.Context, id int) *Invitee {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParty queries the party edge of a Invitee.
func (c *InviteeClient) QueryParty(i *Invitee) *InviteePartyQuery {
	query := &InviteePartyQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(invitee.Table, invitee.FieldID, id),
			sqlgraph.To(inviteeparty.Table, inviteeparty.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, invitee.PartyTable, invitee.PartyColumn),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *InviteeClient) Hooks() []Hook {
	return c.hooks.Invitee
}

// InviteePartyClient is a client for the InviteeParty schema.
type InviteePartyClient struct {
	config
}

// NewInviteePartyClient returns a client for the InviteeParty from the given config.
func NewInviteePartyClient(c config) *InviteePartyClient {
	return &InviteePartyClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `inviteeparty.Hooks(f(g(h())))`.
func (c *InviteePartyClient) Use(hooks ...Hook) {
	c.hooks.InviteeParty = append(c.hooks.InviteeParty, hooks...)
}

// Create returns a create builder for InviteeParty.
func (c *InviteePartyClient) Create() *InviteePartyCreate {
	mutation := newInviteePartyMutation(c.config, OpCreate)
	return &InviteePartyCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of InviteeParty entities.
func (c *InviteePartyClient) CreateBulk(builders ...*InviteePartyCreate) *InviteePartyCreateBulk {
	return &InviteePartyCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for InviteeParty.
func (c *InviteePartyClient) Update() *InviteePartyUpdate {
	mutation := newInviteePartyMutation(c.config, OpUpdate)
	return &InviteePartyUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *InviteePartyClient) UpdateOne(ip *InviteeParty) *InviteePartyUpdateOne {
	mutation := newInviteePartyMutation(c.config, OpUpdateOne, withInviteeParty(ip))
	return &InviteePartyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *InviteePartyClient) UpdateOneID(id int) *InviteePartyUpdateOne {
	mutation := newInviteePartyMutation(c.config, OpUpdateOne, withInviteePartyID(id))
	return &InviteePartyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for InviteeParty.
func (c *InviteePartyClient) Delete() *InviteePartyDelete {
	mutation := newInviteePartyMutation(c.config, OpDelete)
	return &InviteePartyDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *InviteePartyClient) DeleteOne(ip *InviteeParty) *InviteePartyDeleteOne {
	return c.DeleteOneID(ip.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *InviteePartyClient) DeleteOneID(id int) *InviteePartyDeleteOne {
	builder := c.Delete().Where(inviteeparty.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &InviteePartyDeleteOne{builder}
}

// Query returns a query builder for InviteeParty.
func (c *InviteePartyClient) Query() *InviteePartyQuery {
	return &InviteePartyQuery{config: c.config}
}

// Get returns a InviteeParty entity by its id.
func (c *InviteePartyClient) Get(ctx context.Context, id int) (*InviteeParty, error) {
	return c.Query().Where(inviteeparty.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *InviteePartyClient) GetX(ctx context.Context, id int) *InviteeParty {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryInvitees queries the invitees edge of a InviteeParty.
func (c *InviteePartyClient) QueryInvitees(ip *InviteeParty) *InviteeQuery {
	query := &InviteeQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ip.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(inviteeparty.Table, inviteeparty.FieldID, id),
			sqlgraph.To(invitee.Table, invitee.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, inviteeparty.InviteesTable, inviteeparty.InviteesColumn),
		)
		fromV = sqlgraph.Neighbors(ip.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *InviteePartyClient) Hooks() []Hook {
	return c.hooks.InviteeParty
}
