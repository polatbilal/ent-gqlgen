// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"gqlgen-ent/ent/jobdetail"
	"gqlgen-ent/ent/jobprogress"
	"gqlgen-ent/ent/predicate"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// JobProgressQuery is the builder for querying JobProgress entities.
type JobProgressQuery struct {
	config
	ctx               *QueryContext
	order             []jobprogress.OrderOption
	inters            []Interceptor
	predicates        []predicate.JobProgress
	withProgress      *JobDetailQuery
	modifiers         []func(*sql.Selector)
	loadTotal         []func(context.Context, []*JobProgress) error
	withNamedProgress map[string]*JobDetailQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the JobProgressQuery builder.
func (jpq *JobProgressQuery) Where(ps ...predicate.JobProgress) *JobProgressQuery {
	jpq.predicates = append(jpq.predicates, ps...)
	return jpq
}

// Limit the number of records to be returned by this query.
func (jpq *JobProgressQuery) Limit(limit int) *JobProgressQuery {
	jpq.ctx.Limit = &limit
	return jpq
}

// Offset to start from.
func (jpq *JobProgressQuery) Offset(offset int) *JobProgressQuery {
	jpq.ctx.Offset = &offset
	return jpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (jpq *JobProgressQuery) Unique(unique bool) *JobProgressQuery {
	jpq.ctx.Unique = &unique
	return jpq
}

// Order specifies how the records should be ordered.
func (jpq *JobProgressQuery) Order(o ...jobprogress.OrderOption) *JobProgressQuery {
	jpq.order = append(jpq.order, o...)
	return jpq
}

// QueryProgress chains the current query on the "progress" edge.
func (jpq *JobProgressQuery) QueryProgress() *JobDetailQuery {
	query := (&JobDetailClient{config: jpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := jpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := jpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(jobprogress.Table, jobprogress.FieldID, selector),
			sqlgraph.To(jobdetail.Table, jobdetail.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, jobprogress.ProgressTable, jobprogress.ProgressColumn),
		)
		fromU = sqlgraph.SetNeighbors(jpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first JobProgress entity from the query.
// Returns a *NotFoundError when no JobProgress was found.
func (jpq *JobProgressQuery) First(ctx context.Context) (*JobProgress, error) {
	nodes, err := jpq.Limit(1).All(setContextOp(ctx, jpq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{jobprogress.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (jpq *JobProgressQuery) FirstX(ctx context.Context) *JobProgress {
	node, err := jpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first JobProgress ID from the query.
// Returns a *NotFoundError when no JobProgress ID was found.
func (jpq *JobProgressQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = jpq.Limit(1).IDs(setContextOp(ctx, jpq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{jobprogress.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (jpq *JobProgressQuery) FirstIDX(ctx context.Context) int {
	id, err := jpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single JobProgress entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one JobProgress entity is found.
// Returns a *NotFoundError when no JobProgress entities are found.
func (jpq *JobProgressQuery) Only(ctx context.Context) (*JobProgress, error) {
	nodes, err := jpq.Limit(2).All(setContextOp(ctx, jpq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{jobprogress.Label}
	default:
		return nil, &NotSingularError{jobprogress.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (jpq *JobProgressQuery) OnlyX(ctx context.Context) *JobProgress {
	node, err := jpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only JobProgress ID in the query.
// Returns a *NotSingularError when more than one JobProgress ID is found.
// Returns a *NotFoundError when no entities are found.
func (jpq *JobProgressQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = jpq.Limit(2).IDs(setContextOp(ctx, jpq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{jobprogress.Label}
	default:
		err = &NotSingularError{jobprogress.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (jpq *JobProgressQuery) OnlyIDX(ctx context.Context) int {
	id, err := jpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of JobProgresses.
func (jpq *JobProgressQuery) All(ctx context.Context) ([]*JobProgress, error) {
	ctx = setContextOp(ctx, jpq.ctx, "All")
	if err := jpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*JobProgress, *JobProgressQuery]()
	return withInterceptors[[]*JobProgress](ctx, jpq, qr, jpq.inters)
}

// AllX is like All, but panics if an error occurs.
func (jpq *JobProgressQuery) AllX(ctx context.Context) []*JobProgress {
	nodes, err := jpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of JobProgress IDs.
func (jpq *JobProgressQuery) IDs(ctx context.Context) (ids []int, err error) {
	if jpq.ctx.Unique == nil && jpq.path != nil {
		jpq.Unique(true)
	}
	ctx = setContextOp(ctx, jpq.ctx, "IDs")
	if err = jpq.Select(jobprogress.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (jpq *JobProgressQuery) IDsX(ctx context.Context) []int {
	ids, err := jpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (jpq *JobProgressQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, jpq.ctx, "Count")
	if err := jpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, jpq, querierCount[*JobProgressQuery](), jpq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (jpq *JobProgressQuery) CountX(ctx context.Context) int {
	count, err := jpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (jpq *JobProgressQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, jpq.ctx, "Exist")
	switch _, err := jpq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (jpq *JobProgressQuery) ExistX(ctx context.Context) bool {
	exist, err := jpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the JobProgressQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (jpq *JobProgressQuery) Clone() *JobProgressQuery {
	if jpq == nil {
		return nil
	}
	return &JobProgressQuery{
		config:       jpq.config,
		ctx:          jpq.ctx.Clone(),
		order:        append([]jobprogress.OrderOption{}, jpq.order...),
		inters:       append([]Interceptor{}, jpq.inters...),
		predicates:   append([]predicate.JobProgress{}, jpq.predicates...),
		withProgress: jpq.withProgress.Clone(),
		// clone intermediate query.
		sql:  jpq.sql.Clone(),
		path: jpq.path,
	}
}

// WithProgress tells the query-builder to eager-load the nodes that are connected to
// the "progress" edge. The optional arguments are used to configure the query builder of the edge.
func (jpq *JobProgressQuery) WithProgress(opts ...func(*JobDetailQuery)) *JobProgressQuery {
	query := (&JobDetailClient{config: jpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	jpq.withProgress = query
	return jpq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		One int `json:"One,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.JobProgress.Query().
//		GroupBy(jobprogress.FieldOne).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (jpq *JobProgressQuery) GroupBy(field string, fields ...string) *JobProgressGroupBy {
	jpq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &JobProgressGroupBy{build: jpq}
	grbuild.flds = &jpq.ctx.Fields
	grbuild.label = jobprogress.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		One int `json:"One,omitempty"`
//	}
//
//	client.JobProgress.Query().
//		Select(jobprogress.FieldOne).
//		Scan(ctx, &v)
func (jpq *JobProgressQuery) Select(fields ...string) *JobProgressSelect {
	jpq.ctx.Fields = append(jpq.ctx.Fields, fields...)
	sbuild := &JobProgressSelect{JobProgressQuery: jpq}
	sbuild.label = jobprogress.Label
	sbuild.flds, sbuild.scan = &jpq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a JobProgressSelect configured with the given aggregations.
func (jpq *JobProgressQuery) Aggregate(fns ...AggregateFunc) *JobProgressSelect {
	return jpq.Select().Aggregate(fns...)
}

func (jpq *JobProgressQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range jpq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, jpq); err != nil {
				return err
			}
		}
	}
	for _, f := range jpq.ctx.Fields {
		if !jobprogress.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if jpq.path != nil {
		prev, err := jpq.path(ctx)
		if err != nil {
			return err
		}
		jpq.sql = prev
	}
	return nil
}

func (jpq *JobProgressQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*JobProgress, error) {
	var (
		nodes       = []*JobProgress{}
		_spec       = jpq.querySpec()
		loadedTypes = [1]bool{
			jpq.withProgress != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*JobProgress).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &JobProgress{config: jpq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(jpq.modifiers) > 0 {
		_spec.Modifiers = jpq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, jpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := jpq.withProgress; query != nil {
		if err := jpq.loadProgress(ctx, query, nodes,
			func(n *JobProgress) { n.Edges.Progress = []*JobDetail{} },
			func(n *JobProgress, e *JobDetail) { n.Edges.Progress = append(n.Edges.Progress, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range jpq.withNamedProgress {
		if err := jpq.loadProgress(ctx, query, nodes,
			func(n *JobProgress) { n.appendNamedProgress(name) },
			func(n *JobProgress, e *JobDetail) { n.appendNamedProgress(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range jpq.loadTotal {
		if err := jpq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (jpq *JobProgressQuery) loadProgress(ctx context.Context, query *JobDetailQuery, nodes []*JobProgress, init func(*JobProgress), assign func(*JobProgress, *JobDetail)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*JobProgress)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.JobDetail(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(jobprogress.ProgressColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.progress_id
		if fk == nil {
			return fmt.Errorf(`foreign-key "progress_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "progress_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (jpq *JobProgressQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := jpq.querySpec()
	if len(jpq.modifiers) > 0 {
		_spec.Modifiers = jpq.modifiers
	}
	_spec.Node.Columns = jpq.ctx.Fields
	if len(jpq.ctx.Fields) > 0 {
		_spec.Unique = jpq.ctx.Unique != nil && *jpq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, jpq.driver, _spec)
}

func (jpq *JobProgressQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(jobprogress.Table, jobprogress.Columns, sqlgraph.NewFieldSpec(jobprogress.FieldID, field.TypeInt))
	_spec.From = jpq.sql
	if unique := jpq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if jpq.path != nil {
		_spec.Unique = true
	}
	if fields := jpq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, jobprogress.FieldID)
		for i := range fields {
			if fields[i] != jobprogress.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := jpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := jpq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := jpq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := jpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (jpq *JobProgressQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(jpq.driver.Dialect())
	t1 := builder.Table(jobprogress.Table)
	columns := jpq.ctx.Fields
	if len(columns) == 0 {
		columns = jobprogress.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if jpq.sql != nil {
		selector = jpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if jpq.ctx.Unique != nil && *jpq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range jpq.predicates {
		p(selector)
	}
	for _, p := range jpq.order {
		p(selector)
	}
	if offset := jpq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := jpq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedProgress tells the query-builder to eager-load the nodes that are connected to the "progress"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (jpq *JobProgressQuery) WithNamedProgress(name string, opts ...func(*JobDetailQuery)) *JobProgressQuery {
	query := (&JobDetailClient{config: jpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if jpq.withNamedProgress == nil {
		jpq.withNamedProgress = make(map[string]*JobDetailQuery)
	}
	jpq.withNamedProgress[name] = query
	return jpq
}

// JobProgressGroupBy is the group-by builder for JobProgress entities.
type JobProgressGroupBy struct {
	selector
	build *JobProgressQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (jpgb *JobProgressGroupBy) Aggregate(fns ...AggregateFunc) *JobProgressGroupBy {
	jpgb.fns = append(jpgb.fns, fns...)
	return jpgb
}

// Scan applies the selector query and scans the result into the given value.
func (jpgb *JobProgressGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, jpgb.build.ctx, "GroupBy")
	if err := jpgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*JobProgressQuery, *JobProgressGroupBy](ctx, jpgb.build, jpgb, jpgb.build.inters, v)
}

func (jpgb *JobProgressGroupBy) sqlScan(ctx context.Context, root *JobProgressQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(jpgb.fns))
	for _, fn := range jpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*jpgb.flds)+len(jpgb.fns))
		for _, f := range *jpgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*jpgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := jpgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// JobProgressSelect is the builder for selecting fields of JobProgress entities.
type JobProgressSelect struct {
	*JobProgressQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (jps *JobProgressSelect) Aggregate(fns ...AggregateFunc) *JobProgressSelect {
	jps.fns = append(jps.fns, fns...)
	return jps
}

// Scan applies the selector query and scans the result into the given value.
func (jps *JobProgressSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, jps.ctx, "Select")
	if err := jps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*JobProgressQuery, *JobProgressSelect](ctx, jps.JobProgressQuery, jps, jps.inters, v)
}

func (jps *JobProgressSelect) sqlScan(ctx context.Context, root *JobProgressQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(jps.fns))
	for _, fn := range jps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*jps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := jps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
