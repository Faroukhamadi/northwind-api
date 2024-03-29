// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Faroukhamadi/northwind-api/ent/country"
	"github.com/Faroukhamadi/northwind-api/ent/country_language"
	"github.com/Faroukhamadi/northwind-api/ent/predicate"
)

// CountryLanguageUpdate is the builder for updating Country_language entities.
type CountryLanguageUpdate struct {
	config
	hooks    []Hook
	mutation *CountryLanguageMutation
}

// Where appends a list predicates to the CountryLanguageUpdate builder.
func (clu *CountryLanguageUpdate) Where(ps ...predicate.Country_language) *CountryLanguageUpdate {
	clu.mutation.Where(ps...)
	return clu
}

// SetCountryCode sets the "country_code" field.
func (clu *CountryLanguageUpdate) SetCountryCode(s string) *CountryLanguageUpdate {
	clu.mutation.SetCountryCode(s)
	return clu
}

// SetLanguage sets the "language" field.
func (clu *CountryLanguageUpdate) SetLanguage(s string) *CountryLanguageUpdate {
	clu.mutation.SetLanguage(s)
	return clu
}

// SetIsOfficial sets the "is_official" field.
func (clu *CountryLanguageUpdate) SetIsOfficial(b bool) *CountryLanguageUpdate {
	clu.mutation.SetIsOfficial(b)
	return clu
}

// SetPercentage sets the "percentage" field.
func (clu *CountryLanguageUpdate) SetPercentage(f float64) *CountryLanguageUpdate {
	clu.mutation.ResetPercentage()
	clu.mutation.SetPercentage(f)
	return clu
}

// AddPercentage adds f to the "percentage" field.
func (clu *CountryLanguageUpdate) AddPercentage(f float64) *CountryLanguageUpdate {
	clu.mutation.AddPercentage(f)
	return clu
}

// SetCountryID sets the "country" edge to the Country entity by ID.
func (clu *CountryLanguageUpdate) SetCountryID(id string) *CountryLanguageUpdate {
	clu.mutation.SetCountryID(id)
	return clu
}

// SetCountry sets the "country" edge to the Country entity.
func (clu *CountryLanguageUpdate) SetCountry(c *Country) *CountryLanguageUpdate {
	return clu.SetCountryID(c.ID)
}

// Mutation returns the CountryLanguageMutation object of the builder.
func (clu *CountryLanguageUpdate) Mutation() *CountryLanguageMutation {
	return clu.mutation
}

// ClearCountry clears the "country" edge to the Country entity.
func (clu *CountryLanguageUpdate) ClearCountry() *CountryLanguageUpdate {
	clu.mutation.ClearCountry()
	return clu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (clu *CountryLanguageUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(clu.hooks) == 0 {
		if err = clu.check(); err != nil {
			return 0, err
		}
		affected, err = clu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CountryLanguageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = clu.check(); err != nil {
				return 0, err
			}
			clu.mutation = mutation
			affected, err = clu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(clu.hooks) - 1; i >= 0; i-- {
			if clu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = clu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, clu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (clu *CountryLanguageUpdate) SaveX(ctx context.Context) int {
	affected, err := clu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (clu *CountryLanguageUpdate) Exec(ctx context.Context) error {
	_, err := clu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (clu *CountryLanguageUpdate) ExecX(ctx context.Context) {
	if err := clu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (clu *CountryLanguageUpdate) check() error {
	if _, ok := clu.mutation.CountryID(); clu.mutation.CountryCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Country_language.country"`)
	}
	return nil
}

func (clu *CountryLanguageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   country_language.Table,
			Columns: country_language.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: country_language.FieldID,
			},
		},
	}
	if ps := clu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := clu.mutation.CountryCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: country_language.FieldCountryCode,
		})
	}
	if value, ok := clu.mutation.Language(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: country_language.FieldLanguage,
		})
	}
	if value, ok := clu.mutation.IsOfficial(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: country_language.FieldIsOfficial,
		})
	}
	if value, ok := clu.mutation.Percentage(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: country_language.FieldPercentage,
		})
	}
	if value, ok := clu.mutation.AddedPercentage(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: country_language.FieldPercentage,
		})
	}
	if clu.mutation.CountryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   country_language.CountryTable,
			Columns: []string{country_language.CountryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: country.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := clu.mutation.CountryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   country_language.CountryTable,
			Columns: []string{country_language.CountryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: country.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, clu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{country_language.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// CountryLanguageUpdateOne is the builder for updating a single Country_language entity.
type CountryLanguageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CountryLanguageMutation
}

// SetCountryCode sets the "country_code" field.
func (cluo *CountryLanguageUpdateOne) SetCountryCode(s string) *CountryLanguageUpdateOne {
	cluo.mutation.SetCountryCode(s)
	return cluo
}

// SetLanguage sets the "language" field.
func (cluo *CountryLanguageUpdateOne) SetLanguage(s string) *CountryLanguageUpdateOne {
	cluo.mutation.SetLanguage(s)
	return cluo
}

// SetIsOfficial sets the "is_official" field.
func (cluo *CountryLanguageUpdateOne) SetIsOfficial(b bool) *CountryLanguageUpdateOne {
	cluo.mutation.SetIsOfficial(b)
	return cluo
}

// SetPercentage sets the "percentage" field.
func (cluo *CountryLanguageUpdateOne) SetPercentage(f float64) *CountryLanguageUpdateOne {
	cluo.mutation.ResetPercentage()
	cluo.mutation.SetPercentage(f)
	return cluo
}

// AddPercentage adds f to the "percentage" field.
func (cluo *CountryLanguageUpdateOne) AddPercentage(f float64) *CountryLanguageUpdateOne {
	cluo.mutation.AddPercentage(f)
	return cluo
}

// SetCountryID sets the "country" edge to the Country entity by ID.
func (cluo *CountryLanguageUpdateOne) SetCountryID(id string) *CountryLanguageUpdateOne {
	cluo.mutation.SetCountryID(id)
	return cluo
}

// SetCountry sets the "country" edge to the Country entity.
func (cluo *CountryLanguageUpdateOne) SetCountry(c *Country) *CountryLanguageUpdateOne {
	return cluo.SetCountryID(c.ID)
}

// Mutation returns the CountryLanguageMutation object of the builder.
func (cluo *CountryLanguageUpdateOne) Mutation() *CountryLanguageMutation {
	return cluo.mutation
}

// ClearCountry clears the "country" edge to the Country entity.
func (cluo *CountryLanguageUpdateOne) ClearCountry() *CountryLanguageUpdateOne {
	cluo.mutation.ClearCountry()
	return cluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cluo *CountryLanguageUpdateOne) Select(field string, fields ...string) *CountryLanguageUpdateOne {
	cluo.fields = append([]string{field}, fields...)
	return cluo
}

// Save executes the query and returns the updated Country_language entity.
func (cluo *CountryLanguageUpdateOne) Save(ctx context.Context) (*Country_language, error) {
	var (
		err  error
		node *Country_language
	)
	if len(cluo.hooks) == 0 {
		if err = cluo.check(); err != nil {
			return nil, err
		}
		node, err = cluo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CountryLanguageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cluo.check(); err != nil {
				return nil, err
			}
			cluo.mutation = mutation
			node, err = cluo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cluo.hooks) - 1; i >= 0; i-- {
			if cluo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cluo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cluo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cluo *CountryLanguageUpdateOne) SaveX(ctx context.Context) *Country_language {
	node, err := cluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cluo *CountryLanguageUpdateOne) Exec(ctx context.Context) error {
	_, err := cluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cluo *CountryLanguageUpdateOne) ExecX(ctx context.Context) {
	if err := cluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cluo *CountryLanguageUpdateOne) check() error {
	if _, ok := cluo.mutation.CountryID(); cluo.mutation.CountryCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Country_language.country"`)
	}
	return nil
}

func (cluo *CountryLanguageUpdateOne) sqlSave(ctx context.Context) (_node *Country_language, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   country_language.Table,
			Columns: country_language.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: country_language.FieldID,
			},
		},
	}
	id, ok := cluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Country_language.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, country_language.FieldID)
		for _, f := range fields {
			if !country_language.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != country_language.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cluo.mutation.CountryCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: country_language.FieldCountryCode,
		})
	}
	if value, ok := cluo.mutation.Language(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: country_language.FieldLanguage,
		})
	}
	if value, ok := cluo.mutation.IsOfficial(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: country_language.FieldIsOfficial,
		})
	}
	if value, ok := cluo.mutation.Percentage(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: country_language.FieldPercentage,
		})
	}
	if value, ok := cluo.mutation.AddedPercentage(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: country_language.FieldPercentage,
		})
	}
	if cluo.mutation.CountryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   country_language.CountryTable,
			Columns: []string{country_language.CountryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: country.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cluo.mutation.CountryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   country_language.CountryTable,
			Columns: []string{country_language.CountryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: country.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Country_language{config: cluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{country_language.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
