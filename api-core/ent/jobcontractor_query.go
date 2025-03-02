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
	"github.com/polatbilal/gqlgen-ent/api-core/ent/jobcontractor"
	"github.com/polatbilal/gqlgen-ent/api-core/ent/jobrelations"
	"github.com/polatbilal/gqlgen-ent/api-core/ent/predicate"
)

// JobContractorQuery is the builder for querying JobContractor entities.
type JobContractorQuery struct {
	config
	ctx                  *QueryContext
	order                []jobcontractor.OrderOption
	inters               []Interceptor
	predicates           []predicate.JobContractor
	withContractors      *JobRelationsQuery
	modifiers            []func(*sql.Selector)
	loadTotal            []func(context.Context, []*JobContractor) error
	withNamedContractors map[string]*JobRelationsQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the JobContractorQuery builder.
func (jcq *JobContractorQuery) Where(ps ...predicate.JobContractor) *JobContractorQuery {
	jcq.predicates = append(jcq.predicates, ps...)
	return jcq
}

// Limit the number of records to be returned by this query.
func (jcq *JobContractorQuery) Limit(limit int) *JobContractorQuery {
	jcq.ctx.Limit = &limit
	return jcq
}

// Offset to start from.
func (jcq *JobContractorQuery) Offset(offset int) *JobContractorQuery {
	jcq.ctx.Offset = &offset
	return jcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (jcq *JobContractorQuery) Unique(unique bool) *JobContractorQuery {
	jcq.ctx.Unique = &unique
	return jcq
}

// Order specifies how the records should be ordered.
func (jcq *JobContractorQuery) Order(o ...jobcontractor.OrderOption) *JobContractorQuery {
	jcq.order = append(jcq.order, o...)
	return jcq
}

// QueryContractors chains the current query on the "contractors" edge.
func (jcq *JobContractorQuery) QueryContractors() *JobRelationsQuery {
	query := (&JobRelationsClient{config: jcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := jcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := jcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(jobcontractor.Table, jobcontractor.FieldID, selector),
			sqlgraph.To(jobrelations.Table, jobrelations.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, jobcontractor.ContractorsTable, jobcontractor.ContractorsColumn),
		)
		fromU = sqlgraph.SetNeighbors(jcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first JobContractor entity from the query.
// Returns a *NotFoundError when no JobContractor was found.
func (jcq *JobContractorQuery) First(ctx context.Context) (*JobContractor, error) {
	nodes, err := jcq.Limit(1).All(setContextOp(ctx, jcq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{jobcontractor.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (jcq *JobContractorQuery) FirstX(ctx context.Context) *JobContractor {
	node, err := jcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first JobContractor ID from the query.
// Returns a *NotFoundError when no JobContractor ID was found.
func (jcq *JobContractorQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = jcq.Limit(1).IDs(setContextOp(ctx, jcq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{jobcontractor.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (jcq *JobContractorQuery) FirstIDX(ctx context.Context) int {
	id, err := jcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single JobContractor entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one JobContractor entity is found.
// Returns a *NotFoundError when no JobContractor entities are found.
func (jcq *JobContractorQuery) Only(ctx context.Context) (*JobContractor, error) {
	nodes, err := jcq.Limit(2).All(setContextOp(ctx, jcq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{jobcontractor.Label}
	default:
		return nil, &NotSingularError{jobcontractor.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (jcq *JobContractorQuery) OnlyX(ctx context.Context) *JobContractor {
	node, err := jcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only JobContractor ID in the query.
// Returns a *NotSingularError when more than one JobContractor ID is found.
// Returns a *NotFoundError when no entities are found.
func (jcq *JobContractorQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = jcq.Limit(2).IDs(setContextOp(ctx, jcq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{jobcontractor.Label}
	default:
		err = &NotSingularError{jobcontractor.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (jcq *JobContractorQuery) OnlyIDX(ctx context.Context) int {
	id, err := jcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of JobContractors.
func (jcq *JobContractorQuery) All(ctx context.Context) ([]*JobContractor, error) {
	ctx = setContextOp(ctx, jcq.ctx, ent.OpQueryAll)
	if err := jcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*JobContractor, *JobContractorQuery]()
	return withInterceptors[[]*JobContractor](ctx, jcq, qr, jcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (jcq *JobContractorQuery) AllX(ctx context.Context) []*JobContractor {
	nodes, err := jcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of JobContractor IDs.
func (jcq *JobContractorQuery) IDs(ctx context.Context) (ids []int, err error) {
	if jcq.ctx.Unique == nil && jcq.path != nil {
		jcq.Unique(true)
	}
	ctx = setContextOp(ctx, jcq.ctx, ent.OpQueryIDs)
	if err = jcq.Select(jobcontractor.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (jcq *JobContractorQuery) IDsX(ctx context.Context) []int {
	ids, err := jcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (jcq *JobContractorQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, jcq.ctx, ent.OpQueryCount)
	if err := jcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, jcq, querierCount[*JobContractorQuery](), jcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (jcq *JobContractorQuery) CountX(ctx context.Context) int {
	count, err := jcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (jcq *JobContractorQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, jcq.ctx, ent.OpQueryExist)
	switch _, err := jcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (jcq *JobContractorQuery) ExistX(ctx context.Context) bool {
	exist, err := jcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the JobContractorQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (jcq *JobContractorQuery) Clone() *JobContractorQuery {
	if jcq == nil {
		return nil
	}
	return &JobContractorQuery{
		config:          jcq.config,
		ctx:             jcq.ctx.Clone(),
		order:           append([]jobcontractor.OrderOption{}, jcq.order...),
		inters:          append([]Interceptor{}, jcq.inters...),
		predicates:      append([]predicate.JobContractor{}, jcq.predicates...),
		withContractors: jcq.withContractors.Clone(),
		// clone intermediate query.
		sql:  jcq.sql.Clone(),
		path: jcq.path,
	}
}

// WithContractors tells the query-builder to eager-load the nodes that are connected to
// the "contractors" edge. The optional arguments are used to configure the query builder of the edge.
func (jcq *JobContractorQuery) WithContractors(opts ...func(*JobRelationsQuery)) *JobContractorQuery {
	query := (&JobRelationsClient{config: jcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	jcq.withContractors = query
	return jcq
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
//	client.JobContractor.Query().
//		GroupBy(jobcontractor.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (jcq *JobContractorQuery) GroupBy(field string, fields ...string) *JobContractorGroupBy {
	jcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &JobContractorGroupBy{build: jcq}
	grbuild.flds = &jcq.ctx.Fields
	grbuild.label = jobcontractor.Label
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
//	client.JobContractor.Query().
//		Select(jobcontractor.FieldName).
//		Scan(ctx, &v)
func (jcq *JobContractorQuery) Select(fields ...string) *JobContractorSelect {
	jcq.ctx.Fields = append(jcq.ctx.Fields, fields...)
	sbuild := &JobContractorSelect{JobContractorQuery: jcq}
	sbuild.label = jobcontractor.Label
	sbuild.flds, sbuild.scan = &jcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a JobContractorSelect configured with the given aggregations.
func (jcq *JobContractorQuery) Aggregate(fns ...AggregateFunc) *JobContractorSelect {
	return jcq.Select().Aggregate(fns...)
}

func (jcq *JobContractorQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range jcq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, jcq); err != nil {
				return err
			}
		}
	}
	for _, f := range jcq.ctx.Fields {
		if !jobcontractor.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if jcq.path != nil {
		prev, err := jcq.path(ctx)
		if err != nil {
			return err
		}
		jcq.sql = prev
	}
	return nil
}

func (jcq *JobContractorQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*JobContractor, error) {
	var (
		nodes       = []*JobContractor{}
		_spec       = jcq.querySpec()
		loadedTypes = [1]bool{
			jcq.withContractors != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*JobContractor).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &JobContractor{config: jcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(jcq.modifiers) > 0 {
		_spec.Modifiers = jcq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, jcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := jcq.withContractors; query != nil {
		if err := jcq.loadContractors(ctx, query, nodes,
			func(n *JobContractor) { n.Edges.Contractors = []*JobRelations{} },
			func(n *JobContractor, e *JobRelations) { n.Edges.Contractors = append(n.Edges.Contractors, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range jcq.withNamedContractors {
		if err := jcq.loadContractors(ctx, query, nodes,
			func(n *JobContractor) { n.appendNamedContractors(name) },
			func(n *JobContractor, e *JobRelations) { n.appendNamedContractors(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range jcq.loadTotal {
		if err := jcq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (jcq *JobContractorQuery) loadContractors(ctx context.Context, query *JobRelationsQuery, nodes []*JobContractor, init func(*JobContractor), assign func(*JobContractor, *JobRelations)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*JobContractor)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.JobRelations(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(jobcontractor.ContractorsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.contractor_id
		if fk == nil {
			return fmt.Errorf(`foreign-key "contractor_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "contractor_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (jcq *JobContractorQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := jcq.querySpec()
	if len(jcq.modifiers) > 0 {
		_spec.Modifiers = jcq.modifiers
	}
	_spec.Node.Columns = jcq.ctx.Fields
	if len(jcq.ctx.Fields) > 0 {
		_spec.Unique = jcq.ctx.Unique != nil && *jcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, jcq.driver, _spec)
}

func (jcq *JobContractorQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(jobcontractor.Table, jobcontractor.Columns, sqlgraph.NewFieldSpec(jobcontractor.FieldID, field.TypeInt))
	_spec.From = jcq.sql
	if unique := jcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if jcq.path != nil {
		_spec.Unique = true
	}
	if fields := jcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, jobcontractor.FieldID)
		for i := range fields {
			if fields[i] != jobcontractor.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := jcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := jcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := jcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := jcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (jcq *JobContractorQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(jcq.driver.Dialect())
	t1 := builder.Table(jobcontractor.Table)
	columns := jcq.ctx.Fields
	if len(columns) == 0 {
		columns = jobcontractor.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if jcq.sql != nil {
		selector = jcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if jcq.ctx.Unique != nil && *jcq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range jcq.predicates {
		p(selector)
	}
	for _, p := range jcq.order {
		p(selector)
	}
	if offset := jcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := jcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedContractors tells the query-builder to eager-load the nodes that are connected to the "contractors"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (jcq *JobContractorQuery) WithNamedContractors(name string, opts ...func(*JobRelationsQuery)) *JobContractorQuery {
	query := (&JobRelationsClient{config: jcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if jcq.withNamedContractors == nil {
		jcq.withNamedContractors = make(map[string]*JobRelationsQuery)
	}
	jcq.withNamedContractors[name] = query
	return jcq
}

// JobContractorGroupBy is the group-by builder for JobContractor entities.
type JobContractorGroupBy struct {
	selector
	build *JobContractorQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (jcgb *JobContractorGroupBy) Aggregate(fns ...AggregateFunc) *JobContractorGroupBy {
	jcgb.fns = append(jcgb.fns, fns...)
	return jcgb
}

// Scan applies the selector query and scans the result into the given value.
func (jcgb *JobContractorGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, jcgb.build.ctx, ent.OpQueryGroupBy)
	if err := jcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*JobContractorQuery, *JobContractorGroupBy](ctx, jcgb.build, jcgb, jcgb.build.inters, v)
}

func (jcgb *JobContractorGroupBy) sqlScan(ctx context.Context, root *JobContractorQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(jcgb.fns))
	for _, fn := range jcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*jcgb.flds)+len(jcgb.fns))
		for _, f := range *jcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*jcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := jcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// JobContractorSelect is the builder for selecting fields of JobContractor entities.
type JobContractorSelect struct {
	*JobContractorQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (jcs *JobContractorSelect) Aggregate(fns ...AggregateFunc) *JobContractorSelect {
	jcs.fns = append(jcs.fns, fns...)
	return jcs
}

// Scan applies the selector query and scans the result into the given value.
func (jcs *JobContractorSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, jcs.ctx, ent.OpQuerySelect)
	if err := jcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*JobContractorQuery, *JobContractorSelect](ctx, jcs.JobContractorQuery, jcs, jcs.inters, v)
}

func (jcs *JobContractorSelect) sqlScan(ctx context.Context, root *JobContractorQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(jcs.fns))
	for _, fn := range jcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*jcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := jcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
