// Code generated by ent, DO NOT EDIT.

package jobpayments

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the jobpayments type in the database.
	Label = "job_payments"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldYibfNo holds the string denoting the yibfno field in the database.
	FieldYibfNo = "yibf_no"
	// FieldPaymentNo holds the string denoting the paymentno field in the database.
	FieldPaymentNo = "payment_no"
	// FieldPaymentDate holds the string denoting the paymentdate field in the database.
	FieldPaymentDate = "payment_date"
	// FieldPaymentType holds the string denoting the paymenttype field in the database.
	FieldPaymentType = "payment_type"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldTotalPayment holds the string denoting the totalpayment field in the database.
	FieldTotalPayment = "total_payment"
	// FieldLevelRequest holds the string denoting the levelrequest field in the database.
	FieldLevelRequest = "level_request"
	// FieldLevelApprove holds the string denoting the levelapprove field in the database.
	FieldLevelApprove = "level_approve"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgePayments holds the string denoting the payments edge name in mutations.
	EdgePayments = "payments"
	// Table holds the table name of the jobpayments in the database.
	Table = "job_payments"
	// PaymentsTable is the table that holds the payments relation/edge.
	PaymentsTable = "job_payments"
	// PaymentsInverseTable is the table name for the JobRelations entity.
	// It exists in this package in order to avoid circular dependency with the "jobrelations" package.
	PaymentsInverseTable = "job_relations"
	// PaymentsColumn is the table column denoting the payments relation/edge.
	PaymentsColumn = "job_id"
)

// Columns holds all SQL columns for jobpayments fields.
var Columns = []string{
	FieldID,
	FieldYibfNo,
	FieldPaymentNo,
	FieldPaymentDate,
	FieldPaymentType,
	FieldState,
	FieldTotalPayment,
	FieldLevelRequest,
	FieldLevelApprove,
	FieldAmount,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "job_payments"
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
	// DefaultYibfNo holds the default value on creation for the "yibfNo" field.
	DefaultYibfNo int
	// DefaultPaymentNo holds the default value on creation for the "PaymentNo" field.
	DefaultPaymentNo int
	// DefaultPaymentDate holds the default value on creation for the "PaymentDate" field.
	DefaultPaymentDate func() time.Time
	// DefaultPaymentType holds the default value on creation for the "PaymentType" field.
	DefaultPaymentType string
	// DefaultState holds the default value on creation for the "State" field.
	DefaultState string
	// DefaultTotalPayment holds the default value on creation for the "TotalPayment" field.
	DefaultTotalPayment float64
	// DefaultLevelRequest holds the default value on creation for the "LevelRequest" field.
	DefaultLevelRequest float64
	// DefaultLevelApprove holds the default value on creation for the "LevelApprove" field.
	DefaultLevelApprove float64
	// DefaultAmount holds the default value on creation for the "Amount" field.
	DefaultAmount float64
	// DefaultCreatedAt holds the default value on creation for the "CreatedAt" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "UpdatedAt" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "UpdatedAt" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the JobPayments queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByYibfNo orders the results by the yibfNo field.
func ByYibfNo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldYibfNo, opts...).ToFunc()
}

// ByPaymentNo orders the results by the PaymentNo field.
func ByPaymentNo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPaymentNo, opts...).ToFunc()
}

// ByPaymentDate orders the results by the PaymentDate field.
func ByPaymentDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPaymentDate, opts...).ToFunc()
}

// ByPaymentType orders the results by the PaymentType field.
func ByPaymentType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPaymentType, opts...).ToFunc()
}

// ByState orders the results by the State field.
func ByState(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldState, opts...).ToFunc()
}

// ByTotalPayment orders the results by the TotalPayment field.
func ByTotalPayment(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalPayment, opts...).ToFunc()
}

// ByLevelRequest orders the results by the LevelRequest field.
func ByLevelRequest(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLevelRequest, opts...).ToFunc()
}

// ByLevelApprove orders the results by the LevelApprove field.
func ByLevelApprove(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLevelApprove, opts...).ToFunc()
}

// ByAmount orders the results by the Amount field.
func ByAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAmount, opts...).ToFunc()
}

// ByCreatedAt orders the results by the CreatedAt field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the UpdatedAt field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByPaymentsField orders the results by payments field.
func ByPaymentsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPaymentsStep(), sql.OrderByField(field, opts...))
	}
}
func newPaymentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PaymentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, PaymentsTable, PaymentsColumn),
	)
}
