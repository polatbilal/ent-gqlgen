// Code generated by ent, DO NOT EDIT.

package joblayer

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the joblayer type in the database.
	Label = "job_layer"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldMetre holds the string denoting the metre field in the database.
	FieldMetre = "metre"
	// FieldMoldDate holds the string denoting the molddate field in the database.
	FieldMoldDate = "mold_date"
	// FieldConcreteDate holds the string denoting the concretedate field in the database.
	FieldConcreteDate = "concrete_date"
	// FieldSamples holds the string denoting the samples field in the database.
	FieldSamples = "samples"
	// FieldConcreteClass holds the string denoting the concreteclass field in the database.
	FieldConcreteClass = "concrete_class"
	// FieldWeekResult holds the string denoting the weekresult field in the database.
	FieldWeekResult = "week_result"
	// FieldMonthResult holds the string denoting the monthresult field in the database.
	FieldMonthResult = "month_result"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeLayer holds the string denoting the layer edge name in mutations.
	EdgeLayer = "layer"
	// Table holds the table name of the joblayer in the database.
	Table = "job_layers"
	// LayerTable is the table that holds the layer relation/edge.
	LayerTable = "job_layers"
	// LayerInverseTable is the table name for the JobDetail entity.
	// It exists in this package in order to avoid circular dependency with the "jobdetail" package.
	LayerInverseTable = "job_details"
	// LayerColumn is the table column denoting the layer relation/edge.
	LayerColumn = "job_id"
)

// Columns holds all SQL columns for joblayer fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldMetre,
	FieldMoldDate,
	FieldConcreteDate,
	FieldSamples,
	FieldConcreteClass,
	FieldWeekResult,
	FieldMonthResult,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "job_layers"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"job_id",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultName holds the default value on creation for the "Name" field.
	DefaultName string
	// DefaultMetre holds the default value on creation for the "Metre" field.
	DefaultMetre string
	// DefaultCreatedAt holds the default value on creation for the "CreatedAt" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "UpdatedAt" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "UpdatedAt" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the JobLayer queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the Name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByMetre orders the results by the Metre field.
func ByMetre(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMetre, opts...).ToFunc()
}

// ByMoldDate orders the results by the MoldDate field.
func ByMoldDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMoldDate, opts...).ToFunc()
}

// ByConcreteDate orders the results by the ConcreteDate field.
func ByConcreteDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldConcreteDate, opts...).ToFunc()
}

// BySamples orders the results by the Samples field.
func BySamples(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSamples, opts...).ToFunc()
}

// ByConcreteClass orders the results by the ConcreteClass field.
func ByConcreteClass(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldConcreteClass, opts...).ToFunc()
}

// ByWeekResult orders the results by the WeekResult field.
func ByWeekResult(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWeekResult, opts...).ToFunc()
}

// ByMonthResult orders the results by the MonthResult field.
func ByMonthResult(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMonthResult, opts...).ToFunc()
}

// ByCreatedAt orders the results by the CreatedAt field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the UpdatedAt field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByLayerField orders the results by layer field.
func ByLayerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLayerStep(), sql.OrderByField(field, opts...))
	}
}
func newLayerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LayerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, LayerTable, LayerColumn),
	)
}
