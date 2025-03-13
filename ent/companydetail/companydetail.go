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
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldFax holds the string denoting the fax field in the database.
	FieldFax = "fax"
	// FieldMobilePhone holds the string denoting the mobilephone field in the database.
	FieldMobilePhone = "mobile_phone"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldWebsite holds the string denoting the website field in the database.
	FieldWebsite = "website"
	// FieldTaxAdmin holds the string denoting the taxadmin field in the database.
	FieldTaxAdmin = "tax_admin"
	// FieldTaxNo holds the string denoting the taxno field in the database.
	FieldTaxNo = "tax_no"
	// FieldChamberInfo holds the string denoting the chamberinfo field in the database.
	FieldChamberInfo = "chamber_info"
	// FieldChamberRegisterNo holds the string denoting the chamberregisterno field in the database.
	FieldChamberRegisterNo = "chamber_register_no"
	// FieldVisaDate holds the string denoting the visadate field in the database.
	FieldVisaDate = "visa_date"
	// FieldVisaEndDate holds the string denoting the visaenddate field in the database.
	FieldVisaEndDate = "visa_end_date"
	// FieldVisaFinishedFor90Days holds the string denoting the visafinishedfor90days field in the database.
	FieldVisaFinishedFor90Days = "visa_finished_for90days"
	// FieldCorePersonAbsent90Days holds the string denoting the corepersonabsent90days field in the database.
	FieldCorePersonAbsent90Days = "core_person_absent90days"
	// FieldIsClosed holds the string denoting the isclosed field in the database.
	FieldIsClosed = "is_closed"
	// FieldOwnerName holds the string denoting the ownername field in the database.
	FieldOwnerName = "owner_name"
	// FieldOwnerTcNo holds the string denoting the ownertcno field in the database.
	FieldOwnerTcNo = "owner_tc_no"
	// FieldOwnerAddress holds the string denoting the owneraddress field in the database.
	FieldOwnerAddress = "owner_address"
	// FieldOwnerPhone holds the string denoting the ownerphone field in the database.
	FieldOwnerPhone = "owner_phone"
	// FieldOwnerEmail holds the string denoting the owneremail field in the database.
	FieldOwnerEmail = "owner_email"
	// FieldOwnerRegisterNo holds the string denoting the ownerregisterno field in the database.
	FieldOwnerRegisterNo = "owner_register_no"
	// FieldOwnerCareer holds the string denoting the ownercareer field in the database.
	FieldOwnerCareer = "owner_career"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeJobs holds the string denoting the jobs edge name in mutations.
	EdgeJobs = "jobs"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgeTokens holds the string denoting the tokens edge name in mutations.
	EdgeTokens = "tokens"
	// EdgeEngineers holds the string denoting the engineers edge name in mutations.
	EdgeEngineers = "engineers"
	// Table holds the table name of the companydetail in the database.
	Table = "company_details"
	// JobsTable is the table that holds the jobs relation/edge.
	JobsTable = "job_relations"
	// JobsInverseTable is the table name for the JobRelations entity.
	// It exists in this package in order to avoid circular dependency with the "jobrelations" package.
	JobsInverseTable = "job_relations"
	// JobsColumn is the table column denoting the jobs relation/edge.
	JobsColumn = "company_id"
	// UsersTable is the table that holds the users relation/edge.
	UsersTable = "company_users"
	// UsersInverseTable is the table name for the CompanyUser entity.
	// It exists in this package in order to avoid circular dependency with the "companyuser" package.
	UsersInverseTable = "company_users"
	// UsersColumn is the table column denoting the users relation/edge.
	UsersColumn = "company_id"
	// TokensTable is the table that holds the tokens relation/edge.
	TokensTable = "company_tokens"
	// TokensInverseTable is the table name for the CompanyToken entity.
	// It exists in this package in order to avoid circular dependency with the "companytoken" package.
	TokensInverseTable = "company_tokens"
	// TokensColumn is the table column denoting the tokens relation/edge.
	TokensColumn = "company_id"
	// EngineersTable is the table that holds the engineers relation/edge.
	EngineersTable = "company_engineers"
	// EngineersInverseTable is the table name for the CompanyEngineer entity.
	// It exists in this package in order to avoid circular dependency with the "companyengineer" package.
	EngineersInverseTable = "company_engineers"
	// EngineersColumn is the table column denoting the engineers relation/edge.
	EngineersColumn = "company_id"
)

// Columns holds all SQL columns for companydetail fields.
var Columns = []string{
	FieldID,
	FieldCompanyCode,
	FieldName,
	FieldAddress,
	FieldPhone,
	FieldFax,
	FieldMobilePhone,
	FieldEmail,
	FieldWebsite,
	FieldTaxAdmin,
	FieldTaxNo,
	FieldChamberInfo,
	FieldChamberRegisterNo,
	FieldVisaDate,
	FieldVisaEndDate,
	FieldVisaFinishedFor90Days,
	FieldCorePersonAbsent90Days,
	FieldIsClosed,
	FieldOwnerName,
	FieldOwnerTcNo,
	FieldOwnerAddress,
	FieldOwnerPhone,
	FieldOwnerEmail,
	FieldOwnerRegisterNo,
	FieldOwnerCareer,
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
	// DefaultName holds the default value on creation for the "Name" field.
	DefaultName string
	// DefaultVisaFinishedFor90Days holds the default value on creation for the "VisaFinishedFor90Days" field.
	DefaultVisaFinishedFor90Days bool
	// DefaultCorePersonAbsent90Days holds the default value on creation for the "CorePersonAbsent90Days" field.
	DefaultCorePersonAbsent90Days bool
	// DefaultIsClosed holds the default value on creation for the "IsClosed" field.
	DefaultIsClosed bool
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

// ByPhone orders the results by the Phone field.
func ByPhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhone, opts...).ToFunc()
}

// ByFax orders the results by the Fax field.
func ByFax(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFax, opts...).ToFunc()
}

// ByMobilePhone orders the results by the MobilePhone field.
func ByMobilePhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMobilePhone, opts...).ToFunc()
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

// ByChamberInfo orders the results by the ChamberInfo field.
func ByChamberInfo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldChamberInfo, opts...).ToFunc()
}

// ByChamberRegisterNo orders the results by the ChamberRegisterNo field.
func ByChamberRegisterNo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldChamberRegisterNo, opts...).ToFunc()
}

// ByVisaDate orders the results by the VisaDate field.
func ByVisaDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVisaDate, opts...).ToFunc()
}

// ByVisaEndDate orders the results by the VisaEndDate field.
func ByVisaEndDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVisaEndDate, opts...).ToFunc()
}

// ByVisaFinishedFor90Days orders the results by the VisaFinishedFor90Days field.
func ByVisaFinishedFor90Days(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVisaFinishedFor90Days, opts...).ToFunc()
}

// ByCorePersonAbsent90Days orders the results by the CorePersonAbsent90Days field.
func ByCorePersonAbsent90Days(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCorePersonAbsent90Days, opts...).ToFunc()
}

// ByIsClosed orders the results by the IsClosed field.
func ByIsClosed(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsClosed, opts...).ToFunc()
}

// ByOwnerName orders the results by the OwnerName field.
func ByOwnerName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOwnerName, opts...).ToFunc()
}

// ByOwnerTcNo orders the results by the OwnerTcNo field.
func ByOwnerTcNo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOwnerTcNo, opts...).ToFunc()
}

// ByOwnerAddress orders the results by the OwnerAddress field.
func ByOwnerAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOwnerAddress, opts...).ToFunc()
}

// ByOwnerPhone orders the results by the OwnerPhone field.
func ByOwnerPhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOwnerPhone, opts...).ToFunc()
}

// ByOwnerEmail orders the results by the OwnerEmail field.
func ByOwnerEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOwnerEmail, opts...).ToFunc()
}

// ByOwnerRegisterNo orders the results by the OwnerRegisterNo field.
func ByOwnerRegisterNo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOwnerRegisterNo, opts...).ToFunc()
}

// ByOwnerCareer orders the results by the OwnerCareer field.
func ByOwnerCareer(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOwnerCareer, opts...).ToFunc()
}

// ByCreatedAt orders the results by the CreatedAt field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the UpdatedAt field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByJobsCount orders the results by jobs count.
func ByJobsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newJobsStep(), opts...)
	}
}

// ByJobs orders the results by jobs terms.
func ByJobs(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newJobsStep(), append([]sql.OrderTerm{term}, terms...)...)
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

// ByTokensCount orders the results by tokens count.
func ByTokensCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTokensStep(), opts...)
	}
}

// ByTokens orders the results by tokens terms.
func ByTokens(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTokensStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByEngineersCount orders the results by engineers count.
func ByEngineersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newEngineersStep(), opts...)
	}
}

// ByEngineers orders the results by engineers terms.
func ByEngineers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEngineersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newJobsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(JobsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, JobsTable, JobsColumn),
	)
}
func newUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, UsersTable, UsersColumn),
	)
}
func newTokensStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TokensInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TokensTable, TokensColumn),
	)
}
func newEngineersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EngineersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, EngineersTable, EngineersColumn),
	)
}
