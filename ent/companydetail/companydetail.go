// Code generated by ent, DO NOT EDIT.

package companydetail

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the companydetail type in the database.
	Label = "company_detail"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCompanyCode holds the string denoting the companycode field in the database.
	FieldCompanyCode = "company_code"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldCity holds the string denoting the city field in the database.
	FieldCity = "city"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldFax holds the string denoting the fax field in the database.
	FieldFax = "fax"
	// FieldMobile holds the string denoting the mobile field in the database.
	FieldMobile = "mobile"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldWebsite holds the string denoting the website field in the database.
	FieldWebsite = "website"
	// FieldTaxAdmin holds the string denoting the taxadmin field in the database.
	FieldTaxAdmin = "tax_admin"
	// FieldTaxNo holds the string denoting the taxno field in the database.
	FieldTaxNo = "tax_no"
	// FieldCommerce holds the string denoting the commerce field in the database.
	FieldCommerce = "commerce"
	// FieldCommerceReg holds the string denoting the commercereg field in the database.
	FieldCommerceReg = "commerce_reg"
	// FieldVisaDate holds the string denoting the visadate field in the database.
	FieldVisaDate = "visa_date"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeCompanyOwner holds the string denoting the companyowner edge name in mutations.
	EdgeCompanyOwner = "companyOwner"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// Table holds the table name of the companydetail in the database.
	Table = "company_details"
	// CompanyOwnerTable is the table that holds the companyOwner relation/edge.
	CompanyOwnerTable = "company_details"
	// CompanyOwnerInverseTable is the table name for the CompanyEngineer entity.
	// It exists in this package in order to avoid circular dependency with the "companyengineer" package.
	CompanyOwnerInverseTable = "company_engineers"
	// CompanyOwnerColumn is the table column denoting the companyOwner relation/edge.
	CompanyOwnerColumn = "owner_id"
	// UsersTable is the table that holds the users relation/edge.
	UsersTable = "company_users"
	// UsersInverseTable is the table name for the CompanyUser entity.
	// It exists in this package in order to avoid circular dependency with the "companyuser" package.
	UsersInverseTable = "company_users"
	// UsersColumn is the table column denoting the users relation/edge.
	UsersColumn = "company_user_company"
)

// Columns holds all SQL columns for companydetail fields.
var Columns = []string{
	FieldID,
	FieldCompanyCode,
	FieldName,
	FieldAddress,
	FieldCity,
	FieldState,
	FieldPhone,
	FieldFax,
	FieldMobile,
	FieldEmail,
	FieldWebsite,
	FieldTaxAdmin,
	FieldTaxNo,
	FieldCommerce,
	FieldCommerceReg,
	FieldVisaDate,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "company_details"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"owner_id",
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
	// DefaultTaxNo holds the default value on creation for the "TaxNo" field.
	DefaultTaxNo int
	// DefaultCreatedAt holds the default value on creation for the "CreatedAt" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "UpdatedAt" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "UpdatedAt" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the CompanyDetail queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCompanyCode orders the results by the CompanyCode field.
func ByCompanyCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCompanyCode, opts...).ToFunc()
}

// ByName orders the results by the Name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByAddress orders the results by the Address field.
func ByAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddress, opts...).ToFunc()
}

// ByCity orders the results by the City field.
func ByCity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCity, opts...).ToFunc()
}

// ByState orders the results by the State field.
func ByState(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldState, opts...).ToFunc()
}

// ByPhone orders the results by the Phone field.
func ByPhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhone, opts...).ToFunc()
}

// ByFax orders the results by the Fax field.
func ByFax(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFax, opts...).ToFunc()
}

// ByMobile orders the results by the Mobile field.
func ByMobile(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMobile, opts...).ToFunc()
}

// ByEmail orders the results by the Email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByWebsite orders the results by the Website field.
func ByWebsite(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWebsite, opts...).ToFunc()
}

// ByTaxAdmin orders the results by the TaxAdmin field.
func ByTaxAdmin(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTaxAdmin, opts...).ToFunc()
}

// ByTaxNo orders the results by the TaxNo field.
func ByTaxNo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTaxNo, opts...).ToFunc()
}

// ByCommerce orders the results by the Commerce field.
func ByCommerce(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCommerce, opts...).ToFunc()
}

// ByCommerceReg orders the results by the CommerceReg field.
func ByCommerceReg(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCommerceReg, opts...).ToFunc()
}

// ByVisaDate orders the results by the VisaDate field.
func ByVisaDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVisaDate, opts...).ToFunc()
}

// ByCreatedAt orders the results by the CreatedAt field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the UpdatedAt field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByCompanyOwnerField orders the results by companyOwner field.
func ByCompanyOwnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCompanyOwnerStep(), sql.OrderByField(field, opts...))
	}
}

// ByUsersCount orders the results by users count.
func ByUsersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUsersStep(), opts...)
	}
}

// ByUsers orders the results by users terms.
func ByUsers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUsersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newCompanyOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CompanyOwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CompanyOwnerTable, CompanyOwnerColumn),
	)
}
func newUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, UsersTable, UsersColumn),
	)
}
