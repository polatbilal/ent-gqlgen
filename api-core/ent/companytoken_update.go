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
	"github.com/polatbilal/gqlgen-ent/api-core/ent/companydetail"
	"github.com/polatbilal/gqlgen-ent/api-core/ent/companytoken"
	"github.com/polatbilal/gqlgen-ent/api-core/ent/predicate"
)

// CompanyTokenUpdate is the builder for updating CompanyToken entities.
type CompanyTokenUpdate struct {
	config
	hooks    []Hook
	mutation *CompanyTokenMutation
}

// Where appends a list predicates to the CompanyTokenUpdate builder.
func (ctu *CompanyTokenUpdate) Where(ps ...predicate.CompanyToken) *CompanyTokenUpdate {
	ctu.mutation.Where(ps...)
	return ctu
}

// SetYDKUsername sets the "YDKUsername" field.
func (ctu *CompanyTokenUpdate) SetYDKUsername(s string) *CompanyTokenUpdate {
	ctu.mutation.SetYDKUsername(s)
	return ctu
}

// SetNillableYDKUsername sets the "YDKUsername" field if the given value is not nil.
func (ctu *CompanyTokenUpdate) SetNillableYDKUsername(s *string) *CompanyTokenUpdate {
	if s != nil {
		ctu.SetYDKUsername(*s)
	}
	return ctu
}

// ClearYDKUsername clears the value of the "YDKUsername" field.
func (ctu *CompanyTokenUpdate) ClearYDKUsername() *CompanyTokenUpdate {
	ctu.mutation.ClearYDKUsername()
	return ctu
}

// SetYDKPassword sets the "YDKPassword" field.
func (ctu *CompanyTokenUpdate) SetYDKPassword(s string) *CompanyTokenUpdate {
	ctu.mutation.SetYDKPassword(s)
	return ctu
}

// SetNillableYDKPassword sets the "YDKPassword" field if the given value is not nil.
func (ctu *CompanyTokenUpdate) SetNillableYDKPassword(s *string) *CompanyTokenUpdate {
	if s != nil {
		ctu.SetYDKPassword(*s)
	}
	return ctu
}

// ClearYDKPassword clears the value of the "YDKPassword" field.
func (ctu *CompanyTokenUpdate) ClearYDKPassword() *CompanyTokenUpdate {
	ctu.mutation.ClearYDKPassword()
	return ctu
}

// SetToken sets the "Token" field.
func (ctu *CompanyTokenUpdate) SetToken(s string) *CompanyTokenUpdate {
	ctu.mutation.SetToken(s)
	return ctu
}

// SetNillableToken sets the "Token" field if the given value is not nil.
func (ctu *CompanyTokenUpdate) SetNillableToken(s *string) *CompanyTokenUpdate {
	if s != nil {
		ctu.SetToken(*s)
	}
	return ctu
}

// ClearToken clears the value of the "Token" field.
func (ctu *CompanyTokenUpdate) ClearToken() *CompanyTokenUpdate {
	ctu.mutation.ClearToken()
	return ctu
}

// SetDepartmentId sets the "DepartmentId" field.
func (ctu *CompanyTokenUpdate) SetDepartmentId(i int) *CompanyTokenUpdate {
	ctu.mutation.ResetDepartmentId()
	ctu.mutation.SetDepartmentId(i)
	return ctu
}

// SetNillableDepartmentId sets the "DepartmentId" field if the given value is not nil.
func (ctu *CompanyTokenUpdate) SetNillableDepartmentId(i *int) *CompanyTokenUpdate {
	if i != nil {
		ctu.SetDepartmentId(*i)
	}
	return ctu
}

// AddDepartmentId adds i to the "DepartmentId" field.
func (ctu *CompanyTokenUpdate) AddDepartmentId(i int) *CompanyTokenUpdate {
	ctu.mutation.AddDepartmentId(i)
	return ctu
}

// ClearDepartmentId clears the value of the "DepartmentId" field.
func (ctu *CompanyTokenUpdate) ClearDepartmentId() *CompanyTokenUpdate {
	ctu.mutation.ClearDepartmentId()
	return ctu
}

// SetCreatedAt sets the "createdAt" field.
func (ctu *CompanyTokenUpdate) SetCreatedAt(t time.Time) *CompanyTokenUpdate {
	ctu.mutation.SetCreatedAt(t)
	return ctu
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (ctu *CompanyTokenUpdate) SetNillableCreatedAt(t *time.Time) *CompanyTokenUpdate {
	if t != nil {
		ctu.SetCreatedAt(*t)
	}
	return ctu
}

// SetUpdatedAt sets the "updatedAt" field.
func (ctu *CompanyTokenUpdate) SetUpdatedAt(t time.Time) *CompanyTokenUpdate {
	ctu.mutation.SetUpdatedAt(t)
	return ctu
}

// SetCompanyID sets the "company" edge to the CompanyDetail entity by ID.
func (ctu *CompanyTokenUpdate) SetCompanyID(id int) *CompanyTokenUpdate {
	ctu.mutation.SetCompanyID(id)
	return ctu
}

// SetNillableCompanyID sets the "company" edge to the CompanyDetail entity by ID if the given value is not nil.
func (ctu *CompanyTokenUpdate) SetNillableCompanyID(id *int) *CompanyTokenUpdate {
	if id != nil {
		ctu = ctu.SetCompanyID(*id)
	}
	return ctu
}

// SetCompany sets the "company" edge to the CompanyDetail entity.
func (ctu *CompanyTokenUpdate) SetCompany(c *CompanyDetail) *CompanyTokenUpdate {
	return ctu.SetCompanyID(c.ID)
}

// Mutation returns the CompanyTokenMutation object of the builder.
func (ctu *CompanyTokenUpdate) Mutation() *CompanyTokenMutation {
	return ctu.mutation
}

// ClearCompany clears the "company" edge to the CompanyDetail entity.
func (ctu *CompanyTokenUpdate) ClearCompany() *CompanyTokenUpdate {
	ctu.mutation.ClearCompany()
	return ctu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ctu *CompanyTokenUpdate) Save(ctx context.Context) (int, error) {
	ctu.defaults()
	return withHooks(ctx, ctu.sqlSave, ctu.mutation, ctu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ctu *CompanyTokenUpdate) SaveX(ctx context.Context) int {
	affected, err := ctu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ctu *CompanyTokenUpdate) Exec(ctx context.Context) error {
	_, err := ctu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctu *CompanyTokenUpdate) ExecX(ctx context.Context) {
	if err := ctu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ctu *CompanyTokenUpdate) defaults() {
	if _, ok := ctu.mutation.UpdatedAt(); !ok {
		v := companytoken.UpdateDefaultUpdatedAt()
		ctu.mutation.SetUpdatedAt(v)
	}
}

func (ctu *CompanyTokenUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(companytoken.Table, companytoken.Columns, sqlgraph.NewFieldSpec(companytoken.FieldID, field.TypeInt))
	if ps := ctu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ctu.mutation.YDKUsername(); ok {
		_spec.SetField(companytoken.FieldYDKUsername, field.TypeString, value)
	}
	if ctu.mutation.YDKUsernameCleared() {
		_spec.ClearField(companytoken.FieldYDKUsername, field.TypeString)
	}
	if value, ok := ctu.mutation.YDKPassword(); ok {
		_spec.SetField(companytoken.FieldYDKPassword, field.TypeString, value)
	}
	if ctu.mutation.YDKPasswordCleared() {
		_spec.ClearField(companytoken.FieldYDKPassword, field.TypeString)
	}
	if value, ok := ctu.mutation.Token(); ok {
		_spec.SetField(companytoken.FieldToken, field.TypeString, value)
	}
	if ctu.mutation.TokenCleared() {
		_spec.ClearField(companytoken.FieldToken, field.TypeString)
	}
	if value, ok := ctu.mutation.DepartmentId(); ok {
		_spec.SetField(companytoken.FieldDepartmentId, field.TypeInt, value)
	}
	if value, ok := ctu.mutation.AddedDepartmentId(); ok {
		_spec.AddField(companytoken.FieldDepartmentId, field.TypeInt, value)
	}
	if ctu.mutation.DepartmentIdCleared() {
		_spec.ClearField(companytoken.FieldDepartmentId, field.TypeInt)
	}
	if value, ok := ctu.mutation.CreatedAt(); ok {
		_spec.SetField(companytoken.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := ctu.mutation.UpdatedAt(); ok {
		_spec.SetField(companytoken.FieldUpdatedAt, field.TypeTime, value)
	}
	if ctu.mutation.CompanyCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctu.mutation.CompanyIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ctu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{companytoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ctu.mutation.done = true
	return n, nil
}

// CompanyTokenUpdateOne is the builder for updating a single CompanyToken entity.
type CompanyTokenUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CompanyTokenMutation
}

// SetYDKUsername sets the "YDKUsername" field.
func (ctuo *CompanyTokenUpdateOne) SetYDKUsername(s string) *CompanyTokenUpdateOne {
	ctuo.mutation.SetYDKUsername(s)
	return ctuo
}

// SetNillableYDKUsername sets the "YDKUsername" field if the given value is not nil.
func (ctuo *CompanyTokenUpdateOne) SetNillableYDKUsername(s *string) *CompanyTokenUpdateOne {
	if s != nil {
		ctuo.SetYDKUsername(*s)
	}
	return ctuo
}

// ClearYDKUsername clears the value of the "YDKUsername" field.
func (ctuo *CompanyTokenUpdateOne) ClearYDKUsername() *CompanyTokenUpdateOne {
	ctuo.mutation.ClearYDKUsername()
	return ctuo
}

// SetYDKPassword sets the "YDKPassword" field.
func (ctuo *CompanyTokenUpdateOne) SetYDKPassword(s string) *CompanyTokenUpdateOne {
	ctuo.mutation.SetYDKPassword(s)
	return ctuo
}

// SetNillableYDKPassword sets the "YDKPassword" field if the given value is not nil.
func (ctuo *CompanyTokenUpdateOne) SetNillableYDKPassword(s *string) *CompanyTokenUpdateOne {
	if s != nil {
		ctuo.SetYDKPassword(*s)
	}
	return ctuo
}

// ClearYDKPassword clears the value of the "YDKPassword" field.
func (ctuo *CompanyTokenUpdateOne) ClearYDKPassword() *CompanyTokenUpdateOne {
	ctuo.mutation.ClearYDKPassword()
	return ctuo
}

// SetToken sets the "Token" field.
func (ctuo *CompanyTokenUpdateOne) SetToken(s string) *CompanyTokenUpdateOne {
	ctuo.mutation.SetToken(s)
	return ctuo
}

// SetNillableToken sets the "Token" field if the given value is not nil.
func (ctuo *CompanyTokenUpdateOne) SetNillableToken(s *string) *CompanyTokenUpdateOne {
	if s != nil {
		ctuo.SetToken(*s)
	}
	return ctuo
}

// ClearToken clears the value of the "Token" field.
func (ctuo *CompanyTokenUpdateOne) ClearToken() *CompanyTokenUpdateOne {
	ctuo.mutation.ClearToken()
	return ctuo
}

// SetDepartmentId sets the "DepartmentId" field.
func (ctuo *CompanyTokenUpdateOne) SetDepartmentId(i int) *CompanyTokenUpdateOne {
	ctuo.mutation.ResetDepartmentId()
	ctuo.mutation.SetDepartmentId(i)
	return ctuo
}

// SetNillableDepartmentId sets the "DepartmentId" field if the given value is not nil.
func (ctuo *CompanyTokenUpdateOne) SetNillableDepartmentId(i *int) *CompanyTokenUpdateOne {
	if i != nil {
		ctuo.SetDepartmentId(*i)
	}
	return ctuo
}

// AddDepartmentId adds i to the "DepartmentId" field.
func (ctuo *CompanyTokenUpdateOne) AddDepartmentId(i int) *CompanyTokenUpdateOne {
	ctuo.mutation.AddDepartmentId(i)
	return ctuo
}

// ClearDepartmentId clears the value of the "DepartmentId" field.
func (ctuo *CompanyTokenUpdateOne) ClearDepartmentId() *CompanyTokenUpdateOne {
	ctuo.mutation.ClearDepartmentId()
	return ctuo
}

// SetCreatedAt sets the "createdAt" field.
func (ctuo *CompanyTokenUpdateOne) SetCreatedAt(t time.Time) *CompanyTokenUpdateOne {
	ctuo.mutation.SetCreatedAt(t)
	return ctuo
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (ctuo *CompanyTokenUpdateOne) SetNillableCreatedAt(t *time.Time) *CompanyTokenUpdateOne {
	if t != nil {
		ctuo.SetCreatedAt(*t)
	}
	return ctuo
}

// SetUpdatedAt sets the "updatedAt" field.
func (ctuo *CompanyTokenUpdateOne) SetUpdatedAt(t time.Time) *CompanyTokenUpdateOne {
	ctuo.mutation.SetUpdatedAt(t)
	return ctuo
}

// SetCompanyID sets the "company" edge to the CompanyDetail entity by ID.
func (ctuo *CompanyTokenUpdateOne) SetCompanyID(id int) *CompanyTokenUpdateOne {
	ctuo.mutation.SetCompanyID(id)
	return ctuo
}

// SetNillableCompanyID sets the "company" edge to the CompanyDetail entity by ID if the given value is not nil.
func (ctuo *CompanyTokenUpdateOne) SetNillableCompanyID(id *int) *CompanyTokenUpdateOne {
	if id != nil {
		ctuo = ctuo.SetCompanyID(*id)
	}
	return ctuo
}

// SetCompany sets the "company" edge to the CompanyDetail entity.
func (ctuo *CompanyTokenUpdateOne) SetCompany(c *CompanyDetail) *CompanyTokenUpdateOne {
	return ctuo.SetCompanyID(c.ID)
}

// Mutation returns the CompanyTokenMutation object of the builder.
func (ctuo *CompanyTokenUpdateOne) Mutation() *CompanyTokenMutation {
	return ctuo.mutation
}

// ClearCompany clears the "company" edge to the CompanyDetail entity.
func (ctuo *CompanyTokenUpdateOne) ClearCompany() *CompanyTokenUpdateOne {
	ctuo.mutation.ClearCompany()
	return ctuo
}

// Where appends a list predicates to the CompanyTokenUpdate builder.
func (ctuo *CompanyTokenUpdateOne) Where(ps ...predicate.CompanyToken) *CompanyTokenUpdateOne {
	ctuo.mutation.Where(ps...)
	return ctuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ctuo *CompanyTokenUpdateOne) Select(field string, fields ...string) *CompanyTokenUpdateOne {
	ctuo.fields = append([]string{field}, fields...)
	return ctuo
}

// Save executes the query and returns the updated CompanyToken entity.
func (ctuo *CompanyTokenUpdateOne) Save(ctx context.Context) (*CompanyToken, error) {
	ctuo.defaults()
	return withHooks(ctx, ctuo.sqlSave, ctuo.mutation, ctuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ctuo *CompanyTokenUpdateOne) SaveX(ctx context.Context) *CompanyToken {
	node, err := ctuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ctuo *CompanyTokenUpdateOne) Exec(ctx context.Context) error {
	_, err := ctuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctuo *CompanyTokenUpdateOne) ExecX(ctx context.Context) {
	if err := ctuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ctuo *CompanyTokenUpdateOne) defaults() {
	if _, ok := ctuo.mutation.UpdatedAt(); !ok {
		v := companytoken.UpdateDefaultUpdatedAt()
		ctuo.mutation.SetUpdatedAt(v)
	}
}

func (ctuo *CompanyTokenUpdateOne) sqlSave(ctx context.Context) (_node *CompanyToken, err error) {
	_spec := sqlgraph.NewUpdateSpec(companytoken.Table, companytoken.Columns, sqlgraph.NewFieldSpec(companytoken.FieldID, field.TypeInt))
	id, ok := ctuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CompanyToken.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ctuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, companytoken.FieldID)
		for _, f := range fields {
			if !companytoken.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != companytoken.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ctuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ctuo.mutation.YDKUsername(); ok {
		_spec.SetField(companytoken.FieldYDKUsername, field.TypeString, value)
	}
	if ctuo.mutation.YDKUsernameCleared() {
		_spec.ClearField(companytoken.FieldYDKUsername, field.TypeString)
	}
	if value, ok := ctuo.mutation.YDKPassword(); ok {
		_spec.SetField(companytoken.FieldYDKPassword, field.TypeString, value)
	}
	if ctuo.mutation.YDKPasswordCleared() {
		_spec.ClearField(companytoken.FieldYDKPassword, field.TypeString)
	}
	if value, ok := ctuo.mutation.Token(); ok {
		_spec.SetField(companytoken.FieldToken, field.TypeString, value)
	}
	if ctuo.mutation.TokenCleared() {
		_spec.ClearField(companytoken.FieldToken, field.TypeString)
	}
	if value, ok := ctuo.mutation.DepartmentId(); ok {
		_spec.SetField(companytoken.FieldDepartmentId, field.TypeInt, value)
	}
	if value, ok := ctuo.mutation.AddedDepartmentId(); ok {
		_spec.AddField(companytoken.FieldDepartmentId, field.TypeInt, value)
	}
	if ctuo.mutation.DepartmentIdCleared() {
		_spec.ClearField(companytoken.FieldDepartmentId, field.TypeInt)
	}
	if value, ok := ctuo.mutation.CreatedAt(); ok {
		_spec.SetField(companytoken.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := ctuo.mutation.UpdatedAt(); ok {
		_spec.SetField(companytoken.FieldUpdatedAt, field.TypeTime, value)
	}
	if ctuo.mutation.CompanyCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctuo.mutation.CompanyIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &CompanyToken{config: ctuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ctuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{companytoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ctuo.mutation.done = true
	return _node, nil
}
