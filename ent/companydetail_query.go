// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"gqlgen-ent/ent/companydetail"
	"gqlgen-ent/ent/companyengineer"
	"gqlgen-ent/ent/predicate"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CompanyDetailQuery is the builder for querying CompanyDetail entities.
type CompanyDetailQuery struct {
	config
	ctx              *QueryContext
	order            []companydetail.OrderOption
	inters           []Interceptor
	predicates       []predicate.CompanyDetail
	withCompanyOwner *CompanyEngineerQuery
	withFKs          bool
	modifiers        []func(*sql.Selector)
	loadTotal        []func(context.Context, []*CompanyDetail) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CompanyDetailQuery builder.
func (cdq *CompanyDetailQuery) Where(ps ...predicate.CompanyDetail) *CompanyDetailQuery {
	cdq.predicates = append(cdq.predicates, ps...)
	return cdq
}

// Limit the number of records to be returned by this query.
func (cdq *CompanyDetailQuery) Limit(limit int) *CompanyDetailQuery {
	cdq.ctx.Limit = &limit
	return cdq
}

// Offset to start from.
func (cdq *CompanyDetailQuery) Offset(offset int) *CompanyDetailQuery {
	cdq.ctx.Offset = &offset
	return cdq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cdq *CompanyDetailQuery) Unique(unique bool) *CompanyDetailQuery {
	cdq.ctx.Unique = &unique
	return cdq
}

// Order specifies how the records should be ordered.
func (cdq *CompanyDetailQuery) Order(o ...companydetail.OrderOption) *CompanyDetailQuery {
	cdq.order = append(cdq.order, o...)
	return cdq
}

// QueryCompanyOwner chains the current query on the "companyOwner" edge.
func (cdq *CompanyDetailQuery) QueryCompanyOwner() *CompanyEngineerQuery {
	query := (&CompanyEngineerClient{config: cdq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(companydetail.Table, companydetail.FieldID, selector),
			sqlgraph.To(companyengineer.Table, companyengineer.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, companydetail.CompanyOwnerTable, companydetail.CompanyOwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(cdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CompanyDetail entity from the query.
// Returns a *NotFoundError when no CompanyDetail was found.
func (cdq *CompanyDetailQuery) First(ctx context.Context) (*CompanyDetail, error) {
	nodes, err := cdq.Limit(1).All(setContextOp(ctx, cdq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{companydetail.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cdq *CompanyDetailQuery) FirstX(ctx context.Context) *CompanyDetail {
	node, err := cdq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CompanyDetail ID from the query.
// Returns a *NotFoundError when no CompanyDetail ID was found.
func (cdq *CompanyDetailQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cdq.Limit(1).IDs(setContextOp(ctx, cdq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{companydetail.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cdq *CompanyDetailQuery) FirstIDX(ctx context.Context) int {
	id, err := cdq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CompanyDetail entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CompanyDetail entity is found.
// Returns a *NotFoundError when no CompanyDetail entities are found.
func (cdq *CompanyDetailQuery) Only(ctx context.Context) (*CompanyDetail, error) {
	nodes, err := cdq.Limit(2).All(setContextOp(ctx, cdq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{companydetail.Label}
	default:
		return nil, &NotSingularError{companydetail.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cdq *CompanyDetailQuery) OnlyX(ctx context.Context) *CompanyDetail {
	node, err := cdq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CompanyDetail ID in the query.
// Returns a *NotSingularError when more than one CompanyDetail ID is found.
// Returns a *NotFoundError when no entities are found.
func (cdq *CompanyDetailQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cdq.Limit(2).IDs(setContextOp(ctx, cdq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{companydetail.Label}
	default:
		err = &NotSingularError{companydetail.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cdq *CompanyDetailQuery) OnlyIDX(ctx context.Context) int {
	id, err := cdq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CompanyDetails.
func (cdq *CompanyDetailQuery) All(ctx context.Context) ([]*CompanyDetail, error) {
	ctx = setContextOp(ctx, cdq.ctx, ent.OpQueryAll)
	if err := cdq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*CompanyDetail, *CompanyDetailQuery]()
	return withInterceptors[[]*CompanyDetail](ctx, cdq, qr, cdq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cdq *CompanyDetailQuery) AllX(ctx context.Context) []*CompanyDetail {
	nodes, err := cdq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CompanyDetail IDs.
func (cdq *CompanyDetailQuery) IDs(ctx context.Context) (ids []int, err error) {
	if cdq.ctx.Unique == nil && cdq.path != nil {
		cdq.Unique(true)
	}
	ctx = setContextOp(ctx, cdq.ctx, ent.OpQueryIDs)
	if err = cdq.Select(companydetail.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cdq *CompanyDetailQuery) IDsX(ctx context.Context) []int {
	ids, err := cdq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cdq *CompanyDetailQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cdq.ctx, ent.OpQueryCount)
	if err := cdq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cdq, querierCount[*CompanyDetailQuery](), cdq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cdq *CompanyDetailQuery) CountX(ctx context.Context) int {
	count, err := cdq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cdq *CompanyDetailQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cdq.ctx, ent.OpQueryExist)
	switch _, err := cdq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cdq *CompanyDetailQuery) ExistX(ctx context.Context) bool {
	exist, err := cdq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CompanyDetailQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cdq *CompanyDetailQuery) Clone() *CompanyDetailQuery {
	if cdq == nil {
		return nil
	}
	return &CompanyDetailQuery{
		config:           cdq.config,
		ctx:              cdq.ctx.Clone(),
		order:            append([]companydetail.OrderOption{}, cdq.order...),
		inters:           append([]Interceptor{}, cdq.inters...),
		predicates:       append([]predicate.CompanyDetail{}, cdq.predicates...),
		withCompanyOwner: cdq.withCompanyOwner.Clone(),
		// clone intermediate query.
		sql:  cdq.sql.Clone(),
		path: cdq.path,
	}
}

// WithCompanyOwner tells the query-builder to eager-load the nodes that are connected to
// the "companyOwner" edge. The optional arguments are used to configure the query builder of the edge.
func (cdq *CompanyDetailQuery) WithCompanyOwner(opts ...func(*CompanyEngineerQuery)) *CompanyDetailQuery {
	query := (&CompanyEngineerClient{config: cdq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cdq.withCompanyOwner = query
	return cdq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CompanyCode int `json:"CompanyCode,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CompanyDetail.Query().
//		GroupBy(companydetail.FieldCompanyCode).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cdq *CompanyDetailQuery) GroupBy(field string, fields ...string) *CompanyDetailGroupBy {
	cdq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CompanyDetailGroupBy{build: cdq}
	grbuild.flds = &cdq.ctx.Fields
	grbuild.label = companydetail.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CompanyCode int `json:"CompanyCode,omitempty"`
//	}
//
//	client.CompanyDetail.Query().
//		Select(companydetail.FieldCompanyCode).
//		Scan(ctx, &v)
func (cdq *CompanyDetailQuery) Select(fields ...string) *CompanyDetailSelect {
	cdq.ctx.Fields = append(cdq.ctx.Fields, fields...)
	sbuild := &CompanyDetailSelect{CompanyDetailQuery: cdq}
	sbuild.label = companydetail.Label
	sbuild.flds, sbuild.scan = &cdq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CompanyDetailSelect configured with the given aggregations.
func (cdq *CompanyDetailQuery) Aggregate(fns ...AggregateFunc) *CompanyDetailSelect {
	return cdq.Select().Aggregate(fns...)
}

func (cdq *CompanyDetailQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cdq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cdq); err != nil {
				return err
			}
		}
	}
	for _, f := range cdq.ctx.Fields {
		if !companydetail.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cdq.path != nil {
		prev, err := cdq.path(ctx)
		if err != nil {
			return err
		}
		cdq.sql = prev
	}
	return nil
}

func (cdq *CompanyDetailQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CompanyDetail, error) {
	var (
		nodes       = []*CompanyDetail{}
		withFKs     = cdq.withFKs
		_spec       = cdq.querySpec()
		loadedTypes = [1]bool{
			cdq.withCompanyOwner != nil,
		}
	)
	if cdq.withCompanyOwner != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, companydetail.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CompanyDetail).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CompanyDetail{config: cdq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(cdq.modifiers) > 0 {
		_spec.Modifiers = cdq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cdq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cdq.withCompanyOwner; query != nil {
		if err := cdq.loadCompanyOwner(ctx, query, nodes, nil,
			func(n *CompanyDetail, e *CompanyEngineer) { n.Edges.CompanyOwner = e }); err != nil {
			return nil, err
		}
	}
	for i := range cdq.loadTotal {
		if err := cdq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cdq *CompanyDetailQuery) loadCompanyOwner(ctx context.Context, query *CompanyEngineerQuery, nodes []*CompanyDetail, init func(*CompanyDetail), assign func(*CompanyDetail, *CompanyEngineer)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*CompanyDetail)
	for i := range nodes {
		if nodes[i].owner_id == nil {
			continue
		}
		fk := *nodes[i].owner_id
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(companyengineer.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "owner_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (cdq *CompanyDetailQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cdq.querySpec()
	if len(cdq.modifiers) > 0 {
		_spec.Modifiers = cdq.modifiers
	}
	_spec.Node.Columns = cdq.ctx.Fields
	if len(cdq.ctx.Fields) > 0 {
		_spec.Unique = cdq.ctx.Unique != nil && *cdq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cdq.driver, _spec)
}

func (cdq *CompanyDetailQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(companydetail.Table, companydetail.Columns, sqlgraph.NewFieldSpec(companydetail.FieldID, field.TypeInt))
	_spec.From = cdq.sql
	if unique := cdq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cdq.path != nil {
		_spec.Unique = true
	}
	if fields := cdq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, companydetail.FieldID)
		for i := range fields {
			if fields[i] != companydetail.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cdq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cdq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cdq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cdq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cdq *CompanyDetailQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cdq.driver.Dialect())
	t1 := builder.Table(companydetail.Table)
	columns := cdq.ctx.Fields
	if len(columns) == 0 {
		columns = companydetail.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cdq.sql != nil {
		selector = cdq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cdq.ctx.Unique != nil && *cdq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range cdq.predicates {
		p(selector)
	}
	for _, p := range cdq.order {
		p(selector)
	}
	if offset := cdq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cdq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CompanyDetailGroupBy is the group-by builder for CompanyDetail entities.
type CompanyDetailGroupBy struct {
	selector
	build *CompanyDetailQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cdgb *CompanyDetailGroupBy) Aggregate(fns ...AggregateFunc) *CompanyDetailGroupBy {
	cdgb.fns = append(cdgb.fns, fns...)
	return cdgb
}

// Scan applies the selector query and scans the result into the given value.
func (cdgb *CompanyDetailGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cdgb.build.ctx, ent.OpQueryGroupBy)
	if err := cdgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CompanyDetailQuery, *CompanyDetailGroupBy](ctx, cdgb.build, cdgb, cdgb.build.inters, v)
}

func (cdgb *CompanyDetailGroupBy) sqlScan(ctx context.Context, root *CompanyDetailQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cdgb.fns))
	for _, fn := range cdgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cdgb.flds)+len(cdgb.fns))
		for _, f := range *cdgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cdgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cdgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CompanyDetailSelect is the builder for selecting fields of CompanyDetail entities.
type CompanyDetailSelect struct {
	*CompanyDetailQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cds *CompanyDetailSelect) Aggregate(fns ...AggregateFunc) *CompanyDetailSelect {
	cds.fns = append(cds.fns, fns...)
	return cds
}

// Scan applies the selector query and scans the result into the given value.
func (cds *CompanyDetailSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cds.ctx, ent.OpQuerySelect)
	if err := cds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CompanyDetailQuery, *CompanyDetailSelect](ctx, cds.CompanyDetailQuery, cds, cds.inters, v)
}

func (cds *CompanyDetailSelect) sqlScan(ctx context.Context, root *CompanyDetailQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cds.fns))
	for _, fn := range cds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
