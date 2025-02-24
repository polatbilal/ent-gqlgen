// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/polatbilal/gqlgen-ent/ent/jobauthor"
	"github.com/polatbilal/gqlgen-ent/ent/predicate"
)

// JobAuthorDelete is the builder for deleting a JobAuthor entity.
type JobAuthorDelete struct {
	config
	hooks    []Hook
	mutation *JobAuthorMutation
}

// Where appends a list predicates to the JobAuthorDelete builder.
func (jad *JobAuthorDelete) Where(ps ...predicate.JobAuthor) *JobAuthorDelete {
	jad.mutation.Where(ps...)
	return jad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (jad *JobAuthorDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, jad.sqlExec, jad.mutation, jad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (jad *JobAuthorDelete) ExecX(ctx context.Context) int {
	n, err := jad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (jad *JobAuthorDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(jobauthor.Table, sqlgraph.NewFieldSpec(jobauthor.FieldID, field.TypeInt))
	if ps := jad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, jad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	jad.mutation.done = true
	return affected, err
}

// JobAuthorDeleteOne is the builder for deleting a single JobAuthor entity.
type JobAuthorDeleteOne struct {
	jad *JobAuthorDelete
}

// Where appends a list predicates to the JobAuthorDelete builder.
func (jado *JobAuthorDeleteOne) Where(ps ...predicate.JobAuthor) *JobAuthorDeleteOne {
	jado.jad.mutation.Where(ps...)
	return jado
}

// Exec executes the deletion query.
func (jado *JobAuthorDeleteOne) Exec(ctx context.Context) error {
	n, err := jado.jad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{jobauthor.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (jado *JobAuthorDeleteOne) ExecX(ctx context.Context) {
	if err := jado.Exec(ctx); err != nil {
		panic(err)
	}
}
