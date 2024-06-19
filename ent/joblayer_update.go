// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"gqlgen-ent/ent/jobdetail"
	"gqlgen-ent/ent/joblayer"
	"gqlgen-ent/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// JobLayerUpdate is the builder for updating JobLayer entities.
type JobLayerUpdate struct {
	config
	hooks    []Hook
	mutation *JobLayerMutation
}

// Where appends a list predicates to the JobLayerUpdate builder.
func (jlu *JobLayerUpdate) Where(ps ...predicate.JobLayer) *JobLayerUpdate {
	jlu.mutation.Where(ps...)
	return jlu
}

// SetName sets the "Name" field.
func (jlu *JobLayerUpdate) SetName(s string) *JobLayerUpdate {
	jlu.mutation.SetName(s)
	return jlu
}

// SetNillableName sets the "Name" field if the given value is not nil.
func (jlu *JobLayerUpdate) SetNillableName(s *string) *JobLayerUpdate {
	if s != nil {
		jlu.SetName(*s)
	}
	return jlu
}

// SetMetre sets the "Metre" field.
func (jlu *JobLayerUpdate) SetMetre(s string) *JobLayerUpdate {
	jlu.mutation.SetMetre(s)
	return jlu
}

// SetNillableMetre sets the "Metre" field if the given value is not nil.
func (jlu *JobLayerUpdate) SetNillableMetre(s *string) *JobLayerUpdate {
	if s != nil {
		jlu.SetMetre(*s)
	}
	return jlu
}

// SetMoldDate sets the "MoldDate" field.
func (jlu *JobLayerUpdate) SetMoldDate(t time.Time) *JobLayerUpdate {
	jlu.mutation.SetMoldDate(t)
	return jlu
}

// SetNillableMoldDate sets the "MoldDate" field if the given value is not nil.
func (jlu *JobLayerUpdate) SetNillableMoldDate(t *time.Time) *JobLayerUpdate {
	if t != nil {
		jlu.SetMoldDate(*t)
	}
	return jlu
}

// ClearMoldDate clears the value of the "MoldDate" field.
func (jlu *JobLayerUpdate) ClearMoldDate() *JobLayerUpdate {
	jlu.mutation.ClearMoldDate()
	return jlu
}

// SetConcreteDate sets the "ConcreteDate" field.
func (jlu *JobLayerUpdate) SetConcreteDate(t time.Time) *JobLayerUpdate {
	jlu.mutation.SetConcreteDate(t)
	return jlu
}

// SetNillableConcreteDate sets the "ConcreteDate" field if the given value is not nil.
func (jlu *JobLayerUpdate) SetNillableConcreteDate(t *time.Time) *JobLayerUpdate {
	if t != nil {
		jlu.SetConcreteDate(*t)
	}
	return jlu
}

// ClearConcreteDate clears the value of the "ConcreteDate" field.
func (jlu *JobLayerUpdate) ClearConcreteDate() *JobLayerUpdate {
	jlu.mutation.ClearConcreteDate()
	return jlu
}

// SetSamples sets the "Samples" field.
func (jlu *JobLayerUpdate) SetSamples(i int) *JobLayerUpdate {
	jlu.mutation.ResetSamples()
	jlu.mutation.SetSamples(i)
	return jlu
}

// SetNillableSamples sets the "Samples" field if the given value is not nil.
func (jlu *JobLayerUpdate) SetNillableSamples(i *int) *JobLayerUpdate {
	if i != nil {
		jlu.SetSamples(*i)
	}
	return jlu
}

// AddSamples adds i to the "Samples" field.
func (jlu *JobLayerUpdate) AddSamples(i int) *JobLayerUpdate {
	jlu.mutation.AddSamples(i)
	return jlu
}

// ClearSamples clears the value of the "Samples" field.
func (jlu *JobLayerUpdate) ClearSamples() *JobLayerUpdate {
	jlu.mutation.ClearSamples()
	return jlu
}

// SetConcreteClass sets the "ConcreteClass" field.
func (jlu *JobLayerUpdate) SetConcreteClass(s string) *JobLayerUpdate {
	jlu.mutation.SetConcreteClass(s)
	return jlu
}

// SetNillableConcreteClass sets the "ConcreteClass" field if the given value is not nil.
func (jlu *JobLayerUpdate) SetNillableConcreteClass(s *string) *JobLayerUpdate {
	if s != nil {
		jlu.SetConcreteClass(*s)
	}
	return jlu
}

// ClearConcreteClass clears the value of the "ConcreteClass" field.
func (jlu *JobLayerUpdate) ClearConcreteClass() *JobLayerUpdate {
	jlu.mutation.ClearConcreteClass()
	return jlu
}

// SetWeekResult sets the "WeekResult" field.
func (jlu *JobLayerUpdate) SetWeekResult(s string) *JobLayerUpdate {
	jlu.mutation.SetWeekResult(s)
	return jlu
}

// SetNillableWeekResult sets the "WeekResult" field if the given value is not nil.
func (jlu *JobLayerUpdate) SetNillableWeekResult(s *string) *JobLayerUpdate {
	if s != nil {
		jlu.SetWeekResult(*s)
	}
	return jlu
}

// ClearWeekResult clears the value of the "WeekResult" field.
func (jlu *JobLayerUpdate) ClearWeekResult() *JobLayerUpdate {
	jlu.mutation.ClearWeekResult()
	return jlu
}

// SetMonthResult sets the "MonthResult" field.
func (jlu *JobLayerUpdate) SetMonthResult(s string) *JobLayerUpdate {
	jlu.mutation.SetMonthResult(s)
	return jlu
}

// SetNillableMonthResult sets the "MonthResult" field if the given value is not nil.
func (jlu *JobLayerUpdate) SetNillableMonthResult(s *string) *JobLayerUpdate {
	if s != nil {
		jlu.SetMonthResult(*s)
	}
	return jlu
}

// ClearMonthResult clears the value of the "MonthResult" field.
func (jlu *JobLayerUpdate) ClearMonthResult() *JobLayerUpdate {
	jlu.mutation.ClearMonthResult()
	return jlu
}

// SetCreatedAt sets the "created_at" field.
func (jlu *JobLayerUpdate) SetCreatedAt(t time.Time) *JobLayerUpdate {
	jlu.mutation.SetCreatedAt(t)
	return jlu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (jlu *JobLayerUpdate) SetNillableCreatedAt(t *time.Time) *JobLayerUpdate {
	if t != nil {
		jlu.SetCreatedAt(*t)
	}
	return jlu
}

// SetUpdatedAt sets the "updated_at" field.
func (jlu *JobLayerUpdate) SetUpdatedAt(t time.Time) *JobLayerUpdate {
	jlu.mutation.SetUpdatedAt(t)
	return jlu
}

// SetLayerID sets the "layer" edge to the JobDetail entity by ID.
func (jlu *JobLayerUpdate) SetLayerID(id int) *JobLayerUpdate {
	jlu.mutation.SetLayerID(id)
	return jlu
}

// SetNillableLayerID sets the "layer" edge to the JobDetail entity by ID if the given value is not nil.
func (jlu *JobLayerUpdate) SetNillableLayerID(id *int) *JobLayerUpdate {
	if id != nil {
		jlu = jlu.SetLayerID(*id)
	}
	return jlu
}

// SetLayer sets the "layer" edge to the JobDetail entity.
func (jlu *JobLayerUpdate) SetLayer(j *JobDetail) *JobLayerUpdate {
	return jlu.SetLayerID(j.ID)
}

// Mutation returns the JobLayerMutation object of the builder.
func (jlu *JobLayerUpdate) Mutation() *JobLayerMutation {
	return jlu.mutation
}

// ClearLayer clears the "layer" edge to the JobDetail entity.
func (jlu *JobLayerUpdate) ClearLayer() *JobLayerUpdate {
	jlu.mutation.ClearLayer()
	return jlu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (jlu *JobLayerUpdate) Save(ctx context.Context) (int, error) {
	jlu.defaults()
	return withHooks(ctx, jlu.sqlSave, jlu.mutation, jlu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (jlu *JobLayerUpdate) SaveX(ctx context.Context) int {
	affected, err := jlu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (jlu *JobLayerUpdate) Exec(ctx context.Context) error {
	_, err := jlu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jlu *JobLayerUpdate) ExecX(ctx context.Context) {
	if err := jlu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (jlu *JobLayerUpdate) defaults() {
	if _, ok := jlu.mutation.UpdatedAt(); !ok {
		v := joblayer.UpdateDefaultUpdatedAt()
		jlu.mutation.SetUpdatedAt(v)
	}
}

func (jlu *JobLayerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(joblayer.Table, joblayer.Columns, sqlgraph.NewFieldSpec(joblayer.FieldID, field.TypeInt))
	if ps := jlu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := jlu.mutation.Name(); ok {
		_spec.SetField(joblayer.FieldName, field.TypeString, value)
	}
	if value, ok := jlu.mutation.Metre(); ok {
		_spec.SetField(joblayer.FieldMetre, field.TypeString, value)
	}
	if value, ok := jlu.mutation.MoldDate(); ok {
		_spec.SetField(joblayer.FieldMoldDate, field.TypeTime, value)
	}
	if jlu.mutation.MoldDateCleared() {
		_spec.ClearField(joblayer.FieldMoldDate, field.TypeTime)
	}
	if value, ok := jlu.mutation.ConcreteDate(); ok {
		_spec.SetField(joblayer.FieldConcreteDate, field.TypeTime, value)
	}
	if jlu.mutation.ConcreteDateCleared() {
		_spec.ClearField(joblayer.FieldConcreteDate, field.TypeTime)
	}
	if value, ok := jlu.mutation.Samples(); ok {
		_spec.SetField(joblayer.FieldSamples, field.TypeInt, value)
	}
	if value, ok := jlu.mutation.AddedSamples(); ok {
		_spec.AddField(joblayer.FieldSamples, field.TypeInt, value)
	}
	if jlu.mutation.SamplesCleared() {
		_spec.ClearField(joblayer.FieldSamples, field.TypeInt)
	}
	if value, ok := jlu.mutation.ConcreteClass(); ok {
		_spec.SetField(joblayer.FieldConcreteClass, field.TypeString, value)
	}
	if jlu.mutation.ConcreteClassCleared() {
		_spec.ClearField(joblayer.FieldConcreteClass, field.TypeString)
	}
	if value, ok := jlu.mutation.WeekResult(); ok {
		_spec.SetField(joblayer.FieldWeekResult, field.TypeString, value)
	}
	if jlu.mutation.WeekResultCleared() {
		_spec.ClearField(joblayer.FieldWeekResult, field.TypeString)
	}
	if value, ok := jlu.mutation.MonthResult(); ok {
		_spec.SetField(joblayer.FieldMonthResult, field.TypeString, value)
	}
	if jlu.mutation.MonthResultCleared() {
		_spec.ClearField(joblayer.FieldMonthResult, field.TypeString)
	}
	if value, ok := jlu.mutation.CreatedAt(); ok {
		_spec.SetField(joblayer.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := jlu.mutation.UpdatedAt(); ok {
		_spec.SetField(joblayer.FieldUpdatedAt, field.TypeTime, value)
	}
	if jlu.mutation.LayerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   joblayer.LayerTable,
			Columns: []string{joblayer.LayerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(jobdetail.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jlu.mutation.LayerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   joblayer.LayerTable,
			Columns: []string{joblayer.LayerColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, jlu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{joblayer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	jlu.mutation.done = true
	return n, nil
}

// JobLayerUpdateOne is the builder for updating a single JobLayer entity.
type JobLayerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *JobLayerMutation
}

// SetName sets the "Name" field.
func (jluo *JobLayerUpdateOne) SetName(s string) *JobLayerUpdateOne {
	jluo.mutation.SetName(s)
	return jluo
}

// SetNillableName sets the "Name" field if the given value is not nil.
func (jluo *JobLayerUpdateOne) SetNillableName(s *string) *JobLayerUpdateOne {
	if s != nil {
		jluo.SetName(*s)
	}
	return jluo
}

// SetMetre sets the "Metre" field.
func (jluo *JobLayerUpdateOne) SetMetre(s string) *JobLayerUpdateOne {
	jluo.mutation.SetMetre(s)
	return jluo
}

// SetNillableMetre sets the "Metre" field if the given value is not nil.
func (jluo *JobLayerUpdateOne) SetNillableMetre(s *string) *JobLayerUpdateOne {
	if s != nil {
		jluo.SetMetre(*s)
	}
	return jluo
}

// SetMoldDate sets the "MoldDate" field.
func (jluo *JobLayerUpdateOne) SetMoldDate(t time.Time) *JobLayerUpdateOne {
	jluo.mutation.SetMoldDate(t)
	return jluo
}

// SetNillableMoldDate sets the "MoldDate" field if the given value is not nil.
func (jluo *JobLayerUpdateOne) SetNillableMoldDate(t *time.Time) *JobLayerUpdateOne {
	if t != nil {
		jluo.SetMoldDate(*t)
	}
	return jluo
}

// ClearMoldDate clears the value of the "MoldDate" field.
func (jluo *JobLayerUpdateOne) ClearMoldDate() *JobLayerUpdateOne {
	jluo.mutation.ClearMoldDate()
	return jluo
}

// SetConcreteDate sets the "ConcreteDate" field.
func (jluo *JobLayerUpdateOne) SetConcreteDate(t time.Time) *JobLayerUpdateOne {
	jluo.mutation.SetConcreteDate(t)
	return jluo
}

// SetNillableConcreteDate sets the "ConcreteDate" field if the given value is not nil.
func (jluo *JobLayerUpdateOne) SetNillableConcreteDate(t *time.Time) *JobLayerUpdateOne {
	if t != nil {
		jluo.SetConcreteDate(*t)
	}
	return jluo
}

// ClearConcreteDate clears the value of the "ConcreteDate" field.
func (jluo *JobLayerUpdateOne) ClearConcreteDate() *JobLayerUpdateOne {
	jluo.mutation.ClearConcreteDate()
	return jluo
}

// SetSamples sets the "Samples" field.
func (jluo *JobLayerUpdateOne) SetSamples(i int) *JobLayerUpdateOne {
	jluo.mutation.ResetSamples()
	jluo.mutation.SetSamples(i)
	return jluo
}

// SetNillableSamples sets the "Samples" field if the given value is not nil.
func (jluo *JobLayerUpdateOne) SetNillableSamples(i *int) *JobLayerUpdateOne {
	if i != nil {
		jluo.SetSamples(*i)
	}
	return jluo
}

// AddSamples adds i to the "Samples" field.
func (jluo *JobLayerUpdateOne) AddSamples(i int) *JobLayerUpdateOne {
	jluo.mutation.AddSamples(i)
	return jluo
}

// ClearSamples clears the value of the "Samples" field.
func (jluo *JobLayerUpdateOne) ClearSamples() *JobLayerUpdateOne {
	jluo.mutation.ClearSamples()
	return jluo
}

// SetConcreteClass sets the "ConcreteClass" field.
func (jluo *JobLayerUpdateOne) SetConcreteClass(s string) *JobLayerUpdateOne {
	jluo.mutation.SetConcreteClass(s)
	return jluo
}

// SetNillableConcreteClass sets the "ConcreteClass" field if the given value is not nil.
func (jluo *JobLayerUpdateOne) SetNillableConcreteClass(s *string) *JobLayerUpdateOne {
	if s != nil {
		jluo.SetConcreteClass(*s)
	}
	return jluo
}

// ClearConcreteClass clears the value of the "ConcreteClass" field.
func (jluo *JobLayerUpdateOne) ClearConcreteClass() *JobLayerUpdateOne {
	jluo.mutation.ClearConcreteClass()
	return jluo
}

// SetWeekResult sets the "WeekResult" field.
func (jluo *JobLayerUpdateOne) SetWeekResult(s string) *JobLayerUpdateOne {
	jluo.mutation.SetWeekResult(s)
	return jluo
}

// SetNillableWeekResult sets the "WeekResult" field if the given value is not nil.
func (jluo *JobLayerUpdateOne) SetNillableWeekResult(s *string) *JobLayerUpdateOne {
	if s != nil {
		jluo.SetWeekResult(*s)
	}
	return jluo
}

// ClearWeekResult clears the value of the "WeekResult" field.
func (jluo *JobLayerUpdateOne) ClearWeekResult() *JobLayerUpdateOne {
	jluo.mutation.ClearWeekResult()
	return jluo
}

// SetMonthResult sets the "MonthResult" field.
func (jluo *JobLayerUpdateOne) SetMonthResult(s string) *JobLayerUpdateOne {
	jluo.mutation.SetMonthResult(s)
	return jluo
}

// SetNillableMonthResult sets the "MonthResult" field if the given value is not nil.
func (jluo *JobLayerUpdateOne) SetNillableMonthResult(s *string) *JobLayerUpdateOne {
	if s != nil {
		jluo.SetMonthResult(*s)
	}
	return jluo
}

// ClearMonthResult clears the value of the "MonthResult" field.
func (jluo *JobLayerUpdateOne) ClearMonthResult() *JobLayerUpdateOne {
	jluo.mutation.ClearMonthResult()
	return jluo
}

// SetCreatedAt sets the "created_at" field.
func (jluo *JobLayerUpdateOne) SetCreatedAt(t time.Time) *JobLayerUpdateOne {
	jluo.mutation.SetCreatedAt(t)
	return jluo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (jluo *JobLayerUpdateOne) SetNillableCreatedAt(t *time.Time) *JobLayerUpdateOne {
	if t != nil {
		jluo.SetCreatedAt(*t)
	}
	return jluo
}

// SetUpdatedAt sets the "updated_at" field.
func (jluo *JobLayerUpdateOne) SetUpdatedAt(t time.Time) *JobLayerUpdateOne {
	jluo.mutation.SetUpdatedAt(t)
	return jluo
}

// SetLayerID sets the "layer" edge to the JobDetail entity by ID.
func (jluo *JobLayerUpdateOne) SetLayerID(id int) *JobLayerUpdateOne {
	jluo.mutation.SetLayerID(id)
	return jluo
}

// SetNillableLayerID sets the "layer" edge to the JobDetail entity by ID if the given value is not nil.
func (jluo *JobLayerUpdateOne) SetNillableLayerID(id *int) *JobLayerUpdateOne {
	if id != nil {
		jluo = jluo.SetLayerID(*id)
	}
	return jluo
}

// SetLayer sets the "layer" edge to the JobDetail entity.
func (jluo *JobLayerUpdateOne) SetLayer(j *JobDetail) *JobLayerUpdateOne {
	return jluo.SetLayerID(j.ID)
}

// Mutation returns the JobLayerMutation object of the builder.
func (jluo *JobLayerUpdateOne) Mutation() *JobLayerMutation {
	return jluo.mutation
}

// ClearLayer clears the "layer" edge to the JobDetail entity.
func (jluo *JobLayerUpdateOne) ClearLayer() *JobLayerUpdateOne {
	jluo.mutation.ClearLayer()
	return jluo
}

// Where appends a list predicates to the JobLayerUpdate builder.
func (jluo *JobLayerUpdateOne) Where(ps ...predicate.JobLayer) *JobLayerUpdateOne {
	jluo.mutation.Where(ps...)
	return jluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (jluo *JobLayerUpdateOne) Select(field string, fields ...string) *JobLayerUpdateOne {
	jluo.fields = append([]string{field}, fields...)
	return jluo
}

// Save executes the query and returns the updated JobLayer entity.
func (jluo *JobLayerUpdateOne) Save(ctx context.Context) (*JobLayer, error) {
	jluo.defaults()
	return withHooks(ctx, jluo.sqlSave, jluo.mutation, jluo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (jluo *JobLayerUpdateOne) SaveX(ctx context.Context) *JobLayer {
	node, err := jluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (jluo *JobLayerUpdateOne) Exec(ctx context.Context) error {
	_, err := jluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jluo *JobLayerUpdateOne) ExecX(ctx context.Context) {
	if err := jluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (jluo *JobLayerUpdateOne) defaults() {
	if _, ok := jluo.mutation.UpdatedAt(); !ok {
		v := joblayer.UpdateDefaultUpdatedAt()
		jluo.mutation.SetUpdatedAt(v)
	}
}

func (jluo *JobLayerUpdateOne) sqlSave(ctx context.Context) (_node *JobLayer, err error) {
	_spec := sqlgraph.NewUpdateSpec(joblayer.Table, joblayer.Columns, sqlgraph.NewFieldSpec(joblayer.FieldID, field.TypeInt))
	id, ok := jluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "JobLayer.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := jluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, joblayer.FieldID)
		for _, f := range fields {
			if !joblayer.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != joblayer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := jluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := jluo.mutation.Name(); ok {
		_spec.SetField(joblayer.FieldName, field.TypeString, value)
	}
	if value, ok := jluo.mutation.Metre(); ok {
		_spec.SetField(joblayer.FieldMetre, field.TypeString, value)
	}
	if value, ok := jluo.mutation.MoldDate(); ok {
		_spec.SetField(joblayer.FieldMoldDate, field.TypeTime, value)
	}
	if jluo.mutation.MoldDateCleared() {
		_spec.ClearField(joblayer.FieldMoldDate, field.TypeTime)
	}
	if value, ok := jluo.mutation.ConcreteDate(); ok {
		_spec.SetField(joblayer.FieldConcreteDate, field.TypeTime, value)
	}
	if jluo.mutation.ConcreteDateCleared() {
		_spec.ClearField(joblayer.FieldConcreteDate, field.TypeTime)
	}
	if value, ok := jluo.mutation.Samples(); ok {
		_spec.SetField(joblayer.FieldSamples, field.TypeInt, value)
	}
	if value, ok := jluo.mutation.AddedSamples(); ok {
		_spec.AddField(joblayer.FieldSamples, field.TypeInt, value)
	}
	if jluo.mutation.SamplesCleared() {
		_spec.ClearField(joblayer.FieldSamples, field.TypeInt)
	}
	if value, ok := jluo.mutation.ConcreteClass(); ok {
		_spec.SetField(joblayer.FieldConcreteClass, field.TypeString, value)
	}
	if jluo.mutation.ConcreteClassCleared() {
		_spec.ClearField(joblayer.FieldConcreteClass, field.TypeString)
	}
	if value, ok := jluo.mutation.WeekResult(); ok {
		_spec.SetField(joblayer.FieldWeekResult, field.TypeString, value)
	}
	if jluo.mutation.WeekResultCleared() {
		_spec.ClearField(joblayer.FieldWeekResult, field.TypeString)
	}
	if value, ok := jluo.mutation.MonthResult(); ok {
		_spec.SetField(joblayer.FieldMonthResult, field.TypeString, value)
	}
	if jluo.mutation.MonthResultCleared() {
		_spec.ClearField(joblayer.FieldMonthResult, field.TypeString)
	}
	if value, ok := jluo.mutation.CreatedAt(); ok {
		_spec.SetField(joblayer.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := jluo.mutation.UpdatedAt(); ok {
		_spec.SetField(joblayer.FieldUpdatedAt, field.TypeTime, value)
	}
	if jluo.mutation.LayerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   joblayer.LayerTable,
			Columns: []string{joblayer.LayerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(jobdetail.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jluo.mutation.LayerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   joblayer.LayerTable,
			Columns: []string{joblayer.LayerColumn},
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
	_node = &JobLayer{config: jluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, jluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{joblayer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	jluo.mutation.done = true
	return _node, nil
}
