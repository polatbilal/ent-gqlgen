// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/polatbilal/gqlgen-ent/ent/jobcontractor"
	"github.com/polatbilal/gqlgen-ent/ent/jobdetail"
)

// JobContractorCreate is the builder for creating a JobContractor entity.
type JobContractorCreate struct {
	config
	mutation *JobContractorMutation
	hooks    []Hook
}

// SetName sets the "Name" field.
func (jcc *JobContractorCreate) SetName(s string) *JobContractorCreate {
	jcc.mutation.SetName(s)
	return jcc
}

// SetNillableName sets the "Name" field if the given value is not nil.
func (jcc *JobContractorCreate) SetNillableName(s *string) *JobContractorCreate {
	if s != nil {
		jcc.SetName(*s)
	}
	return jcc
}

// SetTcNo sets the "TcNo" field.
func (jcc *JobContractorCreate) SetTcNo(i int) *JobContractorCreate {
	jcc.mutation.SetTcNo(i)
	return jcc
}

// SetNillableTcNo sets the "TcNo" field if the given value is not nil.
func (jcc *JobContractorCreate) SetNillableTcNo(i *int) *JobContractorCreate {
	if i != nil {
		jcc.SetTcNo(*i)
	}
	return jcc
}

// SetAddress sets the "Address" field.
func (jcc *JobContractorCreate) SetAddress(s string) *JobContractorCreate {
	jcc.mutation.SetAddress(s)
	return jcc
}

// SetNillableAddress sets the "Address" field if the given value is not nil.
func (jcc *JobContractorCreate) SetNillableAddress(s *string) *JobContractorCreate {
	if s != nil {
		jcc.SetAddress(*s)
	}
	return jcc
}

// SetRegisterNo sets the "RegisterNo" field.
func (jcc *JobContractorCreate) SetRegisterNo(i int) *JobContractorCreate {
	jcc.mutation.SetRegisterNo(i)
	return jcc
}

// SetNillableRegisterNo sets the "RegisterNo" field if the given value is not nil.
func (jcc *JobContractorCreate) SetNillableRegisterNo(i *int) *JobContractorCreate {
	if i != nil {
		jcc.SetRegisterNo(*i)
	}
	return jcc
}

// SetTaxAdmin sets the "TaxAdmin" field.
func (jcc *JobContractorCreate) SetTaxAdmin(s string) *JobContractorCreate {
	jcc.mutation.SetTaxAdmin(s)
	return jcc
}

// SetNillableTaxAdmin sets the "TaxAdmin" field if the given value is not nil.
func (jcc *JobContractorCreate) SetNillableTaxAdmin(s *string) *JobContractorCreate {
	if s != nil {
		jcc.SetTaxAdmin(*s)
	}
	return jcc
}

// SetTaxNo sets the "TaxNo" field.
func (jcc *JobContractorCreate) SetTaxNo(i int) *JobContractorCreate {
	jcc.mutation.SetTaxNo(i)
	return jcc
}

// SetNillableTaxNo sets the "TaxNo" field if the given value is not nil.
func (jcc *JobContractorCreate) SetNillableTaxNo(i *int) *JobContractorCreate {
	if i != nil {
		jcc.SetTaxNo(*i)
	}
	return jcc
}

// SetPhone sets the "Phone" field.
func (jcc *JobContractorCreate) SetPhone(s string) *JobContractorCreate {
	jcc.mutation.SetPhone(s)
	return jcc
}

// SetNillablePhone sets the "Phone" field if the given value is not nil.
func (jcc *JobContractorCreate) SetNillablePhone(s *string) *JobContractorCreate {
	if s != nil {
		jcc.SetPhone(*s)
	}
	return jcc
}

// SetEmail sets the "Email" field.
func (jcc *JobContractorCreate) SetEmail(s string) *JobContractorCreate {
	jcc.mutation.SetEmail(s)
	return jcc
}

// SetNillableEmail sets the "Email" field if the given value is not nil.
func (jcc *JobContractorCreate) SetNillableEmail(s *string) *JobContractorCreate {
	if s != nil {
		jcc.SetEmail(*s)
	}
	return jcc
}

// SetYdsID sets the "yds_id" field.
func (jcc *JobContractorCreate) SetYdsID(i int) *JobContractorCreate {
	jcc.mutation.SetYdsID(i)
	return jcc
}

// SetNillableYdsID sets the "yds_id" field if the given value is not nil.
func (jcc *JobContractorCreate) SetNillableYdsID(i *int) *JobContractorCreate {
	if i != nil {
		jcc.SetYdsID(*i)
	}
	return jcc
}

// SetNote sets the "Note" field.
func (jcc *JobContractorCreate) SetNote(s string) *JobContractorCreate {
	jcc.mutation.SetNote(s)
	return jcc
}

// SetNillableNote sets the "Note" field if the given value is not nil.
func (jcc *JobContractorCreate) SetNillableNote(s *string) *JobContractorCreate {
	if s != nil {
		jcc.SetNote(*s)
	}
	return jcc
}

// SetCreatedAt sets the "CreatedAt" field.
func (jcc *JobContractorCreate) SetCreatedAt(t time.Time) *JobContractorCreate {
	jcc.mutation.SetCreatedAt(t)
	return jcc
}

// SetNillableCreatedAt sets the "CreatedAt" field if the given value is not nil.
func (jcc *JobContractorCreate) SetNillableCreatedAt(t *time.Time) *JobContractorCreate {
	if t != nil {
		jcc.SetCreatedAt(*t)
	}
	return jcc
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (jcc *JobContractorCreate) SetUpdatedAt(t time.Time) *JobContractorCreate {
	jcc.mutation.SetUpdatedAt(t)
	return jcc
}

// SetNillableUpdatedAt sets the "UpdatedAt" field if the given value is not nil.
func (jcc *JobContractorCreate) SetNillableUpdatedAt(t *time.Time) *JobContractorCreate {
	if t != nil {
		jcc.SetUpdatedAt(*t)
	}
	return jcc
}

// AddContractorIDs adds the "contractors" edge to the JobDetail entity by IDs.
func (jcc *JobContractorCreate) AddContractorIDs(ids ...int) *JobContractorCreate {
	jcc.mutation.AddContractorIDs(ids...)
	return jcc
}

// AddContractors adds the "contractors" edges to the JobDetail entity.
func (jcc *JobContractorCreate) AddContractors(j ...*JobDetail) *JobContractorCreate {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return jcc.AddContractorIDs(ids...)
}

// Mutation returns the JobContractorMutation object of the builder.
func (jcc *JobContractorCreate) Mutation() *JobContractorMutation {
	return jcc.mutation
}

// Save creates the JobContractor in the database.
func (jcc *JobContractorCreate) Save(ctx context.Context) (*JobContractor, error) {
	jcc.defaults()
	return withHooks(ctx, jcc.sqlSave, jcc.mutation, jcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (jcc *JobContractorCreate) SaveX(ctx context.Context) *JobContractor {
	v, err := jcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (jcc *JobContractorCreate) Exec(ctx context.Context) error {
	_, err := jcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jcc *JobContractorCreate) ExecX(ctx context.Context) {
	if err := jcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (jcc *JobContractorCreate) defaults() {
	if _, ok := jcc.mutation.Name(); !ok {
		v := jobcontractor.DefaultName
		jcc.mutation.SetName(v)
	}
	if _, ok := jcc.mutation.CreatedAt(); !ok {
		v := jobcontractor.DefaultCreatedAt()
		jcc.mutation.SetCreatedAt(v)
	}
	if _, ok := jcc.mutation.UpdatedAt(); !ok {
		v := jobcontractor.DefaultUpdatedAt()
		jcc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (jcc *JobContractorCreate) check() error {
	if _, ok := jcc.mutation.Name(); !ok {
		return &ValidationError{Name: "Name", err: errors.New(`ent: missing required field "JobContractor.Name"`)}
	}
	if _, ok := jcc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "CreatedAt", err: errors.New(`ent: missing required field "JobContractor.CreatedAt"`)}
	}
	if _, ok := jcc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "UpdatedAt", err: errors.New(`ent: missing required field "JobContractor.UpdatedAt"`)}
	}
	return nil
}

func (jcc *JobContractorCreate) sqlSave(ctx context.Context) (*JobContractor, error) {
	if err := jcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := jcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, jcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	jcc.mutation.id = &_node.ID
	jcc.mutation.done = true
	return _node, nil
}

func (jcc *JobContractorCreate) createSpec() (*JobContractor, *sqlgraph.CreateSpec) {
	var (
		_node = &JobContractor{config: jcc.config}
		_spec = sqlgraph.NewCreateSpec(jobcontractor.Table, sqlgraph.NewFieldSpec(jobcontractor.FieldID, field.TypeInt))
	)
	if value, ok := jcc.mutation.Name(); ok {
		_spec.SetField(jobcontractor.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := jcc.mutation.TcNo(); ok {
		_spec.SetField(jobcontractor.FieldTcNo, field.TypeInt, value)
		_node.TcNo = value
	}
	if value, ok := jcc.mutation.Address(); ok {
		_spec.SetField(jobcontractor.FieldAddress, field.TypeString, value)
		_node.Address = value
	}
	if value, ok := jcc.mutation.RegisterNo(); ok {
		_spec.SetField(jobcontractor.FieldRegisterNo, field.TypeInt, value)
		_node.RegisterNo = value
	}
	if value, ok := jcc.mutation.TaxAdmin(); ok {
		_spec.SetField(jobcontractor.FieldTaxAdmin, field.TypeString, value)
		_node.TaxAdmin = value
	}
	if value, ok := jcc.mutation.TaxNo(); ok {
		_spec.SetField(jobcontractor.FieldTaxNo, field.TypeInt, value)
		_node.TaxNo = value
	}
	if value, ok := jcc.mutation.Phone(); ok {
		_spec.SetField(jobcontractor.FieldPhone, field.TypeString, value)
		_node.Phone = value
	}
	if value, ok := jcc.mutation.Email(); ok {
		_spec.SetField(jobcontractor.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := jcc.mutation.YdsID(); ok {
		_spec.SetField(jobcontractor.FieldYdsID, field.TypeInt, value)
		_node.YdsID = value
	}
	if value, ok := jcc.mutation.Note(); ok {
		_spec.SetField(jobcontractor.FieldNote, field.TypeString, value)
		_node.Note = value
	}
	if value, ok := jcc.mutation.CreatedAt(); ok {
		_spec.SetField(jobcontractor.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := jcc.mutation.UpdatedAt(); ok {
		_spec.SetField(jobcontractor.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := jcc.mutation.ContractorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   jobcontractor.ContractorsTable,
			Columns: []string{jobcontractor.ContractorsColumn},
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

// JobContractorCreateBulk is the builder for creating many JobContractor entities in bulk.
type JobContractorCreateBulk struct {
	config
	err      error
	builders []*JobContractorCreate
}

// Save creates the JobContractor entities in the database.
func (jccb *JobContractorCreateBulk) Save(ctx context.Context) ([]*JobContractor, error) {
	if jccb.err != nil {
		return nil, jccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(jccb.builders))
	nodes := make([]*JobContractor, len(jccb.builders))
	mutators := make([]Mutator, len(jccb.builders))
	for i := range jccb.builders {
		func(i int, root context.Context) {
			builder := jccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*JobContractorMutation)
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
					_, err = mutators[i+1].Mutate(root, jccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, jccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, jccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (jccb *JobContractorCreateBulk) SaveX(ctx context.Context) []*JobContractor {
	v, err := jccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (jccb *JobContractorCreateBulk) Exec(ctx context.Context) error {
	_, err := jccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jccb *JobContractorCreateBulk) ExecX(ctx context.Context) {
	if err := jccb.Exec(ctx); err != nil {
		panic(err)
	}
}
