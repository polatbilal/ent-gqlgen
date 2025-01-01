// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/polatbilal/gqlgen-ent/ent/jobdetail"
	"github.com/polatbilal/gqlgen-ent/ent/jobowner"
	"github.com/polatbilal/gqlgen-ent/ent/predicate"
)

// JobOwnerUpdate is the builder for updating JobOwner entities.
type JobOwnerUpdate struct {
	config
	hooks    []Hook
	mutation *JobOwnerMutation
}

// Where appends a list predicates to the JobOwnerUpdate builder.
func (jou *JobOwnerUpdate) Where(ps ...predicate.JobOwner) *JobOwnerUpdate {
	jou.mutation.Where(ps...)
	return jou
}

// SetName sets the "Name" field.
func (jou *JobOwnerUpdate) SetName(s string) *JobOwnerUpdate {
	jou.mutation.SetName(s)
	return jou
}

// SetNillableName sets the "Name" field if the given value is not nil.
func (jou *JobOwnerUpdate) SetNillableName(s *string) *JobOwnerUpdate {
	if s != nil {
		jou.SetName(*s)
	}
	return jou
}

// SetTcNo sets the "TcNo" field.
func (jou *JobOwnerUpdate) SetTcNo(i int) *JobOwnerUpdate {
	jou.mutation.ResetTcNo()
	jou.mutation.SetTcNo(i)
	return jou
}

// SetNillableTcNo sets the "TcNo" field if the given value is not nil.
func (jou *JobOwnerUpdate) SetNillableTcNo(i *int) *JobOwnerUpdate {
	if i != nil {
		jou.SetTcNo(*i)
	}
	return jou
}

// AddTcNo adds i to the "TcNo" field.
func (jou *JobOwnerUpdate) AddTcNo(i int) *JobOwnerUpdate {
	jou.mutation.AddTcNo(i)
	return jou
}

// ClearTcNo clears the value of the "TcNo" field.
func (jou *JobOwnerUpdate) ClearTcNo() *JobOwnerUpdate {
	jou.mutation.ClearTcNo()
	return jou
}

// SetAddress sets the "Address" field.
func (jou *JobOwnerUpdate) SetAddress(s string) *JobOwnerUpdate {
	jou.mutation.SetAddress(s)
	return jou
}

// SetNillableAddress sets the "Address" field if the given value is not nil.
func (jou *JobOwnerUpdate) SetNillableAddress(s *string) *JobOwnerUpdate {
	if s != nil {
		jou.SetAddress(*s)
	}
	return jou
}

// ClearAddress clears the value of the "Address" field.
func (jou *JobOwnerUpdate) ClearAddress() *JobOwnerUpdate {
	jou.mutation.ClearAddress()
	return jou
}

// SetTaxAdmin sets the "TaxAdmin" field.
func (jou *JobOwnerUpdate) SetTaxAdmin(s string) *JobOwnerUpdate {
	jou.mutation.SetTaxAdmin(s)
	return jou
}

// SetNillableTaxAdmin sets the "TaxAdmin" field if the given value is not nil.
func (jou *JobOwnerUpdate) SetNillableTaxAdmin(s *string) *JobOwnerUpdate {
	if s != nil {
		jou.SetTaxAdmin(*s)
	}
	return jou
}

// ClearTaxAdmin clears the value of the "TaxAdmin" field.
func (jou *JobOwnerUpdate) ClearTaxAdmin() *JobOwnerUpdate {
	jou.mutation.ClearTaxAdmin()
	return jou
}

// SetTaxNo sets the "TaxNo" field.
func (jou *JobOwnerUpdate) SetTaxNo(i int) *JobOwnerUpdate {
	jou.mutation.ResetTaxNo()
	jou.mutation.SetTaxNo(i)
	return jou
}

// SetNillableTaxNo sets the "TaxNo" field if the given value is not nil.
func (jou *JobOwnerUpdate) SetNillableTaxNo(i *int) *JobOwnerUpdate {
	if i != nil {
		jou.SetTaxNo(*i)
	}
	return jou
}

// AddTaxNo adds i to the "TaxNo" field.
func (jou *JobOwnerUpdate) AddTaxNo(i int) *JobOwnerUpdate {
	jou.mutation.AddTaxNo(i)
	return jou
}

// ClearTaxNo clears the value of the "TaxNo" field.
func (jou *JobOwnerUpdate) ClearTaxNo() *JobOwnerUpdate {
	jou.mutation.ClearTaxNo()
	return jou
}

// SetPhone sets the "Phone" field.
func (jou *JobOwnerUpdate) SetPhone(s string) *JobOwnerUpdate {
	jou.mutation.SetPhone(s)
	return jou
}

// SetNillablePhone sets the "Phone" field if the given value is not nil.
func (jou *JobOwnerUpdate) SetNillablePhone(s *string) *JobOwnerUpdate {
	if s != nil {
		jou.SetPhone(*s)
	}
	return jou
}

// ClearPhone clears the value of the "Phone" field.
func (jou *JobOwnerUpdate) ClearPhone() *JobOwnerUpdate {
	jou.mutation.ClearPhone()
	return jou
}

// SetEmail sets the "Email" field.
func (jou *JobOwnerUpdate) SetEmail(s string) *JobOwnerUpdate {
	jou.mutation.SetEmail(s)
	return jou
}

// SetNillableEmail sets the "Email" field if the given value is not nil.
func (jou *JobOwnerUpdate) SetNillableEmail(s *string) *JobOwnerUpdate {
	if s != nil {
		jou.SetEmail(*s)
	}
	return jou
}

// ClearEmail clears the value of the "Email" field.
func (jou *JobOwnerUpdate) ClearEmail() *JobOwnerUpdate {
	jou.mutation.ClearEmail()
	return jou
}

// SetYdsID sets the "yds_id" field.
func (jou *JobOwnerUpdate) SetYdsID(i int) *JobOwnerUpdate {
	jou.mutation.ResetYdsID()
	jou.mutation.SetYdsID(i)
	return jou
}

// SetNillableYdsID sets the "yds_id" field if the given value is not nil.
func (jou *JobOwnerUpdate) SetNillableYdsID(i *int) *JobOwnerUpdate {
	if i != nil {
		jou.SetYdsID(*i)
	}
	return jou
}

// AddYdsID adds i to the "yds_id" field.
func (jou *JobOwnerUpdate) AddYdsID(i int) *JobOwnerUpdate {
	jou.mutation.AddYdsID(i)
	return jou
}

// ClearYdsID clears the value of the "yds_id" field.
func (jou *JobOwnerUpdate) ClearYdsID() *JobOwnerUpdate {
	jou.mutation.ClearYdsID()
	return jou
}

// SetNote sets the "Note" field.
func (jou *JobOwnerUpdate) SetNote(s string) *JobOwnerUpdate {
	jou.mutation.SetNote(s)
	return jou
}

// SetNillableNote sets the "Note" field if the given value is not nil.
func (jou *JobOwnerUpdate) SetNillableNote(s *string) *JobOwnerUpdate {
	if s != nil {
		jou.SetNote(*s)
	}
	return jou
}

// ClearNote clears the value of the "Note" field.
func (jou *JobOwnerUpdate) ClearNote() *JobOwnerUpdate {
	jou.mutation.ClearNote()
	return jou
}

// SetCreatedAt sets the "CreatedAt" field.
func (jou *JobOwnerUpdate) SetCreatedAt(t time.Time) *JobOwnerUpdate {
	jou.mutation.SetCreatedAt(t)
	return jou
}

// SetNillableCreatedAt sets the "CreatedAt" field if the given value is not nil.
func (jou *JobOwnerUpdate) SetNillableCreatedAt(t *time.Time) *JobOwnerUpdate {
	if t != nil {
		jou.SetCreatedAt(*t)
	}
	return jou
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (jou *JobOwnerUpdate) SetUpdatedAt(t time.Time) *JobOwnerUpdate {
	jou.mutation.SetUpdatedAt(t)
	return jou
}

// AddOwnerIDs adds the "owners" edge to the JobDetail entity by IDs.
func (jou *JobOwnerUpdate) AddOwnerIDs(ids ...int) *JobOwnerUpdate {
	jou.mutation.AddOwnerIDs(ids...)
	return jou
}

// AddOwners adds the "owners" edges to the JobDetail entity.
func (jou *JobOwnerUpdate) AddOwners(j ...*JobDetail) *JobOwnerUpdate {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return jou.AddOwnerIDs(ids...)
}

// Mutation returns the JobOwnerMutation object of the builder.
func (jou *JobOwnerUpdate) Mutation() *JobOwnerMutation {
	return jou.mutation
}

// ClearOwners clears all "owners" edges to the JobDetail entity.
func (jou *JobOwnerUpdate) ClearOwners() *JobOwnerUpdate {
	jou.mutation.ClearOwners()
	return jou
}

// RemoveOwnerIDs removes the "owners" edge to JobDetail entities by IDs.
func (jou *JobOwnerUpdate) RemoveOwnerIDs(ids ...int) *JobOwnerUpdate {
	jou.mutation.RemoveOwnerIDs(ids...)
	return jou
}

// RemoveOwners removes "owners" edges to JobDetail entities.
func (jou *JobOwnerUpdate) RemoveOwners(j ...*JobDetail) *JobOwnerUpdate {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return jou.RemoveOwnerIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (jou *JobOwnerUpdate) Save(ctx context.Context) (int, error) {
	jou.defaults()
	return withHooks(ctx, jou.sqlSave, jou.mutation, jou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (jou *JobOwnerUpdate) SaveX(ctx context.Context) int {
	affected, err := jou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (jou *JobOwnerUpdate) Exec(ctx context.Context) error {
	_, err := jou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jou *JobOwnerUpdate) ExecX(ctx context.Context) {
	if err := jou.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (jou *JobOwnerUpdate) defaults() {
	if _, ok := jou.mutation.UpdatedAt(); !ok {
		v := jobowner.UpdateDefaultUpdatedAt()
		jou.mutation.SetUpdatedAt(v)
	}
}

func (jou *JobOwnerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(jobowner.Table, jobowner.Columns, sqlgraph.NewFieldSpec(jobowner.FieldID, field.TypeInt))
	if ps := jou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := jou.mutation.Name(); ok {
		_spec.SetField(jobowner.FieldName, field.TypeString, value)
	}
	if value, ok := jou.mutation.TcNo(); ok {
		_spec.SetField(jobowner.FieldTcNo, field.TypeInt, value)
	}
	if value, ok := jou.mutation.AddedTcNo(); ok {
		_spec.AddField(jobowner.FieldTcNo, field.TypeInt, value)
	}
	if jou.mutation.TcNoCleared() {
		_spec.ClearField(jobowner.FieldTcNo, field.TypeInt)
	}
	if value, ok := jou.mutation.Address(); ok {
		_spec.SetField(jobowner.FieldAddress, field.TypeString, value)
	}
	if jou.mutation.AddressCleared() {
		_spec.ClearField(jobowner.FieldAddress, field.TypeString)
	}
	if value, ok := jou.mutation.TaxAdmin(); ok {
		_spec.SetField(jobowner.FieldTaxAdmin, field.TypeString, value)
	}
	if jou.mutation.TaxAdminCleared() {
		_spec.ClearField(jobowner.FieldTaxAdmin, field.TypeString)
	}
	if value, ok := jou.mutation.TaxNo(); ok {
		_spec.SetField(jobowner.FieldTaxNo, field.TypeInt, value)
	}
	if value, ok := jou.mutation.AddedTaxNo(); ok {
		_spec.AddField(jobowner.FieldTaxNo, field.TypeInt, value)
	}
	if jou.mutation.TaxNoCleared() {
		_spec.ClearField(jobowner.FieldTaxNo, field.TypeInt)
	}
	if value, ok := jou.mutation.Phone(); ok {
		_spec.SetField(jobowner.FieldPhone, field.TypeString, value)
	}
	if jou.mutation.PhoneCleared() {
		_spec.ClearField(jobowner.FieldPhone, field.TypeString)
	}
	if value, ok := jou.mutation.Email(); ok {
		_spec.SetField(jobowner.FieldEmail, field.TypeString, value)
	}
	if jou.mutation.EmailCleared() {
		_spec.ClearField(jobowner.FieldEmail, field.TypeString)
	}
	if value, ok := jou.mutation.YdsID(); ok {
		_spec.SetField(jobowner.FieldYdsID, field.TypeInt, value)
	}
	if value, ok := jou.mutation.AddedYdsID(); ok {
		_spec.AddField(jobowner.FieldYdsID, field.TypeInt, value)
	}
	if jou.mutation.YdsIDCleared() {
		_spec.ClearField(jobowner.FieldYdsID, field.TypeInt)
	}
	if value, ok := jou.mutation.Note(); ok {
		_spec.SetField(jobowner.FieldNote, field.TypeString, value)
	}
	if jou.mutation.NoteCleared() {
		_spec.ClearField(jobowner.FieldNote, field.TypeString)
	}
	if value, ok := jou.mutation.CreatedAt(); ok {
		_spec.SetField(jobowner.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := jou.mutation.UpdatedAt(); ok {
		_spec.SetField(jobowner.FieldUpdatedAt, field.TypeTime, value)
	}
	if jou.mutation.OwnersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   jobowner.OwnersTable,
			Columns: []string{jobowner.OwnersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(jobdetail.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jou.mutation.RemovedOwnersIDs(); len(nodes) > 0 && !jou.mutation.OwnersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   jobowner.OwnersTable,
			Columns: []string{jobowner.OwnersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(jobdetail.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jou.mutation.OwnersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   jobowner.OwnersTable,
			Columns: []string{jobowner.OwnersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(jobdetail.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, jou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{jobowner.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	jou.mutation.done = true
	return n, nil
}

// JobOwnerUpdateOne is the builder for updating a single JobOwner entity.
type JobOwnerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *JobOwnerMutation
}

// SetName sets the "Name" field.
func (jouo *JobOwnerUpdateOne) SetName(s string) *JobOwnerUpdateOne {
	jouo.mutation.SetName(s)
	return jouo
}

// SetNillableName sets the "Name" field if the given value is not nil.
func (jouo *JobOwnerUpdateOne) SetNillableName(s *string) *JobOwnerUpdateOne {
	if s != nil {
		jouo.SetName(*s)
	}
	return jouo
}

// SetTcNo sets the "TcNo" field.
func (jouo *JobOwnerUpdateOne) SetTcNo(i int) *JobOwnerUpdateOne {
	jouo.mutation.ResetTcNo()
	jouo.mutation.SetTcNo(i)
	return jouo
}

// SetNillableTcNo sets the "TcNo" field if the given value is not nil.
func (jouo *JobOwnerUpdateOne) SetNillableTcNo(i *int) *JobOwnerUpdateOne {
	if i != nil {
		jouo.SetTcNo(*i)
	}
	return jouo
}

// AddTcNo adds i to the "TcNo" field.
func (jouo *JobOwnerUpdateOne) AddTcNo(i int) *JobOwnerUpdateOne {
	jouo.mutation.AddTcNo(i)
	return jouo
}

// ClearTcNo clears the value of the "TcNo" field.
func (jouo *JobOwnerUpdateOne) ClearTcNo() *JobOwnerUpdateOne {
	jouo.mutation.ClearTcNo()
	return jouo
}

// SetAddress sets the "Address" field.
func (jouo *JobOwnerUpdateOne) SetAddress(s string) *JobOwnerUpdateOne {
	jouo.mutation.SetAddress(s)
	return jouo
}

// SetNillableAddress sets the "Address" field if the given value is not nil.
func (jouo *JobOwnerUpdateOne) SetNillableAddress(s *string) *JobOwnerUpdateOne {
	if s != nil {
		jouo.SetAddress(*s)
	}
	return jouo
}

// ClearAddress clears the value of the "Address" field.
func (jouo *JobOwnerUpdateOne) ClearAddress() *JobOwnerUpdateOne {
	jouo.mutation.ClearAddress()
	return jouo
}

// SetTaxAdmin sets the "TaxAdmin" field.
func (jouo *JobOwnerUpdateOne) SetTaxAdmin(s string) *JobOwnerUpdateOne {
	jouo.mutation.SetTaxAdmin(s)
	return jouo
}

// SetNillableTaxAdmin sets the "TaxAdmin" field if the given value is not nil.
func (jouo *JobOwnerUpdateOne) SetNillableTaxAdmin(s *string) *JobOwnerUpdateOne {
	if s != nil {
		jouo.SetTaxAdmin(*s)
	}
	return jouo
}

// ClearTaxAdmin clears the value of the "TaxAdmin" field.
func (jouo *JobOwnerUpdateOne) ClearTaxAdmin() *JobOwnerUpdateOne {
	jouo.mutation.ClearTaxAdmin()
	return jouo
}

// SetTaxNo sets the "TaxNo" field.
func (jouo *JobOwnerUpdateOne) SetTaxNo(i int) *JobOwnerUpdateOne {
	jouo.mutation.ResetTaxNo()
	jouo.mutation.SetTaxNo(i)
	return jouo
}

// SetNillableTaxNo sets the "TaxNo" field if the given value is not nil.
func (jouo *JobOwnerUpdateOne) SetNillableTaxNo(i *int) *JobOwnerUpdateOne {
	if i != nil {
		jouo.SetTaxNo(*i)
	}
	return jouo
}

// AddTaxNo adds i to the "TaxNo" field.
func (jouo *JobOwnerUpdateOne) AddTaxNo(i int) *JobOwnerUpdateOne {
	jouo.mutation.AddTaxNo(i)
	return jouo
}

// ClearTaxNo clears the value of the "TaxNo" field.
func (jouo *JobOwnerUpdateOne) ClearTaxNo() *JobOwnerUpdateOne {
	jouo.mutation.ClearTaxNo()
	return jouo
}

// SetPhone sets the "Phone" field.
func (jouo *JobOwnerUpdateOne) SetPhone(s string) *JobOwnerUpdateOne {
	jouo.mutation.SetPhone(s)
	return jouo
}

// SetNillablePhone sets the "Phone" field if the given value is not nil.
func (jouo *JobOwnerUpdateOne) SetNillablePhone(s *string) *JobOwnerUpdateOne {
	if s != nil {
		jouo.SetPhone(*s)
	}
	return jouo
}

// ClearPhone clears the value of the "Phone" field.
func (jouo *JobOwnerUpdateOne) ClearPhone() *JobOwnerUpdateOne {
	jouo.mutation.ClearPhone()
	return jouo
}

// SetEmail sets the "Email" field.
func (jouo *JobOwnerUpdateOne) SetEmail(s string) *JobOwnerUpdateOne {
	jouo.mutation.SetEmail(s)
	return jouo
}

// SetNillableEmail sets the "Email" field if the given value is not nil.
func (jouo *JobOwnerUpdateOne) SetNillableEmail(s *string) *JobOwnerUpdateOne {
	if s != nil {
		jouo.SetEmail(*s)
	}
	return jouo
}

// ClearEmail clears the value of the "Email" field.
func (jouo *JobOwnerUpdateOne) ClearEmail() *JobOwnerUpdateOne {
	jouo.mutation.ClearEmail()
	return jouo
}

// SetYdsID sets the "yds_id" field.
func (jouo *JobOwnerUpdateOne) SetYdsID(i int) *JobOwnerUpdateOne {
	jouo.mutation.ResetYdsID()
	jouo.mutation.SetYdsID(i)
	return jouo
}

// SetNillableYdsID sets the "yds_id" field if the given value is not nil.
func (jouo *JobOwnerUpdateOne) SetNillableYdsID(i *int) *JobOwnerUpdateOne {
	if i != nil {
		jouo.SetYdsID(*i)
	}
	return jouo
}

// AddYdsID adds i to the "yds_id" field.
func (jouo *JobOwnerUpdateOne) AddYdsID(i int) *JobOwnerUpdateOne {
	jouo.mutation.AddYdsID(i)
	return jouo
}

// ClearYdsID clears the value of the "yds_id" field.
func (jouo *JobOwnerUpdateOne) ClearYdsID() *JobOwnerUpdateOne {
	jouo.mutation.ClearYdsID()
	return jouo
}

// SetNote sets the "Note" field.
func (jouo *JobOwnerUpdateOne) SetNote(s string) *JobOwnerUpdateOne {
	jouo.mutation.SetNote(s)
	return jouo
}

// SetNillableNote sets the "Note" field if the given value is not nil.
func (jouo *JobOwnerUpdateOne) SetNillableNote(s *string) *JobOwnerUpdateOne {
	if s != nil {
		jouo.SetNote(*s)
	}
	return jouo
}

// ClearNote clears the value of the "Note" field.
func (jouo *JobOwnerUpdateOne) ClearNote() *JobOwnerUpdateOne {
	jouo.mutation.ClearNote()
	return jouo
}

// SetCreatedAt sets the "CreatedAt" field.
func (jouo *JobOwnerUpdateOne) SetCreatedAt(t time.Time) *JobOwnerUpdateOne {
	jouo.mutation.SetCreatedAt(t)
	return jouo
}

// SetNillableCreatedAt sets the "CreatedAt" field if the given value is not nil.
func (jouo *JobOwnerUpdateOne) SetNillableCreatedAt(t *time.Time) *JobOwnerUpdateOne {
	if t != nil {
		jouo.SetCreatedAt(*t)
	}
	return jouo
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (jouo *JobOwnerUpdateOne) SetUpdatedAt(t time.Time) *JobOwnerUpdateOne {
	jouo.mutation.SetUpdatedAt(t)
	return jouo
}

// AddOwnerIDs adds the "owners" edge to the JobDetail entity by IDs.
func (jouo *JobOwnerUpdateOne) AddOwnerIDs(ids ...int) *JobOwnerUpdateOne {
	jouo.mutation.AddOwnerIDs(ids...)
	return jouo
}

// AddOwners adds the "owners" edges to the JobDetail entity.
func (jouo *JobOwnerUpdateOne) AddOwners(j ...*JobDetail) *JobOwnerUpdateOne {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return jouo.AddOwnerIDs(ids...)
}

// Mutation returns the JobOwnerMutation object of the builder.
func (jouo *JobOwnerUpdateOne) Mutation() *JobOwnerMutation {
	return jouo.mutation
}

// ClearOwners clears all "owners" edges to the JobDetail entity.
func (jouo *JobOwnerUpdateOne) ClearOwners() *JobOwnerUpdateOne {
	jouo.mutation.ClearOwners()
	return jouo
}

// RemoveOwnerIDs removes the "owners" edge to JobDetail entities by IDs.
func (jouo *JobOwnerUpdateOne) RemoveOwnerIDs(ids ...int) *JobOwnerUpdateOne {
	jouo.mutation.RemoveOwnerIDs(ids...)
	return jouo
}

// RemoveOwners removes "owners" edges to JobDetail entities.
func (jouo *JobOwnerUpdateOne) RemoveOwners(j ...*JobDetail) *JobOwnerUpdateOne {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return jouo.RemoveOwnerIDs(ids...)
}

// Where appends a list predicates to the JobOwnerUpdate builder.
func (jouo *JobOwnerUpdateOne) Where(ps ...predicate.JobOwner) *JobOwnerUpdateOne {
	jouo.mutation.Where(ps...)
	return jouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (jouo *JobOwnerUpdateOne) Select(field string, fields ...string) *JobOwnerUpdateOne {
	jouo.fields = append([]string{field}, fields...)
	return jouo
}

// Save executes the query and returns the updated JobOwner entity.
func (jouo *JobOwnerUpdateOne) Save(ctx context.Context) (*JobOwner, error) {
	jouo.defaults()
	return withHooks(ctx, jouo.sqlSave, jouo.mutation, jouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (jouo *JobOwnerUpdateOne) SaveX(ctx context.Context) *JobOwner {
	node, err := jouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (jouo *JobOwnerUpdateOne) Exec(ctx context.Context) error {
	_, err := jouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jouo *JobOwnerUpdateOne) ExecX(ctx context.Context) {
	if err := jouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (jouo *JobOwnerUpdateOne) defaults() {
	if _, ok := jouo.mutation.UpdatedAt(); !ok {
		v := jobowner.UpdateDefaultUpdatedAt()
		jouo.mutation.SetUpdatedAt(v)
	}
}

func (jouo *JobOwnerUpdateOne) sqlSave(ctx context.Context) (_node *JobOwner, err error) {
	_spec := sqlgraph.NewUpdateSpec(jobowner.Table, jobowner.Columns, sqlgraph.NewFieldSpec(jobowner.FieldID, field.TypeInt))
	id, ok := jouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "JobOwner.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := jouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, jobowner.FieldID)
		for _, f := range fields {
			if !jobowner.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != jobowner.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := jouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := jouo.mutation.Name(); ok {
		_spec.SetField(jobowner.FieldName, field.TypeString, value)
	}
	if value, ok := jouo.mutation.TcNo(); ok {
		_spec.SetField(jobowner.FieldTcNo, field.TypeInt, value)
	}
	if value, ok := jouo.mutation.AddedTcNo(); ok {
		_spec.AddField(jobowner.FieldTcNo, field.TypeInt, value)
	}
	if jouo.mutation.TcNoCleared() {
		_spec.ClearField(jobowner.FieldTcNo, field.TypeInt)
	}
	if value, ok := jouo.mutation.Address(); ok {
		_spec.SetField(jobowner.FieldAddress, field.TypeString, value)
	}
	if jouo.mutation.AddressCleared() {
		_spec.ClearField(jobowner.FieldAddress, field.TypeString)
	}
	if value, ok := jouo.mutation.TaxAdmin(); ok {
		_spec.SetField(jobowner.FieldTaxAdmin, field.TypeString, value)
	}
	if jouo.mutation.TaxAdminCleared() {
		_spec.ClearField(jobowner.FieldTaxAdmin, field.TypeString)
	}
	if value, ok := jouo.mutation.TaxNo(); ok {
		_spec.SetField(jobowner.FieldTaxNo, field.TypeInt, value)
	}
	if value, ok := jouo.mutation.AddedTaxNo(); ok {
		_spec.AddField(jobowner.FieldTaxNo, field.TypeInt, value)
	}
	if jouo.mutation.TaxNoCleared() {
		_spec.ClearField(jobowner.FieldTaxNo, field.TypeInt)
	}
	if value, ok := jouo.mutation.Phone(); ok {
		_spec.SetField(jobowner.FieldPhone, field.TypeString, value)
	}
	if jouo.mutation.PhoneCleared() {
		_spec.ClearField(jobowner.FieldPhone, field.TypeString)
	}
	if value, ok := jouo.mutation.Email(); ok {
		_spec.SetField(jobowner.FieldEmail, field.TypeString, value)
	}
	if jouo.mutation.EmailCleared() {
		_spec.ClearField(jobowner.FieldEmail, field.TypeString)
	}
	if value, ok := jouo.mutation.YdsID(); ok {
		_spec.SetField(jobowner.FieldYdsID, field.TypeInt, value)
	}
	if value, ok := jouo.mutation.AddedYdsID(); ok {
		_spec.AddField(jobowner.FieldYdsID, field.TypeInt, value)
	}
	if jouo.mutation.YdsIDCleared() {
		_spec.ClearField(jobowner.FieldYdsID, field.TypeInt)
	}
	if value, ok := jouo.mutation.Note(); ok {
		_spec.SetField(jobowner.FieldNote, field.TypeString, value)
	}
	if jouo.mutation.NoteCleared() {
		_spec.ClearField(jobowner.FieldNote, field.TypeString)
	}
	if value, ok := jouo.mutation.CreatedAt(); ok {
		_spec.SetField(jobowner.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := jouo.mutation.UpdatedAt(); ok {
		_spec.SetField(jobowner.FieldUpdatedAt, field.TypeTime, value)
	}
	if jouo.mutation.OwnersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   jobowner.OwnersTable,
			Columns: []string{jobowner.OwnersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(jobdetail.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jouo.mutation.RemovedOwnersIDs(); len(nodes) > 0 && !jouo.mutation.OwnersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   jobowner.OwnersTable,
			Columns: []string{jobowner.OwnersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(jobdetail.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jouo.mutation.OwnersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   jobowner.OwnersTable,
			Columns: []string{jobowner.OwnersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(jobdetail.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &JobOwner{config: jouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, jouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{jobowner.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	jouo.mutation.done = true
	return _node, nil
}
