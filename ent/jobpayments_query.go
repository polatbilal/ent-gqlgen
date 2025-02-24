// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/polatbilal/gqlgen-ent/ent/jobdetail"
	"github.com/polatbilal/gqlgen-ent/ent/jobpayments"
	"github.com/polatbilal/gqlgen-ent/ent/predicate"
)

// JobPaymentsQuery is the builder for querying JobPayments entities.
type JobPaymentsQuery struct {
	config
	ctx          *QueryContext
	order        []jobpayments.OrderOption
	inters       []Interceptor
	predicates   []predicate.JobPayments
	withPayments *JobDetailQuery
	withFKs      bool
	modifiers    []func(*sql.Selector)
	loadTotal    []func(context.Context, []*JobPayments) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the JobPaymentsQuery builder.
func (jpq *JobPaymentsQuery) Where(ps ...predicate.JobPayments) *JobPaymentsQuery {
	jpq.predicates = append(jpq.predicates, ps...)
	return jpq
}

// Limit the number of records to be returned by this query.
func (jpq *JobPaymentsQuery) Limit(limit int) *JobPaymentsQuery {
	jpq.ctx.Limit = &limit
	return jpq
}

// Offset to start from.
func (jpq *JobPaymentsQuery) Offset(offset int) *JobPaymentsQuery {
	jpq.ctx.Offset = &offset
	return jpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (jpq *JobPaymentsQuery) Unique(unique bool) *JobPaymentsQuery {
	jpq.ctx.Unique = &unique
	return jpq
}

// Order specifies how the records should be ordered.
func (jpq *JobPaymentsQuery) Order(o ...jobpayments.OrderOption) *JobPaymentsQuery {
	jpq.order = append(jpq.order, o...)
	return jpq
}

// QueryPayments chains the current query on the "payments" edge.
func (jpq *JobPaymentsQuery) QueryPayments() *JobDetailQuery {
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
			sqlgraph.From(jobpayments.Table, jobpayments.FieldID, selector),
			sqlgraph.To(jobdetail.Table, jobdetail.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, jobpayments.PaymentsTable, jobpayments.PaymentsColumn),
		)
		fromU = sqlgraph.SetNeighbors(jpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first JobPayments entity from the query.
// Returns a *NotFoundError when no JobPayments was found.
func (jpq *JobPaymentsQuery) First(ctx context.Context) (*JobPayments, error) {
	nodes, err := jpq.Limit(1).All(setContextOp(ctx, jpq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{jobpayments.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (jpq *JobPaymentsQuery) FirstX(ctx context.Context) *JobPayments {
	node, err := jpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first JobPayments ID from the query.
// Returns a *NotFoundError when no JobPayments ID was found.
func (jpq *JobPaymentsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = jpq.Limit(1).IDs(setContextOp(ctx, jpq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{jobpayments.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (jpq *JobPaymentsQuery) FirstIDX(ctx context.Context) int {
	id, err := jpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single JobPayments entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one JobPayments entity is found.
// Returns a *NotFoundError when no JobPayments entities are found.
func (jpq *JobPaymentsQuery) Only(ctx context.Context) (*JobPayments, error) {
	nodes, err := jpq.Limit(2).All(setContextOp(ctx, jpq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{jobpayments.Label}
	default:
		return nil, &NotSingularError{jobpayments.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (jpq *JobPaymentsQuery) OnlyX(ctx context.Context) *JobPayments {
	node, err := jpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only JobPayments ID in the query.
// Returns a *NotSingularError when more than one JobPayments ID is found.
// Returns a *NotFoundError when no entities are found.
func (jpq *JobPaymentsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = jpq.Limit(2).IDs(setContextOp(ctx, jpq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{jobpayments.Label}
	default:
		err = &NotSingularError{jobpayments.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (jpq *JobPaymentsQuery) OnlyIDX(ctx context.Context) int {
	id, err := jpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of JobPaymentsSlice.
func (jpq *JobPaymentsQuery) All(ctx context.Context) ([]*JobPayments, error) {
	ctx = setContextOp(ctx, jpq.ctx, ent.OpQueryAll)
	if err := jpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*JobPayments, *JobPaymentsQuery]()
	return withInterceptors[[]*JobPayments](ctx, jpq, qr, jpq.inters)
}

// AllX is like All, but panics if an error occurs.
func (jpq *JobPaymentsQuery) AllX(ctx context.Context) []*JobPayments {
	nodes, err := jpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of JobPayments IDs.
func (jpq *JobPaymentsQuery) IDs(ctx context.Context) (ids []int, err error) {
	if jpq.ctx.Unique == nil && jpq.path != nil {
		jpq.Unique(true)
	}
	ctx = setContextOp(ctx, jpq.ctx, ent.OpQueryIDs)
	if err = jpq.Select(jobpayments.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (jpq *JobPaymentsQuery) IDsX(ctx context.Context) []int {
	ids, err := jpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (jpq *JobPaymentsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, jpq.ctx, ent.OpQueryCount)
	if err := jpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, jpq, querierCount[*JobPaymentsQuery](), jpq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (jpq *JobPaymentsQuery) CountX(ctx context.Context) int {
	count, err := jpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (jpq *JobPaymentsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, jpq.ctx, ent.OpQueryExist)
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
func (jpq *JobPaymentsQuery) ExistX(ctx context.Context) bool {
	exist, err := jpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the JobPaymentsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (jpq *JobPaymentsQuery) Clone() *JobPaymentsQuery {
	if jpq == nil {
		return nil
	}
	return &JobPaymentsQuery{
		config:       jpq.config,
		ctx:          jpq.ctx.Clone(),
		order:        append([]jobpayments.OrderOption{}, jpq.order...),
		inters:       append([]Interceptor{}, jpq.inters...),
		predicates:   append([]predicate.JobPayments{}, jpq.predicates...),
		withPayments: jpq.withPayments.Clone(),
		// clone intermediate query.
		sql:  jpq.sql.Clone(),
		path: jpq.path,
	}
}

// WithPayments tells the query-builder to eager-load the nodes that are connected to
// the "payments" edge. The optional arguments are used to configure the query builder of the edge.
func (jpq *JobPaymentsQuery) WithPayments(opts ...func(*JobDetailQuery)) *JobPaymentsQuery {
	query := (&JobDetailClient{config: jpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	jpq.withPayments = query
	return jpq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Date time.Time `json:"Date,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.JobPayments.Query().
//		GroupBy(jobpayments.FieldDate).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (jpq *JobPaymentsQuery) GroupBy(field string, fields ...string) *JobPaymentsGroupBy {
	jpq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &JobPaymentsGroupBy{build: jpq}
	grbuild.flds = &jpq.ctx.Fields
	grbuild.label = jobpayments.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Date time.Time `json:"Date,omitempty"`
//	}
//
//	client.JobPayments.Query().
//		Select(jobpayments.FieldDate).
//		Scan(ctx, &v)
func (jpq *JobPaymentsQuery) Select(fields ...string) *JobPaymentsSelect {
	jpq.ctx.Fields = append(jpq.ctx.Fields, fields...)
	sbuild := &JobPaymentsSelect{JobPaymentsQuery: jpq}
	sbuild.label = jobpayments.Label
	sbuild.flds, sbuild.scan = &jpq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a JobPaymentsSelect configured with the given aggregations.
func (jpq *JobPaymentsQuery) Aggregate(fns ...AggregateFunc) *JobPaymentsSelect {
	return jpq.Select().Aggregate(fns...)
}

func (jpq *JobPaymentsQuery) prepareQuery(ctx context.Context) error {
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
		if !jobpayments.ValidColumn(f) {
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

func (jpq *JobPaymentsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*JobPayments, error) {
	var (
		nodes       = []*JobPayments{}
		withFKs     = jpq.withFKs
		_spec       = jpq.querySpec()
		loadedTypes = [1]bool{
			jpq.withPayments != nil,
		}
	)
	if jpq.withPayments != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, jobpayments.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*JobPayments).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &JobPayments{config: jpq.config}
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
	if query := jpq.withPayments; query != nil {
		if err := jpq.loadPayments(ctx, query, nodes, nil,
			func(n *JobPayments, e *JobDetail) { n.Edges.Payments = e }); err != nil {
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

func (jpq *JobPaymentsQuery) loadPayments(ctx context.Context, query *JobDetailQuery, nodes []*JobPayments, init func(*JobPayments), assign func(*JobPayments, *JobDetail)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*JobPayments)
	for i := range nodes {
		if nodes[i].payments_id == nil {
			continue
		}
		fk := *nodes[i].payments_id
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(jobdetail.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "payments_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (jpq *JobPaymentsQuery) sqlCount(ctx context.Context) (int, error) {
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

func (jpq *JobPaymentsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(jobpayments.Table, jobpayments.Columns, sqlgraph.NewFieldSpec(jobpayments.FieldID, field.TypeInt))
	_spec.From = jpq.sql
	if unique := jpq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if jpq.path != nil {
		_spec.Unique = true
	}
	if fields := jpq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, jobpayments.FieldID)
		for i := range fields {
			if fields[i] != jobpayments.FieldID {
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

func (jpq *JobPaymentsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(jpq.driver.Dialect())
	t1 := builder.Table(jobpayments.Table)
	columns := jpq.ctx.Fields
	if len(columns) == 0 {
		columns = jobpayments.Columns
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

// JobPaymentsGroupBy is the group-by builder for JobPayments entities.
type JobPaymentsGroupBy struct {
	selector
	build *JobPaymentsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (jpgb *JobPaymentsGroupBy) Aggregate(fns ...AggregateFunc) *JobPaymentsGroupBy {
	jpgb.fns = append(jpgb.fns, fns...)
	return jpgb
}

// Scan applies the selector query and scans the result into the given value.
func (jpgb *JobPaymentsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, jpgb.build.ctx, ent.OpQueryGroupBy)
	if err := jpgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*JobPaymentsQuery, *JobPaymentsGroupBy](ctx, jpgb.build, jpgb, jpgb.build.inters, v)
}

func (jpgb *JobPaymentsGroupBy) sqlScan(ctx context.Context, root *JobPaymentsQuery, v any) error {
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

// JobPaymentsSelect is the builder for selecting fields of JobPayments entities.
type JobPaymentsSelect struct {
	*JobPaymentsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (jps *JobPaymentsSelect) Aggregate(fns ...AggregateFunc) *JobPaymentsSelect {
	jps.fns = append(jps.fns, fns...)
	return jps
}

// Scan applies the selector query and scans the result into the given value.
func (jps *JobPaymentsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, jps.ctx, ent.OpQuerySelect)
	if err := jps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*JobPaymentsQuery, *JobPaymentsSelect](ctx, jps.JobPaymentsQuery, jps, jps.inters, v)
}

func (jps *JobPaymentsSelect) sqlScan(ctx context.Context, root *JobPaymentsQuery, v any) error {
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
