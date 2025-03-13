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
	"github.com/polatbilal/gqlgen-ent/ent/jobcontractor"
	"github.com/polatbilal/gqlgen-ent/ent/jobrelations"
	"github.com/polatbilal/gqlgen-ent/ent/predicate"
)

// JobContractorUpdate is the builder for updating JobContractor entities.
type JobContractorUpdate struct {
	config
	hooks    []Hook
	mutation *JobContractorMutation
}

// Where appends a list predicates to the JobContractorUpdate builder.
func (jcu *JobContractorUpdate) Where(ps ...predicate.JobContractor) *JobContractorUpdate {
	jcu.mutation.Where(ps...)
	return jcu
}

// SetName sets the "Name" field.
func (jcu *JobContractorUpdate) SetName(s string) *JobContractorUpdate {
	jcu.mutation.SetName(s)
	return jcu
}

// SetNillableName sets the "Name" field if the given value is not nil.
func (jcu *JobContractorUpdate) SetNillableName(s *string) *JobContractorUpdate {
	if s != nil {
		jcu.SetName(*s)
	}
	return jcu
}

// SetTcNo sets the "TcNo" field.
func (jcu *JobContractorUpdate) SetTcNo(s string) *JobContractorUpdate {
	jcu.mutation.SetTcNo(s)
	return jcu
}

// SetNillableTcNo sets the "TcNo" field if the given value is not nil.
func (jcu *JobContractorUpdate) SetNillableTcNo(s *string) *JobContractorUpdate {
	if s != nil {
		jcu.SetTcNo(*s)
	}
	return jcu
}

// ClearTcNo clears the value of the "TcNo" field.
func (jcu *JobContractorUpdate) ClearTcNo() *JobContractorUpdate {
	jcu.mutation.ClearTcNo()
	return jcu
}

// SetRegisterNo sets the "RegisterNo" field.
func (jcu *JobContractorUpdate) SetRegisterNo(s string) *JobContractorUpdate {
	jcu.mutation.SetRegisterNo(s)
	return jcu
}

// SetNillableRegisterNo sets the "RegisterNo" field if the given value is not nil.
func (jcu *JobContractorUpdate) SetNillableRegisterNo(s *string) *JobContractorUpdate {
	if s != nil {
		jcu.SetRegisterNo(*s)
	}
	return jcu
}

// ClearRegisterNo clears the value of the "RegisterNo" field.
func (jcu *JobContractorUpdate) ClearRegisterNo() *JobContractorUpdate {
	jcu.mutation.ClearRegisterNo()
	return jcu
}

// SetAddress sets the "Address" field.
func (jcu *JobContractorUpdate) SetAddress(s string) *JobContractorUpdate {
	jcu.mutation.SetAddress(s)
	return jcu
}

// SetNillableAddress sets the "Address" field if the given value is not nil.
func (jcu *JobContractorUpdate) SetNillableAddress(s *string) *JobContractorUpdate {
	if s != nil {
		jcu.SetAddress(*s)
	}
	return jcu
}

// ClearAddress clears the value of the "Address" field.
func (jcu *JobContractorUpdate) ClearAddress() *JobContractorUpdate {
	jcu.mutation.ClearAddress()
	return jcu
}

// SetTaxNo sets the "TaxNo" field.
func (jcu *JobContractorUpdate) SetTaxNo(s string) *JobContractorUpdate {
	jcu.mutation.SetTaxNo(s)
	return jcu
}

// SetNillableTaxNo sets the "TaxNo" field if the given value is not nil.
func (jcu *JobContractorUpdate) SetNillableTaxNo(s *string) *JobContractorUpdate {
	if s != nil {
		jcu.SetTaxNo(*s)
	}
	return jcu
}

// ClearTaxNo clears the value of the "TaxNo" field.
func (jcu *JobContractorUpdate) ClearTaxNo() *JobContractorUpdate {
	jcu.mutation.ClearTaxNo()
	return jcu
}

// SetMobilePhone sets the "MobilePhone" field.
func (jcu *JobContractorUpdate) SetMobilePhone(s string) *JobContractorUpdate {
	jcu.mutation.SetMobilePhone(s)
	return jcu
}

// SetNillableMobilePhone sets the "MobilePhone" field if the given value is not nil.
func (jcu *JobContractorUpdate) SetNillableMobilePhone(s *string) *JobContractorUpdate {
	if s != nil {
		jcu.SetMobilePhone(*s)
	}
	return jcu
}

// ClearMobilePhone clears the value of the "MobilePhone" field.
func (jcu *JobContractorUpdate) ClearMobilePhone() *JobContractorUpdate {
	jcu.mutation.ClearMobilePhone()
	return jcu
}

// SetPhone sets the "Phone" field.
func (jcu *JobContractorUpdate) SetPhone(s string) *JobContractorUpdate {
	jcu.mutation.SetPhone(s)
	return jcu
}

// SetNillablePhone sets the "Phone" field if the given value is not nil.
func (jcu *JobContractorUpdate) SetNillablePhone(s *string) *JobContractorUpdate {
	if s != nil {
		jcu.SetPhone(*s)
	}
	return jcu
}

// ClearPhone clears the value of the "Phone" field.
func (jcu *JobContractorUpdate) ClearPhone() *JobContractorUpdate {
	jcu.mutation.ClearPhone()
	return jcu
}

// SetEmail sets the "Email" field.
func (jcu *JobContractorUpdate) SetEmail(s string) *JobContractorUpdate {
	jcu.mutation.SetEmail(s)
	return jcu
}

// SetNillableEmail sets the "Email" field if the given value is not nil.
func (jcu *JobContractorUpdate) SetNillableEmail(s *string) *JobContractorUpdate {
	if s != nil {
		jcu.SetEmail(*s)
	}
	return jcu
}

// ClearEmail clears the value of the "Email" field.
func (jcu *JobContractorUpdate) ClearEmail() *JobContractorUpdate {
	jcu.mutation.ClearEmail()
	return jcu
}

// SetPersonType sets the "PersonType" field.
func (jcu *JobContractorUpdate) SetPersonType(s string) *JobContractorUpdate {
	jcu.mutation.SetPersonType(s)
	return jcu
}

// SetNillablePersonType sets the "PersonType" field if the given value is not nil.
func (jcu *JobContractorUpdate) SetNillablePersonType(s *string) *JobContractorUpdate {
	if s != nil {
		jcu.SetPersonType(*s)
	}
	return jcu
}

// ClearPersonType clears the value of the "PersonType" field.
func (jcu *JobContractorUpdate) ClearPersonType() *JobContractorUpdate {
	jcu.mutation.ClearPersonType()
	return jcu
}

// SetYDSID sets the "YDSID" field.
func (jcu *JobContractorUpdate) SetYDSID(i int) *JobContractorUpdate {
	jcu.mutation.ResetYDSID()
	jcu.mutation.SetYDSID(i)
	return jcu
}

// SetNillableYDSID sets the "YDSID" field if the given value is not nil.
func (jcu *JobContractorUpdate) SetNillableYDSID(i *int) *JobContractorUpdate {
	if i != nil {
		jcu.SetYDSID(*i)
	}
	return jcu
}

// AddYDSID adds i to the "YDSID" field.
func (jcu *JobContractorUpdate) AddYDSID(i int) *JobContractorUpdate {
	jcu.mutation.AddYDSID(i)
	return jcu
}

// ClearYDSID clears the value of the "YDSID" field.
func (jcu *JobContractorUpdate) ClearYDSID() *JobContractorUpdate {
	jcu.mutation.ClearYDSID()
	return jcu
}

// SetNote sets the "Note" field.
func (jcu *JobContractorUpdate) SetNote(s string) *JobContractorUpdate {
	jcu.mutation.SetNote(s)
	return jcu
}

// SetNillableNote sets the "Note" field if the given value is not nil.
func (jcu *JobContractorUpdate) SetNillableNote(s *string) *JobContractorUpdate {
	if s != nil {
		jcu.SetNote(*s)
	}
	return jcu
}

// ClearNote clears the value of the "Note" field.
func (jcu *JobContractorUpdate) ClearNote() *JobContractorUpdate {
	jcu.mutation.ClearNote()
	return jcu
}

// SetCreatedAt sets the "CreatedAt" field.
func (jcu *JobContractorUpdate) SetCreatedAt(t time.Time) *JobContractorUpdate {
	jcu.mutation.SetCreatedAt(t)
	return jcu
}

// SetNillableCreatedAt sets the "CreatedAt" field if the given value is not nil.
func (jcu *JobContractorUpdate) SetNillableCreatedAt(t *time.Time) *JobContractorUpdate {
	if t != nil {
		jcu.SetCreatedAt(*t)
	}
	return jcu
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (jcu *JobContractorUpdate) SetUpdatedAt(t time.Time) *JobContractorUpdate {
	jcu.mutation.SetUpdatedAt(t)
	return jcu
}

// AddContractorIDs adds the "contractors" edge to the JobRelations entity by IDs.
func (jcu *JobContractorUpdate) AddContractorIDs(ids ...int) *JobContractorUpdate {
	jcu.mutation.AddContractorIDs(ids...)
	return jcu
}

// AddContractors adds the "contractors" edges to the JobRelations entity.
func (jcu *JobContractorUpdate) AddContractors(j ...*JobRelations) *JobContractorUpdate {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return jcu.AddContractorIDs(ids...)
}

// Mutation returns the JobContractorMutation object of the builder.
func (jcu *JobContractorUpdate) Mutation() *JobContractorMutation {
	return jcu.mutation
}

// ClearContractors clears all "contractors" edges to the JobRelations entity.
func (jcu *JobContractorUpdate) ClearContractors() *JobContractorUpdate {
	jcu.mutation.ClearContractors()
	return jcu
}

// RemoveContractorIDs removes the "contractors" edge to JobRelations entities by IDs.
func (jcu *JobContractorUpdate) RemoveContractorIDs(ids ...int) *JobContractorUpdate {
	jcu.mutation.RemoveContractorIDs(ids...)
	return jcu
}

// RemoveContractors removes "contractors" edges to JobRelations entities.
func (jcu *JobContractorUpdate) RemoveContractors(j ...*JobRelations) *JobContractorUpdate {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return jcu.RemoveContractorIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (jcu *JobContractorUpdate) Save(ctx context.Context) (int, error) {
	jcu.defaults()
	return withHooks(ctx, jcu.sqlSave, jcu.mutation, jcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (jcu *JobContractorUpdate) SaveX(ctx context.Context) int {
	affected, err := jcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (jcu *JobContractorUpdate) Exec(ctx context.Context) error {
	_, err := jcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jcu *JobContractorUpdate) ExecX(ctx context.Context) {
	if err := jcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (jcu *JobContractorUpdate) defaults() {
	if _, ok := jcu.mutation.UpdatedAt(); !ok {
		v := jobcontractor.UpdateDefaultUpdatedAt()
		jcu.mutation.SetUpdatedAt(v)
	}
}

func (jcu *JobContractorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(jobcontractor.Table, jobcontractor.Columns, sqlgraph.NewFieldSpec(jobcontractor.FieldID, field.TypeInt))
	if ps := jcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := jcu.mutation.Name(); ok {
		_spec.SetField(jobcontractor.FieldName, field.TypeString, value)
	}
	if value, ok := jcu.mutation.TcNo(); ok {
		_spec.SetField(jobcontractor.FieldTcNo, field.TypeString, value)
	}
	if jcu.mutation.TcNoCleared() {
		_spec.ClearField(jobcontractor.FieldTcNo, field.TypeString)
	}
	if value, ok := jcu.mutation.RegisterNo(); ok {
		_spec.SetField(jobcontractor.FieldRegisterNo, field.TypeString, value)
	}
	if jcu.mutation.RegisterNoCleared() {
		_spec.ClearField(jobcontractor.FieldRegisterNo, field.TypeString)
	}
	if value, ok := jcu.mutation.Address(); ok {
		_spec.SetField(jobcontractor.FieldAddress, field.TypeString, value)
	}
	if jcu.mutation.AddressCleared() {
		_spec.ClearField(jobcontractor.FieldAddress, field.TypeString)
	}
	if value, ok := jcu.mutation.TaxNo(); ok {
		_spec.SetField(jobcontractor.FieldTaxNo, field.TypeString, value)
	}
	if jcu.mutation.TaxNoCleared() {
		_spec.ClearField(jobcontractor.FieldTaxNo, field.TypeString)
	}
	if value, ok := jcu.mutation.MobilePhone(); ok {
		_spec.SetField(jobcontractor.FieldMobilePhone, field.TypeString, value)
	}
	if jcu.mutation.MobilePhoneCleared() {
		_spec.ClearField(jobcontractor.FieldMobilePhone, field.TypeString)
	}
	if value, ok := jcu.mutation.Phone(); ok {
		_spec.SetField(jobcontractor.FieldPhone, field.TypeString, value)
	}
	if jcu.mutation.PhoneCleared() {
		_spec.ClearField(jobcontractor.FieldPhone, field.TypeString)
	}
	if value, ok := jcu.mutation.Email(); ok {
		_spec.SetField(jobcontractor.FieldEmail, field.TypeString, value)
	}
	if jcu.mutation.EmailCleared() {
		_spec.ClearField(jobcontractor.FieldEmail, field.TypeString)
	}
	if value, ok := jcu.mutation.PersonType(); ok {
		_spec.SetField(jobcontractor.FieldPersonType, field.TypeString, value)
	}
	if jcu.mutation.PersonTypeCleared() {
		_spec.ClearField(jobcontractor.FieldPersonType, field.TypeString)
	}
	if value, ok := jcu.mutation.YDSID(); ok {
		_spec.SetField(jobcontractor.FieldYDSID, field.TypeInt, value)
	}
	if value, ok := jcu.mutation.AddedYDSID(); ok {
		_spec.AddField(jobcontractor.FieldYDSID, field.TypeInt, value)
	}
	if jcu.mutation.YDSIDCleared() {
		_spec.ClearField(jobcontractor.FieldYDSID, field.TypeInt)
	}
	if value, ok := jcu.mutation.Note(); ok {
		_spec.SetField(jobcontractor.FieldNote, field.TypeString, value)
	}
	if jcu.mutation.NoteCleared() {
		_spec.ClearField(jobcontractor.FieldNote, field.TypeString)
	}
	if value, ok := jcu.mutation.CreatedAt(); ok {
		_spec.SetField(jobcontractor.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := jcu.mutation.UpdatedAt(); ok {
		_spec.SetField(jobcontractor.FieldUpdatedAt, field.TypeTime, value)
	}
	if jcu.mutation.ContractorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   jobcontractor.ContractorsTable,
			Columns: []string{jobcontractor.ContractorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(jobrelations.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jcu.mutation.RemovedContractorsIDs(); len(nodes) > 0 && !jcu.mutation.ContractorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   jobcontractor.ContractorsTable,
			Columns: []string{jobcontractor.ContractorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(jobrelations.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jcu.mutation.ContractorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   jobcontractor.ContractorsTable,
			Columns: []string{jobcontractor.ContractorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(jobrelations.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, jcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{jobcontractor.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	jcu.mutation.done = true
	return n, nil
}

// JobContractorUpdateOne is the builder for updating a single JobContractor entity.
type JobContractorUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *JobContractorMutation
}

// SetName sets the "Name" field.
func (jcuo *JobContractorUpdateOne) SetName(s string) *JobContractorUpdateOne {
	jcuo.mutation.SetName(s)
	return jcuo
}

// SetNillableName sets the "Name" field if the given value is not nil.
func (jcuo *JobContractorUpdateOne) SetNillableName(s *string) *JobContractorUpdateOne {
	if s != nil {
		jcuo.SetName(*s)
	}
	return jcuo
}

// SetTcNo sets the "TcNo" field.
func (jcuo *JobContractorUpdateOne) SetTcNo(s string) *JobContractorUpdateOne {
	jcuo.mutation.SetTcNo(s)
	return jcuo
}

// SetNillableTcNo sets the "TcNo" field if the given value is not nil.
func (jcuo *JobContractorUpdateOne) SetNillableTcNo(s *string) *JobContractorUpdateOne {
	if s != nil {
		jcuo.SetTcNo(*s)
	}
	return jcuo
}

// ClearTcNo clears the value of the "TcNo" field.
func (jcuo *JobContractorUpdateOne) ClearTcNo() *JobContractorUpdateOne {
	jcuo.mutation.ClearTcNo()
	return jcuo
}

// SetRegisterNo sets the "RegisterNo" field.
func (jcuo *JobContractorUpdateOne) SetRegisterNo(s string) *JobContractorUpdateOne {
	jcuo.mutation.SetRegisterNo(s)
	return jcuo
}

// SetNillableRegisterNo sets the "RegisterNo" field if the given value is not nil.
func (jcuo *JobContractorUpdateOne) SetNillableRegisterNo(s *string) *JobContractorUpdateOne {
	if s != nil {
		jcuo.SetRegisterNo(*s)
	}
	return jcuo
}

// ClearRegisterNo clears the value of the "RegisterNo" field.
func (jcuo *JobContractorUpdateOne) ClearRegisterNo() *JobContractorUpdateOne {
	jcuo.mutation.ClearRegisterNo()
	return jcuo
}

// SetAddress sets the "Address" field.
func (jcuo *JobContractorUpdateOne) SetAddress(s string) *JobContractorUpdateOne {
	jcuo.mutation.SetAddress(s)
	return jcuo
}

// SetNillableAddress sets the "Address" field if the given value is not nil.
func (jcuo *JobContractorUpdateOne) SetNillableAddress(s *string) *JobContractorUpdateOne {
	if s != nil {
		jcuo.SetAddress(*s)
	}
	return jcuo
}

// ClearAddress clears the value of the "Address" field.
func (jcuo *JobContractorUpdateOne) ClearAddress() *JobContractorUpdateOne {
	jcuo.mutation.ClearAddress()
	return jcuo
}

// SetTaxNo sets the "TaxNo" field.
func (jcuo *JobContractorUpdateOne) SetTaxNo(s string) *JobContractorUpdateOne {
	jcuo.mutation.SetTaxNo(s)
	return jcuo
}

// SetNillableTaxNo sets the "TaxNo" field if the given value is not nil.
func (jcuo *JobContractorUpdateOne) SetNillableTaxNo(s *string) *JobContractorUpdateOne {
	if s != nil {
		jcuo.SetTaxNo(*s)
	}
	return jcuo
}

// ClearTaxNo clears the value of the "TaxNo" field.
func (jcuo *JobContractorUpdateOne) ClearTaxNo() *JobContractorUpdateOne {
	jcuo.mutation.ClearTaxNo()
	return jcuo
}

// SetMobilePhone sets the "MobilePhone" field.
func (jcuo *JobContractorUpdateOne) SetMobilePhone(s string) *JobContractorUpdateOne {
	jcuo.mutation.SetMobilePhone(s)
	return jcuo
}

// SetNillableMobilePhone sets the "MobilePhone" field if the given value is not nil.
func (jcuo *JobContractorUpdateOne) SetNillableMobilePhone(s *string) *JobContractorUpdateOne {
	if s != nil {
		jcuo.SetMobilePhone(*s)
	}
	return jcuo
}

// ClearMobilePhone clears the value of the "MobilePhone" field.
func (jcuo *JobContractorUpdateOne) ClearMobilePhone() *JobContractorUpdateOne {
	jcuo.mutation.ClearMobilePhone()
	return jcuo
}

// SetPhone sets the "Phone" field.
func (jcuo *JobContractorUpdateOne) SetPhone(s string) *JobContractorUpdateOne {
	jcuo.mutation.SetPhone(s)
	return jcuo
}

// SetNillablePhone sets the "Phone" field if the given value is not nil.
func (jcuo *JobContractorUpdateOne) SetNillablePhone(s *string) *JobContractorUpdateOne {
	if s != nil {
		jcuo.SetPhone(*s)
	}
	return jcuo
}

// ClearPhone clears the value of the "Phone" field.
func (jcuo *JobContractorUpdateOne) ClearPhone() *JobContractorUpdateOne {
	jcuo.mutation.ClearPhone()
	return jcuo
}

// SetEmail sets the "Email" field.
func (jcuo *JobContractorUpdateOne) SetEmail(s string) *JobContractorUpdateOne {
	jcuo.mutation.SetEmail(s)
	return jcuo
}

// SetNillableEmail sets the "Email" field if the given value is not nil.
func (jcuo *JobContractorUpdateOne) SetNillableEmail(s *string) *JobContractorUpdateOne {
	if s != nil {
		jcuo.SetEmail(*s)
	}
	return jcuo
}

// ClearEmail clears the value of the "Email" field.
func (jcuo *JobContractorUpdateOne) ClearEmail() *JobContractorUpdateOne {
	jcuo.mutation.ClearEmail()
	return jcuo
}

// SetPersonType sets the "PersonType" field.
func (jcuo *JobContractorUpdateOne) SetPersonType(s string) *JobContractorUpdateOne {
	jcuo.mutation.SetPersonType(s)
	return jcuo
}

// SetNillablePersonType sets the "PersonType" field if the given value is not nil.
func (jcuo *JobContractorUpdateOne) SetNillablePersonType(s *string) *JobContractorUpdateOne {
	if s != nil {
		jcuo.SetPersonType(*s)
	}
	return jcuo
}

// ClearPersonType clears the value of the "PersonType" field.
func (jcuo *JobContractorUpdateOne) ClearPersonType() *JobContractorUpdateOne {
	jcuo.mutation.ClearPersonType()
	return jcuo
}

// SetYDSID sets the "YDSID" field.
func (jcuo *JobContractorUpdateOne) SetYDSID(i int) *JobContractorUpdateOne {
	jcuo.mutation.ResetYDSID()
	jcuo.mutation.SetYDSID(i)
	return jcuo
}

// SetNillableYDSID sets the "YDSID" field if the given value is not nil.
func (jcuo *JobContractorUpdateOne) SetNillableYDSID(i *int) *JobContractorUpdateOne {
	if i != nil {
		jcuo.SetYDSID(*i)
	}
	return jcuo
}

// AddYDSID adds i to the "YDSID" field.
func (jcuo *JobContractorUpdateOne) AddYDSID(i int) *JobContractorUpdateOne {
	jcuo.mutation.AddYDSID(i)
	return jcuo
}

// ClearYDSID clears the value of the "YDSID" field.
func (jcuo *JobContractorUpdateOne) ClearYDSID() *JobContractorUpdateOne {
	jcuo.mutation.ClearYDSID()
	return jcuo
}

// SetNote sets the "Note" field.
func (jcuo *JobContractorUpdateOne) SetNote(s string) *JobContractorUpdateOne {
	jcuo.mutation.SetNote(s)
	return jcuo
}

// SetNillableNote sets the "Note" field if the given value is not nil.
func (jcuo *JobContractorUpdateOne) SetNillableNote(s *string) *JobContractorUpdateOne {
	if s != nil {
		jcuo.SetNote(*s)
	}
	return jcuo
}

// ClearNote clears the value of the "Note" field.
func (jcuo *JobContractorUpdateOne) ClearNote() *JobContractorUpdateOne {
	jcuo.mutation.ClearNote()
	return jcuo
}

// SetCreatedAt sets the "CreatedAt" field.
func (jcuo *JobContractorUpdateOne) SetCreatedAt(t time.Time) *JobContractorUpdateOne {
	jcuo.mutation.SetCreatedAt(t)
	return jcuo
}

// SetNillableCreatedAt sets the "CreatedAt" field if the given value is not nil.
func (jcuo *JobContractorUpdateOne) SetNillableCreatedAt(t *time.Time) *JobContractorUpdateOne {
	if t != nil {
		jcuo.SetCreatedAt(*t)
	}
	return jcuo
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (jcuo *JobContractorUpdateOne) SetUpdatedAt(t time.Time) *JobContractorUpdateOne {
	jcuo.mutation.SetUpdatedAt(t)
	return jcuo
}

// AddContractorIDs adds the "contractors" edge to the JobRelations entity by IDs.
func (jcuo *JobContractorUpdateOne) AddContractorIDs(ids ...int) *JobContractorUpdateOne {
	jcuo.mutation.AddContractorIDs(ids...)
	return jcuo
}

// AddContractors adds the "contractors" edges to the JobRelations entity.
func (jcuo *JobContractorUpdateOne) AddContractors(j ...*JobRelations) *JobContractorUpdateOne {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return jcuo.AddContractorIDs(ids...)
}

// Mutation returns the JobContractorMutation object of the builder.
func (jcuo *JobContractorUpdateOne) Mutation() *JobContractorMutation {
	return jcuo.mutation
}

// ClearContractors clears all "contractors" edges to the JobRelations entity.
func (jcuo *JobContractorUpdateOne) ClearContractors() *JobContractorUpdateOne {
	jcuo.mutation.ClearContractors()
	return jcuo
}

// RemoveContractorIDs removes the "contractors" edge to JobRelations entities by IDs.
func (jcuo *JobContractorUpdateOne) RemoveContractorIDs(ids ...int) *JobContractorUpdateOne {
	jcuo.mutation.RemoveContractorIDs(ids...)
	return jcuo
}

// RemoveContractors removes "contractors" edges to JobRelations entities.
func (jcuo *JobContractorUpdateOne) RemoveContractors(j ...*JobRelations) *JobContractorUpdateOne {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return jcuo.RemoveContractorIDs(ids...)
}

// Where appends a list predicates to the JobContractorUpdate builder.
func (jcuo *JobContractorUpdateOne) Where(ps ...predicate.JobContractor) *JobContractorUpdateOne {
	jcuo.mutation.Where(ps...)
	return jcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (jcuo *JobContractorUpdateOne) Select(field string, fields ...string) *JobContractorUpdateOne {
	jcuo.fields = append([]string{field}, fields...)
	return jcuo
}

// Save executes the query and returns the updated JobContractor entity.
func (jcuo *JobContractorUpdateOne) Save(ctx context.Context) (*JobContractor, error) {
	jcuo.defaults()
	return withHooks(ctx, jcuo.sqlSave, jcuo.mutation, jcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (jcuo *JobContractorUpdateOne) SaveX(ctx context.Context) *JobContractor {
	node, err := jcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (jcuo *JobContractorUpdateOne) Exec(ctx context.Context) error {
	_, err := jcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jcuo *JobContractorUpdateOne) ExecX(ctx context.Context) {
	if err := jcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (jcuo *JobContractorUpdateOne) defaults() {
	if _, ok := jcuo.mutation.UpdatedAt(); !ok {
		v := jobcontractor.UpdateDefaultUpdatedAt()
		jcuo.mutation.SetUpdatedAt(v)
	}
}

func (jcuo *JobContractorUpdateOne) sqlSave(ctx context.Context) (_node *JobContractor, err error) {
	_spec := sqlgraph.NewUpdateSpec(jobcontractor.Table, jobcontractor.Columns, sqlgraph.NewFieldSpec(jobcontractor.FieldID, field.TypeInt))
	id, ok := jcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "JobContractor.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := jcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, jobcontractor.FieldID)
		for _, f := range fields {
			if !jobcontractor.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != jobcontractor.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := jcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := jcuo.mutation.Name(); ok {
		_spec.SetField(jobcontractor.FieldName, field.TypeString, value)
	}
	if value, ok := jcuo.mutation.TcNo(); ok {
		_spec.SetField(jobcontractor.FieldTcNo, field.TypeString, value)
	}
	if jcuo.mutation.TcNoCleared() {
		_spec.ClearField(jobcontractor.FieldTcNo, field.TypeString)
	}
	if value, ok := jcuo.mutation.RegisterNo(); ok {
		_spec.SetField(jobcontractor.FieldRegisterNo, field.TypeString, value)
	}
	if jcuo.mutation.RegisterNoCleared() {
		_spec.ClearField(jobcontractor.FieldRegisterNo, field.TypeString)
	}
	if value, ok := jcuo.mutation.Address(); ok {
		_spec.SetField(jobcontractor.FieldAddress, field.TypeString, value)
	}
	if jcuo.mutation.AddressCleared() {
		_spec.ClearField(jobcontractor.FieldAddress, field.TypeString)
	}
	if value, ok := jcuo.mutation.TaxNo(); ok {
		_spec.SetField(jobcontractor.FieldTaxNo, field.TypeString, value)
	}
	if jcuo.mutation.TaxNoCleared() {
		_spec.ClearField(jobcontractor.FieldTaxNo, field.TypeString)
	}
	if value, ok := jcuo.mutation.MobilePhone(); ok {
		_spec.SetField(jobcontractor.FieldMobilePhone, field.TypeString, value)
	}
	if jcuo.mutation.MobilePhoneCleared() {
		_spec.ClearField(jobcontractor.FieldMobilePhone, field.TypeString)
	}
	if value, ok := jcuo.mutation.Phone(); ok {
		_spec.SetField(jobcontractor.FieldPhone, field.TypeString, value)
	}
	if jcuo.mutation.PhoneCleared() {
		_spec.ClearField(jobcontractor.FieldPhone, field.TypeString)
	}
	if value, ok := jcuo.mutation.Email(); ok {
		_spec.SetField(jobcontractor.FieldEmail, field.TypeString, value)
	}
	if jcuo.mutation.EmailCleared() {
		_spec.ClearField(jobcontractor.FieldEmail, field.TypeString)
	}
	if value, ok := jcuo.mutation.PersonType(); ok {
		_spec.SetField(jobcontractor.FieldPersonType, field.TypeString, value)
	}
	if jcuo.mutation.PersonTypeCleared() {
		_spec.ClearField(jobcontractor.FieldPersonType, field.TypeString)
	}
	if value, ok := jcuo.mutation.YDSID(); ok {
		_spec.SetField(jobcontractor.FieldYDSID, field.TypeInt, value)
	}
	if value, ok := jcuo.mutation.AddedYDSID(); ok {
		_spec.AddField(jobcontractor.FieldYDSID, field.TypeInt, value)
	}
	if jcuo.mutation.YDSIDCleared() {
		_spec.ClearField(jobcontractor.FieldYDSID, field.TypeInt)
	}
	if value, ok := jcuo.mutation.Note(); ok {
		_spec.SetField(jobcontractor.FieldNote, field.TypeString, value)
	}
	if jcuo.mutation.NoteCleared() {
		_spec.ClearField(jobcontractor.FieldNote, field.TypeString)
	}
	if value, ok := jcuo.mutation.CreatedAt(); ok {
		_spec.SetField(jobcontractor.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := jcuo.mutation.UpdatedAt(); ok {
		_spec.SetField(jobcontractor.FieldUpdatedAt, field.TypeTime, value)
	}
	if jcuo.mutation.ContractorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   jobcontractor.ContractorsTable,
			Columns: []string{jobcontractor.ContractorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(jobrelations.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jcuo.mutation.RemovedContractorsIDs(); len(nodes) > 0 && !jcuo.mutation.ContractorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   jobcontractor.ContractorsTable,
			Columns: []string{jobcontractor.ContractorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(jobrelations.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jcuo.mutation.ContractorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   jobcontractor.ContractorsTable,
			Columns: []string{jobcontractor.ContractorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(jobrelations.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &JobContractor{config: jcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, jcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{jobcontractor.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	jcuo.mutation.done = true
	return _node, nil
}
