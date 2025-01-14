// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CompanyDetailsColumns holds the columns for the "company_details" table.
	CompanyDetailsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "company_code", Type: field.TypeInt, Unique: true},
		{Name: "name", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "address", Type: field.TypeString, Nullable: true},
		{Name: "phone", Type: field.TypeString, Nullable: true},
		{Name: "fax", Type: field.TypeString, Nullable: true},
		{Name: "mobile_phone", Type: field.TypeString, Nullable: true},
		{Name: "email", Type: field.TypeString, Nullable: true},
		{Name: "website", Type: field.TypeString, Nullable: true},
		{Name: "tax_admin", Type: field.TypeString, Nullable: true},
		{Name: "tax_no", Type: field.TypeInt, Nullable: true, Default: 0},
		{Name: "chamber_info", Type: field.TypeString, Nullable: true},
		{Name: "chamber_register_no", Type: field.TypeString, Nullable: true},
		{Name: "visa_date", Type: field.TypeTime, Nullable: true},
		{Name: "visa_end_date", Type: field.TypeTime, Nullable: true},
		{Name: "visa_finished_for90days", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "core_person_absent90days", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "is_closed", Type: field.TypeBool, Default: false},
		{Name: "owner_name", Type: field.TypeString, Nullable: true},
		{Name: "owner_tc_no", Type: field.TypeInt, Nullable: true},
		{Name: "owner_address", Type: field.TypeString, Nullable: true},
		{Name: "owner_phone", Type: field.TypeString, Nullable: true},
		{Name: "owner_email", Type: field.TypeString, Nullable: true},
		{Name: "owner_register_no", Type: field.TypeInt, Nullable: true},
		{Name: "owner_career", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// CompanyDetailsTable holds the schema information for the "company_details" table.
	CompanyDetailsTable = &schema.Table{
		Name:       "company_details",
		Columns:    CompanyDetailsColumns,
		PrimaryKey: []*schema.Column{CompanyDetailsColumns[0]},
	}
	// CompanyEngineersColumns holds the columns for the "company_engineers" table.
	CompanyEngineersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Default: ""},
		{Name: "tc_no", Type: field.TypeInt, Nullable: true},
		{Name: "phone", Type: field.TypeString, Nullable: true},
		{Name: "email", Type: field.TypeString, Nullable: true},
		{Name: "address", Type: field.TypeString, Nullable: true},
		{Name: "career", Type: field.TypeString, Nullable: true},
		{Name: "position", Type: field.TypeString, Nullable: true},
		{Name: "register_no", Type: field.TypeInt, Nullable: true},
		{Name: "cert_no", Type: field.TypeInt, Nullable: true},
		{Name: "ydsid", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "employment", Type: field.TypeTime, Nullable: true},
		{Name: "status", Type: field.TypeInt, Default: 1},
		{Name: "note", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "company_id", Type: field.TypeInt, Nullable: true},
	}
	// CompanyEngineersTable holds the schema information for the "company_engineers" table.
	CompanyEngineersTable = &schema.Table{
		Name:       "company_engineers",
		Columns:    CompanyEngineersColumns,
		PrimaryKey: []*schema.Column{CompanyEngineersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "company_engineers_company_details_engineers",
				Columns:    []*schema.Column{CompanyEngineersColumns[16]},
				RefColumns: []*schema.Column{CompanyDetailsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CompanyTokensColumns holds the columns for the "company_tokens" table.
	CompanyTokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "token", Type: field.TypeString, Nullable: true},
		{Name: "department_id", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "expire", Type: field.TypeInt, Nullable: true},
		{Name: "refresh_token", Type: field.TypeString, Nullable: true},
		{Name: "secret_key", Type: field.TypeString, Nullable: true},
		{Name: "secure_secret_key", Type: field.TypeString, Nullable: true},
		{Name: "otp_uri", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "company_id", Type: field.TypeInt, Nullable: true},
	}
	// CompanyTokensTable holds the schema information for the "company_tokens" table.
	CompanyTokensTable = &schema.Table{
		Name:       "company_tokens",
		Columns:    CompanyTokensColumns,
		PrimaryKey: []*schema.Column{CompanyTokensColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "company_tokens_company_details_tokens",
				Columns:    []*schema.Column{CompanyTokensColumns[10]},
				RefColumns: []*schema.Column{CompanyDetailsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CompanyUsersColumns holds the columns for the "company_users" table.
	CompanyUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "company_id", Type: field.TypeInt, Nullable: true},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
	}
	// CompanyUsersTable holds the schema information for the "company_users" table.
	CompanyUsersTable = &schema.Table{
		Name:       "company_users",
		Columns:    CompanyUsersColumns,
		PrimaryKey: []*schema.Column{CompanyUsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "company_users_company_details_users",
				Columns:    []*schema.Column{CompanyUsersColumns[1]},
				RefColumns: []*schema.Column{CompanyDetailsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "company_users_users_companies",
				Columns:    []*schema.Column{CompanyUsersColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "companyuser_user_id_company_id",
				Unique:  true,
				Columns: []*schema.Column{CompanyUsersColumns[2], CompanyUsersColumns[1]},
			},
		},
	}
	// JobAuthorsColumns holds the columns for the "job_authors" table.
	JobAuthorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "static", Type: field.TypeString, Nullable: true},
		{Name: "mechanic", Type: field.TypeString, Nullable: true},
		{Name: "electric", Type: field.TypeString, Nullable: true},
		{Name: "architect", Type: field.TypeString, Nullable: true},
		{Name: "geotechnical_engineer", Type: field.TypeString, Nullable: true},
		{Name: "geotechnical_geologist", Type: field.TypeString, Nullable: true},
		{Name: "geotechnical_geophysicist", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// JobAuthorsTable holds the schema information for the "job_authors" table.
	JobAuthorsTable = &schema.Table{
		Name:       "job_authors",
		Columns:    JobAuthorsColumns,
		PrimaryKey: []*schema.Column{JobAuthorsColumns[0]},
	}
	// JobContractorsColumns holds the columns for the "job_contractors" table.
	JobContractorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Default: ""},
		{Name: "tc_no", Type: field.TypeInt, Nullable: true},
		{Name: "register_no", Type: field.TypeInt, Nullable: true},
		{Name: "address", Type: field.TypeString, Nullable: true},
		{Name: "tax_no", Type: field.TypeInt, Nullable: true},
		{Name: "mobile_phone", Type: field.TypeString, Nullable: true},
		{Name: "phone", Type: field.TypeString, Nullable: true},
		{Name: "email", Type: field.TypeString, Nullable: true},
		{Name: "person_type", Type: field.TypeString, Nullable: true},
		{Name: "ydsid", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "note", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// JobContractorsTable holds the schema information for the "job_contractors" table.
	JobContractorsTable = &schema.Table{
		Name:       "job_contractors",
		Columns:    JobContractorsColumns,
		PrimaryKey: []*schema.Column{JobContractorsColumns[0]},
	}
	// JobDetailsColumns holds the columns for the "job_details" table.
	JobDetailsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "yibf_no", Type: field.TypeInt, Unique: true},
		{Name: "title", Type: field.TypeString, Nullable: true},
		{Name: "administration", Type: field.TypeString, Nullable: true},
		{Name: "state", Type: field.TypeString, Nullable: true},
		{Name: "island", Type: field.TypeString, Nullable: true},
		{Name: "parcel", Type: field.TypeString, Nullable: true},
		{Name: "sheet", Type: field.TypeString, Nullable: true},
		{Name: "contract_date", Type: field.TypeTime, Nullable: true},
		{Name: "start_date", Type: field.TypeTime, Nullable: true},
		{Name: "license_date", Type: field.TypeTime, Nullable: true},
		{Name: "license_no", Type: field.TypeString, Nullable: true},
		{Name: "completion_date", Type: field.TypeTime, Nullable: true},
		{Name: "land_area", Type: field.TypeFloat64, Nullable: true},
		{Name: "total_area", Type: field.TypeFloat64, Nullable: true},
		{Name: "construction_area", Type: field.TypeFloat64, Nullable: true},
		{Name: "left_area", Type: field.TypeFloat64, Nullable: true},
		{Name: "yds_address", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "address", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "building_class", Type: field.TypeString, Nullable: true},
		{Name: "building_type", Type: field.TypeString, Nullable: true},
		{Name: "level", Type: field.TypeFloat64, Nullable: true},
		{Name: "unit_price", Type: field.TypeFloat64, Nullable: true},
		{Name: "floor_count", Type: field.TypeInt, Nullable: true},
		{Name: "bks_reference_no", Type: field.TypeInt, Nullable: true},
		{Name: "coordinates", Type: field.TypeString, Nullable: true},
		{Name: "folder_no", Type: field.TypeString, Nullable: true},
		{Name: "uploaded_file", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "industry_area", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "cluster_structure", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "is_license_expired", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "is_completed", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "note", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "company_id", Type: field.TypeInt, Nullable: true},
		{Name: "static_id", Type: field.TypeInt, Nullable: true},
		{Name: "mechanic_id", Type: field.TypeInt, Nullable: true},
		{Name: "electric_id", Type: field.TypeInt, Nullable: true},
		{Name: "inspector_id", Type: field.TypeInt, Nullable: true},
		{Name: "architect_id", Type: field.TypeInt, Nullable: true},
		{Name: "controller_id", Type: field.TypeInt, Nullable: true},
		{Name: "mechaniccontroller_id", Type: field.TypeInt, Nullable: true},
		{Name: "electriccontroller_id", Type: field.TypeInt, Nullable: true},
		{Name: "author_id", Type: field.TypeInt, Nullable: true},
		{Name: "contractor_id", Type: field.TypeInt, Nullable: true},
		{Name: "owner_id", Type: field.TypeInt, Nullable: true},
		{Name: "progress_id", Type: field.TypeInt, Nullable: true},
		{Name: "supervisor_id", Type: field.TypeInt, Nullable: true},
	}
	// JobDetailsTable holds the schema information for the "job_details" table.
	JobDetailsTable = &schema.Table{
		Name:       "job_details",
		Columns:    JobDetailsColumns,
		PrimaryKey: []*schema.Column{JobDetailsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "job_details_company_details_jobs",
				Columns:    []*schema.Column{JobDetailsColumns[35]},
				RefColumns: []*schema.Column{CompanyDetailsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "job_details_company_engineers_statics",
				Columns:    []*schema.Column{JobDetailsColumns[36]},
				RefColumns: []*schema.Column{CompanyEngineersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "job_details_company_engineers_mechanics",
				Columns:    []*schema.Column{JobDetailsColumns[37]},
				RefColumns: []*schema.Column{CompanyEngineersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "job_details_company_engineers_electrics",
				Columns:    []*schema.Column{JobDetailsColumns[38]},
				RefColumns: []*schema.Column{CompanyEngineersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "job_details_company_engineers_inspectors",
				Columns:    []*schema.Column{JobDetailsColumns[39]},
				RefColumns: []*schema.Column{CompanyEngineersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "job_details_company_engineers_architects",
				Columns:    []*schema.Column{JobDetailsColumns[40]},
				RefColumns: []*schema.Column{CompanyEngineersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "job_details_company_engineers_controllers",
				Columns:    []*schema.Column{JobDetailsColumns[41]},
				RefColumns: []*schema.Column{CompanyEngineersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "job_details_company_engineers_mechaniccontrollers",
				Columns:    []*schema.Column{JobDetailsColumns[42]},
				RefColumns: []*schema.Column{CompanyEngineersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "job_details_company_engineers_electriccontrollers",
				Columns:    []*schema.Column{JobDetailsColumns[43]},
				RefColumns: []*schema.Column{CompanyEngineersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "job_details_job_authors_authors",
				Columns:    []*schema.Column{JobDetailsColumns[44]},
				RefColumns: []*schema.Column{JobAuthorsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "job_details_job_contractors_contractors",
				Columns:    []*schema.Column{JobDetailsColumns[45]},
				RefColumns: []*schema.Column{JobContractorsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "job_details_job_owners_owners",
				Columns:    []*schema.Column{JobDetailsColumns[46]},
				RefColumns: []*schema.Column{JobOwnersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "job_details_job_progresses_progress",
				Columns:    []*schema.Column{JobDetailsColumns[47]},
				RefColumns: []*schema.Column{JobProgressesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "job_details_job_supervisors_supervisors",
				Columns:    []*schema.Column{JobDetailsColumns[48]},
				RefColumns: []*schema.Column{JobSupervisorsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// JobLayersColumns holds the columns for the "job_layers" table.
	JobLayersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Default: ""},
		{Name: "metre", Type: field.TypeString, Default: ""},
		{Name: "mold_date", Type: field.TypeTime, Nullable: true},
		{Name: "concrete_date", Type: field.TypeTime, Nullable: true},
		{Name: "samples", Type: field.TypeInt, Nullable: true},
		{Name: "concrete_class", Type: field.TypeString, Nullable: true},
		{Name: "week_result", Type: field.TypeString, Nullable: true},
		{Name: "month_result", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "job_id", Type: field.TypeInt, Nullable: true},
	}
	// JobLayersTable holds the schema information for the "job_layers" table.
	JobLayersTable = &schema.Table{
		Name:       "job_layers",
		Columns:    JobLayersColumns,
		PrimaryKey: []*schema.Column{JobLayersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "job_layers_job_details_layers",
				Columns:    []*schema.Column{JobLayersColumns[11]},
				RefColumns: []*schema.Column{JobDetailsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// JobOwnersColumns holds the columns for the "job_owners" table.
	JobOwnersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Default: ""},
		{Name: "address", Type: field.TypeString, Nullable: true},
		{Name: "tc_no", Type: field.TypeInt, Nullable: true},
		{Name: "tax_admin", Type: field.TypeString, Nullable: true},
		{Name: "tax_no", Type: field.TypeInt, Nullable: true},
		{Name: "phone", Type: field.TypeString, Nullable: true},
		{Name: "email", Type: field.TypeString, Nullable: true},
		{Name: "ydsid", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "shareholder", Type: field.TypeBool, Default: false},
		{Name: "note", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// JobOwnersTable holds the schema information for the "job_owners" table.
	JobOwnersTable = &schema.Table{
		Name:       "job_owners",
		Columns:    JobOwnersColumns,
		PrimaryKey: []*schema.Column{JobOwnersColumns[0]},
	}
	// JobPaymentsColumns holds the columns for the "job_payments" table.
	JobPaymentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "date", Type: field.TypeTime},
		{Name: "amount", Type: field.TypeInt, Nullable: true, Default: 0},
		{Name: "description", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "status", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "percentage", Type: field.TypeFloat64, Nullable: true, Default: 0.5},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "payments_id", Type: field.TypeInt, Nullable: true},
	}
	// JobPaymentsTable holds the schema information for the "job_payments" table.
	JobPaymentsTable = &schema.Table{
		Name:       "job_payments",
		Columns:    JobPaymentsColumns,
		PrimaryKey: []*schema.Column{JobPaymentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "job_payments_job_details_payments",
				Columns:    []*schema.Column{JobPaymentsColumns[8]},
				RefColumns: []*schema.Column{JobDetailsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// JobProgressesColumns holds the columns for the "job_progresses" table.
	JobProgressesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "one", Type: field.TypeInt, Nullable: true, Default: 0},
		{Name: "two", Type: field.TypeInt, Nullable: true, Default: 0},
		{Name: "three", Type: field.TypeInt, Nullable: true, Default: 0},
		{Name: "four", Type: field.TypeInt, Nullable: true, Default: 0},
		{Name: "five", Type: field.TypeInt, Nullable: true, Default: 0},
		{Name: "six", Type: field.TypeInt, Nullable: true, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// JobProgressesTable holds the schema information for the "job_progresses" table.
	JobProgressesTable = &schema.Table{
		Name:       "job_progresses",
		Columns:    JobProgressesColumns,
		PrimaryKey: []*schema.Column{JobProgressesColumns[0]},
	}
	// JobSupervisorsColumns holds the columns for the "job_supervisors" table.
	JobSupervisorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Nullable: true},
		{Name: "address", Type: field.TypeString, Nullable: true},
		{Name: "phone", Type: field.TypeString, Nullable: true},
		{Name: "email", Type: field.TypeString, Nullable: true},
		{Name: "tc_no", Type: field.TypeInt, Nullable: true},
		{Name: "position", Type: field.TypeString, Nullable: true},
		{Name: "career", Type: field.TypeString, Nullable: true},
		{Name: "register_no", Type: field.TypeInt, Nullable: true},
		{Name: "social_security_no", Type: field.TypeInt, Nullable: true},
		{Name: "school_graduation", Type: field.TypeString, Nullable: true},
		{Name: "ydsid", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// JobSupervisorsTable holds the schema information for the "job_supervisors" table.
	JobSupervisorsTable = &schema.Table{
		Name:       "job_supervisors",
		Columns:    JobSupervisorsColumns,
		PrimaryKey: []*schema.Column{JobSupervisorsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString, Default: ""},
		{Name: "email", Type: field.TypeString, Default: ""},
		{Name: "phone", Type: field.TypeString, Nullable: true},
		{Name: "password", Type: field.TypeString},
		{Name: "role", Type: field.TypeString, Default: "User"},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CompanyDetailsTable,
		CompanyEngineersTable,
		CompanyTokensTable,
		CompanyUsersTable,
		JobAuthorsTable,
		JobContractorsTable,
		JobDetailsTable,
		JobLayersTable,
		JobOwnersTable,
		JobPaymentsTable,
		JobProgressesTable,
		JobSupervisorsTable,
		UsersTable,
	}
)

func init() {
	CompanyEngineersTable.ForeignKeys[0].RefTable = CompanyDetailsTable
	CompanyTokensTable.ForeignKeys[0].RefTable = CompanyDetailsTable
	CompanyUsersTable.ForeignKeys[0].RefTable = CompanyDetailsTable
	CompanyUsersTable.ForeignKeys[1].RefTable = UsersTable
	JobDetailsTable.ForeignKeys[0].RefTable = CompanyDetailsTable
	JobDetailsTable.ForeignKeys[1].RefTable = CompanyEngineersTable
	JobDetailsTable.ForeignKeys[2].RefTable = CompanyEngineersTable
	JobDetailsTable.ForeignKeys[3].RefTable = CompanyEngineersTable
	JobDetailsTable.ForeignKeys[4].RefTable = CompanyEngineersTable
	JobDetailsTable.ForeignKeys[5].RefTable = CompanyEngineersTable
	JobDetailsTable.ForeignKeys[6].RefTable = CompanyEngineersTable
	JobDetailsTable.ForeignKeys[7].RefTable = CompanyEngineersTable
	JobDetailsTable.ForeignKeys[8].RefTable = CompanyEngineersTable
	JobDetailsTable.ForeignKeys[9].RefTable = JobAuthorsTable
	JobDetailsTable.ForeignKeys[10].RefTable = JobContractorsTable
	JobDetailsTable.ForeignKeys[11].RefTable = JobOwnersTable
	JobDetailsTable.ForeignKeys[12].RefTable = JobProgressesTable
	JobDetailsTable.ForeignKeys[13].RefTable = JobSupervisorsTable
	JobLayersTable.ForeignKeys[0].RefTable = JobDetailsTable
	JobPaymentsTable.ForeignKeys[0].RefTable = JobDetailsTable
}
