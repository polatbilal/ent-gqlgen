// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"gqlgen-ent/ent/companydetail"
	"gqlgen-ent/ent/companyengineer"
	"gqlgen-ent/ent/companyuser"
	"gqlgen-ent/ent/jobdetail"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CompanyDetailCreate is the builder for creating a CompanyDetail entity.
type CompanyDetailCreate struct {
	config
	mutation *CompanyDetailMutation
	hooks    []Hook
}

// SetCompanyCode sets the "CompanyCode" field.
func (cdc *CompanyDetailCreate) SetCompanyCode(i int) *CompanyDetailCreate {
	cdc.mutation.SetCompanyCode(i)
	return cdc
}

// SetName sets the "Name" field.
func (cdc *CompanyDetailCreate) SetName(s string) *CompanyDetailCreate {
	cdc.mutation.SetName(s)
	return cdc
}

// SetNillableName sets the "Name" field if the given value is not nil.
func (cdc *CompanyDetailCreate) SetNillableName(s *string) *CompanyDetailCreate {
	if s != nil {
		cdc.SetName(*s)
	}
	return cdc
}

// SetAddress sets the "Address" field.
func (cdc *CompanyDetailCreate) SetAddress(s string) *CompanyDetailCreate {
	cdc.mutation.SetAddress(s)
	return cdc
}

// SetNillableAddress sets the "Address" field if the given value is not nil.
func (cdc *CompanyDetailCreate) SetNillableAddress(s *string) *CompanyDetailCreate {
	if s != nil {
		cdc.SetAddress(*s)
	}
	return cdc
}

// SetPhone sets the "Phone" field.
func (cdc *CompanyDetailCreate) SetPhone(s string) *CompanyDetailCreate {
	cdc.mutation.SetPhone(s)
	return cdc
}

// SetNillablePhone sets the "Phone" field if the given value is not nil.
func (cdc *CompanyDetailCreate) SetNillablePhone(s *string) *CompanyDetailCreate {
	if s != nil {
		cdc.SetPhone(*s)
	}
	return cdc
}

// SetEmail sets the "Email" field.
func (cdc *CompanyDetailCreate) SetEmail(s string) *CompanyDetailCreate {
	cdc.mutation.SetEmail(s)
	return cdc
}

// SetNillableEmail sets the "Email" field if the given value is not nil.
func (cdc *CompanyDetailCreate) SetNillableEmail(s *string) *CompanyDetailCreate {
	if s != nil {
		cdc.SetEmail(*s)
	}
	return cdc
}

// SetWebsite sets the "Website" field.
func (cdc *CompanyDetailCreate) SetWebsite(s string) *CompanyDetailCreate {
	cdc.mutation.SetWebsite(s)
	return cdc
}

// SetNillableWebsite sets the "Website" field if the given value is not nil.
func (cdc *CompanyDetailCreate) SetNillableWebsite(s *string) *CompanyDetailCreate {
	if s != nil {
		cdc.SetWebsite(*s)
	}
	return cdc
}

// SetTaxAdmin sets the "TaxAdmin" field.
func (cdc *CompanyDetailCreate) SetTaxAdmin(s string) *CompanyDetailCreate {
	cdc.mutation.SetTaxAdmin(s)
	return cdc
}

// SetNillableTaxAdmin sets the "TaxAdmin" field if the given value is not nil.
func (cdc *CompanyDetailCreate) SetNillableTaxAdmin(s *string) *CompanyDetailCreate {
	if s != nil {
		cdc.SetTaxAdmin(*s)
	}
	return cdc
}

// SetTaxNo sets the "TaxNo" field.
func (cdc *CompanyDetailCreate) SetTaxNo(i int) *CompanyDetailCreate {
	cdc.mutation.SetTaxNo(i)
	return cdc
}

// SetNillableTaxNo sets the "TaxNo" field if the given value is not nil.
func (cdc *CompanyDetailCreate) SetNillableTaxNo(i *int) *CompanyDetailCreate {
	if i != nil {
		cdc.SetTaxNo(*i)
	}
	return cdc
}

// SetChamberInfo sets the "ChamberInfo" field.
func (cdc *CompanyDetailCreate) SetChamberInfo(s string) *CompanyDetailCreate {
	cdc.mutation.SetChamberInfo(s)
	return cdc
}

// SetNillableChamberInfo sets the "ChamberInfo" field if the given value is not nil.
func (cdc *CompanyDetailCreate) SetNillableChamberInfo(s *string) *CompanyDetailCreate {
	if s != nil {
		cdc.SetChamberInfo(*s)
	}
	return cdc
}

// SetChamberRegNo sets the "ChamberRegNo" field.
func (cdc *CompanyDetailCreate) SetChamberRegNo(s string) *CompanyDetailCreate {
	cdc.mutation.SetChamberRegNo(s)
	return cdc
}

// SetNillableChamberRegNo sets the "ChamberRegNo" field if the given value is not nil.
func (cdc *CompanyDetailCreate) SetNillableChamberRegNo(s *string) *CompanyDetailCreate {
	if s != nil {
		cdc.SetChamberRegNo(*s)
	}
	return cdc
}

// SetVisaDate sets the "VisaDate" field.
func (cdc *CompanyDetailCreate) SetVisaDate(t time.Time) *CompanyDetailCreate {
	cdc.mutation.SetVisaDate(t)
	return cdc
}

// SetNillableVisaDate sets the "VisaDate" field if the given value is not nil.
func (cdc *CompanyDetailCreate) SetNillableVisaDate(t *time.Time) *CompanyDetailCreate {
	if t != nil {
		cdc.SetVisaDate(*t)
	}
	return cdc
}

// SetVisaEndDate sets the "VisaEndDate" field.
func (cdc *CompanyDetailCreate) SetVisaEndDate(t time.Time) *CompanyDetailCreate {
	cdc.mutation.SetVisaEndDate(t)
	return cdc
}

// SetNillableVisaEndDate sets the "VisaEndDate" field if the given value is not nil.
func (cdc *CompanyDetailCreate) SetNillableVisaEndDate(t *time.Time) *CompanyDetailCreate {
	if t != nil {
		cdc.SetVisaEndDate(*t)
	}
	return cdc
}

// SetOwnerName sets the "OwnerName" field.
func (cdc *CompanyDetailCreate) SetOwnerName(s string) *CompanyDetailCreate {
	cdc.mutation.SetOwnerName(s)
	return cdc
}

// SetNillableOwnerName sets the "OwnerName" field if the given value is not nil.
func (cdc *CompanyDetailCreate) SetNillableOwnerName(s *string) *CompanyDetailCreate {
	if s != nil {
		cdc.SetOwnerName(*s)
	}
	return cdc
}

// SetOwnerTcNo sets the "OwnerTcNo" field.
func (cdc *CompanyDetailCreate) SetOwnerTcNo(s string) *CompanyDetailCreate {
	cdc.mutation.SetOwnerTcNo(s)
	return cdc
}

// SetNillableOwnerTcNo sets the "OwnerTcNo" field if the given value is not nil.
func (cdc *CompanyDetailCreate) SetNillableOwnerTcNo(s *string) *CompanyDetailCreate {
	if s != nil {
		cdc.SetOwnerTcNo(*s)
	}
	return cdc
}

// SetOwnerAddress sets the "OwnerAddress" field.
func (cdc *CompanyDetailCreate) SetOwnerAddress(s string) *CompanyDetailCreate {
	cdc.mutation.SetOwnerAddress(s)
	return cdc
}

// SetNillableOwnerAddress sets the "OwnerAddress" field if the given value is not nil.
func (cdc *CompanyDetailCreate) SetNillableOwnerAddress(s *string) *CompanyDetailCreate {
	if s != nil {
		cdc.SetOwnerAddress(*s)
	}
	return cdc
}

// SetOwnerPhone sets the "OwnerPhone" field.
func (cdc *CompanyDetailCreate) SetOwnerPhone(s string) *CompanyDetailCreate {
	cdc.mutation.SetOwnerPhone(s)
	return cdc
}

// SetNillableOwnerPhone sets the "OwnerPhone" field if the given value is not nil.
func (cdc *CompanyDetailCreate) SetNillableOwnerPhone(s *string) *CompanyDetailCreate {
	if s != nil {
		cdc.SetOwnerPhone(*s)
	}
	return cdc
}

// SetOwnerEmail sets the "OwnerEmail" field.
func (cdc *CompanyDetailCreate) SetOwnerEmail(s string) *CompanyDetailCreate {
	cdc.mutation.SetOwnerEmail(s)
	return cdc
}

// SetNillableOwnerEmail sets the "OwnerEmail" field if the given value is not nil.
func (cdc *CompanyDetailCreate) SetNillableOwnerEmail(s *string) *CompanyDetailCreate {
	if s != nil {
		cdc.SetOwnerEmail(*s)
	}
	return cdc
}

// SetOwnerRegNo sets the "OwnerRegNo" field.
func (cdc *CompanyDetailCreate) SetOwnerRegNo(s string) *CompanyDetailCreate {
	cdc.mutation.SetOwnerRegNo(s)
	return cdc
}

// SetNillableOwnerRegNo sets the "OwnerRegNo" field if the given value is not nil.
func (cdc *CompanyDetailCreate) SetNillableOwnerRegNo(s *string) *CompanyDetailCreate {
	if s != nil {
		cdc.SetOwnerRegNo(*s)
	}
	return cdc
}

// SetOwnerCareer sets the "OwnerCareer" field.
func (cdc *CompanyDetailCreate) SetOwnerCareer(s string) *CompanyDetailCreate {
	cdc.mutation.SetOwnerCareer(s)
	return cdc
}

// SetNillableOwnerCareer sets the "OwnerCareer" field if the given value is not nil.
func (cdc *CompanyDetailCreate) SetNillableOwnerCareer(s *string) *CompanyDetailCreate {
	if s != nil {
		cdc.SetOwnerCareer(*s)
	}
	return cdc
}

// SetOwnerBirthDate sets the "OwnerBirthDate" field.
func (cdc *CompanyDetailCreate) SetOwnerBirthDate(s string) *CompanyDetailCreate {
	cdc.mutation.SetOwnerBirthDate(s)
	return cdc
}

// SetNillableOwnerBirthDate sets the "OwnerBirthDate" field if the given value is not nil.
func (cdc *CompanyDetailCreate) SetNillableOwnerBirthDate(s *string) *CompanyDetailCreate {
	if s != nil {
		cdc.SetOwnerBirthDate(*s)
	}
	return cdc
}

// SetVisaFinishedFor90Days sets the "VisaFinishedFor90Days" field.
func (cdc *CompanyDetailCreate) SetVisaFinishedFor90Days(b bool) *CompanyDetailCreate {
	cdc.mutation.SetVisaFinishedFor90Days(b)
	return cdc
}

// SetNillableVisaFinishedFor90Days sets the "VisaFinishedFor90Days" field if the given value is not nil.
func (cdc *CompanyDetailCreate) SetNillableVisaFinishedFor90Days(b *bool) *CompanyDetailCreate {
	if b != nil {
		cdc.SetVisaFinishedFor90Days(*b)
	}
	return cdc
}

// SetCorePersonAbsent90Days sets the "CorePersonAbsent90Days" field.
func (cdc *CompanyDetailCreate) SetCorePersonAbsent90Days(b bool) *CompanyDetailCreate {
	cdc.mutation.SetCorePersonAbsent90Days(b)
	return cdc
}

// SetNillableCorePersonAbsent90Days sets the "CorePersonAbsent90Days" field if the given value is not nil.
func (cdc *CompanyDetailCreate) SetNillableCorePersonAbsent90Days(b *bool) *CompanyDetailCreate {
	if b != nil {
		cdc.SetCorePersonAbsent90Days(*b)
	}
	return cdc
}

// SetIsClosed sets the "IsClosed" field.
func (cdc *CompanyDetailCreate) SetIsClosed(b bool) *CompanyDetailCreate {
	cdc.mutation.SetIsClosed(b)
	return cdc
}

// SetNillableIsClosed sets the "IsClosed" field if the given value is not nil.
func (cdc *CompanyDetailCreate) SetNillableIsClosed(b *bool) *CompanyDetailCreate {
	if b != nil {
		cdc.SetIsClosed(*b)
	}
	return cdc
}

// SetCreatedAt sets the "CreatedAt" field.
func (cdc *CompanyDetailCreate) SetCreatedAt(t time.Time) *CompanyDetailCreate {
	cdc.mutation.SetCreatedAt(t)
	return cdc
}

// SetNillableCreatedAt sets the "CreatedAt" field if the given value is not nil.
func (cdc *CompanyDetailCreate) SetNillableCreatedAt(t *time.Time) *CompanyDetailCreate {
	if t != nil {
		cdc.SetCreatedAt(*t)
	}
	return cdc
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (cdc *CompanyDetailCreate) SetUpdatedAt(t time.Time) *CompanyDetailCreate {
	cdc.mutation.SetUpdatedAt(t)
	return cdc
}

// SetNillableUpdatedAt sets the "UpdatedAt" field if the given value is not nil.
func (cdc *CompanyDetailCreate) SetNillableUpdatedAt(t *time.Time) *CompanyDetailCreate {
	if t != nil {
		cdc.SetUpdatedAt(*t)
	}
	return cdc
}

// AddEngineerIDs adds the "engineers" edge to the CompanyEngineer entity by IDs.
func (cdc *CompanyDetailCreate) AddEngineerIDs(ids ...int) *CompanyDetailCreate {
	cdc.mutation.AddEngineerIDs(ids...)
	return cdc
}

// AddEngineers adds the "engineers" edges to the CompanyEngineer entity.
func (cdc *CompanyDetailCreate) AddEngineers(c ...*CompanyEngineer) *CompanyDetailCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cdc.AddEngineerIDs(ids...)
}

// AddUserIDs adds the "users" edge to the CompanyUser entity by IDs.
func (cdc *CompanyDetailCreate) AddUserIDs(ids ...int) *CompanyDetailCreate {
	cdc.mutation.AddUserIDs(ids...)
	return cdc
}

// AddUsers adds the "users" edges to the CompanyUser entity.
func (cdc *CompanyDetailCreate) AddUsers(c ...*CompanyUser) *CompanyDetailCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cdc.AddUserIDs(ids...)
}

// AddJobIDs adds the "jobs" edge to the JobDetail entity by IDs.
func (cdc *CompanyDetailCreate) AddJobIDs(ids ...int) *CompanyDetailCreate {
	cdc.mutation.AddJobIDs(ids...)
	return cdc
}

// AddJobs adds the "jobs" edges to the JobDetail entity.
func (cdc *CompanyDetailCreate) AddJobs(j ...*JobDetail) *CompanyDetailCreate {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return cdc.AddJobIDs(ids...)
}

// Mutation returns the CompanyDetailMutation object of the builder.
func (cdc *CompanyDetailCreate) Mutation() *CompanyDetailMutation {
	return cdc.mutation
}

// Save creates the CompanyDetail in the database.
func (cdc *CompanyDetailCreate) Save(ctx context.Context) (*CompanyDetail, error) {
	cdc.defaults()
	return withHooks(ctx, cdc.sqlSave, cdc.mutation, cdc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cdc *CompanyDetailCreate) SaveX(ctx context.Context) *CompanyDetail {
	v, err := cdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cdc *CompanyDetailCreate) Exec(ctx context.Context) error {
	_, err := cdc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cdc *CompanyDetailCreate) ExecX(ctx context.Context) {
	if err := cdc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cdc *CompanyDetailCreate) defaults() {
	if _, ok := cdc.mutation.Name(); !ok {
		v := companydetail.DefaultName
		cdc.mutation.SetName(v)
	}
	if _, ok := cdc.mutation.TaxNo(); !ok {
		v := companydetail.DefaultTaxNo
		cdc.mutation.SetTaxNo(v)
	}
	if _, ok := cdc.mutation.VisaFinishedFor90Days(); !ok {
		v := companydetail.DefaultVisaFinishedFor90Days
		cdc.mutation.SetVisaFinishedFor90Days(v)
	}
	if _, ok := cdc.mutation.CorePersonAbsent90Days(); !ok {
		v := companydetail.DefaultCorePersonAbsent90Days
		cdc.mutation.SetCorePersonAbsent90Days(v)
	}
	if _, ok := cdc.mutation.IsClosed(); !ok {
		v := companydetail.DefaultIsClosed
		cdc.mutation.SetIsClosed(v)
	}
	if _, ok := cdc.mutation.CreatedAt(); !ok {
		v := companydetail.DefaultCreatedAt()
		cdc.mutation.SetCreatedAt(v)
	}
	if _, ok := cdc.mutation.UpdatedAt(); !ok {
		v := companydetail.DefaultUpdatedAt()
		cdc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cdc *CompanyDetailCreate) check() error {
	if _, ok := cdc.mutation.CompanyCode(); !ok {
		return &ValidationError{Name: "CompanyCode", err: errors.New(`ent: missing required field "CompanyDetail.CompanyCode"`)}
	}
	if _, ok := cdc.mutation.Name(); !ok {
		return &ValidationError{Name: "Name", err: errors.New(`ent: missing required field "CompanyDetail.Name"`)}
	}
	if _, ok := cdc.mutation.IsClosed(); !ok {
		return &ValidationError{Name: "IsClosed", err: errors.New(`ent: missing required field "CompanyDetail.IsClosed"`)}
	}
	if _, ok := cdc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "CreatedAt", err: errors.New(`ent: missing required field "CompanyDetail.CreatedAt"`)}
	}
	if _, ok := cdc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "UpdatedAt", err: errors.New(`ent: missing required field "CompanyDetail.UpdatedAt"`)}
	}
	return nil
}

func (cdc *CompanyDetailCreate) sqlSave(ctx context.Context) (*CompanyDetail, error) {
	if err := cdc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cdc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cdc.mutation.id = &_node.ID
	cdc.mutation.done = true
	return _node, nil
}

func (cdc *CompanyDetailCreate) createSpec() (*CompanyDetail, *sqlgraph.CreateSpec) {
	var (
		_node = &CompanyDetail{config: cdc.config}
		_spec = sqlgraph.NewCreateSpec(companydetail.Table, sqlgraph.NewFieldSpec(companydetail.FieldID, field.TypeInt))
	)
	if value, ok := cdc.mutation.CompanyCode(); ok {
		_spec.SetField(companydetail.FieldCompanyCode, field.TypeInt, value)
		_node.CompanyCode = value
	}
	if value, ok := cdc.mutation.Name(); ok {
		_spec.SetField(companydetail.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cdc.mutation.Address(); ok {
		_spec.SetField(companydetail.FieldAddress, field.TypeString, value)
		_node.Address = value
	}
	if value, ok := cdc.mutation.Phone(); ok {
		_spec.SetField(companydetail.FieldPhone, field.TypeString, value)
		_node.Phone = value
	}
	if value, ok := cdc.mutation.Email(); ok {
		_spec.SetField(companydetail.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := cdc.mutation.Website(); ok {
		_spec.SetField(companydetail.FieldWebsite, field.TypeString, value)
		_node.Website = value
	}
	if value, ok := cdc.mutation.TaxAdmin(); ok {
		_spec.SetField(companydetail.FieldTaxAdmin, field.TypeString, value)
		_node.TaxAdmin = value
	}
	if value, ok := cdc.mutation.TaxNo(); ok {
		_spec.SetField(companydetail.FieldTaxNo, field.TypeInt, value)
		_node.TaxNo = value
	}
	if value, ok := cdc.mutation.ChamberInfo(); ok {
		_spec.SetField(companydetail.FieldChamberInfo, field.TypeString, value)
		_node.ChamberInfo = value
	}
	if value, ok := cdc.mutation.ChamberRegNo(); ok {
		_spec.SetField(companydetail.FieldChamberRegNo, field.TypeString, value)
		_node.ChamberRegNo = value
	}
	if value, ok := cdc.mutation.VisaDate(); ok {
		_spec.SetField(companydetail.FieldVisaDate, field.TypeTime, value)
		_node.VisaDate = value
	}
	if value, ok := cdc.mutation.VisaEndDate(); ok {
		_spec.SetField(companydetail.FieldVisaEndDate, field.TypeTime, value)
		_node.VisaEndDate = value
	}
	if value, ok := cdc.mutation.OwnerName(); ok {
		_spec.SetField(companydetail.FieldOwnerName, field.TypeString, value)
		_node.OwnerName = value
	}
	if value, ok := cdc.mutation.OwnerTcNo(); ok {
		_spec.SetField(companydetail.FieldOwnerTcNo, field.TypeString, value)
		_node.OwnerTcNo = value
	}
	if value, ok := cdc.mutation.OwnerAddress(); ok {
		_spec.SetField(companydetail.FieldOwnerAddress, field.TypeString, value)
		_node.OwnerAddress = value
	}
	if value, ok := cdc.mutation.OwnerPhone(); ok {
		_spec.SetField(companydetail.FieldOwnerPhone, field.TypeString, value)
		_node.OwnerPhone = value
	}
	if value, ok := cdc.mutation.OwnerEmail(); ok {
		_spec.SetField(companydetail.FieldOwnerEmail, field.TypeString, value)
		_node.OwnerEmail = value
	}
	if value, ok := cdc.mutation.OwnerRegNo(); ok {
		_spec.SetField(companydetail.FieldOwnerRegNo, field.TypeString, value)
		_node.OwnerRegNo = value
	}
	if value, ok := cdc.mutation.OwnerCareer(); ok {
		_spec.SetField(companydetail.FieldOwnerCareer, field.TypeString, value)
		_node.OwnerCareer = value
	}
	if value, ok := cdc.mutation.OwnerBirthDate(); ok {
		_spec.SetField(companydetail.FieldOwnerBirthDate, field.TypeString, value)
		_node.OwnerBirthDate = value
	}
	if value, ok := cdc.mutation.VisaFinishedFor90Days(); ok {
		_spec.SetField(companydetail.FieldVisaFinishedFor90Days, field.TypeBool, value)
		_node.VisaFinishedFor90Days = value
	}
	if value, ok := cdc.mutation.CorePersonAbsent90Days(); ok {
		_spec.SetField(companydetail.FieldCorePersonAbsent90Days, field.TypeBool, value)
		_node.CorePersonAbsent90Days = value
	}
	if value, ok := cdc.mutation.IsClosed(); ok {
		_spec.SetField(companydetail.FieldIsClosed, field.TypeBool, value)
		_node.IsClosed = value
	}
	if value, ok := cdc.mutation.CreatedAt(); ok {
		_spec.SetField(companydetail.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cdc.mutation.UpdatedAt(); ok {
		_spec.SetField(companydetail.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := cdc.mutation.EngineersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   companydetail.EngineersTable,
			Columns: []string{companydetail.EngineersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(companyengineer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cdc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   companydetail.UsersTable,
			Columns: []string{companydetail.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(companyuser.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cdc.mutation.JobsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   companydetail.JobsTable,
			Columns: []string{companydetail.JobsColumn},
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

// CompanyDetailCreateBulk is the builder for creating many CompanyDetail entities in bulk.
type CompanyDetailCreateBulk struct {
	config
	err      error
	builders []*CompanyDetailCreate
}

// Save creates the CompanyDetail entities in the database.
func (cdcb *CompanyDetailCreateBulk) Save(ctx context.Context) ([]*CompanyDetail, error) {
	if cdcb.err != nil {
		return nil, cdcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(cdcb.builders))
	nodes := make([]*CompanyDetail, len(cdcb.builders))
	mutators := make([]Mutator, len(cdcb.builders))
	for i := range cdcb.builders {
		func(i int, root context.Context) {
			builder := cdcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CompanyDetailMutation)
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
					_, err = mutators[i+1].Mutate(root, cdcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cdcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cdcb *CompanyDetailCreateBulk) SaveX(ctx context.Context) []*CompanyDetail {
	v, err := cdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cdcb *CompanyDetailCreateBulk) Exec(ctx context.Context) error {
	_, err := cdcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cdcb *CompanyDetailCreateBulk) ExecX(ctx context.Context) {
	if err := cdcb.Exec(ctx); err != nil {
		panic(err)
	}
}
