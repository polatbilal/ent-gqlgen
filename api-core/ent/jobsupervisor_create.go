// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/polatbilal/gqlgen-ent/api-core/ent/jobrelations"
	"github.com/polatbilal/gqlgen-ent/api-core/ent/jobsupervisor"
)

// JobSupervisorCreate is the builder for creating a JobSupervisor entity.
type JobSupervisorCreate struct {
	config
	mutation *JobSupervisorMutation
	hooks    []Hook
}

// SetName sets the "Name" field.
func (jsc *JobSupervisorCreate) SetName(s string) *JobSupervisorCreate {
	jsc.mutation.SetName(s)
	return jsc
}

// SetNillableName sets the "Name" field if the given value is not nil.
func (jsc *JobSupervisorCreate) SetNillableName(s *string) *JobSupervisorCreate {
	if s != nil {
		jsc.SetName(*s)
	}
	return jsc
}

// SetAddress sets the "Address" field.
func (jsc *JobSupervisorCreate) SetAddress(s string) *JobSupervisorCreate {
	jsc.mutation.SetAddress(s)
	return jsc
}

// SetNillableAddress sets the "Address" field if the given value is not nil.
func (jsc *JobSupervisorCreate) SetNillableAddress(s *string) *JobSupervisorCreate {
	if s != nil {
		jsc.SetAddress(*s)
	}
	return jsc
}

// SetPhone sets the "Phone" field.
func (jsc *JobSupervisorCreate) SetPhone(s string) *JobSupervisorCreate {
	jsc.mutation.SetPhone(s)
	return jsc
}

// SetNillablePhone sets the "Phone" field if the given value is not nil.
func (jsc *JobSupervisorCreate) SetNillablePhone(s *string) *JobSupervisorCreate {
	if s != nil {
		jsc.SetPhone(*s)
	}
	return jsc
}

// SetEmail sets the "Email" field.
func (jsc *JobSupervisorCreate) SetEmail(s string) *JobSupervisorCreate {
	jsc.mutation.SetEmail(s)
	return jsc
}

// SetNillableEmail sets the "Email" field if the given value is not nil.
func (jsc *JobSupervisorCreate) SetNillableEmail(s *string) *JobSupervisorCreate {
	if s != nil {
		jsc.SetEmail(*s)
	}
	return jsc
}

// SetTcNo sets the "TcNo" field.
func (jsc *JobSupervisorCreate) SetTcNo(s string) *JobSupervisorCreate {
	jsc.mutation.SetTcNo(s)
	return jsc
}

// SetNillableTcNo sets the "TcNo" field if the given value is not nil.
func (jsc *JobSupervisorCreate) SetNillableTcNo(s *string) *JobSupervisorCreate {
	if s != nil {
		jsc.SetTcNo(*s)
	}
	return jsc
}

// SetPosition sets the "Position" field.
func (jsc *JobSupervisorCreate) SetPosition(s string) *JobSupervisorCreate {
	jsc.mutation.SetPosition(s)
	return jsc
}

// SetNillablePosition sets the "Position" field if the given value is not nil.
func (jsc *JobSupervisorCreate) SetNillablePosition(s *string) *JobSupervisorCreate {
	if s != nil {
		jsc.SetPosition(*s)
	}
	return jsc
}

// SetCareer sets the "Career" field.
func (jsc *JobSupervisorCreate) SetCareer(s string) *JobSupervisorCreate {
	jsc.mutation.SetCareer(s)
	return jsc
}

// SetNillableCareer sets the "Career" field if the given value is not nil.
func (jsc *JobSupervisorCreate) SetNillableCareer(s *string) *JobSupervisorCreate {
	if s != nil {
		jsc.SetCareer(*s)
	}
	return jsc
}

// SetRegisterNo sets the "RegisterNo" field.
func (jsc *JobSupervisorCreate) SetRegisterNo(s string) *JobSupervisorCreate {
	jsc.mutation.SetRegisterNo(s)
	return jsc
}

// SetNillableRegisterNo sets the "RegisterNo" field if the given value is not nil.
func (jsc *JobSupervisorCreate) SetNillableRegisterNo(s *string) *JobSupervisorCreate {
	if s != nil {
		jsc.SetRegisterNo(*s)
	}
	return jsc
}

// SetSocialSecurityNo sets the "SocialSecurityNo" field.
func (jsc *JobSupervisorCreate) SetSocialSecurityNo(s string) *JobSupervisorCreate {
	jsc.mutation.SetSocialSecurityNo(s)
	return jsc
}

// SetNillableSocialSecurityNo sets the "SocialSecurityNo" field if the given value is not nil.
func (jsc *JobSupervisorCreate) SetNillableSocialSecurityNo(s *string) *JobSupervisorCreate {
	if s != nil {
		jsc.SetSocialSecurityNo(*s)
	}
	return jsc
}

// SetSchoolGraduation sets the "SchoolGraduation" field.
func (jsc *JobSupervisorCreate) SetSchoolGraduation(s string) *JobSupervisorCreate {
	jsc.mutation.SetSchoolGraduation(s)
	return jsc
}

// SetNillableSchoolGraduation sets the "SchoolGraduation" field if the given value is not nil.
func (jsc *JobSupervisorCreate) SetNillableSchoolGraduation(s *string) *JobSupervisorCreate {
	if s != nil {
		jsc.SetSchoolGraduation(*s)
	}
	return jsc
}

// SetYDSID sets the "YDSID" field.
func (jsc *JobSupervisorCreate) SetYDSID(i int) *JobSupervisorCreate {
	jsc.mutation.SetYDSID(i)
	return jsc
}

// SetNillableYDSID sets the "YDSID" field if the given value is not nil.
func (jsc *JobSupervisorCreate) SetNillableYDSID(i *int) *JobSupervisorCreate {
	if i != nil {
		jsc.SetYDSID(*i)
	}
	return jsc
}

// SetCreatedAt sets the "CreatedAt" field.
func (jsc *JobSupervisorCreate) SetCreatedAt(t time.Time) *JobSupervisorCreate {
	jsc.mutation.SetCreatedAt(t)
	return jsc
}

// SetNillableCreatedAt sets the "CreatedAt" field if the given value is not nil.
func (jsc *JobSupervisorCreate) SetNillableCreatedAt(t *time.Time) *JobSupervisorCreate {
	if t != nil {
		jsc.SetCreatedAt(*t)
	}
	return jsc
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (jsc *JobSupervisorCreate) SetUpdatedAt(t time.Time) *JobSupervisorCreate {
	jsc.mutation.SetUpdatedAt(t)
	return jsc
}

// SetNillableUpdatedAt sets the "UpdatedAt" field if the given value is not nil.
func (jsc *JobSupervisorCreate) SetNillableUpdatedAt(t *time.Time) *JobSupervisorCreate {
	if t != nil {
		jsc.SetUpdatedAt(*t)
	}
	return jsc
}

// AddSupervisorIDs adds the "supervisors" edge to the JobRelations entity by IDs.
func (jsc *JobSupervisorCreate) AddSupervisorIDs(ids ...int) *JobSupervisorCreate {
	jsc.mutation.AddSupervisorIDs(ids...)
	return jsc
}

// AddSupervisors adds the "supervisors" edges to the JobRelations entity.
func (jsc *JobSupervisorCreate) AddSupervisors(j ...*JobRelations) *JobSupervisorCreate {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return jsc.AddSupervisorIDs(ids...)
}

// Mutation returns the JobSupervisorMutation object of the builder.
func (jsc *JobSupervisorCreate) Mutation() *JobSupervisorMutation {
	return jsc.mutation
}

// Save creates the JobSupervisor in the database.
func (jsc *JobSupervisorCreate) Save(ctx context.Context) (*JobSupervisor, error) {
	jsc.defaults()
	return withHooks(ctx, jsc.sqlSave, jsc.mutation, jsc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (jsc *JobSupervisorCreate) SaveX(ctx context.Context) *JobSupervisor {
	v, err := jsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (jsc *JobSupervisorCreate) Exec(ctx context.Context) error {
	_, err := jsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jsc *JobSupervisorCreate) ExecX(ctx context.Context) {
	if err := jsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (jsc *JobSupervisorCreate) defaults() {
	if _, ok := jsc.mutation.CreatedAt(); !ok {
		v := jobsupervisor.DefaultCreatedAt()
		jsc.mutation.SetCreatedAt(v)
	}
	if _, ok := jsc.mutation.UpdatedAt(); !ok {
		v := jobsupervisor.DefaultUpdatedAt()
		jsc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (jsc *JobSupervisorCreate) check() error {
	if _, ok := jsc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "CreatedAt", err: errors.New(`ent: missing required field "JobSupervisor.CreatedAt"`)}
	}
	if _, ok := jsc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "UpdatedAt", err: errors.New(`ent: missing required field "JobSupervisor.UpdatedAt"`)}
	}
	return nil
}

func (jsc *JobSupervisorCreate) sqlSave(ctx context.Context) (*JobSupervisor, error) {
	if err := jsc.check(); err != nil {
		return nil, err
	}
	_node, _spec := jsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, jsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	jsc.mutation.id = &_node.ID
	jsc.mutation.done = true
	return _node, nil
}

func (jsc *JobSupervisorCreate) createSpec() (*JobSupervisor, *sqlgraph.CreateSpec) {
	var (
		_node = &JobSupervisor{config: jsc.config}
		_spec = sqlgraph.NewCreateSpec(jobsupervisor.Table, sqlgraph.NewFieldSpec(jobsupervisor.FieldID, field.TypeInt))
	)
	if value, ok := jsc.mutation.Name(); ok {
		_spec.SetField(jobsupervisor.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := jsc.mutation.Address(); ok {
		_spec.SetField(jobsupervisor.FieldAddress, field.TypeString, value)
		_node.Address = value
	}
	if value, ok := jsc.mutation.Phone(); ok {
		_spec.SetField(jobsupervisor.FieldPhone, field.TypeString, value)
		_node.Phone = value
	}
	if value, ok := jsc.mutation.Email(); ok {
		_spec.SetField(jobsupervisor.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := jsc.mutation.TcNo(); ok {
		_spec.SetField(jobsupervisor.FieldTcNo, field.TypeString, value)
		_node.TcNo = value
	}
	if value, ok := jsc.mutation.Position(); ok {
		_spec.SetField(jobsupervisor.FieldPosition, field.TypeString, value)
		_node.Position = value
	}
	if value, ok := jsc.mutation.Career(); ok {
		_spec.SetField(jobsupervisor.FieldCareer, field.TypeString, value)
		_node.Career = value
	}
	if value, ok := jsc.mutation.RegisterNo(); ok {
		_spec.SetField(jobsupervisor.FieldRegisterNo, field.TypeString, value)
		_node.RegisterNo = value
	}
	if value, ok := jsc.mutation.SocialSecurityNo(); ok {
		_spec.SetField(jobsupervisor.FieldSocialSecurityNo, field.TypeString, value)
		_node.SocialSecurityNo = value
	}
	if value, ok := jsc.mutation.SchoolGraduation(); ok {
		_spec.SetField(jobsupervisor.FieldSchoolGraduation, field.TypeString, value)
		_node.SchoolGraduation = value
	}
	if value, ok := jsc.mutation.YDSID(); ok {
		_spec.SetField(jobsupervisor.FieldYDSID, field.TypeInt, value)
		_node.YDSID = value
	}
	if value, ok := jsc.mutation.CreatedAt(); ok {
		_spec.SetField(jobsupervisor.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := jsc.mutation.UpdatedAt(); ok {
		_spec.SetField(jobsupervisor.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := jsc.mutation.SupervisorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   jobsupervisor.SupervisorsTable,
			Columns: []string{jobsupervisor.SupervisorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(jobrelations.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// JobSupervisorCreateBulk is the builder for creating many JobSupervisor entities in bulk.
type JobSupervisorCreateBulk struct {
	config
	err      error
	builders []*JobSupervisorCreate
}

// Save creates the JobSupervisor entities in the database.
func (jscb *JobSupervisorCreateBulk) Save(ctx context.Context) ([]*JobSupervisor, error) {
	if jscb.err != nil {
		return nil, jscb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(jscb.builders))
	nodes := make([]*JobSupervisor, len(jscb.builders))
	mutators := make([]Mutator, len(jscb.builders))
	for i := range jscb.builders {
		func(i int, root context.Context) {
			builder := jscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*JobSupervisorMutation)
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
					_, err = mutators[i+1].Mutate(root, jscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, jscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, jscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (jscb *JobSupervisorCreateBulk) SaveX(ctx context.Context) []*JobSupervisor {
	v, err := jscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (jscb *JobSupervisorCreateBulk) Exec(ctx context.Context) error {
	_, err := jscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jscb *JobSupervisorCreateBulk) ExecX(ctx context.Context) {
	if err := jscb.Exec(ctx); err != nil {
		panic(err)
	}
}
