// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/nuomizi-fw/stargazer/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/nuomizi-fw/stargazer/ent/bangumi"
	"github.com/nuomizi-fw/stargazer/ent/castmember"
	"github.com/nuomizi-fw/stargazer/ent/episode"
	"github.com/nuomizi-fw/stargazer/ent/season"
	"github.com/nuomizi-fw/stargazer/ent/user"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Bangumi is the client for interacting with the Bangumi builders.
	Bangumi *BangumiClient
	// CastMember is the client for interacting with the CastMember builders.
	CastMember *CastMemberClient
	// Episode is the client for interacting with the Episode builders.
	Episode *EpisodeClient
	// Season is the client for interacting with the Season builders.
	Season *SeasonClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Bangumi = NewBangumiClient(c.config)
	c.CastMember = NewCastMemberClient(c.config)
	c.Episode = NewEpisodeClient(c.config)
	c.Season = NewSeasonClient(c.config)
	c.User = NewUserClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
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

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Bangumi:    NewBangumiClient(cfg),
		CastMember: NewCastMemberClient(cfg),
		Episode:    NewEpisodeClient(cfg),
		Season:     NewSeasonClient(cfg),
		User:       NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Bangumi:    NewBangumiClient(cfg),
		CastMember: NewCastMemberClient(cfg),
		Episode:    NewEpisodeClient(cfg),
		Season:     NewSeasonClient(cfg),
		User:       NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Bangumi.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
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
	c.Bangumi.Use(hooks...)
	c.CastMember.Use(hooks...)
	c.Episode.Use(hooks...)
	c.Season.Use(hooks...)
	c.User.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Bangumi.Intercept(interceptors...)
	c.CastMember.Intercept(interceptors...)
	c.Episode.Intercept(interceptors...)
	c.Season.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *BangumiMutation:
		return c.Bangumi.mutate(ctx, m)
	case *CastMemberMutation:
		return c.CastMember.mutate(ctx, m)
	case *EpisodeMutation:
		return c.Episode.mutate(ctx, m)
	case *SeasonMutation:
		return c.Season.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// BangumiClient is a client for the Bangumi schema.
type BangumiClient struct {
	config
}

// NewBangumiClient returns a client for the Bangumi from the given config.
func NewBangumiClient(c config) *BangumiClient {
	return &BangumiClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `bangumi.Hooks(f(g(h())))`.
func (c *BangumiClient) Use(hooks ...Hook) {
	c.hooks.Bangumi = append(c.hooks.Bangumi, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `bangumi.Intercept(f(g(h())))`.
func (c *BangumiClient) Intercept(interceptors ...Interceptor) {
	c.inters.Bangumi = append(c.inters.Bangumi, interceptors...)
}

// Create returns a builder for creating a Bangumi entity.
func (c *BangumiClient) Create() *BangumiCreate {
	mutation := newBangumiMutation(c.config, OpCreate)
	return &BangumiCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Bangumi entities.
func (c *BangumiClient) CreateBulk(builders ...*BangumiCreate) *BangumiCreateBulk {
	return &BangumiCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *BangumiClient) MapCreateBulk(slice any, setFunc func(*BangumiCreate, int)) *BangumiCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &BangumiCreateBulk{err: fmt.Errorf("calling to BangumiClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*BangumiCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &BangumiCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Bangumi.
func (c *BangumiClient) Update() *BangumiUpdate {
	mutation := newBangumiMutation(c.config, OpUpdate)
	return &BangumiUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BangumiClient) UpdateOne(b *Bangumi) *BangumiUpdateOne {
	mutation := newBangumiMutation(c.config, OpUpdateOne, withBangumi(b))
	return &BangumiUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BangumiClient) UpdateOneID(id int) *BangumiUpdateOne {
	mutation := newBangumiMutation(c.config, OpUpdateOne, withBangumiID(id))
	return &BangumiUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Bangumi.
func (c *BangumiClient) Delete() *BangumiDelete {
	mutation := newBangumiMutation(c.config, OpDelete)
	return &BangumiDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *BangumiClient) DeleteOne(b *Bangumi) *BangumiDeleteOne {
	return c.DeleteOneID(b.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *BangumiClient) DeleteOneID(id int) *BangumiDeleteOne {
	builder := c.Delete().Where(bangumi.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BangumiDeleteOne{builder}
}

// Query returns a query builder for Bangumi.
func (c *BangumiClient) Query() *BangumiQuery {
	return &BangumiQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeBangumi},
		inters: c.Interceptors(),
	}
}

// Get returns a Bangumi entity by its id.
func (c *BangumiClient) Get(ctx context.Context, id int) (*Bangumi, error) {
	return c.Query().Where(bangumi.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BangumiClient) GetX(ctx context.Context, id int) *Bangumi {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QuerySeasons queries the seasons edge of a Bangumi.
func (c *BangumiClient) QuerySeasons(b *Bangumi) *SeasonQuery {
	query := (&SeasonClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(bangumi.Table, bangumi.FieldID, id),
			sqlgraph.To(season.Table, season.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, bangumi.SeasonsTable, bangumi.SeasonsColumn),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCastMembers queries the cast_members edge of a Bangumi.
func (c *BangumiClient) QueryCastMembers(b *Bangumi) *CastMemberQuery {
	query := (&CastMemberClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(bangumi.Table, bangumi.FieldID, id),
			sqlgraph.To(castmember.Table, castmember.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, bangumi.CastMembersTable, bangumi.CastMembersColumn),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BangumiClient) Hooks() []Hook {
	return c.hooks.Bangumi
}

// Interceptors returns the client interceptors.
func (c *BangumiClient) Interceptors() []Interceptor {
	return c.inters.Bangumi
}

func (c *BangumiClient) mutate(ctx context.Context, m *BangumiMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&BangumiCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&BangumiUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&BangumiUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&BangumiDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Bangumi mutation op: %q", m.Op())
	}
}

// CastMemberClient is a client for the CastMember schema.
type CastMemberClient struct {
	config
}

// NewCastMemberClient returns a client for the CastMember from the given config.
func NewCastMemberClient(c config) *CastMemberClient {
	return &CastMemberClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `castmember.Hooks(f(g(h())))`.
func (c *CastMemberClient) Use(hooks ...Hook) {
	c.hooks.CastMember = append(c.hooks.CastMember, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `castmember.Intercept(f(g(h())))`.
func (c *CastMemberClient) Intercept(interceptors ...Interceptor) {
	c.inters.CastMember = append(c.inters.CastMember, interceptors...)
}

// Create returns a builder for creating a CastMember entity.
func (c *CastMemberClient) Create() *CastMemberCreate {
	mutation := newCastMemberMutation(c.config, OpCreate)
	return &CastMemberCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CastMember entities.
func (c *CastMemberClient) CreateBulk(builders ...*CastMemberCreate) *CastMemberCreateBulk {
	return &CastMemberCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *CastMemberClient) MapCreateBulk(slice any, setFunc func(*CastMemberCreate, int)) *CastMemberCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &CastMemberCreateBulk{err: fmt.Errorf("calling to CastMemberClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*CastMemberCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &CastMemberCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CastMember.
func (c *CastMemberClient) Update() *CastMemberUpdate {
	mutation := newCastMemberMutation(c.config, OpUpdate)
	return &CastMemberUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CastMemberClient) UpdateOne(cm *CastMember) *CastMemberUpdateOne {
	mutation := newCastMemberMutation(c.config, OpUpdateOne, withCastMember(cm))
	return &CastMemberUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CastMemberClient) UpdateOneID(id int) *CastMemberUpdateOne {
	mutation := newCastMemberMutation(c.config, OpUpdateOne, withCastMemberID(id))
	return &CastMemberUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CastMember.
func (c *CastMemberClient) Delete() *CastMemberDelete {
	mutation := newCastMemberMutation(c.config, OpDelete)
	return &CastMemberDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CastMemberClient) DeleteOne(cm *CastMember) *CastMemberDeleteOne {
	return c.DeleteOneID(cm.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CastMemberClient) DeleteOneID(id int) *CastMemberDeleteOne {
	builder := c.Delete().Where(castmember.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CastMemberDeleteOne{builder}
}

// Query returns a query builder for CastMember.
func (c *CastMemberClient) Query() *CastMemberQuery {
	return &CastMemberQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCastMember},
		inters: c.Interceptors(),
	}
}

// Get returns a CastMember entity by its id.
func (c *CastMemberClient) Get(ctx context.Context, id int) (*CastMember, error) {
	return c.Query().Where(castmember.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CastMemberClient) GetX(ctx context.Context, id int) *CastMember {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QuerySeason queries the season edge of a CastMember.
func (c *CastMemberClient) QuerySeason(cm *CastMember) *SeasonQuery {
	query := (&SeasonClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := cm.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(castmember.Table, castmember.FieldID, id),
			sqlgraph.To(season.Table, season.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, castmember.SeasonTable, castmember.SeasonColumn),
		)
		fromV = sqlgraph.Neighbors(cm.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CastMemberClient) Hooks() []Hook {
	return c.hooks.CastMember
}

// Interceptors returns the client interceptors.
func (c *CastMemberClient) Interceptors() []Interceptor {
	return c.inters.CastMember
}

func (c *CastMemberClient) mutate(ctx context.Context, m *CastMemberMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CastMemberCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CastMemberUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CastMemberUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CastMemberDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown CastMember mutation op: %q", m.Op())
	}
}

// EpisodeClient is a client for the Episode schema.
type EpisodeClient struct {
	config
}

// NewEpisodeClient returns a client for the Episode from the given config.
func NewEpisodeClient(c config) *EpisodeClient {
	return &EpisodeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `episode.Hooks(f(g(h())))`.
func (c *EpisodeClient) Use(hooks ...Hook) {
	c.hooks.Episode = append(c.hooks.Episode, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `episode.Intercept(f(g(h())))`.
func (c *EpisodeClient) Intercept(interceptors ...Interceptor) {
	c.inters.Episode = append(c.inters.Episode, interceptors...)
}

// Create returns a builder for creating a Episode entity.
func (c *EpisodeClient) Create() *EpisodeCreate {
	mutation := newEpisodeMutation(c.config, OpCreate)
	return &EpisodeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Episode entities.
func (c *EpisodeClient) CreateBulk(builders ...*EpisodeCreate) *EpisodeCreateBulk {
	return &EpisodeCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *EpisodeClient) MapCreateBulk(slice any, setFunc func(*EpisodeCreate, int)) *EpisodeCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &EpisodeCreateBulk{err: fmt.Errorf("calling to EpisodeClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*EpisodeCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &EpisodeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Episode.
func (c *EpisodeClient) Update() *EpisodeUpdate {
	mutation := newEpisodeMutation(c.config, OpUpdate)
	return &EpisodeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *EpisodeClient) UpdateOne(e *Episode) *EpisodeUpdateOne {
	mutation := newEpisodeMutation(c.config, OpUpdateOne, withEpisode(e))
	return &EpisodeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *EpisodeClient) UpdateOneID(id int) *EpisodeUpdateOne {
	mutation := newEpisodeMutation(c.config, OpUpdateOne, withEpisodeID(id))
	return &EpisodeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Episode.
func (c *EpisodeClient) Delete() *EpisodeDelete {
	mutation := newEpisodeMutation(c.config, OpDelete)
	return &EpisodeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *EpisodeClient) DeleteOne(e *Episode) *EpisodeDeleteOne {
	return c.DeleteOneID(e.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *EpisodeClient) DeleteOneID(id int) *EpisodeDeleteOne {
	builder := c.Delete().Where(episode.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &EpisodeDeleteOne{builder}
}

// Query returns a query builder for Episode.
func (c *EpisodeClient) Query() *EpisodeQuery {
	return &EpisodeQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeEpisode},
		inters: c.Interceptors(),
	}
}

// Get returns a Episode entity by its id.
func (c *EpisodeClient) Get(ctx context.Context, id int) (*Episode, error) {
	return c.Query().Where(episode.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *EpisodeClient) GetX(ctx context.Context, id int) *Episode {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QuerySeason queries the season edge of a Episode.
func (c *EpisodeClient) QuerySeason(e *Episode) *SeasonQuery {
	query := (&SeasonClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := e.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(episode.Table, episode.FieldID, id),
			sqlgraph.To(season.Table, season.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, episode.SeasonTable, episode.SeasonColumn),
		)
		fromV = sqlgraph.Neighbors(e.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCastMembers queries the cast_members edge of a Episode.
func (c *EpisodeClient) QueryCastMembers(e *Episode) *CastMemberQuery {
	query := (&CastMemberClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := e.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(episode.Table, episode.FieldID, id),
			sqlgraph.To(castmember.Table, castmember.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, episode.CastMembersTable, episode.CastMembersColumn),
		)
		fromV = sqlgraph.Neighbors(e.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *EpisodeClient) Hooks() []Hook {
	return c.hooks.Episode
}

// Interceptors returns the client interceptors.
func (c *EpisodeClient) Interceptors() []Interceptor {
	return c.inters.Episode
}

func (c *EpisodeClient) mutate(ctx context.Context, m *EpisodeMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&EpisodeCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&EpisodeUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&EpisodeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&EpisodeDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Episode mutation op: %q", m.Op())
	}
}

// SeasonClient is a client for the Season schema.
type SeasonClient struct {
	config
}

// NewSeasonClient returns a client for the Season from the given config.
func NewSeasonClient(c config) *SeasonClient {
	return &SeasonClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `season.Hooks(f(g(h())))`.
func (c *SeasonClient) Use(hooks ...Hook) {
	c.hooks.Season = append(c.hooks.Season, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `season.Intercept(f(g(h())))`.
func (c *SeasonClient) Intercept(interceptors ...Interceptor) {
	c.inters.Season = append(c.inters.Season, interceptors...)
}

// Create returns a builder for creating a Season entity.
func (c *SeasonClient) Create() *SeasonCreate {
	mutation := newSeasonMutation(c.config, OpCreate)
	return &SeasonCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Season entities.
func (c *SeasonClient) CreateBulk(builders ...*SeasonCreate) *SeasonCreateBulk {
	return &SeasonCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *SeasonClient) MapCreateBulk(slice any, setFunc func(*SeasonCreate, int)) *SeasonCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &SeasonCreateBulk{err: fmt.Errorf("calling to SeasonClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*SeasonCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &SeasonCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Season.
func (c *SeasonClient) Update() *SeasonUpdate {
	mutation := newSeasonMutation(c.config, OpUpdate)
	return &SeasonUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SeasonClient) UpdateOne(s *Season) *SeasonUpdateOne {
	mutation := newSeasonMutation(c.config, OpUpdateOne, withSeason(s))
	return &SeasonUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SeasonClient) UpdateOneID(id int) *SeasonUpdateOne {
	mutation := newSeasonMutation(c.config, OpUpdateOne, withSeasonID(id))
	return &SeasonUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Season.
func (c *SeasonClient) Delete() *SeasonDelete {
	mutation := newSeasonMutation(c.config, OpDelete)
	return &SeasonDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *SeasonClient) DeleteOne(s *Season) *SeasonDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *SeasonClient) DeleteOneID(id int) *SeasonDeleteOne {
	builder := c.Delete().Where(season.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SeasonDeleteOne{builder}
}

// Query returns a query builder for Season.
func (c *SeasonClient) Query() *SeasonQuery {
	return &SeasonQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeSeason},
		inters: c.Interceptors(),
	}
}

// Get returns a Season entity by its id.
func (c *SeasonClient) Get(ctx context.Context, id int) (*Season, error) {
	return c.Query().Where(season.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SeasonClient) GetX(ctx context.Context, id int) *Season {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryBangumi queries the bangumi edge of a Season.
func (c *SeasonClient) QueryBangumi(s *Season) *BangumiQuery {
	query := (&BangumiClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(season.Table, season.FieldID, id),
			sqlgraph.To(bangumi.Table, bangumi.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, season.BangumiTable, season.BangumiColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryEpisodes queries the episodes edge of a Season.
func (c *SeasonClient) QueryEpisodes(s *Season) *EpisodeQuery {
	query := (&EpisodeClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(season.Table, season.FieldID, id),
			sqlgraph.To(episode.Table, episode.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, season.EpisodesTable, season.EpisodesColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCastMembers queries the cast_members edge of a Season.
func (c *SeasonClient) QueryCastMembers(s *Season) *CastMemberQuery {
	query := (&CastMemberClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(season.Table, season.FieldID, id),
			sqlgraph.To(castmember.Table, castmember.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, season.CastMembersTable, season.CastMembersColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *SeasonClient) Hooks() []Hook {
	return c.hooks.Season
}

// Interceptors returns the client interceptors.
func (c *SeasonClient) Interceptors() []Interceptor {
	return c.inters.Season
}

func (c *SeasonClient) mutate(ctx context.Context, m *SeasonMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&SeasonCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&SeasonUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&SeasonUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&SeasonDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Season mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *UserClient) MapCreateBulk(slice any, setFunc func(*UserCreate, int)) *UserCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &UserCreateBulk{err: fmt.Errorf("calling to UserClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*UserCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Bangumi, CastMember, Episode, Season, User []ent.Hook
	}
	inters struct {
		Bangumi, CastMember, Episode, Season, User []ent.Interceptor
	}
)