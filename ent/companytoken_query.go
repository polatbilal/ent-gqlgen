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
	"github.com/polatbilal/gqlgen-ent/ent/companydetail"
	"github.com/polatbilal/gqlgen-ent/ent/companytoken"
	"github.com/polatbilal/gqlgen-ent/ent/predicate"
)

// CompanyTokenQuery is the builder for querying CompanyToken entities.
type CompanyTokenQuery struct {
	config
	ctx         *QueryContext
	order       []companytoken.OrderOption
	inters      []Interceptor
	predicates  []predicate.CompanyToken
	withCompany *CompanyDetailQuery
	withFKs     bool
	modifiers   []func(*sql.Selector)
	loadTotal   []func(context.Context, []*CompanyToken) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CompanyTokenQuery builder.
func (ctq *CompanyTokenQuery) Where(ps ...predicate.CompanyToken) *CompanyTokenQuery {
	ctq.predicates = append(ctq.predicates, ps...)
	return ctq
}

// Limit the number of records to be returned by this query.
func (ctq *CompanyTokenQuery) Limit(limit int) *CompanyTokenQuery {
	ctq.ctx.Limit = &limit
	return ctq
}

// Offset to start from.
func (ctq *CompanyTokenQuery) Offset(offset int) *CompanyTokenQuery {
	ctq.ctx.Offset = &offset
	return ctq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ctq *CompanyTokenQuery) Unique(unique bool) *CompanyTokenQuery {
	ctq.ctx.Unique = &unique
	return ctq
}

// Order specifies how the records should be ordered.
func (ctq *CompanyTokenQuery) Order(o ...companytoken.OrderOption) *CompanyTokenQuery {
	ctq.order = append(ctq.order, o...)
	return ctq
}

// QueryCompany chains the current query on the "company" edge.
func (ctq *CompanyTokenQuery) QueryCompany() *CompanyDetailQuery {
	query := (&CompanyDetailClient{config: ctq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ctq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ctq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(companytoken.Table, companytoken.FieldID, selector),
			sqlgraph.To(companydetail.Table, companydetail.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, companytoken.CompanyTable, companytoken.CompanyColumn),
		)
		fromU = sqlgraph.SetNeighbors(ctq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CompanyToken entity from the query.
// Returns a *NotFoundError when no CompanyToken was found.
func (ctq *CompanyTokenQuery) First(ctx context.Context) (*CompanyToken, error) {
	nodes, err := ctq.Limit(1).All(setContextOp(ctx, ctq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{companytoken.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ctq *CompanyTokenQuery) FirstX(ctx context.Context) *CompanyToken {
	node, err := ctq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CompanyToken ID from the query.
// Returns a *NotFoundError when no CompanyToken ID was found.
func (ctq *CompanyTokenQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ctq.Limit(1).IDs(setContextOp(ctx, ctq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{companytoken.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ctq *CompanyTokenQuery) FirstIDX(ctx context.Context) int {
	id, err := ctq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CompanyToken entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CompanyToken entity is found.
// Returns a *NotFoundError when no CompanyToken entities are found.
func (ctq *CompanyTokenQuery) Only(ctx context.Context) (*CompanyToken, error) {
	nodes, err := ctq.Limit(2).All(setContextOp(ctx, ctq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{companytoken.Label}
	default:
		return nil, &NotSingularError{companytoken.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ctq *CompanyTokenQuery) OnlyX(ctx context.Context) *CompanyToken {
	node, err := ctq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CompanyToken ID in the query.
// Returns a *NotSingularError when more than one CompanyToken ID is found.
// Returns a *NotFoundError when no entities are found.
func (ctq *CompanyTokenQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ctq.Limit(2).IDs(setContextOp(ctx, ctq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{companytoken.Label}
	default:
		err = &NotSingularError{companytoken.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ctq *CompanyTokenQuery) OnlyIDX(ctx context.Context) int {
	id, err := ctq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CompanyTokens.
func (ctq *CompanyTokenQuery) All(ctx context.Context) ([]*CompanyToken, error) {
	ctx = setContextOp(ctx, ctq.ctx, ent.OpQueryAll)
	if err := ctq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*CompanyToken, *CompanyTokenQuery]()
	return withInterceptors[[]*CompanyToken](ctx, ctq, qr, ctq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ctq *CompanyTokenQuery) AllX(ctx context.Context) []*CompanyToken {
	nodes, err := ctq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CompanyToken IDs.
func (ctq *CompanyTokenQuery) IDs(ctx context.Context) (ids []int, err error) {
	if ctq.ctx.Unique == nil && ctq.path != nil {
		ctq.Unique(true)
	}
	ctx = setContextOp(ctx, ctq.ctx, ent.OpQueryIDs)
	if err = ctq.Select(companytoken.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ctq *CompanyTokenQuery) IDsX(ctx context.Context) []int {
	ids, err := ctq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ctq *CompanyTokenQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ctq.ctx, ent.OpQueryCount)
	if err := ctq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ctq, querierCount[*CompanyTokenQuery](), ctq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ctq *CompanyTokenQuery) CountX(ctx context.Context) int {
	count, err := ctq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ctq *CompanyTokenQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ctq.ctx, ent.OpQueryExist)
	switch _, err := ctq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ctq *CompanyTokenQuery) ExistX(ctx context.Context) bool {
	exist, err := ctq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CompanyTokenQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ctq *CompanyTokenQuery) Clone() *CompanyTokenQuery {
	if ctq == nil {
		return nil
	}
	return &CompanyTokenQuery{
		config:      ctq.config,
		ctx:         ctq.ctx.Clone(),
		order:       append([]companytoken.OrderOption{}, ctq.order...),
		inters:      append([]Interceptor{}, ctq.inters...),
		predicates:  append([]predicate.CompanyToken{}, ctq.predicates...),
		withCompany: ctq.withCompany.Clone(),
		// clone intermediate query.
		sql:  ctq.sql.Clone(),
		path: ctq.path,
	}
}

// WithCompany tells the query-builder to eager-load the nodes that are connected to
// the "company" edge. The optional arguments are used to configure the query builder of the edge.
func (ctq *CompanyTokenQuery) WithCompany(opts ...func(*CompanyDetailQuery)) *CompanyTokenQuery {
	query := (&CompanyDetailClient{config: ctq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ctq.withCompany = query
	return ctq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		YDKUsername string `json:"YDKUsername,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CompanyToken.Query().
//		GroupBy(companytoken.FieldYDKUsername).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ctq *CompanyTokenQuery) GroupBy(field string, fields ...string) *CompanyTokenGroupBy {
	ctq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CompanyTokenGroupBy{build: ctq}
	grbuild.flds = &ctq.ctx.Fields
	grbuild.label = companytoken.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		YDKUsername string `json:"YDKUsername,omitempty"`
//	}
//
//	client.CompanyToken.Query().
//		Select(companytoken.FieldYDKUsername).
//		Scan(ctx, &v)
func (ctq *CompanyTokenQuery) Select(fields ...string) *CompanyTokenSelect {
	ctq.ctx.Fields = append(ctq.ctx.Fields, fields...)
	sbuild := &CompanyTokenSelect{CompanyTokenQuery: ctq}
	sbuild.label = companytoken.Label
	sbuild.flds, sbuild.scan = &ctq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CompanyTokenSelect configured with the given aggregations.
func (ctq *CompanyTokenQuery) Aggregate(fns ...AggregateFunc) *CompanyTokenSelect {
	return ctq.Select().Aggregate(fns...)
}

func (ctq *CompanyTokenQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ctq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ctq); err != nil {
				return err
			}
		}
	}
	for _, f := range ctq.ctx.Fields {
		if !companytoken.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ctq.path != nil {
		prev, err := ctq.path(ctx)
		if err != nil {
			return err
		}
		ctq.sql = prev
	}
	return nil
}

func (ctq *CompanyTokenQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CompanyToken, error) {
	var (
		nodes       = []*CompanyToken{}
		withFKs     = ctq.withFKs
		_spec       = ctq.querySpec()
		loadedTypes = [1]bool{
			ctq.withCompany != nil,
		}
	)
	if ctq.withCompany != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, companytoken.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CompanyToken).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CompanyToken{config: ctq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(ctq.modifiers) > 0 {
		_spec.Modifiers = ctq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ctq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ctq.withCompany; query != nil {
		if err := ctq.loadCompany(ctx, query, nodes, nil,
			func(n *CompanyToken, e *CompanyDetail) { n.Edges.Company = e }); err != nil {
			return nil, err
		}
	}
	for i := range ctq.loadTotal {
		if err := ctq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ctq *CompanyTokenQuery) loadCompany(ctx context.Context, query *CompanyDetailQuery, nodes []*CompanyToken, init func(*CompanyToken), assign func(*CompanyToken, *CompanyDetail)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*CompanyToken)
	for i := range nodes {
		if nodes[i].company_id == nil {
			continue
		}
		fk := *nodes[i].company_id
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(companydetail.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "company_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ctq *CompanyTokenQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ctq.querySpec()
	if len(ctq.modifiers) > 0 {
		_spec.Modifiers = ctq.modifiers
	}
	_spec.Node.Columns = ctq.ctx.Fields
	if len(ctq.ctx.Fields) > 0 {
		_spec.Unique = ctq.ctx.Unique != nil && *ctq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ctq.driver, _spec)
}

func (ctq *CompanyTokenQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(companytoken.Table, companytoken.Columns, sqlgraph.NewFieldSpec(companytoken.FieldID, field.TypeInt))
	_spec.From = ctq.sql
	if unique := ctq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ctq.path != nil {
		_spec.Unique = true
	}
	if fields := ctq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, companytoken.FieldID)
		for i := range fields {
			if fields[i] != companytoken.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ctq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ctq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ctq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ctq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ctq *CompanyTokenQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ctq.driver.Dialect())
	t1 := builder.Table(companytoken.Table)
	columns := ctq.ctx.Fields
	if len(columns) == 0 {
		columns = companytoken.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ctq.sql != nil {
		selector = ctq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ctq.ctx.Unique != nil && *ctq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ctq.predicates {
		p(selector)
	}
	for _, p := range ctq.order {
		p(selector)
	}
	if offset := ctq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ctq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CompanyTokenGroupBy is the group-by builder for CompanyToken entities.
type CompanyTokenGroupBy struct {
	selector
	build *CompanyTokenQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ctgb *CompanyTokenGroupBy) Aggregate(fns ...AggregateFunc) *CompanyTokenGroupBy {
	ctgb.fns = append(ctgb.fns, fns...)
	return ctgb
}

// Scan applies the selector query and scans the result into the given value.
func (ctgb *CompanyTokenGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ctgb.build.ctx, ent.OpQueryGroupBy)
	if err := ctgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CompanyTokenQuery, *CompanyTokenGroupBy](ctx, ctgb.build, ctgb, ctgb.build.inters, v)
}

func (ctgb *CompanyTokenGroupBy) sqlScan(ctx context.Context, root *CompanyTokenQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ctgb.fns))
	for _, fn := range ctgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ctgb.flds)+len(ctgb.fns))
		for _, f := range *ctgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ctgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ctgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CompanyTokenSelect is the builder for selecting fields of CompanyToken entities.
type CompanyTokenSelect struct {
	*CompanyTokenQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cts *CompanyTokenSelect) Aggregate(fns ...AggregateFunc) *CompanyTokenSelect {
	cts.fns = append(cts.fns, fns...)
	return cts
}

// Scan applies the selector query and scans the result into the given value.
func (cts *CompanyTokenSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cts.ctx, ent.OpQuerySelect)
	if err := cts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CompanyTokenQuery, *CompanyTokenSelect](ctx, cts.CompanyTokenQuery, cts, cts.inters, v)
}

func (cts *CompanyTokenSelect) sqlScan(ctx context.Context, root *CompanyTokenQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cts.fns))
	for _, fn := range cts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
