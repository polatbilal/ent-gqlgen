// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"gqlgen-ent/ent/companydetail"
	"gqlgen-ent/ent/companyengineer"
	"gqlgen-ent/ent/jobauthor"
	"gqlgen-ent/ent/jobcontractor"
	"gqlgen-ent/ent/jobdetail"
	"gqlgen-ent/ent/joblayer"
	"gqlgen-ent/ent/jobowner"
	"gqlgen-ent/ent/jobprogress"
	"gqlgen-ent/ent/user"
	"sync"
	"sync/atomic"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/99designs/gqlgen/graphql"
	"github.com/hashicorp/go-multierror"
	"golang.org/x/sync/semaphore"
)

// Noder wraps the basic Node method.
type Noder interface {
	IsNode()
}

var companydetailImplementors = []string{"CompanyDetail", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*CompanyDetail) IsNode() {}

var companyengineerImplementors = []string{"CompanyEngineer", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*CompanyEngineer) IsNode() {}

var jobauthorImplementors = []string{"JobAuthor", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*JobAuthor) IsNode() {}

var jobcontractorImplementors = []string{"JobContractor", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*JobContractor) IsNode() {}

var jobdetailImplementors = []string{"JobDetail", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*JobDetail) IsNode() {}

var joblayerImplementors = []string{"JobLayer", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*JobLayer) IsNode() {}

var jobownerImplementors = []string{"JobOwner", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*JobOwner) IsNode() {}

var jobprogressImplementors = []string{"JobProgress", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*JobProgress) IsNode() {}

var userImplementors = []string{"User", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*User) IsNode() {}

var errNodeInvalidID = &NotFoundError{"node"}

// NodeOption allows configuring the Noder execution using functional options.
type NodeOption func(*nodeOptions)

// WithNodeType sets the node Type resolver function (i.e. the table to query).
// If was not provided, the table will be derived from the universal-id
// configuration as described in: https://entgo.io/docs/migrate/#universal-ids.
func WithNodeType(f func(context.Context, int) (string, error)) NodeOption {
	return func(o *nodeOptions) {
		o.nodeType = f
	}
}

// WithFixedNodeType sets the Type of the node to a fixed value.
func WithFixedNodeType(t string) NodeOption {
	return WithNodeType(func(context.Context, int) (string, error) {
		return t, nil
	})
}

type nodeOptions struct {
	nodeType func(context.Context, int) (string, error)
}

func (c *Client) newNodeOpts(opts []NodeOption) *nodeOptions {
	nopts := &nodeOptions{}
	for _, opt := range opts {
		opt(nopts)
	}
	if nopts.nodeType == nil {
		nopts.nodeType = func(ctx context.Context, id int) (string, error) {
			return c.tables.nodeType(ctx, c.driver, id)
		}
	}
	return nopts
}

// Noder returns a Node by its id. If the NodeType was not provided, it will
// be derived from the id value according to the universal-id configuration.
//
//	c.Noder(ctx, id)
//	c.Noder(ctx, id, ent.WithNodeType(typeResolver))
func (c *Client) Noder(ctx context.Context, id int, opts ...NodeOption) (_ Noder, err error) {
	defer func() {
		if IsNotFound(err) {
			err = multierror.Append(err, entgql.ErrNodeNotFound(id))
		}
	}()
	table, err := c.newNodeOpts(opts).nodeType(ctx, id)
	if err != nil {
		return nil, err
	}
	return c.noder(ctx, table, id)
}

func (c *Client) noder(ctx context.Context, table string, id int) (Noder, error) {
	switch table {
	case companydetail.Table:
		query := c.CompanyDetail.Query().
			Where(companydetail.ID(id))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, companydetailImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	case companyengineer.Table:
		query := c.CompanyEngineer.Query().
			Where(companyengineer.ID(id))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, companyengineerImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	case jobauthor.Table:
		query := c.JobAuthor.Query().
			Where(jobauthor.ID(id))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, jobauthorImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	case jobcontractor.Table:
		query := c.JobContractor.Query().
			Where(jobcontractor.ID(id))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, jobcontractorImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	case jobdetail.Table:
		query := c.JobDetail.Query().
			Where(jobdetail.ID(id))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, jobdetailImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	case joblayer.Table:
		query := c.JobLayer.Query().
			Where(joblayer.ID(id))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, joblayerImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	case jobowner.Table:
		query := c.JobOwner.Query().
			Where(jobowner.ID(id))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, jobownerImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	case jobprogress.Table:
		query := c.JobProgress.Query().
			Where(jobprogress.ID(id))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, jobprogressImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	case user.Table:
		query := c.User.Query().
			Where(user.ID(id))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, userImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	default:
		return nil, fmt.Errorf("cannot resolve noder from table %q: %w", table, errNodeInvalidID)
	}
}

func (c *Client) Noders(ctx context.Context, ids []int, opts ...NodeOption) ([]Noder, error) {
	switch len(ids) {
	case 1:
		noder, err := c.Noder(ctx, ids[0], opts...)
		if err != nil {
			return nil, err
		}
		return []Noder{noder}, nil
	case 0:
		return []Noder{}, nil
	}

	noders := make([]Noder, len(ids))
	errors := make([]error, len(ids))
	tables := make(map[string][]int)
	id2idx := make(map[int][]int, len(ids))
	nopts := c.newNodeOpts(opts)
	for i, id := range ids {
		table, err := nopts.nodeType(ctx, id)
		if err != nil {
			errors[i] = err
			continue
		}
		tables[table] = append(tables[table], id)
		id2idx[id] = append(id2idx[id], i)
	}

	for table, ids := range tables {
		nodes, err := c.noders(ctx, table, ids)
		if err != nil {
			for _, id := range ids {
				for _, idx := range id2idx[id] {
					errors[idx] = err
				}
			}
		} else {
			for i, id := range ids {
				for _, idx := range id2idx[id] {
					noders[idx] = nodes[i]
				}
			}
		}
	}

	for i, id := range ids {
		if errors[i] == nil {
			if noders[i] != nil {
				continue
			}
			errors[i] = entgql.ErrNodeNotFound(id)
		} else if IsNotFound(errors[i]) {
			errors[i] = multierror.Append(errors[i], entgql.ErrNodeNotFound(id))
		}
		ctx := graphql.WithPathContext(ctx,
			graphql.NewPathWithIndex(i),
		)
		graphql.AddError(ctx, errors[i])
	}
	return noders, nil
}

func (c *Client) noders(ctx context.Context, table string, ids []int) ([]Noder, error) {
	noders := make([]Noder, len(ids))
	idmap := make(map[int][]*Noder, len(ids))
	for i, id := range ids {
		idmap[id] = append(idmap[id], &noders[i])
	}
	switch table {
	case companydetail.Table:
		query := c.CompanyDetail.Query().
			Where(companydetail.IDIn(ids...))
		query, err := query.CollectFields(ctx, companydetailImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case companyengineer.Table:
		query := c.CompanyEngineer.Query().
			Where(companyengineer.IDIn(ids...))
		query, err := query.CollectFields(ctx, companyengineerImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case jobauthor.Table:
		query := c.JobAuthor.Query().
			Where(jobauthor.IDIn(ids...))
		query, err := query.CollectFields(ctx, jobauthorImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case jobcontractor.Table:
		query := c.JobContractor.Query().
			Where(jobcontractor.IDIn(ids...))
		query, err := query.CollectFields(ctx, jobcontractorImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case jobdetail.Table:
		query := c.JobDetail.Query().
			Where(jobdetail.IDIn(ids...))
		query, err := query.CollectFields(ctx, jobdetailImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case joblayer.Table:
		query := c.JobLayer.Query().
			Where(joblayer.IDIn(ids...))
		query, err := query.CollectFields(ctx, joblayerImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case jobowner.Table:
		query := c.JobOwner.Query().
			Where(jobowner.IDIn(ids...))
		query, err := query.CollectFields(ctx, jobownerImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case jobprogress.Table:
		query := c.JobProgress.Query().
			Where(jobprogress.IDIn(ids...))
		query, err := query.CollectFields(ctx, jobprogressImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case user.Table:
		query := c.User.Query().
			Where(user.IDIn(ids...))
		query, err := query.CollectFields(ctx, userImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	default:
		return nil, fmt.Errorf("cannot resolve noders from table %q: %w", table, errNodeInvalidID)
	}
	return noders, nil
}

type tables struct {
	once  sync.Once
	sem   *semaphore.Weighted
	value atomic.Value
}

func (t *tables) nodeType(ctx context.Context, drv dialect.Driver, id int) (string, error) {
	tables, err := t.Load(ctx, drv)
	if err != nil {
		return "", err
	}
	idx := int(id / (1<<32 - 1))
	if idx < 0 || idx >= len(tables) {
		return "", fmt.Errorf("cannot resolve table from id %v: %w", id, errNodeInvalidID)
	}
	return tables[idx], nil
}

func (t *tables) Load(ctx context.Context, drv dialect.Driver) ([]string, error) {
	if tables := t.value.Load(); tables != nil {
		return tables.([]string), nil
	}
	t.once.Do(func() { t.sem = semaphore.NewWeighted(1) })
	if err := t.sem.Acquire(ctx, 1); err != nil {
		return nil, err
	}
	defer t.sem.Release(1)
	if tables := t.value.Load(); tables != nil {
		return tables.([]string), nil
	}
	tables, err := t.load(ctx, drv)
	if err == nil {
		t.value.Store(tables)
	}
	return tables, err
}

func (*tables) load(ctx context.Context, drv dialect.Driver) ([]string, error) {
	rows := &sql.Rows{}
	query, args := sql.Dialect(drv.Dialect()).
		Select("type").
		From(sql.Table(schema.TypeTable)).
		OrderBy(sql.Asc("id")).
		Query()
	if err := drv.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var tables []string
	return tables, sql.ScanSlice(rows, &tables)
}
