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
	"github.com/polatbilal/gqlgen-ent/api-core/ent/jobrelations"
	"github.com/polatbilal/gqlgen-ent/api-core/ent/jobsupervisor"
	"github.com/polatbilal/gqlgen-ent/api-core/ent/predicate"
)

// JobSupervisorUpdate is the builder for updating JobSupervisor entities.
type JobSupervisorUpdate struct {
	config
	hooks    []Hook
	mutation *JobSupervisorMutation
}

// Where appends a list predicates to the JobSupervisorUpdate builder.
func (jsu *JobSupervisorUpdate) Where(ps ...predicate.JobSupervisor) *JobSupervisorUpdate {
	jsu.mutation.Where(ps...)
	return jsu
}

// SetName sets the "Name" field.
func (jsu *JobSupervisorUpdate) SetName(s string) *JobSupervisorUpdate {
	jsu.mutation.SetName(s)
	return jsu
}

// SetNillableName sets the "Name" field if the given value is not nil.
func (jsu *JobSupervisorUpdate) SetNillableName(s *string) *JobSupervisorUpdate {
	if s != nil {
		jsu.SetName(*s)
	}
	return jsu
}

// ClearName clears the value of the "Name" field.
func (jsu *JobSupervisorUpdate) ClearName() *JobSupervisorUpdate {
	jsu.mutation.ClearName()
	return jsu
}

// SetAddress sets the "Address" field.
func (jsu *JobSupervisorUpdate) SetAddress(s string) *JobSupervisorUpdate {
	jsu.mutation.SetAddress(s)
	return jsu
}

// SetNillableAddress sets the "Address" field if the given value is not nil.
func (jsu *JobSupervisorUpdate) SetNillableAddress(s *string) *JobSupervisorUpdate {
	if s != nil {
		jsu.SetAddress(*s)
	}
	return jsu
}

// ClearAddress clears the value of the "Address" field.
func (jsu *JobSupervisorUpdate) ClearAddress() *JobSupervisorUpdate {
	jsu.mutation.ClearAddress()
	return jsu
}

// SetPhone sets the "Phone" field.
func (jsu *JobSupervisorUpdate) SetPhone(s string) *JobSupervisorUpdate {
	jsu.mutation.SetPhone(s)
	return jsu
}

// SetNillablePhone sets the "Phone" field if the given value is not nil.
func (jsu *JobSupervisorUpdate) SetNillablePhone(s *string) *JobSupervisorUpdate {
	if s != nil {
		jsu.SetPhone(*s)
	}
	return jsu
}

// ClearPhone clears the value of the "Phone" field.
func (jsu *JobSupervisorUpdate) ClearPhone() *JobSupervisorUpdate {
	jsu.mutation.ClearPhone()
	return jsu
}

// SetEmail sets the "Email" field.
func (jsu *JobSupervisorUpdate) SetEmail(s string) *JobSupervisorUpdate {
	jsu.mutation.SetEmail(s)
	return jsu
}

// SetNillableEmail sets the "Email" field if the given value is not nil.
func (jsu *JobSupervisorUpdate) SetNillableEmail(s *string) *JobSupervisorUpdate {
	if s != nil {
		jsu.SetEmail(*s)
	}
	return jsu
}

// ClearEmail clears the value of the "Email" field.
func (jsu *JobSupervisorUpdate) ClearEmail() *JobSupervisorUpdate {
	jsu.mutation.ClearEmail()
	return jsu
}

// SetTcNo sets the "TcNo" field.
func (jsu *JobSupervisorUpdate) SetTcNo(s string) *JobSupervisorUpdate {
	jsu.mutation.SetTcNo(s)
	return jsu
}

// SetNillableTcNo sets the "TcNo" field if the given value is not nil.
func (jsu *JobSupervisorUpdate) SetNillableTcNo(s *string) *JobSupervisorUpdate {
	if s != nil {
		jsu.SetTcNo(*s)
	}
	return jsu
}

// ClearTcNo clears the value of the "TcNo" field.
func (jsu *JobSupervisorUpdate) ClearTcNo() *JobSupervisorUpdate {
	jsu.mutation.ClearTcNo()
	return jsu
}

// SetPosition sets the "Position" field.
func (jsu *JobSupervisorUpdate) SetPosition(s string) *JobSupervisorUpdate {
	jsu.mutation.SetPosition(s)
	return jsu
}

// SetNillablePosition sets the "Position" field if the given value is not nil.
func (jsu *JobSupervisorUpdate) SetNillablePosition(s *string) *JobSupervisorUpdate {
	if s != nil {
		jsu.SetPosition(*s)
	}
	return jsu
}

// ClearPosition clears the value of the "Position" field.
func (jsu *JobSupervisorUpdate) ClearPosition() *JobSupervisorUpdate {
	jsu.mutation.ClearPosition()
	return jsu
}

// SetCareer sets the "Career" field.
func (jsu *JobSupervisorUpdate) SetCareer(s string) *JobSupervisorUpdate {
	jsu.mutation.SetCareer(s)
	return jsu
}

// SetNillableCareer sets the "Career" field if the given value is not nil.
func (jsu *JobSupervisorUpdate) SetNillableCareer(s *string) *JobSupervisorUpdate {
	if s != nil {
		jsu.SetCareer(*s)
	}
	return jsu
}

// ClearCareer clears the value of the "Career" field.
func (jsu *JobSupervisorUpdate) ClearCareer() *JobSupervisorUpdate {
	jsu.mutation.ClearCareer()
	return jsu
}

// SetRegisterNo sets the "RegisterNo" field.
func (jsu *JobSupervisorUpdate) SetRegisterNo(s string) *JobSupervisorUpdate {
	jsu.mutation.SetRegisterNo(s)
	return jsu
}

// SetNillableRegisterNo sets the "RegisterNo" field if the given value is not nil.
func (jsu *JobSupervisorUpdate) SetNillableRegisterNo(s *string) *JobSupervisorUpdate {
	if s != nil {
		jsu.SetRegisterNo(*s)
	}
	return jsu
}

// ClearRegisterNo clears the value of the "RegisterNo" field.
func (jsu *JobSupervisorUpdate) ClearRegisterNo() *JobSupervisorUpdate {
	jsu.mutation.ClearRegisterNo()
	return jsu
}

// SetSocialSecurityNo sets the "SocialSecurityNo" field.
func (jsu *JobSupervisorUpdate) SetSocialSecurityNo(s string) *JobSupervisorUpdate {
	jsu.mutation.SetSocialSecurityNo(s)
	return jsu
}

// SetNillableSocialSecurityNo sets the "SocialSecurityNo" field if the given value is not nil.
func (jsu *JobSupervisorUpdate) SetNillableSocialSecurityNo(s *string) *JobSupervisorUpdate {
	if s != nil {
		jsu.SetSocialSecurityNo(*s)
	}
	return jsu
}

// ClearSocialSecurityNo clears the value of the "SocialSecurityNo" field.
func (jsu *JobSupervisorUpdate) ClearSocialSecurityNo() *JobSupervisorUpdate {
	jsu.mutation.ClearSocialSecurityNo()
	return jsu
}

// SetSchoolGraduation sets the "SchoolGraduation" field.
func (jsu *JobSupervisorUpdate) SetSchoolGraduation(s string) *JobSupervisorUpdate {
	jsu.mutation.SetSchoolGraduation(s)
	return jsu
}

// SetNillableSchoolGraduation sets the "SchoolGraduation" field if the given value is not nil.
func (jsu *JobSupervisorUpdate) SetNillableSchoolGraduation(s *string) *JobSupervisorUpdate {
	if s != nil {
		jsu.SetSchoolGraduation(*s)
	}
	return jsu
}

// ClearSchoolGraduation clears the value of the "SchoolGraduation" field.
func (jsu *JobSupervisorUpdate) ClearSchoolGraduation() *JobSupervisorUpdate {
	jsu.mutation.ClearSchoolGraduation()
	return jsu
}

// SetYDSID sets the "YDSID" field.
func (jsu *JobSupervisorUpdate) SetYDSID(i int) *JobSupervisorUpdate {
	jsu.mutation.ResetYDSID()
	jsu.mutation.SetYDSID(i)
	return jsu
}

// SetNillableYDSID sets the "YDSID" field if the given value is not nil.
func (jsu *JobSupervisorUpdate) SetNillableYDSID(i *int) *JobSupervisorUpdate {
	if i != nil {
		jsu.SetYDSID(*i)
	}
	return jsu
}

// AddYDSID adds i to the "YDSID" field.
func (jsu *JobSupervisorUpdate) AddYDSID(i int) *JobSupervisorUpdate {
	jsu.mutation.AddYDSID(i)
	return jsu
}

// ClearYDSID clears the value of the "YDSID" field.
func (jsu *JobSupervisorUpdate) ClearYDSID() *JobSupervisorUpdate {
	jsu.mutation.ClearYDSID()
	return jsu
}

// SetCreatedAt sets the "CreatedAt" field.
func (jsu *JobSupervisorUpdate) SetCreatedAt(t time.Time) *JobSupervisorUpdate {
	jsu.mutation.SetCreatedAt(t)
	return jsu
}

// SetNillableCreatedAt sets the "CreatedAt" field if the given value is not nil.
func (jsu *JobSupervisorUpdate) SetNillableCreatedAt(t *time.Time) *JobSupervisorUpdate {
	if t != nil {
		jsu.SetCreatedAt(*t)
	}
	return jsu
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (jsu *JobSupervisorUpdate) SetUpdatedAt(t time.Time) *JobSupervisorUpdate {
	jsu.mutation.SetUpdatedAt(t)
	return jsu
}

// AddSupervisorIDs adds the "supervisors" edge to the JobRelations entity by IDs.
func (jsu *JobSupervisorUpdate) AddSupervisorIDs(ids ...int) *JobSupervisorUpdate {
	jsu.mutation.AddSupervisorIDs(ids...)
	return jsu
}

// AddSupervisors adds the "supervisors" edges to the JobRelations entity.
func (jsu *JobSupervisorUpdate) AddSupervisors(j ...*JobRelations) *JobSupervisorUpdate {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return jsu.AddSupervisorIDs(ids...)
}

// Mutation returns the JobSupervisorMutation object of the builder.
func (jsu *JobSupervisorUpdate) Mutation() *JobSupervisorMutation {
	return jsu.mutation
}

// ClearSupervisors clears all "supervisors" edges to the JobRelations entity.
func (jsu *JobSupervisorUpdate) ClearSupervisors() *JobSupervisorUpdate {
	jsu.mutation.ClearSupervisors()
	return jsu
}

// RemoveSupervisorIDs removes the "supervisors" edge to JobRelations entities by IDs.
func (jsu *JobSupervisorUpdate) RemoveSupervisorIDs(ids ...int) *JobSupervisorUpdate {
	jsu.mutation.RemoveSupervisorIDs(ids...)
	return jsu
}

// RemoveSupervisors removes "supervisors" edges to JobRelations entities.
func (jsu *JobSupervisorUpdate) RemoveSupervisors(j ...*JobRelations) *JobSupervisorUpdate {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return jsu.RemoveSupervisorIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (jsu *JobSupervisorUpdate) Save(ctx context.Context) (int, error) {
	jsu.defaults()
	return withHooks(ctx, jsu.sqlSave, jsu.mutation, jsu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (jsu *JobSupervisorUpdate) SaveX(ctx context.Context) int {
	affected, err := jsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (jsu *JobSupervisorUpdate) Exec(ctx context.Context) error {
	_, err := jsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jsu *JobSupervisorUpdate) ExecX(ctx context.Context) {
	if err := jsu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (jsu *JobSupervisorUpdate) defaults() {
	if _, ok := jsu.mutation.UpdatedAt(); !ok {
		v := jobsupervisor.UpdateDefaultUpdatedAt()
		jsu.mutation.SetUpdatedAt(v)
	}
}

func (jsu *JobSupervisorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(jobsupervisor.Table, jobsupervisor.Columns, sqlgraph.NewFieldSpec(jobsupervisor.FieldID, field.TypeInt))
	if ps := jsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := jsu.mutation.Name(); ok {
		_spec.SetField(jobsupervisor.FieldName, field.TypeString, value)
	}
	if jsu.mutation.NameCleared() {
		_spec.ClearField(jobsupervisor.FieldName, field.TypeString)
	}
	if value, ok := jsu.mutation.Address(); ok {
		_spec.SetField(jobsupervisor.FieldAddress, field.TypeString, value)
	}
	if jsu.mutation.AddressCleared() {
		_spec.ClearField(jobsupervisor.FieldAddress, field.TypeString)
	}
	if value, ok := jsu.mutation.Phone(); ok {
		_spec.SetField(jobsupervisor.FieldPhone, field.TypeString, value)
	}
	if jsu.mutation.PhoneCleared() {
		_spec.ClearField(jobsupervisor.FieldPhone, field.TypeString)
	}
	if value, ok := jsu.mutation.Email(); ok {
		_spec.SetField(jobsupervisor.FieldEmail, field.TypeString, value)
	}
	if jsu.mutation.EmailCleared() {
		_spec.ClearField(jobsupervisor.FieldEmail, field.TypeString)
	}
	if value, ok := jsu.mutation.TcNo(); ok {
		_spec.SetField(jobsupervisor.FieldTcNo, field.TypeString, value)
	}
	if jsu.mutation.TcNoCleared() {
		_spec.ClearField(jobsupervisor.FieldTcNo, field.TypeString)
	}
	if value, ok := jsu.mutation.Position(); ok {
		_spec.SetField(jobsupervisor.FieldPosition, field.TypeString, value)
	}
	if jsu.mutation.PositionCleared() {
		_spec.ClearField(jobsupervisor.FieldPosition, field.TypeString)
	}
	if value, ok := jsu.mutation.Career(); ok {
		_spec.SetField(jobsupervisor.FieldCareer, field.TypeString, value)
	}
	if jsu.mutation.CareerCleared() {
		_spec.ClearField(jobsupervisor.FieldCareer, field.TypeString)
	}
	if value, ok := jsu.mutation.RegisterNo(); ok {
		_spec.SetField(jobsupervisor.FieldRegisterNo, field.TypeString, value)
	}
	if jsu.mutation.RegisterNoCleared() {
		_spec.ClearField(jobsupervisor.FieldRegisterNo, field.TypeString)
	}
	if value, ok := jsu.mutation.SocialSecurityNo(); ok {
		_spec.SetField(jobsupervisor.FieldSocialSecurityNo, field.TypeString, value)
	}
	if jsu.mutation.SocialSecurityNoCleared() {
		_spec.ClearField(jobsupervisor.FieldSocialSecurityNo, field.TypeString)
	}
	if value, ok := jsu.mutation.SchoolGraduation(); ok {
		_spec.SetField(jobsupervisor.FieldSchoolGraduation, field.TypeString, value)
	}
	if jsu.mutation.SchoolGraduationCleared() {
		_spec.ClearField(jobsupervisor.FieldSchoolGraduation, field.TypeString)
	}
	if value, ok := jsu.mutation.YDSID(); ok {
		_spec.SetField(jobsupervisor.FieldYDSID, field.TypeInt, value)
	}
	if value, ok := jsu.mutation.AddedYDSID(); ok {
		_spec.AddField(jobsupervisor.FieldYDSID, field.TypeInt, value)
	}
	if jsu.mutation.YDSIDCleared() {
		_spec.ClearField(jobsupervisor.FieldYDSID, field.TypeInt)
	}
	if value, ok := jsu.mutation.CreatedAt(); ok {
		_spec.SetField(jobsupervisor.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := jsu.mutation.UpdatedAt(); ok {
		_spec.SetField(jobsupervisor.FieldUpdatedAt, field.TypeTime, value)
	}
	if jsu.mutation.SupervisorsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jsu.mutation.RemovedSupervisorsIDs(); len(nodes) > 0 && !jsu.mutation.SupervisorsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jsu.mutation.SupervisorsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, jsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{jobsupervisor.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	jsu.mutation.done = true
	return n, nil
}

// JobSupervisorUpdateOne is the builder for updating a single JobSupervisor entity.
type JobSupervisorUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *JobSupervisorMutation
}

// SetName sets the "Name" field.
func (jsuo *JobSupervisorUpdateOne) SetName(s string) *JobSupervisorUpdateOne {
	jsuo.mutation.SetName(s)
	return jsuo
}

// SetNillableName sets the "Name" field if the given value is not nil.
func (jsuo *JobSupervisorUpdateOne) SetNillableName(s *string) *JobSupervisorUpdateOne {
	if s != nil {
		jsuo.SetName(*s)
	}
	return jsuo
}

// ClearName clears the value of the "Name" field.
func (jsuo *JobSupervisorUpdateOne) ClearName() *JobSupervisorUpdateOne {
	jsuo.mutation.ClearName()
	return jsuo
}

// SetAddress sets the "Address" field.
func (jsuo *JobSupervisorUpdateOne) SetAddress(s string) *JobSupervisorUpdateOne {
	jsuo.mutation.SetAddress(s)
	return jsuo
}

// SetNillableAddress sets the "Address" field if the given value is not nil.
func (jsuo *JobSupervisorUpdateOne) SetNillableAddress(s *string) *JobSupervisorUpdateOne {
	if s != nil {
		jsuo.SetAddress(*s)
	}
	return jsuo
}

// ClearAddress clears the value of the "Address" field.
func (jsuo *JobSupervisorUpdateOne) ClearAddress() *JobSupervisorUpdateOne {
	jsuo.mutation.ClearAddress()
	return jsuo
}

// SetPhone sets the "Phone" field.
func (jsuo *JobSupervisorUpdateOne) SetPhone(s string) *JobSupervisorUpdateOne {
	jsuo.mutation.SetPhone(s)
	return jsuo
}

// SetNillablePhone sets the "Phone" field if the given value is not nil.
func (jsuo *JobSupervisorUpdateOne) SetNillablePhone(s *string) *JobSupervisorUpdateOne {
	if s != nil {
		jsuo.SetPhone(*s)
	}
	return jsuo
}

// ClearPhone clears the value of the "Phone" field.
func (jsuo *JobSupervisorUpdateOne) ClearPhone() *JobSupervisorUpdateOne {
	jsuo.mutation.ClearPhone()
	return jsuo
}

// SetEmail sets the "Email" field.
func (jsuo *JobSupervisorUpdateOne) SetEmail(s string) *JobSupervisorUpdateOne {
	jsuo.mutation.SetEmail(s)
	return jsuo
}

// SetNillableEmail sets the "Email" field if the given value is not nil.
func (jsuo *JobSupervisorUpdateOne) SetNillableEmail(s *string) *JobSupervisorUpdateOne {
	if s != nil {
		jsuo.SetEmail(*s)
	}
	return jsuo
}

// ClearEmail clears the value of the "Email" field.
func (jsuo *JobSupervisorUpdateOne) ClearEmail() *JobSupervisorUpdateOne {
	jsuo.mutation.ClearEmail()
	return jsuo
}

// SetTcNo sets the "TcNo" field.
func (jsuo *JobSupervisorUpdateOne) SetTcNo(s string) *JobSupervisorUpdateOne {
	jsuo.mutation.SetTcNo(s)
	return jsuo
}

// SetNillableTcNo sets the "TcNo" field if the given value is not nil.
func (jsuo *JobSupervisorUpdateOne) SetNillableTcNo(s *string) *JobSupervisorUpdateOne {
	if s != nil {
		jsuo.SetTcNo(*s)
	}
	return jsuo
}

// ClearTcNo clears the value of the "TcNo" field.
func (jsuo *JobSupervisorUpdateOne) ClearTcNo() *JobSupervisorUpdateOne {
	jsuo.mutation.ClearTcNo()
	return jsuo
}

// SetPosition sets the "Position" field.
func (jsuo *JobSupervisorUpdateOne) SetPosition(s string) *JobSupervisorUpdateOne {
	jsuo.mutation.SetPosition(s)
	return jsuo
}

// SetNillablePosition sets the "Position" field if the given value is not nil.
func (jsuo *JobSupervisorUpdateOne) SetNillablePosition(s *string) *JobSupervisorUpdateOne {
	if s != nil {
		jsuo.SetPosition(*s)
	}
	return jsuo
}

// ClearPosition clears the value of the "Position" field.
func (jsuo *JobSupervisorUpdateOne) ClearPosition() *JobSupervisorUpdateOne {
	jsuo.mutation.ClearPosition()
	return jsuo
}

// SetCareer sets the "Career" field.
func (jsuo *JobSupervisorUpdateOne) SetCareer(s string) *JobSupervisorUpdateOne {
	jsuo.mutation.SetCareer(s)
	return jsuo
}

// SetNillableCareer sets the "Career" field if the given value is not nil.
func (jsuo *JobSupervisorUpdateOne) SetNillableCareer(s *string) *JobSupervisorUpdateOne {
	if s != nil {
		jsuo.SetCareer(*s)
	}
	return jsuo
}

// ClearCareer clears the value of the "Career" field.
func (jsuo *JobSupervisorUpdateOne) ClearCareer() *JobSupervisorUpdateOne {
	jsuo.mutation.ClearCareer()
	return jsuo
}

// SetRegisterNo sets the "RegisterNo" field.
func (jsuo *JobSupervisorUpdateOne) SetRegisterNo(s string) *JobSupervisorUpdateOne {
	jsuo.mutation.SetRegisterNo(s)
	return jsuo
}

// SetNillableRegisterNo sets the "RegisterNo" field if the given value is not nil.
func (jsuo *JobSupervisorUpdateOne) SetNillableRegisterNo(s *string) *JobSupervisorUpdateOne {
	if s != nil {
		jsuo.SetRegisterNo(*s)
	}
	return jsuo
}

// ClearRegisterNo clears the value of the "RegisterNo" field.
func (jsuo *JobSupervisorUpdateOne) ClearRegisterNo() *JobSupervisorUpdateOne {
	jsuo.mutation.ClearRegisterNo()
	return jsuo
}

// SetSocialSecurityNo sets the "SocialSecurityNo" field.
func (jsuo *JobSupervisorUpdateOne) SetSocialSecurityNo(s string) *JobSupervisorUpdateOne {
	jsuo.mutation.SetSocialSecurityNo(s)
	return jsuo
}

// SetNillableSocialSecurityNo sets the "SocialSecurityNo" field if the given value is not nil.
func (jsuo *JobSupervisorUpdateOne) SetNillableSocialSecurityNo(s *string) *JobSupervisorUpdateOne {
	if s != nil {
		jsuo.SetSocialSecurityNo(*s)
	}
	return jsuo
}

// ClearSocialSecurityNo clears the value of the "SocialSecurityNo" field.
func (jsuo *JobSupervisorUpdateOne) ClearSocialSecurityNo() *JobSupervisorUpdateOne {
	jsuo.mutation.ClearSocialSecurityNo()
	return jsuo
}

// SetSchoolGraduation sets the "SchoolGraduation" field.
func (jsuo *JobSupervisorUpdateOne) SetSchoolGraduation(s string) *JobSupervisorUpdateOne {
	jsuo.mutation.SetSchoolGraduation(s)
	return jsuo
}

// SetNillableSchoolGraduation sets the "SchoolGraduation" field if the given value is not nil.
func (jsuo *JobSupervisorUpdateOne) SetNillableSchoolGraduation(s *string) *JobSupervisorUpdateOne {
	if s != nil {
		jsuo.SetSchoolGraduation(*s)
	}
	return jsuo
}

// ClearSchoolGraduation clears the value of the "SchoolGraduation" field.
func (jsuo *JobSupervisorUpdateOne) ClearSchoolGraduation() *JobSupervisorUpdateOne {
	jsuo.mutation.ClearSchoolGraduation()
	return jsuo
}

// SetYDSID sets the "YDSID" field.
func (jsuo *JobSupervisorUpdateOne) SetYDSID(i int) *JobSupervisorUpdateOne {
	jsuo.mutation.ResetYDSID()
	jsuo.mutation.SetYDSID(i)
	return jsuo
}

// SetNillableYDSID sets the "YDSID" field if the given value is not nil.
func (jsuo *JobSupervisorUpdateOne) SetNillableYDSID(i *int) *JobSupervisorUpdateOne {
	if i != nil {
		jsuo.SetYDSID(*i)
	}
	return jsuo
}

// AddYDSID adds i to the "YDSID" field.
func (jsuo *JobSupervisorUpdateOne) AddYDSID(i int) *JobSupervisorUpdateOne {
	jsuo.mutation.AddYDSID(i)
	return jsuo
}

// ClearYDSID clears the value of the "YDSID" field.
func (jsuo *JobSupervisorUpdateOne) ClearYDSID() *JobSupervisorUpdateOne {
	jsuo.mutation.ClearYDSID()
	return jsuo
}

// SetCreatedAt sets the "CreatedAt" field.
func (jsuo *JobSupervisorUpdateOne) SetCreatedAt(t time.Time) *JobSupervisorUpdateOne {
	jsuo.mutation.SetCreatedAt(t)
	return jsuo
}

// SetNillableCreatedAt sets the "CreatedAt" field if the given value is not nil.
func (jsuo *JobSupervisorUpdateOne) SetNillableCreatedAt(t *time.Time) *JobSupervisorUpdateOne {
	if t != nil {
		jsuo.SetCreatedAt(*t)
	}
	return jsuo
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (jsuo *JobSupervisorUpdateOne) SetUpdatedAt(t time.Time) *JobSupervisorUpdateOne {
	jsuo.mutation.SetUpdatedAt(t)
	return jsuo
}

// AddSupervisorIDs adds the "supervisors" edge to the JobRelations entity by IDs.
func (jsuo *JobSupervisorUpdateOne) AddSupervisorIDs(ids ...int) *JobSupervisorUpdateOne {
	jsuo.mutation.AddSupervisorIDs(ids...)
	return jsuo
}

// AddSupervisors adds the "supervisors" edges to the JobRelations entity.
func (jsuo *JobSupervisorUpdateOne) AddSupervisors(j ...*JobRelations) *JobSupervisorUpdateOne {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return jsuo.AddSupervisorIDs(ids...)
}

// Mutation returns the JobSupervisorMutation object of the builder.
func (jsuo *JobSupervisorUpdateOne) Mutation() *JobSupervisorMutation {
	return jsuo.mutation
}

// ClearSupervisors clears all "supervisors" edges to the JobRelations entity.
func (jsuo *JobSupervisorUpdateOne) ClearSupervisors() *JobSupervisorUpdateOne {
	jsuo.mutation.ClearSupervisors()
	return jsuo
}

// RemoveSupervisorIDs removes the "supervisors" edge to JobRelations entities by IDs.
func (jsuo *JobSupervisorUpdateOne) RemoveSupervisorIDs(ids ...int) *JobSupervisorUpdateOne {
	jsuo.mutation.RemoveSupervisorIDs(ids...)
	return jsuo
}

// RemoveSupervisors removes "supervisors" edges to JobRelations entities.
func (jsuo *JobSupervisorUpdateOne) RemoveSupervisors(j ...*JobRelations) *JobSupervisorUpdateOne {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return jsuo.RemoveSupervisorIDs(ids...)
}

// Where appends a list predicates to the JobSupervisorUpdate builder.
func (jsuo *JobSupervisorUpdateOne) Where(ps ...predicate.JobSupervisor) *JobSupervisorUpdateOne {
	jsuo.mutation.Where(ps...)
	return jsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (jsuo *JobSupervisorUpdateOne) Select(field string, fields ...string) *JobSupervisorUpdateOne {
	jsuo.fields = append([]string{field}, fields...)
	return jsuo
}

// Save executes the query and returns the updated JobSupervisor entity.
func (jsuo *JobSupervisorUpdateOne) Save(ctx context.Context) (*JobSupervisor, error) {
	jsuo.defaults()
	return withHooks(ctx, jsuo.sqlSave, jsuo.mutation, jsuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (jsuo *JobSupervisorUpdateOne) SaveX(ctx context.Context) *JobSupervisor {
	node, err := jsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (jsuo *JobSupervisorUpdateOne) Exec(ctx context.Context) error {
	_, err := jsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jsuo *JobSupervisorUpdateOne) ExecX(ctx context.Context) {
	if err := jsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (jsuo *JobSupervisorUpdateOne) defaults() {
	if _, ok := jsuo.mutation.UpdatedAt(); !ok {
		v := jobsupervisor.UpdateDefaultUpdatedAt()
		jsuo.mutation.SetUpdatedAt(v)
	}
}

func (jsuo *JobSupervisorUpdateOne) sqlSave(ctx context.Context) (_node *JobSupervisor, err error) {
	_spec := sqlgraph.NewUpdateSpec(jobsupervisor.Table, jobsupervisor.Columns, sqlgraph.NewFieldSpec(jobsupervisor.FieldID, field.TypeInt))
	id, ok := jsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "JobSupervisor.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := jsuo.fields; len(fields) > 0 {
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
	if ps := jsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := jsuo.mutation.Name(); ok {
		_spec.SetField(jobsupervisor.FieldName, field.TypeString, value)
	}
	if jsuo.mutation.NameCleared() {
		_spec.ClearField(jobsupervisor.FieldName, field.TypeString)
	}
	if value, ok := jsuo.mutation.Address(); ok {
		_spec.SetField(jobsupervisor.FieldAddress, field.TypeString, value)
	}
	if jsuo.mutation.AddressCleared() {
		_spec.ClearField(jobsupervisor.FieldAddress, field.TypeString)
	}
	if value, ok := jsuo.mutation.Phone(); ok {
		_spec.SetField(jobsupervisor.FieldPhone, field.TypeString, value)
	}
	if jsuo.mutation.PhoneCleared() {
		_spec.ClearField(jobsupervisor.FieldPhone, field.TypeString)
	}
	if value, ok := jsuo.mutation.Email(); ok {
		_spec.SetField(jobsupervisor.FieldEmail, field.TypeString, value)
	}
	if jsuo.mutation.EmailCleared() {
		_spec.ClearField(jobsupervisor.FieldEmail, field.TypeString)
	}
	if value, ok := jsuo.mutation.TcNo(); ok {
		_spec.SetField(jobsupervisor.FieldTcNo, field.TypeString, value)
	}
	if jsuo.mutation.TcNoCleared() {
		_spec.ClearField(jobsupervisor.FieldTcNo, field.TypeString)
	}
	if value, ok := jsuo.mutation.Position(); ok {
		_spec.SetField(jobsupervisor.FieldPosition, field.TypeString, value)
	}
	if jsuo.mutation.PositionCleared() {
		_spec.ClearField(jobsupervisor.FieldPosition, field.TypeString)
	}
	if value, ok := jsuo.mutation.Career(); ok {
		_spec.SetField(jobsupervisor.FieldCareer, field.TypeString, value)
	}
	if jsuo.mutation.CareerCleared() {
		_spec.ClearField(jobsupervisor.FieldCareer, field.TypeString)
	}
	if value, ok := jsuo.mutation.RegisterNo(); ok {
		_spec.SetField(jobsupervisor.FieldRegisterNo, field.TypeString, value)
	}
	if jsuo.mutation.RegisterNoCleared() {
		_spec.ClearField(jobsupervisor.FieldRegisterNo, field.TypeString)
	}
	if value, ok := jsuo.mutation.SocialSecurityNo(); ok {
		_spec.SetField(jobsupervisor.FieldSocialSecurityNo, field.TypeString, value)
	}
	if jsuo.mutation.SocialSecurityNoCleared() {
		_spec.ClearField(jobsupervisor.FieldSocialSecurityNo, field.TypeString)
	}
	if value, ok := jsuo.mutation.SchoolGraduation(); ok {
		_spec.SetField(jobsupervisor.FieldSchoolGraduation, field.TypeString, value)
	}
	if jsuo.mutation.SchoolGraduationCleared() {
		_spec.ClearField(jobsupervisor.FieldSchoolGraduation, field.TypeString)
	}
	if value, ok := jsuo.mutation.YDSID(); ok {
		_spec.SetField(jobsupervisor.FieldYDSID, field.TypeInt, value)
	}
	if value, ok := jsuo.mutation.AddedYDSID(); ok {
		_spec.AddField(jobsupervisor.FieldYDSID, field.TypeInt, value)
	}
	if jsuo.mutation.YDSIDCleared() {
		_spec.ClearField(jobsupervisor.FieldYDSID, field.TypeInt)
	}
	if value, ok := jsuo.mutation.CreatedAt(); ok {
		_spec.SetField(jobsupervisor.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := jsuo.mutation.UpdatedAt(); ok {
		_spec.SetField(jobsupervisor.FieldUpdatedAt, field.TypeTime, value)
	}
	if jsuo.mutation.SupervisorsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jsuo.mutation.RemovedSupervisorsIDs(); len(nodes) > 0 && !jsuo.mutation.SupervisorsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jsuo.mutation.SupervisorsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &JobSupervisor{config: jsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, jsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{jobsupervisor.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	jsuo.mutation.done = true
	return _node, nil
}
