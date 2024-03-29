// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Faroukhamadi/northwind-api/ent/country_language"
	"github.com/Faroukhamadi/northwind-api/ent/predicate"
)

// CountryLanguageDelete is the builder for deleting a Country_language entity.
type CountryLanguageDelete struct {
	config
	hooks    []Hook
	mutation *CountryLanguageMutation
}

// Where appends a list predicates to the CountryLanguageDelete builder.
func (cld *CountryLanguageDelete) Where(ps ...predicate.Country_language) *CountryLanguageDelete {
	cld.mutation.Where(ps...)
	return cld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cld *CountryLanguageDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cld.hooks) == 0 {
		affected, err = cld.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CountryLanguageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cld.mutation = mutation
			affected, err = cld.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cld.hooks) - 1; i >= 0; i-- {
			if cld.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cld.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cld.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (cld *CountryLanguageDelete) ExecX(ctx context.Context) int {
	n, err := cld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cld *CountryLanguageDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: country_language.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: country_language.FieldID,
			},
		},
	}
	if ps := cld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, cld.driver, _spec)
}

// CountryLanguageDeleteOne is the builder for deleting a single Country_language entity.
type CountryLanguageDeleteOne struct {
	cld *CountryLanguageDelete
}

// Exec executes the deletion query.
func (cldo *CountryLanguageDeleteOne) Exec(ctx context.Context) error {
	n, err := cldo.cld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{country_language.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cldo *CountryLanguageDeleteOne) ExecX(ctx context.Context) {
	cldo.cld.ExecX(ctx)
}
