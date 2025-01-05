// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/polatbilal/gqlgen-ent/ent/companytoken"
	"github.com/polatbilal/gqlgen-ent/ent/predicate"
)

// CompanyTokenDelete is the builder for deleting a CompanyToken entity.
type CompanyTokenDelete struct {
	config
	hooks    []Hook
	mutation *CompanyTokenMutation
}

// Where appends a list predicates to the CompanyTokenDelete builder.
func (ctd *CompanyTokenDelete) Where(ps ...predicate.CompanyToken) *CompanyTokenDelete {
	ctd.mutation.Where(ps...)
	return ctd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ctd *CompanyTokenDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ctd.sqlExec, ctd.mutation, ctd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ctd *CompanyTokenDelete) ExecX(ctx context.Context) int {
	n, err := ctd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ctd *CompanyTokenDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(companytoken.Table, sqlgraph.NewFieldSpec(companytoken.FieldID, field.TypeInt))
	if ps := ctd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ctd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ctd.mutation.done = true
	return affected, err
}

// CompanyTokenDeleteOne is the builder for deleting a single CompanyToken entity.
type CompanyTokenDeleteOne struct {
	ctd *CompanyTokenDelete
}

// Where appends a list predicates to the CompanyTokenDelete builder.
func (ctdo *CompanyTokenDeleteOne) Where(ps ...predicate.CompanyToken) *CompanyTokenDeleteOne {
	ctdo.ctd.mutation.Where(ps...)
	return ctdo
}

// Exec executes the deletion query.
func (ctdo *CompanyTokenDeleteOne) Exec(ctx context.Context) error {
	n, err := ctdo.ctd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{companytoken.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ctdo *CompanyTokenDeleteOne) ExecX(ctx context.Context) {
	if err := ctdo.Exec(ctx); err != nil {
		panic(err)
	}
}
