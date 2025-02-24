// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/polatbilal/gqlgen-ent/ent/jobauthor"
	"github.com/polatbilal/gqlgen-ent/ent/jobdetail"
)

// JobAuthorCreate is the builder for creating a JobAuthor entity.
type JobAuthorCreate struct {
	config
	mutation *JobAuthorMutation
	hooks    []Hook
}

// SetArchitect sets the "Architect" field.
func (jac *JobAuthorCreate) SetArchitect(s string) *JobAuthorCreate {
	jac.mutation.SetArchitect(s)
	return jac
}

// SetNillableArchitect sets the "Architect" field if the given value is not nil.
func (jac *JobAuthorCreate) SetNillableArchitect(s *string) *JobAuthorCreate {
	if s != nil {
		jac.SetArchitect(*s)
	}
	return jac
}

// SetStatic sets the "Static" field.
func (jac *JobAuthorCreate) SetStatic(s string) *JobAuthorCreate {
	jac.mutation.SetStatic(s)
	return jac
}

// SetNillableStatic sets the "Static" field if the given value is not nil.
func (jac *JobAuthorCreate) SetNillableStatic(s *string) *JobAuthorCreate {
	if s != nil {
		jac.SetStatic(*s)
	}
	return jac
}

// SetMechanic sets the "Mechanic" field.
func (jac *JobAuthorCreate) SetMechanic(s string) *JobAuthorCreate {
	jac.mutation.SetMechanic(s)
	return jac
}

// SetNillableMechanic sets the "Mechanic" field if the given value is not nil.
func (jac *JobAuthorCreate) SetNillableMechanic(s *string) *JobAuthorCreate {
	if s != nil {
		jac.SetMechanic(*s)
	}
	return jac
}

// SetElectric sets the "Electric" field.
func (jac *JobAuthorCreate) SetElectric(s string) *JobAuthorCreate {
	jac.mutation.SetElectric(s)
	return jac
}

// SetNillableElectric sets the "Electric" field if the given value is not nil.
func (jac *JobAuthorCreate) SetNillableElectric(s *string) *JobAuthorCreate {
	if s != nil {
		jac.SetElectric(*s)
	}
	return jac
}

// SetFloor sets the "Floor" field.
func (jac *JobAuthorCreate) SetFloor(s string) *JobAuthorCreate {
	jac.mutation.SetFloor(s)
	return jac
}

// SetNillableFloor sets the "Floor" field if the given value is not nil.
func (jac *JobAuthorCreate) SetNillableFloor(s *string) *JobAuthorCreate {
	if s != nil {
		jac.SetFloor(*s)
	}
	return jac
}

// SetCreatedAt sets the "CreatedAt" field.
func (jac *JobAuthorCreate) SetCreatedAt(t time.Time) *JobAuthorCreate {
	jac.mutation.SetCreatedAt(t)
	return jac
}

// SetNillableCreatedAt sets the "CreatedAt" field if the given value is not nil.
func (jac *JobAuthorCreate) SetNillableCreatedAt(t *time.Time) *JobAuthorCreate {
	if t != nil {
		jac.SetCreatedAt(*t)
	}
	return jac
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (jac *JobAuthorCreate) SetUpdatedAt(t time.Time) *JobAuthorCreate {
	jac.mutation.SetUpdatedAt(t)
	return jac
}

// SetNillableUpdatedAt sets the "UpdatedAt" field if the given value is not nil.
func (jac *JobAuthorCreate) SetNillableUpdatedAt(t *time.Time) *JobAuthorCreate {
	if t != nil {
		jac.SetUpdatedAt(*t)
	}
	return jac
}

// AddAuthorIDs adds the "authors" edge to the JobDetail entity by IDs.
func (jac *JobAuthorCreate) AddAuthorIDs(ids ...int) *JobAuthorCreate {
	jac.mutation.AddAuthorIDs(ids...)
	return jac
}

// AddAuthors adds the "authors" edges to the JobDetail entity.
func (jac *JobAuthorCreate) AddAuthors(j ...*JobDetail) *JobAuthorCreate {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return jac.AddAuthorIDs(ids...)
}

// Mutation returns the JobAuthorMutation object of the builder.
func (jac *JobAuthorCreate) Mutation() *JobAuthorMutation {
	return jac.mutation
}

// Save creates the JobAuthor in the database.
func (jac *JobAuthorCreate) Save(ctx context.Context) (*JobAuthor, error) {
	jac.defaults()
	return withHooks(ctx, jac.sqlSave, jac.mutation, jac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (jac *JobAuthorCreate) SaveX(ctx context.Context) *JobAuthor {
	v, err := jac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (jac *JobAuthorCreate) Exec(ctx context.Context) error {
	_, err := jac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jac *JobAuthorCreate) ExecX(ctx context.Context) {
	if err := jac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (jac *JobAuthorCreate) defaults() {
	if _, ok := jac.mutation.Architect(); !ok {
		v := jobauthor.DefaultArchitect
		jac.mutation.SetArchitect(v)
	}
	if _, ok := jac.mutation.Static(); !ok {
		v := jobauthor.DefaultStatic
		jac.mutation.SetStatic(v)
	}
	if _, ok := jac.mutation.Mechanic(); !ok {
		v := jobauthor.DefaultMechanic
		jac.mutation.SetMechanic(v)
	}
	if _, ok := jac.mutation.Electric(); !ok {
		v := jobauthor.DefaultElectric
		jac.mutation.SetElectric(v)
	}
	if _, ok := jac.mutation.Floor(); !ok {
		v := jobauthor.DefaultFloor
		jac.mutation.SetFloor(v)
	}
	if _, ok := jac.mutation.CreatedAt(); !ok {
		v := jobauthor.DefaultCreatedAt()
		jac.mutation.SetCreatedAt(v)
	}
	if _, ok := jac.mutation.UpdatedAt(); !ok {
		v := jobauthor.DefaultUpdatedAt()
		jac.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (jac *JobAuthorCreate) check() error {
	if _, ok := jac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "CreatedAt", err: errors.New(`ent: missing required field "JobAuthor.CreatedAt"`)}
	}
	if _, ok := jac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "UpdatedAt", err: errors.New(`ent: missing required field "JobAuthor.UpdatedAt"`)}
	}
	return nil
}

func (jac *JobAuthorCreate) sqlSave(ctx context.Context) (*JobAuthor, error) {
	if err := jac.check(); err != nil {
		return nil, err
	}
	_node, _spec := jac.createSpec()
	if err := sqlgraph.CreateNode(ctx, jac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	jac.mutation.id = &_node.ID
	jac.mutation.done = true
	return _node, nil
}

func (jac *JobAuthorCreate) createSpec() (*JobAuthor, *sqlgraph.CreateSpec) {
	var (
		_node = &JobAuthor{config: jac.config}
		_spec = sqlgraph.NewCreateSpec(jobauthor.Table, sqlgraph.NewFieldSpec(jobauthor.FieldID, field.TypeInt))
	)
	if value, ok := jac.mutation.Architect(); ok {
		_spec.SetField(jobauthor.FieldArchitect, field.TypeString, value)
		_node.Architect = value
	}
	if value, ok := jac.mutation.Static(); ok {
		_spec.SetField(jobauthor.FieldStatic, field.TypeString, value)
		_node.Static = value
	}
	if value, ok := jac.mutation.Mechanic(); ok {
		_spec.SetField(jobauthor.FieldMechanic, field.TypeString, value)
		_node.Mechanic = value
	}
	if value, ok := jac.mutation.Electric(); ok {
		_spec.SetField(jobauthor.FieldElectric, field.TypeString, value)
		_node.Electric = value
	}
	if value, ok := jac.mutation.Floor(); ok {
		_spec.SetField(jobauthor.FieldFloor, field.TypeString, value)
		_node.Floor = value
	}
	if value, ok := jac.mutation.CreatedAt(); ok {
		_spec.SetField(jobauthor.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := jac.mutation.UpdatedAt(); ok {
		_spec.SetField(jobauthor.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := jac.mutation.AuthorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   jobauthor.AuthorsTable,
			Columns: []string{jobauthor.AuthorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(jobdetail.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// JobAuthorCreateBulk is the builder for creating many JobAuthor entities in bulk.
type JobAuthorCreateBulk struct {
	config
	err      error
	builders []*JobAuthorCreate
}

// Save creates the JobAuthor entities in the database.
func (jacb *JobAuthorCreateBulk) Save(ctx context.Context) ([]*JobAuthor, error) {
	if jacb.err != nil {
		return nil, jacb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(jacb.builders))
	nodes := make([]*JobAuthor, len(jacb.builders))
	mutators := make([]Mutator, len(jacb.builders))
	for i := range jacb.builders {
		func(i int, root context.Context) {
			builder := jacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*JobAuthorMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, jacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, jacb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, jacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (jacb *JobAuthorCreateBulk) SaveX(ctx context.Context) []*JobAuthor {
	v, err := jacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (jacb *JobAuthorCreateBulk) Exec(ctx context.Context) error {
	_, err := jacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jacb *JobAuthorCreateBulk) ExecX(ctx context.Context) {
	if err := jacb.Exec(ctx); err != nil {
		panic(err)
	}
}
