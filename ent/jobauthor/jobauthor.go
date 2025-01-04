// Code generated by ent, DO NOT EDIT.

package jobauthor

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the jobauthor type in the database.
	Label = "job_author"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStatic holds the string denoting the static field in the database.
	FieldStatic = "static"
	// FieldMechanic holds the string denoting the mechanic field in the database.
	FieldMechanic = "mechanic"
	// FieldElectric holds the string denoting the electric field in the database.
	FieldElectric = "electric"
	// FieldArchitect holds the string denoting the architect field in the database.
	FieldArchitect = "architect"
	// FieldGeotechnicalEngineer holds the string denoting the geotechnicalengineer field in the database.
	FieldGeotechnicalEngineer = "geotechnical_engineer"
	// FieldGeotechnicalGeologist holds the string denoting the geotechnicalgeologist field in the database.
	FieldGeotechnicalGeologist = "geotechnical_geologist"
	// FieldGeotechnicalGeophysicist holds the string denoting the geotechnicalgeophysicist field in the database.
	FieldGeotechnicalGeophysicist = "geotechnical_geophysicist"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeAuthors holds the string denoting the authors edge name in mutations.
	EdgeAuthors = "authors"
	// Table holds the table name of the jobauthor in the database.
	Table = "job_authors"
	// AuthorsTable is the table that holds the authors relation/edge.
	AuthorsTable = "job_details"
	// AuthorsInverseTable is the table name for the JobDetail entity.
	// It exists in this package in order to avoid circular dependency with the "jobdetail" package.
	AuthorsInverseTable = "job_details"
	// AuthorsColumn is the table column denoting the authors relation/edge.
	AuthorsColumn = "author_id"
)

// Columns holds all SQL columns for jobauthor fields.
var Columns = []string{
	FieldID,
	FieldStatic,
	FieldMechanic,
	FieldElectric,
	FieldArchitect,
	FieldGeotechnicalEngineer,
	FieldGeotechnicalGeologist,
	FieldGeotechnicalGeophysicist,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "CreatedAt" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "UpdatedAt" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "UpdatedAt" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the JobAuthor queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByStatic orders the results by the Static field.
func ByStatic(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatic, opts...).ToFunc()
}

// ByMechanic orders the results by the Mechanic field.
func ByMechanic(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMechanic, opts...).ToFunc()
}

// ByElectric orders the results by the Electric field.
func ByElectric(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldElectric, opts...).ToFunc()
}

// ByArchitect orders the results by the Architect field.
func ByArchitect(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldArchitect, opts...).ToFunc()
}

// ByGeotechnicalEngineer orders the results by the GeotechnicalEngineer field.
func ByGeotechnicalEngineer(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGeotechnicalEngineer, opts...).ToFunc()
}

// ByGeotechnicalGeologist orders the results by the GeotechnicalGeologist field.
func ByGeotechnicalGeologist(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGeotechnicalGeologist, opts...).ToFunc()
}

// ByGeotechnicalGeophysicist orders the results by the GeotechnicalGeophysicist field.
func ByGeotechnicalGeophysicist(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGeotechnicalGeophysicist, opts...).ToFunc()
}

// ByCreatedAt orders the results by the CreatedAt field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the UpdatedAt field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByAuthorsCount orders the results by authors count.
func ByAuthorsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAuthorsStep(), opts...)
	}
}

// ByAuthors orders the results by authors terms.
func ByAuthors(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAuthorsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAuthorsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AuthorsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AuthorsTable, AuthorsColumn),
	)
}
