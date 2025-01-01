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
	"github.com/polatbilal/gqlgen-ent/ent/jobsupervisor"
	"github.com/polatbilal/gqlgen-ent/ent/predicate"
)

// JobSuperVisorUpdate is the builder for updating JobSuperVisor entities.
type JobSuperVisorUpdate struct {
	config
	hooks    []Hook
	mutation *JobSuperVisorMutation
}

// Where appends a list predicates to the JobSuperVisorUpdate builder.
func (jsvu *JobSuperVisorUpdate) Where(ps ...predicate.JobSuperVisor) *JobSuperVisorUpdate {
	jsvu.mutation.Where(ps...)
	return jsvu
}

// SetName sets the "Name" field.
func (jsvu *JobSuperVisorUpdate) SetName(s string) *JobSuperVisorUpdate {
	jsvu.mutation.SetName(s)
	return jsvu
}

// SetNillableName sets the "Name" field if the given value is not nil.
func (jsvu *JobSuperVisorUpdate) SetNillableName(s *string) *JobSuperVisorUpdate {
	if s != nil {
		jsvu.SetName(*s)
	}
	return jsvu
}

// ClearName clears the value of the "Name" field.
func (jsvu *JobSuperVisorUpdate) ClearName() *JobSuperVisorUpdate {
	jsvu.mutation.ClearName()
	return jsvu
}

// SetAddress sets the "Address" field.
func (jsvu *JobSuperVisorUpdate) SetAddress(s string) *JobSuperVisorUpdate {
	jsvu.mutation.SetAddress(s)
	return jsvu
}

// SetNillableAddress sets the "Address" field if the given value is not nil.
func (jsvu *JobSuperVisorUpdate) SetNillableAddress(s *string) *JobSuperVisorUpdate {
	if s != nil {
		jsvu.SetAddress(*s)
	}
	return jsvu
}

// ClearAddress clears the value of the "Address" field.
func (jsvu *JobSuperVisorUpdate) ClearAddress() *JobSuperVisorUpdate {
	jsvu.mutation.ClearAddress()
	return jsvu
}

// SetPhone sets the "Phone" field.
func (jsvu *JobSuperVisorUpdate) SetPhone(s string) *JobSuperVisorUpdate {
	jsvu.mutation.SetPhone(s)
	return jsvu
}

// SetNillablePhone sets the "Phone" field if the given value is not nil.
func (jsvu *JobSuperVisorUpdate) SetNillablePhone(s *string) *JobSuperVisorUpdate {
	if s != nil {
		jsvu.SetPhone(*s)
	}
	return jsvu
}

// ClearPhone clears the value of the "Phone" field.
func (jsvu *JobSuperVisorUpdate) ClearPhone() *JobSuperVisorUpdate {
	jsvu.mutation.ClearPhone()
	return jsvu
}

// SetEmail sets the "Email" field.
func (jsvu *JobSuperVisorUpdate) SetEmail(s string) *JobSuperVisorUpdate {
	jsvu.mutation.SetEmail(s)
	return jsvu
}

// SetNillableEmail sets the "Email" field if the given value is not nil.
func (jsvu *JobSuperVisorUpdate) SetNillableEmail(s *string) *JobSuperVisorUpdate {
	if s != nil {
		jsvu.SetEmail(*s)
	}
	return jsvu
}

// ClearEmail clears the value of the "Email" field.
func (jsvu *JobSuperVisorUpdate) ClearEmail() *JobSuperVisorUpdate {
	jsvu.mutation.ClearEmail()
	return jsvu
}

// SetTCNO sets the "TCNO" field.
func (jsvu *JobSuperVisorUpdate) SetTCNO(i int) *JobSuperVisorUpdate {
	jsvu.mutation.ResetTCNO()
	jsvu.mutation.SetTCNO(i)
	return jsvu
}

// SetNillableTCNO sets the "TCNO" field if the given value is not nil.
func (jsvu *JobSuperVisorUpdate) SetNillableTCNO(i *int) *JobSuperVisorUpdate {
	if i != nil {
		jsvu.SetTCNO(*i)
	}
	return jsvu
}

// AddTCNO adds i to the "TCNO" field.
func (jsvu *JobSuperVisorUpdate) AddTCNO(i int) *JobSuperVisorUpdate {
	jsvu.mutation.AddTCNO(i)
	return jsvu
}

// ClearTCNO clears the value of the "TCNO" field.
func (jsvu *JobSuperVisorUpdate) ClearTCNO() *JobSuperVisorUpdate {
	jsvu.mutation.ClearTCNO()
	return jsvu
}

// SetPosition sets the "Position" field.
func (jsvu *JobSuperVisorUpdate) SetPosition(s string) *JobSuperVisorUpdate {
	jsvu.mutation.SetPosition(s)
	return jsvu
}

// SetNillablePosition sets the "Position" field if the given value is not nil.
func (jsvu *JobSuperVisorUpdate) SetNillablePosition(s *string) *JobSuperVisorUpdate {
	if s != nil {
		jsvu.SetPosition(*s)
	}
	return jsvu
}

// ClearPosition clears the value of the "Position" field.
func (jsvu *JobSuperVisorUpdate) ClearPosition() *JobSuperVisorUpdate {
	jsvu.mutation.ClearPosition()
	return jsvu
}

// SetCareer sets the "Career" field.
func (jsvu *JobSuperVisorUpdate) SetCareer(s string) *JobSuperVisorUpdate {
	jsvu.mutation.SetCareer(s)
	return jsvu
}

// SetNillableCareer sets the "Career" field if the given value is not nil.
func (jsvu *JobSuperVisorUpdate) SetNillableCareer(s *string) *JobSuperVisorUpdate {
	if s != nil {
		jsvu.SetCareer(*s)
	}
	return jsvu
}

// ClearCareer clears the value of the "Career" field.
func (jsvu *JobSuperVisorUpdate) ClearCareer() *JobSuperVisorUpdate {
	jsvu.mutation.ClearCareer()
	return jsvu
}

// SetRegNo sets the "RegNo" field.
func (jsvu *JobSuperVisorUpdate) SetRegNo(i int) *JobSuperVisorUpdate {
	jsvu.mutation.ResetRegNo()
	jsvu.mutation.SetRegNo(i)
	return jsvu
}

// SetNillableRegNo sets the "RegNo" field if the given value is not nil.
func (jsvu *JobSuperVisorUpdate) SetNillableRegNo(i *int) *JobSuperVisorUpdate {
	if i != nil {
		jsvu.SetRegNo(*i)
	}
	return jsvu
}

// AddRegNo adds i to the "RegNo" field.
func (jsvu *JobSuperVisorUpdate) AddRegNo(i int) *JobSuperVisorUpdate {
	jsvu.mutation.AddRegNo(i)
	return jsvu
}

// ClearRegNo clears the value of the "RegNo" field.
func (jsvu *JobSuperVisorUpdate) ClearRegNo() *JobSuperVisorUpdate {
	jsvu.mutation.ClearRegNo()
	return jsvu
}

// SetSocialSecurityNo sets the "socialSecurityNo" field.
func (jsvu *JobSuperVisorUpdate) SetSocialSecurityNo(i int) *JobSuperVisorUpdate {
	jsvu.mutation.ResetSocialSecurityNo()
	jsvu.mutation.SetSocialSecurityNo(i)
	return jsvu
}

// SetNillableSocialSecurityNo sets the "socialSecurityNo" field if the given value is not nil.
func (jsvu *JobSuperVisorUpdate) SetNillableSocialSecurityNo(i *int) *JobSuperVisorUpdate {
	if i != nil {
		jsvu.SetSocialSecurityNo(*i)
	}
	return jsvu
}

// AddSocialSecurityNo adds i to the "socialSecurityNo" field.
func (jsvu *JobSuperVisorUpdate) AddSocialSecurityNo(i int) *JobSuperVisorUpdate {
	jsvu.mutation.AddSocialSecurityNo(i)
	return jsvu
}

// ClearSocialSecurityNo clears the value of the "socialSecurityNo" field.
func (jsvu *JobSuperVisorUpdate) ClearSocialSecurityNo() *JobSuperVisorUpdate {
	jsvu.mutation.ClearSocialSecurityNo()
	return jsvu
}

// SetSchoolGraduation sets the "schoolGraduation" field.
func (jsvu *JobSuperVisorUpdate) SetSchoolGraduation(s string) *JobSuperVisorUpdate {
	jsvu.mutation.SetSchoolGraduation(s)
	return jsvu
}

// SetNillableSchoolGraduation sets the "schoolGraduation" field if the given value is not nil.
func (jsvu *JobSuperVisorUpdate) SetNillableSchoolGraduation(s *string) *JobSuperVisorUpdate {
	if s != nil {
		jsvu.SetSchoolGraduation(*s)
	}
	return jsvu
}

// ClearSchoolGraduation clears the value of the "schoolGraduation" field.
func (jsvu *JobSuperVisorUpdate) ClearSchoolGraduation() *JobSuperVisorUpdate {
	jsvu.mutation.ClearSchoolGraduation()
	return jsvu
}

// SetYDSID sets the "YDSID" field.
func (jsvu *JobSuperVisorUpdate) SetYDSID(i int) *JobSuperVisorUpdate {
	jsvu.mutation.ResetYDSID()
	jsvu.mutation.SetYDSID(i)
	return jsvu
}

// SetNillableYDSID sets the "YDSID" field if the given value is not nil.
func (jsvu *JobSuperVisorUpdate) SetNillableYDSID(i *int) *JobSuperVisorUpdate {
	if i != nil {
		jsvu.SetYDSID(*i)
	}
	return jsvu
}

// AddYDSID adds i to the "YDSID" field.
func (jsvu *JobSuperVisorUpdate) AddYDSID(i int) *JobSuperVisorUpdate {
	jsvu.mutation.AddYDSID(i)
	return jsvu
}

// ClearYDSID clears the value of the "YDSID" field.
func (jsvu *JobSuperVisorUpdate) ClearYDSID() *JobSuperVisorUpdate {
	jsvu.mutation.ClearYDSID()
	return jsvu
}

// SetCreatedAt sets the "CreatedAt" field.
func (jsvu *JobSuperVisorUpdate) SetCreatedAt(t time.Time) *JobSuperVisorUpdate {
	jsvu.mutation.SetCreatedAt(t)
	return jsvu
}

// SetNillableCreatedAt sets the "CreatedAt" field if the given value is not nil.
func (jsvu *JobSuperVisorUpdate) SetNillableCreatedAt(t *time.Time) *JobSuperVisorUpdate {
	if t != nil {
		jsvu.SetCreatedAt(*t)
	}
	return jsvu
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (jsvu *JobSuperVisorUpdate) SetUpdatedAt(t time.Time) *JobSuperVisorUpdate {
	jsvu.mutation.SetUpdatedAt(t)
	return jsvu
}

// AddSupervisorIDs adds the "supervisors" edge to the JobDetail entity by IDs.
func (jsvu *JobSuperVisorUpdate) AddSupervisorIDs(ids ...int) *JobSuperVisorUpdate {
	jsvu.mutation.AddSupervisorIDs(ids...)
	return jsvu
}

// AddSupervisors adds the "supervisors" edges to the JobDetail entity.
func (jsvu *JobSuperVisorUpdate) AddSupervisors(j ...*JobDetail) *JobSuperVisorUpdate {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return jsvu.AddSupervisorIDs(ids...)
}

// Mutation returns the JobSuperVisorMutation object of the builder.
func (jsvu *JobSuperVisorUpdate) Mutation() *JobSuperVisorMutation {
	return jsvu.mutation
}

// ClearSupervisors clears all "supervisors" edges to the JobDetail entity.
func (jsvu *JobSuperVisorUpdate) ClearSupervisors() *JobSuperVisorUpdate {
	jsvu.mutation.ClearSupervisors()
	return jsvu
}

// RemoveSupervisorIDs removes the "supervisors" edge to JobDetail entities by IDs.
func (jsvu *JobSuperVisorUpdate) RemoveSupervisorIDs(ids ...int) *JobSuperVisorUpdate {
	jsvu.mutation.RemoveSupervisorIDs(ids...)
	return jsvu
}

// RemoveSupervisors removes "supervisors" edges to JobDetail entities.
func (jsvu *JobSuperVisorUpdate) RemoveSupervisors(j ...*JobDetail) *JobSuperVisorUpdate {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return jsvu.RemoveSupervisorIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (jsvu *JobSuperVisorUpdate) Save(ctx context.Context) (int, error) {
	jsvu.defaults()
	return withHooks(ctx, jsvu.sqlSave, jsvu.mutation, jsvu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (jsvu *JobSuperVisorUpdate) SaveX(ctx context.Context) int {
	affected, err := jsvu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (jsvu *JobSuperVisorUpdate) Exec(ctx context.Context) error {
	_, err := jsvu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jsvu *JobSuperVisorUpdate) ExecX(ctx context.Context) {
	if err := jsvu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (jsvu *JobSuperVisorUpdate) defaults() {
	if _, ok := jsvu.mutation.UpdatedAt(); !ok {
		v := jobsupervisor.UpdateDefaultUpdatedAt()
		jsvu.mutation.SetUpdatedAt(v)
	}
}

func (jsvu *JobSuperVisorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(jobsupervisor.Table, jobsupervisor.Columns, sqlgraph.NewFieldSpec(jobsupervisor.FieldID, field.TypeInt))
	if ps := jsvu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := jsvu.mutation.Name(); ok {
		_spec.SetField(jobsupervisor.FieldName, field.TypeString, value)
	}
	if jsvu.mutation.NameCleared() {
		_spec.ClearField(jobsupervisor.FieldName, field.TypeString)
	}
	if value, ok := jsvu.mutation.Address(); ok {
		_spec.SetField(jobsupervisor.FieldAddress, field.TypeString, value)
	}
	if jsvu.mutation.AddressCleared() {
		_spec.ClearField(jobsupervisor.FieldAddress, field.TypeString)
	}
	if value, ok := jsvu.mutation.Phone(); ok {
		_spec.SetField(jobsupervisor.FieldPhone, field.TypeString, value)
	}
	if jsvu.mutation.PhoneCleared() {
		_spec.ClearField(jobsupervisor.FieldPhone, field.TypeString)
	}
	if value, ok := jsvu.mutation.Email(); ok {
		_spec.SetField(jobsupervisor.FieldEmail, field.TypeString, value)
	}
	if jsvu.mutation.EmailCleared() {
		_spec.ClearField(jobsupervisor.FieldEmail, field.TypeString)
	}
	if value, ok := jsvu.mutation.TCNO(); ok {
		_spec.SetField(jobsupervisor.FieldTCNO, field.TypeInt, value)
	}
	if value, ok := jsvu.mutation.AddedTCNO(); ok {
		_spec.AddField(jobsupervisor.FieldTCNO, field.TypeInt, value)
	}
	if jsvu.mutation.TCNOCleared() {
		_spec.ClearField(jobsupervisor.FieldTCNO, field.TypeInt)
	}
	if value, ok := jsvu.mutation.Position(); ok {
		_spec.SetField(jobsupervisor.FieldPosition, field.TypeString, value)
	}
	if jsvu.mutation.PositionCleared() {
		_spec.ClearField(jobsupervisor.FieldPosition, field.TypeString)
	}
	if value, ok := jsvu.mutation.Career(); ok {
		_spec.SetField(jobsupervisor.FieldCareer, field.TypeString, value)
	}
	if jsvu.mutation.CareerCleared() {
		_spec.ClearField(jobsupervisor.FieldCareer, field.TypeString)
	}
	if value, ok := jsvu.mutation.RegNo(); ok {
		_spec.SetField(jobsupervisor.FieldRegNo, field.TypeInt, value)
	}
	if value, ok := jsvu.mutation.AddedRegNo(); ok {
		_spec.AddField(jobsupervisor.FieldRegNo, field.TypeInt, value)
	}
	if jsvu.mutation.RegNoCleared() {
		_spec.ClearField(jobsupervisor.FieldRegNo, field.TypeInt)
	}
	if value, ok := jsvu.mutation.SocialSecurityNo(); ok {
		_spec.SetField(jobsupervisor.FieldSocialSecurityNo, field.TypeInt, value)
	}
	if value, ok := jsvu.mutation.AddedSocialSecurityNo(); ok {
		_spec.AddField(jobsupervisor.FieldSocialSecurityNo, field.TypeInt, value)
	}
	if jsvu.mutation.SocialSecurityNoCleared() {
		_spec.ClearField(jobsupervisor.FieldSocialSecurityNo, field.TypeInt)
	}
	if value, ok := jsvu.mutation.SchoolGraduation(); ok {
		_spec.SetField(jobsupervisor.FieldSchoolGraduation, field.TypeString, value)
	}
	if jsvu.mutation.SchoolGraduationCleared() {
		_spec.ClearField(jobsupervisor.FieldSchoolGraduation, field.TypeString)
	}
	if value, ok := jsvu.mutation.YDSID(); ok {
		_spec.SetField(jobsupervisor.FieldYDSID, field.TypeInt, value)
	}
	if value, ok := jsvu.mutation.AddedYDSID(); ok {
		_spec.AddField(jobsupervisor.FieldYDSID, field.TypeInt, value)
	}
	if jsvu.mutation.YDSIDCleared() {
		_spec.ClearField(jobsupervisor.FieldYDSID, field.TypeInt)
	}
	if value, ok := jsvu.mutation.CreatedAt(); ok {
		_spec.SetField(jobsupervisor.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := jsvu.mutation.UpdatedAt(); ok {
		_spec.SetField(jobsupervisor.FieldUpdatedAt, field.TypeTime, value)
	}
	if jsvu.mutation.SupervisorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   jobsupervisor.SupervisorsTable,
			Columns: []string{jobsupervisor.SupervisorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(jobdetail.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jsvu.mutation.RemovedSupervisorsIDs(); len(nodes) > 0 && !jsvu.mutation.SupervisorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   jobsupervisor.SupervisorsTable,
			Columns: []string{jobsupervisor.SupervisorsColumn},
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
	if nodes := jsvu.mutation.SupervisorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   jobsupervisor.SupervisorsTable,
			Columns: []string{jobsupervisor.SupervisorsColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, jsvu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{jobsupervisor.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	jsvu.mutation.done = true
	return n, nil
}

// JobSuperVisorUpdateOne is the builder for updating a single JobSuperVisor entity.
type JobSuperVisorUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *JobSuperVisorMutation
}

// SetName sets the "Name" field.
func (jsvuo *JobSuperVisorUpdateOne) SetName(s string) *JobSuperVisorUpdateOne {
	jsvuo.mutation.SetName(s)
	return jsvuo
}

// SetNillableName sets the "Name" field if the given value is not nil.
func (jsvuo *JobSuperVisorUpdateOne) SetNillableName(s *string) *JobSuperVisorUpdateOne {
	if s != nil {
		jsvuo.SetName(*s)
	}
	return jsvuo
}

// ClearName clears the value of the "Name" field.
func (jsvuo *JobSuperVisorUpdateOne) ClearName() *JobSuperVisorUpdateOne {
	jsvuo.mutation.ClearName()
	return jsvuo
}

// SetAddress sets the "Address" field.
func (jsvuo *JobSuperVisorUpdateOne) SetAddress(s string) *JobSuperVisorUpdateOne {
	jsvuo.mutation.SetAddress(s)
	return jsvuo
}

// SetNillableAddress sets the "Address" field if the given value is not nil.
func (jsvuo *JobSuperVisorUpdateOne) SetNillableAddress(s *string) *JobSuperVisorUpdateOne {
	if s != nil {
		jsvuo.SetAddress(*s)
	}
	return jsvuo
}

// ClearAddress clears the value of the "Address" field.
func (jsvuo *JobSuperVisorUpdateOne) ClearAddress() *JobSuperVisorUpdateOne {
	jsvuo.mutation.ClearAddress()
	return jsvuo
}

// SetPhone sets the "Phone" field.
func (jsvuo *JobSuperVisorUpdateOne) SetPhone(s string) *JobSuperVisorUpdateOne {
	jsvuo.mutation.SetPhone(s)
	return jsvuo
}

// SetNillablePhone sets the "Phone" field if the given value is not nil.
func (jsvuo *JobSuperVisorUpdateOne) SetNillablePhone(s *string) *JobSuperVisorUpdateOne {
	if s != nil {
		jsvuo.SetPhone(*s)
	}
	return jsvuo
}

// ClearPhone clears the value of the "Phone" field.
func (jsvuo *JobSuperVisorUpdateOne) ClearPhone() *JobSuperVisorUpdateOne {
	jsvuo.mutation.ClearPhone()
	return jsvuo
}

// SetEmail sets the "Email" field.
func (jsvuo *JobSuperVisorUpdateOne) SetEmail(s string) *JobSuperVisorUpdateOne {
	jsvuo.mutation.SetEmail(s)
	return jsvuo
}

// SetNillableEmail sets the "Email" field if the given value is not nil.
func (jsvuo *JobSuperVisorUpdateOne) SetNillableEmail(s *string) *JobSuperVisorUpdateOne {
	if s != nil {
		jsvuo.SetEmail(*s)
	}
	return jsvuo
}

// ClearEmail clears the value of the "Email" field.
func (jsvuo *JobSuperVisorUpdateOne) ClearEmail() *JobSuperVisorUpdateOne {
	jsvuo.mutation.ClearEmail()
	return jsvuo
}

// SetTCNO sets the "TCNO" field.
func (jsvuo *JobSuperVisorUpdateOne) SetTCNO(i int) *JobSuperVisorUpdateOne {
	jsvuo.mutation.ResetTCNO()
	jsvuo.mutation.SetTCNO(i)
	return jsvuo
}

// SetNillableTCNO sets the "TCNO" field if the given value is not nil.
func (jsvuo *JobSuperVisorUpdateOne) SetNillableTCNO(i *int) *JobSuperVisorUpdateOne {
	if i != nil {
		jsvuo.SetTCNO(*i)
	}
	return jsvuo
}

// AddTCNO adds i to the "TCNO" field.
func (jsvuo *JobSuperVisorUpdateOne) AddTCNO(i int) *JobSuperVisorUpdateOne {
	jsvuo.mutation.AddTCNO(i)
	return jsvuo
}

// ClearTCNO clears the value of the "TCNO" field.
func (jsvuo *JobSuperVisorUpdateOne) ClearTCNO() *JobSuperVisorUpdateOne {
	jsvuo.mutation.ClearTCNO()
	return jsvuo
}

// SetPosition sets the "Position" field.
func (jsvuo *JobSuperVisorUpdateOne) SetPosition(s string) *JobSuperVisorUpdateOne {
	jsvuo.mutation.SetPosition(s)
	return jsvuo
}

// SetNillablePosition sets the "Position" field if the given value is not nil.
func (jsvuo *JobSuperVisorUpdateOne) SetNillablePosition(s *string) *JobSuperVisorUpdateOne {
	if s != nil {
		jsvuo.SetPosition(*s)
	}
	return jsvuo
}

// ClearPosition clears the value of the "Position" field.
func (jsvuo *JobSuperVisorUpdateOne) ClearPosition() *JobSuperVisorUpdateOne {
	jsvuo.mutation.ClearPosition()
	return jsvuo
}

// SetCareer sets the "Career" field.
func (jsvuo *JobSuperVisorUpdateOne) SetCareer(s string) *JobSuperVisorUpdateOne {
	jsvuo.mutation.SetCareer(s)
	return jsvuo
}

// SetNillableCareer sets the "Career" field if the given value is not nil.
func (jsvuo *JobSuperVisorUpdateOne) SetNillableCareer(s *string) *JobSuperVisorUpdateOne {
	if s != nil {
		jsvuo.SetCareer(*s)
	}
	return jsvuo
}

// ClearCareer clears the value of the "Career" field.
func (jsvuo *JobSuperVisorUpdateOne) ClearCareer() *JobSuperVisorUpdateOne {
	jsvuo.mutation.ClearCareer()
	return jsvuo
}

// SetRegNo sets the "RegNo" field.
func (jsvuo *JobSuperVisorUpdateOne) SetRegNo(i int) *JobSuperVisorUpdateOne {
	jsvuo.mutation.ResetRegNo()
	jsvuo.mutation.SetRegNo(i)
	return jsvuo
}

// SetNillableRegNo sets the "RegNo" field if the given value is not nil.
func (jsvuo *JobSuperVisorUpdateOne) SetNillableRegNo(i *int) *JobSuperVisorUpdateOne {
	if i != nil {
		jsvuo.SetRegNo(*i)
	}
	return jsvuo
}

// AddRegNo adds i to the "RegNo" field.
func (jsvuo *JobSuperVisorUpdateOne) AddRegNo(i int) *JobSuperVisorUpdateOne {
	jsvuo.mutation.AddRegNo(i)
	return jsvuo
}

// ClearRegNo clears the value of the "RegNo" field.
func (jsvuo *JobSuperVisorUpdateOne) ClearRegNo() *JobSuperVisorUpdateOne {
	jsvuo.mutation.ClearRegNo()
	return jsvuo
}

// SetSocialSecurityNo sets the "socialSecurityNo" field.
func (jsvuo *JobSuperVisorUpdateOne) SetSocialSecurityNo(i int) *JobSuperVisorUpdateOne {
	jsvuo.mutation.ResetSocialSecurityNo()
	jsvuo.mutation.SetSocialSecurityNo(i)
	return jsvuo
}

// SetNillableSocialSecurityNo sets the "socialSecurityNo" field if the given value is not nil.
func (jsvuo *JobSuperVisorUpdateOne) SetNillableSocialSecurityNo(i *int) *JobSuperVisorUpdateOne {
	if i != nil {
		jsvuo.SetSocialSecurityNo(*i)
	}
	return jsvuo
}

// AddSocialSecurityNo adds i to the "socialSecurityNo" field.
func (jsvuo *JobSuperVisorUpdateOne) AddSocialSecurityNo(i int) *JobSuperVisorUpdateOne {
	jsvuo.mutation.AddSocialSecurityNo(i)
	return jsvuo
}

// ClearSocialSecurityNo clears the value of the "socialSecurityNo" field.
func (jsvuo *JobSuperVisorUpdateOne) ClearSocialSecurityNo() *JobSuperVisorUpdateOne {
	jsvuo.mutation.ClearSocialSecurityNo()
	return jsvuo
}

// SetSchoolGraduation sets the "schoolGraduation" field.
func (jsvuo *JobSuperVisorUpdateOne) SetSchoolGraduation(s string) *JobSuperVisorUpdateOne {
	jsvuo.mutation.SetSchoolGraduation(s)
	return jsvuo
}

// SetNillableSchoolGraduation sets the "schoolGraduation" field if the given value is not nil.
func (jsvuo *JobSuperVisorUpdateOne) SetNillableSchoolGraduation(s *string) *JobSuperVisorUpdateOne {
	if s != nil {
		jsvuo.SetSchoolGraduation(*s)
	}
	return jsvuo
}

// ClearSchoolGraduation clears the value of the "schoolGraduation" field.
func (jsvuo *JobSuperVisorUpdateOne) ClearSchoolGraduation() *JobSuperVisorUpdateOne {
	jsvuo.mutation.ClearSchoolGraduation()
	return jsvuo
}

// SetYDSID sets the "YDSID" field.
func (jsvuo *JobSuperVisorUpdateOne) SetYDSID(i int) *JobSuperVisorUpdateOne {
	jsvuo.mutation.ResetYDSID()
	jsvuo.mutation.SetYDSID(i)
	return jsvuo
}

// SetNillableYDSID sets the "YDSID" field if the given value is not nil.
func (jsvuo *JobSuperVisorUpdateOne) SetNillableYDSID(i *int) *JobSuperVisorUpdateOne {
	if i != nil {
		jsvuo.SetYDSID(*i)
	}
	return jsvuo
}

// AddYDSID adds i to the "YDSID" field.
func (jsvuo *JobSuperVisorUpdateOne) AddYDSID(i int) *JobSuperVisorUpdateOne {
	jsvuo.mutation.AddYDSID(i)
	return jsvuo
}

// ClearYDSID clears the value of the "YDSID" field.
func (jsvuo *JobSuperVisorUpdateOne) ClearYDSID() *JobSuperVisorUpdateOne {
	jsvuo.mutation.ClearYDSID()
	return jsvuo
}

// SetCreatedAt sets the "CreatedAt" field.
func (jsvuo *JobSuperVisorUpdateOne) SetCreatedAt(t time.Time) *JobSuperVisorUpdateOne {
	jsvuo.mutation.SetCreatedAt(t)
	return jsvuo
}

// SetNillableCreatedAt sets the "CreatedAt" field if the given value is not nil.
func (jsvuo *JobSuperVisorUpdateOne) SetNillableCreatedAt(t *time.Time) *JobSuperVisorUpdateOne {
	if t != nil {
		jsvuo.SetCreatedAt(*t)
	}
	return jsvuo
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (jsvuo *JobSuperVisorUpdateOne) SetUpdatedAt(t time.Time) *JobSuperVisorUpdateOne {
	jsvuo.mutation.SetUpdatedAt(t)
	return jsvuo
}

// AddSupervisorIDs adds the "supervisors" edge to the JobDetail entity by IDs.
func (jsvuo *JobSuperVisorUpdateOne) AddSupervisorIDs(ids ...int) *JobSuperVisorUpdateOne {
	jsvuo.mutation.AddSupervisorIDs(ids...)
	return jsvuo
}

// AddSupervisors adds the "supervisors" edges to the JobDetail entity.
func (jsvuo *JobSuperVisorUpdateOne) AddSupervisors(j ...*JobDetail) *JobSuperVisorUpdateOne {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return jsvuo.AddSupervisorIDs(ids...)
}

// Mutation returns the JobSuperVisorMutation object of the builder.
func (jsvuo *JobSuperVisorUpdateOne) Mutation() *JobSuperVisorMutation {
	return jsvuo.mutation
}

// ClearSupervisors clears all "supervisors" edges to the JobDetail entity.
func (jsvuo *JobSuperVisorUpdateOne) ClearSupervisors() *JobSuperVisorUpdateOne {
	jsvuo.mutation.ClearSupervisors()
	return jsvuo
}

// RemoveSupervisorIDs removes the "supervisors" edge to JobDetail entities by IDs.
func (jsvuo *JobSuperVisorUpdateOne) RemoveSupervisorIDs(ids ...int) *JobSuperVisorUpdateOne {
	jsvuo.mutation.RemoveSupervisorIDs(ids...)
	return jsvuo
}

// RemoveSupervisors removes "supervisors" edges to JobDetail entities.
func (jsvuo *JobSuperVisorUpdateOne) RemoveSupervisors(j ...*JobDetail) *JobSuperVisorUpdateOne {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return jsvuo.RemoveSupervisorIDs(ids...)
}

// Where appends a list predicates to the JobSuperVisorUpdate builder.
func (jsvuo *JobSuperVisorUpdateOne) Where(ps ...predicate.JobSuperVisor) *JobSuperVisorUpdateOne {
	jsvuo.mutation.Where(ps...)
	return jsvuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (jsvuo *JobSuperVisorUpdateOne) Select(field string, fields ...string) *JobSuperVisorUpdateOne {
	jsvuo.fields = append([]string{field}, fields...)
	return jsvuo
}

// Save executes the query and returns the updated JobSuperVisor entity.
func (jsvuo *JobSuperVisorUpdateOne) Save(ctx context.Context) (*JobSuperVisor, error) {
	jsvuo.defaults()
	return withHooks(ctx, jsvuo.sqlSave, jsvuo.mutation, jsvuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (jsvuo *JobSuperVisorUpdateOne) SaveX(ctx context.Context) *JobSuperVisor {
	node, err := jsvuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (jsvuo *JobSuperVisorUpdateOne) Exec(ctx context.Context) error {
	_, err := jsvuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jsvuo *JobSuperVisorUpdateOne) ExecX(ctx context.Context) {
	if err := jsvuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (jsvuo *JobSuperVisorUpdateOne) defaults() {
	if _, ok := jsvuo.mutation.UpdatedAt(); !ok {
		v := jobsupervisor.UpdateDefaultUpdatedAt()
		jsvuo.mutation.SetUpdatedAt(v)
	}
}

func (jsvuo *JobSuperVisorUpdateOne) sqlSave(ctx context.Context) (_node *JobSuperVisor, err error) {
	_spec := sqlgraph.NewUpdateSpec(jobsupervisor.Table, jobsupervisor.Columns, sqlgraph.NewFieldSpec(jobsupervisor.FieldID, field.TypeInt))
	id, ok := jsvuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "JobSuperVisor.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := jsvuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, jobsupervisor.FieldID)
		for _, f := range fields {
			if !jobsupervisor.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != jobsupervisor.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := jsvuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := jsvuo.mutation.Name(); ok {
		_spec.SetField(jobsupervisor.FieldName, field.TypeString, value)
	}
	if jsvuo.mutation.NameCleared() {
		_spec.ClearField(jobsupervisor.FieldName, field.TypeString)
	}
	if value, ok := jsvuo.mutation.Address(); ok {
		_spec.SetField(jobsupervisor.FieldAddress, field.TypeString, value)
	}
	if jsvuo.mutation.AddressCleared() {
		_spec.ClearField(jobsupervisor.FieldAddress, field.TypeString)
	}
	if value, ok := jsvuo.mutation.Phone(); ok {
		_spec.SetField(jobsupervisor.FieldPhone, field.TypeString, value)
	}
	if jsvuo.mutation.PhoneCleared() {
		_spec.ClearField(jobsupervisor.FieldPhone, field.TypeString)
	}
	if value, ok := jsvuo.mutation.Email(); ok {
		_spec.SetField(jobsupervisor.FieldEmail, field.TypeString, value)
	}
	if jsvuo.mutation.EmailCleared() {
		_spec.ClearField(jobsupervisor.FieldEmail, field.TypeString)
	}
	if value, ok := jsvuo.mutation.TCNO(); ok {
		_spec.SetField(jobsupervisor.FieldTCNO, field.TypeInt, value)
	}
	if value, ok := jsvuo.mutation.AddedTCNO(); ok {
		_spec.AddField(jobsupervisor.FieldTCNO, field.TypeInt, value)
	}
	if jsvuo.mutation.TCNOCleared() {
		_spec.ClearField(jobsupervisor.FieldTCNO, field.TypeInt)
	}
	if value, ok := jsvuo.mutation.Position(); ok {
		_spec.SetField(jobsupervisor.FieldPosition, field.TypeString, value)
	}
	if jsvuo.mutation.PositionCleared() {
		_spec.ClearField(jobsupervisor.FieldPosition, field.TypeString)
	}
	if value, ok := jsvuo.mutation.Career(); ok {
		_spec.SetField(jobsupervisor.FieldCareer, field.TypeString, value)
	}
	if jsvuo.mutation.CareerCleared() {
		_spec.ClearField(jobsupervisor.FieldCareer, field.TypeString)
	}
	if value, ok := jsvuo.mutation.RegNo(); ok {
		_spec.SetField(jobsupervisor.FieldRegNo, field.TypeInt, value)
	}
	if value, ok := jsvuo.mutation.AddedRegNo(); ok {
		_spec.AddField(jobsupervisor.FieldRegNo, field.TypeInt, value)
	}
	if jsvuo.mutation.RegNoCleared() {
		_spec.ClearField(jobsupervisor.FieldRegNo, field.TypeInt)
	}
	if value, ok := jsvuo.mutation.SocialSecurityNo(); ok {
		_spec.SetField(jobsupervisor.FieldSocialSecurityNo, field.TypeInt, value)
	}
	if value, ok := jsvuo.mutation.AddedSocialSecurityNo(); ok {
		_spec.AddField(jobsupervisor.FieldSocialSecurityNo, field.TypeInt, value)
	}
	if jsvuo.mutation.SocialSecurityNoCleared() {
		_spec.ClearField(jobsupervisor.FieldSocialSecurityNo, field.TypeInt)
	}
	if value, ok := jsvuo.mutation.SchoolGraduation(); ok {
		_spec.SetField(jobsupervisor.FieldSchoolGraduation, field.TypeString, value)
	}
	if jsvuo.mutation.SchoolGraduationCleared() {
		_spec.ClearField(jobsupervisor.FieldSchoolGraduation, field.TypeString)
	}
	if value, ok := jsvuo.mutation.YDSID(); ok {
		_spec.SetField(jobsupervisor.FieldYDSID, field.TypeInt, value)
	}
	if value, ok := jsvuo.mutation.AddedYDSID(); ok {
		_spec.AddField(jobsupervisor.FieldYDSID, field.TypeInt, value)
	}
	if jsvuo.mutation.YDSIDCleared() {
		_spec.ClearField(jobsupervisor.FieldYDSID, field.TypeInt)
	}
	if value, ok := jsvuo.mutation.CreatedAt(); ok {
		_spec.SetField(jobsupervisor.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := jsvuo.mutation.UpdatedAt(); ok {
		_spec.SetField(jobsupervisor.FieldUpdatedAt, field.TypeTime, value)
	}
	if jsvuo.mutation.SupervisorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   jobsupervisor.SupervisorsTable,
			Columns: []string{jobsupervisor.SupervisorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(jobdetail.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jsvuo.mutation.RemovedSupervisorsIDs(); len(nodes) > 0 && !jsvuo.mutation.SupervisorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   jobsupervisor.SupervisorsTable,
			Columns: []string{jobsupervisor.SupervisorsColumn},
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
	if nodes := jsvuo.mutation.SupervisorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   jobsupervisor.SupervisorsTable,
			Columns: []string{jobsupervisor.SupervisorsColumn},
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
	_node = &JobSuperVisor{config: jsvuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, jsvuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{jobsupervisor.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	jsvuo.mutation.done = true
	return _node, nil
}
