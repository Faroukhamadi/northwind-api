// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/Faroukhamadi/northwind-api/ent/migrate"

	"github.com/Faroukhamadi/northwind-api/ent/city"
	"github.com/Faroukhamadi/northwind-api/ent/country"
	"github.com/Faroukhamadi/northwind-api/ent/country_language"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// City is the client for interacting with the City builders.
	City *CityClient
	// Country is the client for interacting with the Country builders.
	Country *CountryClient
	// Country_language is the client for interacting with the Country_language builders.
	Country_language *Country_languageClient
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
	c.City = NewCityClient(c.config)
	c.Country = NewCountryClient(c.config)
	c.Country_language = NewCountry_languageClient(c.config)
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
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:              ctx,
		config:           cfg,
		City:             NewCityClient(cfg),
		Country:          NewCountryClient(cfg),
		Country_language: NewCountry_languageClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
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
		ctx:              ctx,
		config:           cfg,
		City:             NewCityClient(cfg),
		Country:          NewCountryClient(cfg),
		Country_language: NewCountry_languageClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		City.
//		Query().
//		Count(ctx)
//
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
	c.City.Use(hooks...)
	c.Country.Use(hooks...)
	c.Country_language.Use(hooks...)
}

// CityClient is a client for the City schema.
type CityClient struct {
	config
}

// NewCityClient returns a client for the City from the given config.
func NewCityClient(c config) *CityClient {
	return &CityClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `city.Hooks(f(g(h())))`.
func (c *CityClient) Use(hooks ...Hook) {
	c.hooks.City = append(c.hooks.City, hooks...)
}

// Create returns a create builder for City.
func (c *CityClient) Create() *CityCreate {
	mutation := newCityMutation(c.config, OpCreate)
	return &CityCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of City entities.
func (c *CityClient) CreateBulk(builders ...*CityCreate) *CityCreateBulk {
	return &CityCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for City.
func (c *CityClient) Update() *CityUpdate {
	mutation := newCityMutation(c.config, OpUpdate)
	return &CityUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CityClient) UpdateOne(ci *City) *CityUpdateOne {
	mutation := newCityMutation(c.config, OpUpdateOne, withCity(ci))
	return &CityUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CityClient) UpdateOneID(id int) *CityUpdateOne {
	mutation := newCityMutation(c.config, OpUpdateOne, withCityID(id))
	return &CityUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for City.
func (c *CityClient) Delete() *CityDelete {
	mutation := newCityMutation(c.config, OpDelete)
	return &CityDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CityClient) DeleteOne(ci *City) *CityDeleteOne {
	return c.DeleteOneID(ci.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CityClient) DeleteOneID(id int) *CityDeleteOne {
	builder := c.Delete().Where(city.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CityDeleteOne{builder}
}

// Query returns a query builder for City.
func (c *CityClient) Query() *CityQuery {
	return &CityQuery{
		config: c.config,
	}
}

// Get returns a City entity by its id.
func (c *CityClient) Get(ctx context.Context, id int) (*City, error) {
	return c.Query().Where(city.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CityClient) GetX(ctx context.Context, id int) *City {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCountry queries the country edge of a City.
func (c *CityClient) QueryCountry(ci *City) *CountryQuery {
	query := &CountryQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ci.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(city.Table, city.FieldID, id),
			sqlgraph.To(country.Table, country.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, city.CountryTable, city.CountryColumn),
		)
		fromV = sqlgraph.Neighbors(ci.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CityClient) Hooks() []Hook {
	return c.hooks.City
}

// CountryClient is a client for the Country schema.
type CountryClient struct {
	config
}

// NewCountryClient returns a client for the Country from the given config.
func NewCountryClient(c config) *CountryClient {
	return &CountryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `country.Hooks(f(g(h())))`.
func (c *CountryClient) Use(hooks ...Hook) {
	c.hooks.Country = append(c.hooks.Country, hooks...)
}

// Create returns a create builder for Country.
func (c *CountryClient) Create() *CountryCreate {
	mutation := newCountryMutation(c.config, OpCreate)
	return &CountryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Country entities.
func (c *CountryClient) CreateBulk(builders ...*CountryCreate) *CountryCreateBulk {
	return &CountryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Country.
func (c *CountryClient) Update() *CountryUpdate {
	mutation := newCountryMutation(c.config, OpUpdate)
	return &CountryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CountryClient) UpdateOne(co *Country) *CountryUpdateOne {
	mutation := newCountryMutation(c.config, OpUpdateOne, withCountry(co))
	return &CountryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CountryClient) UpdateOneID(id string) *CountryUpdateOne {
	mutation := newCountryMutation(c.config, OpUpdateOne, withCountryID(id))
	return &CountryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Country.
func (c *CountryClient) Delete() *CountryDelete {
	mutation := newCountryMutation(c.config, OpDelete)
	return &CountryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CountryClient) DeleteOne(co *Country) *CountryDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CountryClient) DeleteOneID(id string) *CountryDeleteOne {
	builder := c.Delete().Where(country.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CountryDeleteOne{builder}
}

// Query returns a query builder for Country.
func (c *CountryClient) Query() *CountryQuery {
	return &CountryQuery{
		config: c.config,
	}
}

// Get returns a Country entity by its id.
func (c *CountryClient) Get(ctx context.Context, id string) (*Country, error) {
	return c.Query().Where(country.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CountryClient) GetX(ctx context.Context, id string) *Country {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCapital queries the capital edge of a Country.
func (c *CountryClient) QueryCapital(co *Country) *CityQuery {
	query := &CityQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(country.Table, country.FieldID, id),
			sqlgraph.To(city.Table, city.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, country.CapitalTable, country.CapitalColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryLanguage queries the language edge of a Country.
func (c *CountryClient) QueryLanguage(co *Country) *CountryLanguageQuery {
	query := &CountryLanguageQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(country.Table, country.FieldID, id),
			sqlgraph.To(country_language.Table, country_language.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, country.LanguageTable, country.LanguageColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CountryClient) Hooks() []Hook {
	return c.hooks.Country
}

// Country_languageClient is a client for the Country_language schema.
type Country_languageClient struct {
	config
}

// NewCountry_languageClient returns a client for the Country_language from the given config.
func NewCountry_languageClient(c config) *Country_languageClient {
	return &Country_languageClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `country_language.Hooks(f(g(h())))`.
func (c *Country_languageClient) Use(hooks ...Hook) {
	c.hooks.Country_language = append(c.hooks.Country_language, hooks...)
}

// Create returns a create builder for Country_language.
func (c *Country_languageClient) Create() *CountryLanguageCreate {
	mutation := newCountryLanguageMutation(c.config, OpCreate)
	return &CountryLanguageCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Country_language entities.
func (c *Country_languageClient) CreateBulk(builders ...*CountryLanguageCreate) *CountryLanguageCreateBulk {
	return &CountryLanguageCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Country_language.
func (c *Country_languageClient) Update() *CountryLanguageUpdate {
	mutation := newCountryLanguageMutation(c.config, OpUpdate)
	return &CountryLanguageUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *Country_languageClient) UpdateOne(cl *Country_language) *CountryLanguageUpdateOne {
	mutation := newCountryLanguageMutation(c.config, OpUpdateOne, withCountry_language(cl))
	return &CountryLanguageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *Country_languageClient) UpdateOneID(id int) *CountryLanguageUpdateOne {
	mutation := newCountryLanguageMutation(c.config, OpUpdateOne, withCountry_languageID(id))
	return &CountryLanguageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Country_language.
func (c *Country_languageClient) Delete() *CountryLanguageDelete {
	mutation := newCountryLanguageMutation(c.config, OpDelete)
	return &CountryLanguageDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *Country_languageClient) DeleteOne(cl *Country_language) *CountryLanguageDeleteOne {
	return c.DeleteOneID(cl.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *Country_languageClient) DeleteOneID(id int) *CountryLanguageDeleteOne {
	builder := c.Delete().Where(country_language.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CountryLanguageDeleteOne{builder}
}

// Query returns a query builder for Country_language.
func (c *Country_languageClient) Query() *CountryLanguageQuery {
	return &CountryLanguageQuery{
		config: c.config,
	}
}

// Get returns a Country_language entity by its id.
func (c *Country_languageClient) Get(ctx context.Context, id int) (*Country_language, error) {
	return c.Query().Where(country_language.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *Country_languageClient) GetX(ctx context.Context, id int) *Country_language {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCountry queries the country edge of a Country_language.
func (c *Country_languageClient) QueryCountry(cl *Country_language) *CountryQuery {
	query := &CountryQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := cl.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(country_language.Table, country_language.FieldID, id),
			sqlgraph.To(country.Table, country.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, country_language.CountryTable, country_language.CountryColumn),
		)
		fromV = sqlgraph.Neighbors(cl.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *Country_languageClient) Hooks() []Hook {
	return c.hooks.Country_language
}
