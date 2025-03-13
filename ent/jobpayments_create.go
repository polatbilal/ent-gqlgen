// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/polatbilal/gqlgen-ent/ent/jobpayments"
	"github.com/polatbilal/gqlgen-ent/ent/jobrelations"
)

// JobPaymentsCreate is the builder for creating a JobPayments entity.
type JobPaymentsCreate struct {
	config
	mutation *JobPaymentsMutation
	hooks    []Hook
}

// SetYibfNo sets the "yibfNo" field.
func (jpc *JobPaymentsCreate) SetYibfNo(i int) *JobPaymentsCreate {
	jpc.mutation.SetYibfNo(i)
	return jpc
}

// SetNillableYibfNo sets the "yibfNo" field if the given value is not nil.
func (jpc *JobPaymentsCreate) SetNillableYibfNo(i *int) *JobPaymentsCreate {
	if i != nil {
		jpc.SetYibfNo(*i)
	}
	return jpc
}

// SetPaymentNo sets the "PaymentNo" field.
func (jpc *JobPaymentsCreate) SetPaymentNo(i int) *JobPaymentsCreate {
	jpc.mutation.SetPaymentNo(i)
	return jpc
}

// SetNillablePaymentNo sets the "PaymentNo" field if the given value is not nil.
func (jpc *JobPaymentsCreate) SetNillablePaymentNo(i *int) *JobPaymentsCreate {
	if i != nil {
		jpc.SetPaymentNo(*i)
	}
	return jpc
}

// SetPaymentDate sets the "PaymentDate" field.
func (jpc *JobPaymentsCreate) SetPaymentDate(t time.Time) *JobPaymentsCreate {
	jpc.mutation.SetPaymentDate(t)
	return jpc
}

// SetNillablePaymentDate sets the "PaymentDate" field if the given value is not nil.
func (jpc *JobPaymentsCreate) SetNillablePaymentDate(t *time.Time) *JobPaymentsCreate {
	if t != nil {
		jpc.SetPaymentDate(*t)
	}
	return jpc
}

// SetPaymentType sets the "PaymentType" field.
func (jpc *JobPaymentsCreate) SetPaymentType(s string) *JobPaymentsCreate {
	jpc.mutation.SetPaymentType(s)
	return jpc
}

// SetNillablePaymentType sets the "PaymentType" field if the given value is not nil.
func (jpc *JobPaymentsCreate) SetNillablePaymentType(s *string) *JobPaymentsCreate {
	if s != nil {
		jpc.SetPaymentType(*s)
	}
	return jpc
}

// SetState sets the "State" field.
func (jpc *JobPaymentsCreate) SetState(s string) *JobPaymentsCreate {
	jpc.mutation.SetState(s)
	return jpc
}

// SetNillableState sets the "State" field if the given value is not nil.
func (jpc *JobPaymentsCreate) SetNillableState(s *string) *JobPaymentsCreate {
	if s != nil {
		jpc.SetState(*s)
	}
	return jpc
}

// SetTotalPayment sets the "TotalPayment" field.
func (jpc *JobPaymentsCreate) SetTotalPayment(f float64) *JobPaymentsCreate {
	jpc.mutation.SetTotalPayment(f)
	return jpc
}

// SetNillableTotalPayment sets the "TotalPayment" field if the given value is not nil.
func (jpc *JobPaymentsCreate) SetNillableTotalPayment(f *float64) *JobPaymentsCreate {
	if f != nil {
		jpc.SetTotalPayment(*f)
	}
	return jpc
}

// SetLevelRequest sets the "LevelRequest" field.
func (jpc *JobPaymentsCreate) SetLevelRequest(f float64) *JobPaymentsCreate {
	jpc.mutation.SetLevelRequest(f)
	return jpc
}

// SetNillableLevelRequest sets the "LevelRequest" field if the given value is not nil.
func (jpc *JobPaymentsCreate) SetNillableLevelRequest(f *float64) *JobPaymentsCreate {
	if f != nil {
		jpc.SetLevelRequest(*f)
	}
	return jpc
}

// SetLevelApprove sets the "LevelApprove" field.
func (jpc *JobPaymentsCreate) SetLevelApprove(f float64) *JobPaymentsCreate {
	jpc.mutation.SetLevelApprove(f)
	return jpc
}

// SetNillableLevelApprove sets the "LevelApprove" field if the given value is not nil.
func (jpc *JobPaymentsCreate) SetNillableLevelApprove(f *float64) *JobPaymentsCreate {
	if f != nil {
		jpc.SetLevelApprove(*f)
	}
	return jpc
}

// SetAmount sets the "Amount" field.
func (jpc *JobPaymentsCreate) SetAmount(f float64) *JobPaymentsCreate {
	jpc.mutation.SetAmount(f)
	return jpc
}

// SetNillableAmount sets the "Amount" field if the given value is not nil.
func (jpc *JobPaymentsCreate) SetNillableAmount(f *float64) *JobPaymentsCreate {
	if f != nil {
		jpc.SetAmount(*f)
	}
	return jpc
}

// SetCreatedAt sets the "CreatedAt" field.
func (jpc *JobPaymentsCreate) SetCreatedAt(t time.Time) *JobPaymentsCreate {
	jpc.mutation.SetCreatedAt(t)
	return jpc
}

// SetNillableCreatedAt sets the "CreatedAt" field if the given value is not nil.
func (jpc *JobPaymentsCreate) SetNillableCreatedAt(t *time.Time) *JobPaymentsCreate {
	if t != nil {
		jpc.SetCreatedAt(*t)
	}
	return jpc
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (jpc *JobPaymentsCreate) SetUpdatedAt(t time.Time) *JobPaymentsCreate {
	jpc.mutation.SetUpdatedAt(t)
	return jpc
}

// SetNillableUpdatedAt sets the "UpdatedAt" field if the given value is not nil.
func (jpc *JobPaymentsCreate) SetNillableUpdatedAt(t *time.Time) *JobPaymentsCreate {
	if t != nil {
		jpc.SetUpdatedAt(*t)
	}
	return jpc
}

// SetPaymentsID sets the "payments" edge to the JobRelations entity by ID.
func (jpc *JobPaymentsCreate) SetPaymentsID(id int) *JobPaymentsCreate {
	jpc.mutation.SetPaymentsID(id)
	return jpc
}

// SetNillablePaymentsID sets the "payments" edge to the JobRelations entity by ID if the given value is not nil.
func (jpc *JobPaymentsCreate) SetNillablePaymentsID(id *int) *JobPaymentsCreate {
	if id != nil {
		jpc = jpc.SetPaymentsID(*id)
	}
	return jpc
}

// SetPayments sets the "payments" edge to the JobRelations entity.
func (jpc *JobPaymentsCreate) SetPayments(j *JobRelations) *JobPaymentsCreate {
	return jpc.SetPaymentsID(j.ID)
}

// Mutation returns the JobPaymentsMutation object of the builder.
func (jpc *JobPaymentsCreate) Mutation() *JobPaymentsMutation {
	return jpc.mutation
}

// Save creates the JobPayments in the database.
func (jpc *JobPaymentsCreate) Save(ctx context.Context) (*JobPayments, error) {
	jpc.defaults()
	return withHooks(ctx, jpc.sqlSave, jpc.mutation, jpc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (jpc *JobPaymentsCreate) SaveX(ctx context.Context) *JobPayments {
	v, err := jpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (jpc *JobPaymentsCreate) Exec(ctx context.Context) error {
	_, err := jpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jpc *JobPaymentsCreate) ExecX(ctx context.Context) {
	if err := jpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (jpc *JobPaymentsCreate) defaults() {
	if _, ok := jpc.mutation.YibfNo(); !ok {
		v := jobpayments.DefaultYibfNo
		jpc.mutation.SetYibfNo(v)
	}
	if _, ok := jpc.mutation.PaymentNo(); !ok {
		v := jobpayments.DefaultPaymentNo
		jpc.mutation.SetPaymentNo(v)
	}
	if _, ok := jpc.mutation.PaymentDate(); !ok {
		v := jobpayments.DefaultPaymentDate()
		jpc.mutation.SetPaymentDate(v)
	}
	if _, ok := jpc.mutation.PaymentType(); !ok {
		v := jobpayments.DefaultPaymentType
		jpc.mutation.SetPaymentType(v)
	}
	if _, ok := jpc.mutation.State(); !ok {
		v := jobpayments.DefaultState
		jpc.mutation.SetState(v)
	}
	if _, ok := jpc.mutation.TotalPayment(); !ok {
		v := jobpayments.DefaultTotalPayment
		jpc.mutation.SetTotalPayment(v)
	}
	if _, ok := jpc.mutation.LevelRequest(); !ok {
		v := jobpayments.DefaultLevelRequest
		jpc.mutation.SetLevelRequest(v)
	}
	if _, ok := jpc.mutation.LevelApprove(); !ok {
		v := jobpayments.DefaultLevelApprove
		jpc.mutation.SetLevelApprove(v)
	}
	if _, ok := jpc.mutation.Amount(); !ok {
		v := jobpayments.DefaultAmount
		jpc.mutation.SetAmount(v)
	}
	if _, ok := jpc.mutation.CreatedAt(); !ok {
		v := jobpayments.DefaultCreatedAt()
		jpc.mutation.SetCreatedAt(v)
	}
	if _, ok := jpc.mutation.UpdatedAt(); !ok {
		v := jobpayments.DefaultUpdatedAt()
		jpc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (jpc *JobPaymentsCreate) check() error {
	if _, ok := jpc.mutation.YibfNo(); !ok {
		return &ValidationError{Name: "yibfNo", err: errors.New(`ent: missing required field "JobPayments.yibfNo"`)}
	}
	if _, ok := jpc.mutation.PaymentDate(); !ok {
		return &ValidationError{Name: "PaymentDate", err: errors.New(`ent: missing required field "JobPayments.PaymentDate"`)}
	}
	if _, ok := jpc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "CreatedAt", err: errors.New(`ent: missing required field "JobPayments.CreatedAt"`)}
	}
	if _, ok := jpc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "UpdatedAt", err: errors.New(`ent: missing required field "JobPayments.UpdatedAt"`)}
	}
	return nil
}

func (jpc *JobPaymentsCreate) sqlSave(ctx context.Context) (*JobPayments, error) {
	if err := jpc.check(); err != nil {
		return nil, err
	}
	_node, _spec := jpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, jpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	jpc.mutation.id = &_node.ID
	jpc.mutation.done = true
	return _node, nil
}

func (jpc *JobPaymentsCreate) createSpec() (*JobPayments, *sqlgraph.CreateSpec) {
	var (
		_node = &JobPayments{config: jpc.config}
		_spec = sqlgraph.NewCreateSpec(jobpayments.Table, sqlgraph.NewFieldSpec(jobpayments.FieldID, field.TypeInt))
	)
	if value, ok := jpc.mutation.YibfNo(); ok {
		_spec.SetField(jobpayments.FieldYibfNo, field.TypeInt, value)
		_node.YibfNo = value
	}
	if value, ok := jpc.mutation.PaymentNo(); ok {
		_spec.SetField(jobpayments.FieldPaymentNo, field.TypeInt, value)
		_node.PaymentNo = value
	}
	if value, ok := jpc.mutation.PaymentDate(); ok {
		_spec.SetField(jobpayments.FieldPaymentDate, field.TypeTime, value)
		_node.PaymentDate = value
	}
	if value, ok := jpc.mutation.PaymentType(); ok {
		_spec.SetField(jobpayments.FieldPaymentType, field.TypeString, value)
		_node.PaymentType = value
	}
	if value, ok := jpc.mutation.State(); ok {
		_spec.SetField(jobpayments.FieldState, field.TypeString, value)
		_node.State = value
	}
	if value, ok := jpc.mutation.TotalPayment(); ok {
		_spec.SetField(jobpayments.FieldTotalPayment, field.TypeFloat64, value)
		_node.TotalPayment = value
	}
	if value, ok := jpc.mutation.LevelRequest(); ok {
		_spec.SetField(jobpayments.FieldLevelRequest, field.TypeFloat64, value)
		_node.LevelRequest = value
	}
	if value, ok := jpc.mutation.LevelApprove(); ok {
		_spec.SetField(jobpayments.FieldLevelApprove, field.TypeFloat64, value)
		_node.LevelApprove = value
	}
	if value, ok := jpc.mutation.Amount(); ok {
		_spec.SetField(jobpayments.FieldAmount, field.TypeFloat64, value)
		_node.Amount = value
	}
	if value, ok := jpc.mutation.CreatedAt(); ok {
		_spec.SetField(jobpayments.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := jpc.mutation.UpdatedAt(); ok {
		_spec.SetField(jobpayments.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := jpc.mutation.PaymentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobpayments.PaymentsTable,
			Columns: []string{jobpayments.PaymentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(jobrelations.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.relations_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// JobPaymentsCreateBulk is the builder for creating many JobPayments entities in bulk.
type JobPaymentsCreateBulk struct {
	config
	err      error
	builders []*JobPaymentsCreate
}

// Save creates the JobPayments entities in the database.
func (jpcb *JobPaymentsCreateBulk) Save(ctx context.Context) ([]*JobPayments, error) {
	if jpcb.err != nil {
		return nil, jpcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(jpcb.builders))
	nodes := make([]*JobPayments, len(jpcb.builders))
	mutators := make([]Mutator, len(jpcb.builders))
	for i := range jpcb.builders {
		func(i int, root context.Context) {
			builder := jpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*JobPaymentsMutation)
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
					_, err = mutators[i+1].Mutate(root, jpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, jpcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, jpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (jpcb *JobPaymentsCreateBulk) SaveX(ctx context.Context) []*JobPayments {
	v, err := jpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (jpcb *JobPaymentsCreateBulk) Exec(ctx context.Context) error {
	_, err := jpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jpcb *JobPaymentsCreateBulk) ExecX(ctx context.Context) {
	if err := jpcb.Exec(ctx); err != nil {
		panic(err)
	}
}
