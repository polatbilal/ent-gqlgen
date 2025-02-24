// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/polatbilal/gqlgen-ent/ent/jobdetail"
	"github.com/polatbilal/gqlgen-ent/ent/jobsupervisor"
	"github.com/polatbilal/gqlgen-ent/ent/predicate"
)

// JobSuperVisorQuery is the builder for querying JobSuperVisor entities.
type JobSuperVisorQuery struct {
	config
	ctx                  *QueryContext
	order                []jobsupervisor.OrderOption
	inters               []Interceptor
	predicates           []predicate.JobSuperVisor
	withSupervisors      *JobDetailQuery
	modifiers            []func(*sql.Selector)
	loadTotal            []func(context.Context, []*JobSuperVisor) error
	withNamedSupervisors map[string]*JobDetailQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the JobSuperVisorQuery builder.
func (jsvq *JobSuperVisorQuery) Where(ps ...predicate.JobSuperVisor) *JobSuperVisorQuery {
	jsvq.predicates = append(jsvq.predicates, ps...)
	return jsvq
}

// Limit the number of records to be returned by this query.
func (jsvq *JobSuperVisorQuery) Limit(limit int) *JobSuperVisorQuery {
	jsvq.ctx.Limit = &limit
	return jsvq
}

// Offset to start from.
func (jsvq *JobSuperVisorQuery) Offset(offset int) *JobSuperVisorQuery {
	jsvq.ctx.Offset = &offset
	return jsvq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (jsvq *JobSuperVisorQuery) Unique(unique bool) *JobSuperVisorQuery {
	jsvq.ctx.Unique = &unique
	return jsvq
}

// Order specifies how the records should be ordered.
func (jsvq *JobSuperVisorQuery) Order(o ...jobsupervisor.OrderOption) *JobSuperVisorQuery {
	jsvq.order = append(jsvq.order, o...)
	return jsvq
}

// QuerySupervisors chains the current query on the "supervisors" edge.
func (jsvq *JobSuperVisorQuery) QuerySupervisors() *JobDetailQuery {
	query := (&JobDetailClient{config: jsvq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := jsvq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := jsvq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(jobsupervisor.Table, jobsupervisor.FieldID, selector),
			sqlgraph.To(jobdetail.Table, jobdetail.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, jobsupervisor.SupervisorsTable, jobsupervisor.SupervisorsColumn),
		)
		fromU = sqlgraph.SetNeighbors(jsvq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first JobSuperVisor entity from the query.
// Returns a *NotFoundError when no JobSuperVisor was found.
func (jsvq *JobSuperVisorQuery) First(ctx context.Context) (*JobSuperVisor, error) {
	nodes, err := jsvq.Limit(1).All(setContextOp(ctx, jsvq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{jobsupervisor.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (jsvq *JobSuperVisorQuery) FirstX(ctx context.Context) *JobSuperVisor {
	node, err := jsvq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first JobSuperVisor ID from the query.
// Returns a *NotFoundError when no JobSuperVisor ID was found.
func (jsvq *JobSuperVisorQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = jsvq.Limit(1).IDs(setContextOp(ctx, jsvq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{jobsupervisor.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (jsvq *JobSuperVisorQuery) FirstIDX(ctx context.Context) int {
	id, err := jsvq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single JobSuperVisor entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one JobSuperVisor entity is found.
// Returns a *NotFoundError when no JobSuperVisor entities are found.
func (jsvq *JobSuperVisorQuery) Only(ctx context.Context) (*JobSuperVisor, error) {
	nodes, err := jsvq.Limit(2).All(setContextOp(ctx, jsvq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{jobsupervisor.Label}
	default:
		return nil, &NotSingularError{jobsupervisor.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (jsvq *JobSuperVisorQuery) OnlyX(ctx context.Context) *JobSuperVisor {
	node, err := jsvq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only JobSuperVisor ID in the query.
// Returns a *NotSingularError when more than one JobSuperVisor ID is found.
// Returns a *NotFoundError when no entities are found.
func (jsvq *JobSuperVisorQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = jsvq.Limit(2).IDs(setContextOp(ctx, jsvq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{jobsupervisor.Label}
	default:
		err = &NotSingularError{jobsupervisor.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (jsvq *JobSuperVisorQuery) OnlyIDX(ctx context.Context) int {
	id, err := jsvq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of JobSuperVisors.
func (jsvq *JobSuperVisorQuery) All(ctx context.Context) ([]*JobSuperVisor, error) {
	ctx = setContextOp(ctx, jsvq.ctx, ent.OpQueryAll)
	if err := jsvq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*JobSuperVisor, *JobSuperVisorQuery]()
	return withInterceptors[[]*JobSuperVisor](ctx, jsvq, qr, jsvq.inters)
}

// AllX is like All, but panics if an error occurs.
func (jsvq *JobSuperVisorQuery) AllX(ctx context.Context) []*JobSuperVisor {
	nodes, err := jsvq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of JobSuperVisor IDs.
func (jsvq *JobSuperVisorQuery) IDs(ctx context.Context) (ids []int, err error) {
	if jsvq.ctx.Unique == nil && jsvq.path != nil {
		jsvq.Unique(true)
	}
	ctx = setContextOp(ctx, jsvq.ctx, ent.OpQueryIDs)
	if err = jsvq.Select(jobsupervisor.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (jsvq *JobSuperVisorQuery) IDsX(ctx context.Context) []int {
	ids, err := jsvq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (jsvq *JobSuperVisorQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, jsvq.ctx, ent.OpQueryCount)
	if err := jsvq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, jsvq, querierCount[*JobSuperVisorQuery](), jsvq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (jsvq *JobSuperVisorQuery) CountX(ctx context.Context) int {
	count, err := jsvq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (jsvq *JobSuperVisorQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, jsvq.ctx, ent.OpQueryExist)
	switch _, err := jsvq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (jsvq *JobSuperVisorQuery) ExistX(ctx context.Context) bool {
	exist, err := jsvq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the JobSuperVisorQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (jsvq *JobSuperVisorQuery) Clone() *JobSuperVisorQuery {
	if jsvq == nil {
		return nil
	}
	return &JobSuperVisorQuery{
		config:          jsvq.config,
		ctx:             jsvq.ctx.Clone(),
		order:           append([]jobsupervisor.OrderOption{}, jsvq.order...),
		inters:          append([]Interceptor{}, jsvq.inters...),
		predicates:      append([]predicate.JobSuperVisor{}, jsvq.predicates...),
		withSupervisors: jsvq.withSupervisors.Clone(),
		// clone intermediate query.
		sql:  jsvq.sql.Clone(),
		path: jsvq.path,
	}
}

// WithSupervisors tells the query-builder to eager-load the nodes that are connected to
// the "supervisors" edge. The optional arguments are used to configure the query builder of the edge.
func (jsvq *JobSuperVisorQuery) WithSupervisors(opts ...func(*JobDetailQuery)) *JobSuperVisorQuery {
	query := (&JobDetailClient{config: jsvq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	jsvq.withSupervisors = query
	return jsvq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"Name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.JobSuperVisor.Query().
//		GroupBy(jobsupervisor.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (jsvq *JobSuperVisorQuery) GroupBy(field string, fields ...string) *JobSuperVisorGroupBy {
	jsvq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &JobSuperVisorGroupBy{build: jsvq}
	grbuild.flds = &jsvq.ctx.Fields
	grbuild.label = jobsupervisor.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"Name,omitempty"`
//	}
//
//	client.JobSuperVisor.Query().
//		Select(jobsupervisor.FieldName).
//		Scan(ctx, &v)
func (jsvq *JobSuperVisorQuery) Select(fields ...string) *JobSuperVisorSelect {
	jsvq.ctx.Fields = append(jsvq.ctx.Fields, fields...)
	sbuild := &JobSuperVisorSelect{JobSuperVisorQuery: jsvq}
	sbuild.label = jobsupervisor.Label
	sbuild.flds, sbuild.scan = &jsvq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a JobSuperVisorSelect configured with the given aggregations.
func (jsvq *JobSuperVisorQuery) Aggregate(fns ...AggregateFunc) *JobSuperVisorSelect {
	return jsvq.Select().Aggregate(fns...)
}

func (jsvq *JobSuperVisorQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range jsvq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, jsvq); err != nil {
				return err
			}
		}
	}
	for _, f := range jsvq.ctx.Fields {
		if !jobsupervisor.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if jsvq.path != nil {
		prev, err := jsvq.path(ctx)
		if err != nil {
			return err
		}
		jsvq.sql = prev
	}
	return nil
}

func (jsvq *JobSuperVisorQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*JobSuperVisor, error) {
	var (
		nodes       = []*JobSuperVisor{}
		_spec       = jsvq.querySpec()
		loadedTypes = [1]bool{
			jsvq.withSupervisors != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*JobSuperVisor).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &JobSuperVisor{config: jsvq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(jsvq.modifiers) > 0 {
		_spec.Modifiers = jsvq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, jsvq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := jsvq.withSupervisors; query != nil {
		if err := jsvq.loadSupervisors(ctx, query, nodes,
			func(n *JobSuperVisor) { n.Edges.Supervisors = []*JobDetail{} },
			func(n *JobSuperVisor, e *JobDetail) { n.Edges.Supervisors = append(n.Edges.Supervisors, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range jsvq.withNamedSupervisors {
		if err := jsvq.loadSupervisors(ctx, query, nodes,
			func(n *JobSuperVisor) { n.appendNamedSupervisors(name) },
			func(n *JobSuperVisor, e *JobDetail) { n.appendNamedSupervisors(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range jsvq.loadTotal {
		if err := jsvq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (jsvq *JobSuperVisorQuery) loadSupervisors(ctx context.Context, query *JobDetailQuery, nodes []*JobSuperVisor, init func(*JobSuperVisor), assign func(*JobSuperVisor, *JobDetail)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*JobSuperVisor)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.JobDetail(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(jobsupervisor.SupervisorsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.supervisor_id
		if fk == nil {
			return fmt.Errorf(`foreign-key "supervisor_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "supervisor_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (jsvq *JobSuperVisorQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := jsvq.querySpec()
	if len(jsvq.modifiers) > 0 {
		_spec.Modifiers = jsvq.modifiers
	}
	_spec.Node.Columns = jsvq.ctx.Fields
	if len(jsvq.ctx.Fields) > 0 {
		_spec.Unique = jsvq.ctx.Unique != nil && *jsvq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, jsvq.driver, _spec)
}

func (jsvq *JobSuperVisorQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(jobsupervisor.Table, jobsupervisor.Columns, sqlgraph.NewFieldSpec(jobsupervisor.FieldID, field.TypeInt))
	_spec.From = jsvq.sql
	if unique := jsvq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if jsvq.path != nil {
		_spec.Unique = true
	}
	if fields := jsvq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, jobsupervisor.FieldID)
		for i := range fields {
			if fields[i] != jobsupervisor.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := jsvq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := jsvq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := jsvq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := jsvq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (jsvq *JobSuperVisorQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(jsvq.driver.Dialect())
	t1 := builder.Table(jobsupervisor.Table)
	columns := jsvq.ctx.Fields
	if len(columns) == 0 {
		columns = jobsupervisor.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if jsvq.sql != nil {
		selector = jsvq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if jsvq.ctx.Unique != nil && *jsvq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range jsvq.predicates {
		p(selector)
	}
	for _, p := range jsvq.order {
		p(selector)
	}
	if offset := jsvq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := jsvq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedSupervisors tells the query-builder to eager-load the nodes that are connected to the "supervisors"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (jsvq *JobSuperVisorQuery) WithNamedSupervisors(name string, opts ...func(*JobDetailQuery)) *JobSuperVisorQuery {
	query := (&JobDetailClient{config: jsvq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if jsvq.withNamedSupervisors == nil {
		jsvq.withNamedSupervisors = make(map[string]*JobDetailQuery)
	}
	jsvq.withNamedSupervisors[name] = query
	return jsvq
}

// JobSuperVisorGroupBy is the group-by builder for JobSuperVisor entities.
type JobSuperVisorGroupBy struct {
	selector
	build *JobSuperVisorQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (jsvgb *JobSuperVisorGroupBy) Aggregate(fns ...AggregateFunc) *JobSuperVisorGroupBy {
	jsvgb.fns = append(jsvgb.fns, fns...)
	return jsvgb
}

// Scan applies the selector query and scans the result into the given value.
func (jsvgb *JobSuperVisorGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, jsvgb.build.ctx, ent.OpQueryGroupBy)
	if err := jsvgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*JobSuperVisorQuery, *JobSuperVisorGroupBy](ctx, jsvgb.build, jsvgb, jsvgb.build.inters, v)
}

func (jsvgb *JobSuperVisorGroupBy) sqlScan(ctx context.Context, root *JobSuperVisorQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(jsvgb.fns))
	for _, fn := range jsvgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*jsvgb.flds)+len(jsvgb.fns))
		for _, f := range *jsvgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*jsvgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := jsvgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// JobSuperVisorSelect is the builder for selecting fields of JobSuperVisor entities.
type JobSuperVisorSelect struct {
	*JobSuperVisorQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (jsvs *JobSuperVisorSelect) Aggregate(fns ...AggregateFunc) *JobSuperVisorSelect {
	jsvs.fns = append(jsvs.fns, fns...)
	return jsvs
}

// Scan applies the selector query and scans the result into the given value.
func (jsvs *JobSuperVisorSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, jsvs.ctx, ent.OpQuerySelect)
	if err := jsvs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*JobSuperVisorQuery, *JobSuperVisorSelect](ctx, jsvs.JobSuperVisorQuery, jsvs, jsvs.inters, v)
}

func (jsvs *JobSuperVisorSelect) sqlScan(ctx context.Context, root *JobSuperVisorQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(jsvs.fns))
	for _, fn := range jsvs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*jsvs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := jsvs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
