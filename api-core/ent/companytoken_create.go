// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/polatbilal/gqlgen-ent/api-core/ent/companydetail"
	"github.com/polatbilal/gqlgen-ent/api-core/ent/companytoken"
)

// CompanyTokenCreate is the builder for creating a CompanyToken entity.
type CompanyTokenCreate struct {
	config
	mutation *CompanyTokenMutation
	hooks    []Hook
}

// SetYDKUsername sets the "YDKUsername" field.
func (ctc *CompanyTokenCreate) SetYDKUsername(s string) *CompanyTokenCreate {
	ctc.mutation.SetYDKUsername(s)
	return ctc
}

// SetNillableYDKUsername sets the "YDKUsername" field if the given value is not nil.
func (ctc *CompanyTokenCreate) SetNillableYDKUsername(s *string) *CompanyTokenCreate {
	if s != nil {
		ctc.SetYDKUsername(*s)
	}
	return ctc
}

// SetYDKPassword sets the "YDKPassword" field.
func (ctc *CompanyTokenCreate) SetYDKPassword(s string) *CompanyTokenCreate {
	ctc.mutation.SetYDKPassword(s)
	return ctc
}

// SetNillableYDKPassword sets the "YDKPassword" field if the given value is not nil.
func (ctc *CompanyTokenCreate) SetNillableYDKPassword(s *string) *CompanyTokenCreate {
	if s != nil {
		ctc.SetYDKPassword(*s)
	}
	return ctc
}

// SetToken sets the "Token" field.
func (ctc *CompanyTokenCreate) SetToken(s string) *CompanyTokenCreate {
	ctc.mutation.SetToken(s)
	return ctc
}

// SetNillableToken sets the "Token" field if the given value is not nil.
func (ctc *CompanyTokenCreate) SetNillableToken(s *string) *CompanyTokenCreate {
	if s != nil {
		ctc.SetToken(*s)
	}
	return ctc
}

// SetDepartmentId sets the "DepartmentId" field.
func (ctc *CompanyTokenCreate) SetDepartmentId(i int) *CompanyTokenCreate {
	ctc.mutation.SetDepartmentId(i)
	return ctc
}

// SetNillableDepartmentId sets the "DepartmentId" field if the given value is not nil.
func (ctc *CompanyTokenCreate) SetNillableDepartmentId(i *int) *CompanyTokenCreate {
	if i != nil {
		ctc.SetDepartmentId(*i)
	}
	return ctc
}

// SetCreatedAt sets the "createdAt" field.
func (ctc *CompanyTokenCreate) SetCreatedAt(t time.Time) *CompanyTokenCreate {
	ctc.mutation.SetCreatedAt(t)
	return ctc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (ctc *CompanyTokenCreate) SetNillableCreatedAt(t *time.Time) *CompanyTokenCreate {
	if t != nil {
		ctc.SetCreatedAt(*t)
	}
	return ctc
}

// SetUpdatedAt sets the "updatedAt" field.
func (ctc *CompanyTokenCreate) SetUpdatedAt(t time.Time) *CompanyTokenCreate {
	ctc.mutation.SetUpdatedAt(t)
	return ctc
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (ctc *CompanyTokenCreate) SetNillableUpdatedAt(t *time.Time) *CompanyTokenCreate {
	if t != nil {
		ctc.SetUpdatedAt(*t)
	}
	return ctc
}

// SetCompanyID sets the "company" edge to the CompanyDetail entity by ID.
func (ctc *CompanyTokenCreate) SetCompanyID(id int) *CompanyTokenCreate {
	ctc.mutation.SetCompanyID(id)
	return ctc
}

// SetNillableCompanyID sets the "company" edge to the CompanyDetail entity by ID if the given value is not nil.
func (ctc *CompanyTokenCreate) SetNillableCompanyID(id *int) *CompanyTokenCreate {
	if id != nil {
		ctc = ctc.SetCompanyID(*id)
	}
	return ctc
}

// SetCompany sets the "company" edge to the CompanyDetail entity.
func (ctc *CompanyTokenCreate) SetCompany(c *CompanyDetail) *CompanyTokenCreate {
	return ctc.SetCompanyID(c.ID)
}

// Mutation returns the CompanyTokenMutation object of the builder.
func (ctc *CompanyTokenCreate) Mutation() *CompanyTokenMutation {
	return ctc.mutation
}

// Save creates the CompanyToken in the database.
func (ctc *CompanyTokenCreate) Save(ctx context.Context) (*CompanyToken, error) {
	ctc.defaults()
	return withHooks(ctx, ctc.sqlSave, ctc.mutation, ctc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ctc *CompanyTokenCreate) SaveX(ctx context.Context) *CompanyToken {
	v, err := ctc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ctc *CompanyTokenCreate) Exec(ctx context.Context) error {
	_, err := ctc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctc *CompanyTokenCreate) ExecX(ctx context.Context) {
	if err := ctc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ctc *CompanyTokenCreate) defaults() {
	if _, ok := ctc.mutation.CreatedAt(); !ok {
		v := companytoken.DefaultCreatedAt()
		ctc.mutation.SetCreatedAt(v)
	}
	if _, ok := ctc.mutation.UpdatedAt(); !ok {
		v := companytoken.DefaultUpdatedAt()
		ctc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ctc *CompanyTokenCreate) check() error {
	if _, ok := ctc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "CompanyToken.createdAt"`)}
	}
	if _, ok := ctc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "CompanyToken.updatedAt"`)}
	}
	return nil
}

func (ctc *CompanyTokenCreate) sqlSave(ctx context.Context) (*CompanyToken, error) {
	if err := ctc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ctc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ctc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ctc.mutation.id = &_node.ID
	ctc.mutation.done = true
	return _node, nil
}

func (ctc *CompanyTokenCreate) createSpec() (*CompanyToken, *sqlgraph.CreateSpec) {
	var (
		_node = &CompanyToken{config: ctc.config}
		_spec = sqlgraph.NewCreateSpec(companytoken.Table, sqlgraph.NewFieldSpec(companytoken.FieldID, field.TypeInt))
	)
	if value, ok := ctc.mutation.YDKUsername(); ok {
		_spec.SetField(companytoken.FieldYDKUsername, field.TypeString, value)
		_node.YDKUsername = value
	}
	if value, ok := ctc.mutation.YDKPassword(); ok {
		_spec.SetField(companytoken.FieldYDKPassword, field.TypeString, value)
		_node.YDKPassword = value
	}
	if value, ok := ctc.mutation.Token(); ok {
		_spec.SetField(companytoken.FieldToken, field.TypeString, value)
		_node.Token = value
	}
	if value, ok := ctc.mutation.DepartmentId(); ok {
		_spec.SetField(companytoken.FieldDepartmentId, field.TypeInt, value)
		_node.DepartmentId = value
	}
	if value, ok := ctc.mutation.CreatedAt(); ok {
		_spec.SetField(companytoken.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ctc.mutation.UpdatedAt(); ok {
		_spec.SetField(companytoken.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := ctc.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   companytoken.CompanyTable,
			Columns: []string{companytoken.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(companydetail.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.company_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CompanyTokenCreateBulk is the builder for creating many CompanyToken entities in bulk.
type CompanyTokenCreateBulk struct {
	config
	err      error
	builders []*CompanyTokenCreate
}

// Save creates the CompanyToken entities in the database.
func (ctcb *CompanyTokenCreateBulk) Save(ctx context.Context) ([]*CompanyToken, error) {
	if ctcb.err != nil {
		return nil, ctcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ctcb.builders))
	nodes := make([]*CompanyToken, len(ctcb.builders))
	mutators := make([]Mutator, len(ctcb.builders))
	for i := range ctcb.builders {
		func(i int, root context.Context) {
			builder := ctcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CompanyTokenMutation)
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
					_, err = mutators[i+1].Mutate(root, ctcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ctcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ctcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ctcb *CompanyTokenCreateBulk) SaveX(ctx context.Context) []*CompanyToken {
	v, err := ctcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ctcb *CompanyTokenCreateBulk) Exec(ctx context.Context) error {
	_, err := ctcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctcb *CompanyTokenCreateBulk) ExecX(ctx context.Context) {
	if err := ctcb.Exec(ctx); err != nil {
		panic(err)
	}
}
