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
		{Name: "company_code", Type: field.TypeInt},
		{Name: "name", Type: field.TypeString, Default: ""},
		{Name: "address", Type: field.TypeString, Nullable: true},
		{Name: "city", Type: field.TypeString, Nullable: true},
		{Name: "state", Type: field.TypeString, Nullable: true},
		{Name: "phone", Type: field.TypeString, Nullable: true},
		{Name: "fax", Type: field.TypeString, Nullable: true},
		{Name: "mobile", Type: field.TypeString, Nullable: true},
		{Name: "email", Type: field.TypeString, Nullable: true},
		{Name: "website", Type: field.TypeString, Nullable: true},
		{Name: "tax_admin", Type: field.TypeString, Nullable: true},
		{Name: "tax_no", Type: field.TypeInt, Nullable: true, Default: 0},
		{Name: "commerce", Type: field.TypeString, Nullable: true},
		{Name: "commerce_reg", Type: field.TypeString, Nullable: true},
		{Name: "visa_date", Type: field.TypeTime, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "owner_id", Type: field.TypeInt, Nullable: true},
	}
	// CompanyDetailsTable holds the schema information for the "company_details" table.
	CompanyDetailsTable = &schema.Table{
		Name:       "company_details",
		Columns:    CompanyDetailsColumns,
		PrimaryKey: []*schema.Column{CompanyDetailsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "company_details_company_engineers_companyOwners",
				Columns:    []*schema.Column{CompanyDetailsColumns[18]},
				RefColumns: []*schema.Column{CompanyEngineersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CompanyEngineersColumns holds the columns for the "company_engineers" table.
	CompanyEngineersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Default: ""},
		{Name: "address", Type: field.TypeString, Nullable: true},
		{Name: "email", Type: field.TypeString, Nullable: true},
		{Name: "tc_no", Type: field.TypeInt, Nullable: true},
		{Name: "phone", Type: field.TypeString, Nullable: true},
		{Name: "reg_no", Type: field.TypeInt, Nullable: true},
		{Name: "cert_no", Type: field.TypeInt, Nullable: true},
		{Name: "career", Type: field.TypeString, Nullable: true},
		{Name: "position", Type: field.TypeString, Nullable: true},
		{Name: "note", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeInt, Default: 1},
		{Name: "deleted", Type: field.TypeInt, Default: 0},
		{Name: "employment", Type: field.TypeTime, Nullable: true},
		{Name: "dismissal", Type: field.TypeTime, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// CompanyEngineersTable holds the schema information for the "company_engineers" table.
	CompanyEngineersTable = &schema.Table{
		Name:       "company_engineers",
		Columns:    CompanyEngineersColumns,
		PrimaryKey: []*schema.Column{CompanyEngineersColumns[0]},
	}
	// JobAuthorsColumns holds the columns for the "job_authors" table.
	JobAuthorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "architect", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "static", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "mechanic", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "electric", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "floor", Type: field.TypeString, Nullable: true, Default: ""},
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
		{Name: "address", Type: field.TypeString, Nullable: true},
		{Name: "register_no", Type: field.TypeInt, Nullable: true},
		{Name: "tax_admin", Type: field.TypeString, Nullable: true},
		{Name: "tax_no", Type: field.TypeInt, Nullable: true},
		{Name: "phone", Type: field.TypeString, Nullable: true},
		{Name: "email", Type: field.TypeString, Nullable: true},
		{Name: "note", Type: field.TypeString, Nullable: true},
		{Name: "deleted", Type: field.TypeInt, Nullable: true, Default: 0},
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
		{Name: "idare", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "pafta", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "ada", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "parsel", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "folder_no", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "status", Type: field.TypeInt, Default: 0},
		{Name: "contract_date", Type: field.TypeTime, Nullable: true},
		{Name: "completion_date", Type: field.TypeTime, Nullable: true},
		{Name: "start_date", Type: field.TypeTime, Nullable: true},
		{Name: "license_date", Type: field.TypeTime, Nullable: true},
		{Name: "license_no", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "construction_area", Type: field.TypeString, Nullable: true},
		{Name: "city", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "district", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "village", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "street", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "building_class", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "building_type", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "building_block", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "land_area", Type: field.TypeString, Nullable: true},
		{Name: "floors", Type: field.TypeInt, Nullable: true},
		{Name: "usage_purpose", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "note", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "started", Type: field.TypeInt, Default: 0},
		{Name: "deleted", Type: field.TypeInt, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "inspector_id", Type: field.TypeInt, Nullable: true},
		{Name: "architect_id", Type: field.TypeInt, Nullable: true},
		{Name: "static_id", Type: field.TypeInt, Nullable: true},
		{Name: "mechanic_id", Type: field.TypeInt, Nullable: true},
		{Name: "electric_id", Type: field.TypeInt, Nullable: true},
		{Name: "controller_id", Type: field.TypeInt, Nullable: true},
		{Name: "mechaniccontroller_id", Type: field.TypeInt, Nullable: true},
		{Name: "electriccontroller_id", Type: field.TypeInt, Nullable: true},
		{Name: "author_id", Type: field.TypeInt, Nullable: true},
		{Name: "contractor_id", Type: field.TypeInt, Nullable: true},
		{Name: "owner_id", Type: field.TypeInt, Nullable: true},
		{Name: "progress_id", Type: field.TypeInt, Nullable: true},
	}
	// JobDetailsTable holds the schema information for the "job_details" table.
	JobDetailsTable = &schema.Table{
		Name:       "job_details",
		Columns:    JobDetailsColumns,
		PrimaryKey: []*schema.Column{JobDetailsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "job_details_company_engineers_inspectors",
				Columns:    []*schema.Column{JobDetailsColumns[29]},
				RefColumns: []*schema.Column{CompanyEngineersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "job_details_company_engineers_architects",
				Columns:    []*schema.Column{JobDetailsColumns[30]},
				RefColumns: []*schema.Column{CompanyEngineersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "job_details_company_engineers_statics",
				Columns:    []*schema.Column{JobDetailsColumns[31]},
				RefColumns: []*schema.Column{CompanyEngineersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "job_details_company_engineers_mechanics",
				Columns:    []*schema.Column{JobDetailsColumns[32]},
				RefColumns: []*schema.Column{CompanyEngineersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "job_details_company_engineers_electrics",
				Columns:    []*schema.Column{JobDetailsColumns[33]},
				RefColumns: []*schema.Column{CompanyEngineersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "job_details_company_engineers_controllers",
				Columns:    []*schema.Column{JobDetailsColumns[34]},
				RefColumns: []*schema.Column{CompanyEngineersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "job_details_company_engineers_mechaniccontrollers",
				Columns:    []*schema.Column{JobDetailsColumns[35]},
				RefColumns: []*schema.Column{CompanyEngineersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "job_details_company_engineers_electriccontrollers",
				Columns:    []*schema.Column{JobDetailsColumns[36]},
				RefColumns: []*schema.Column{CompanyEngineersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "job_details_job_authors_authors",
				Columns:    []*schema.Column{JobDetailsColumns[37]},
				RefColumns: []*schema.Column{JobAuthorsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "job_details_job_contractors_contractors",
				Columns:    []*schema.Column{JobDetailsColumns[38]},
				RefColumns: []*schema.Column{JobContractorsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "job_details_job_owners_owners",
				Columns:    []*schema.Column{JobDetailsColumns[39]},
				RefColumns: []*schema.Column{JobOwnersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "job_details_job_progresses_progress",
				Columns:    []*schema.Column{JobDetailsColumns[40]},
				RefColumns: []*schema.Column{JobProgressesColumns[0]},
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
		{Name: "tc_no", Type: field.TypeInt, Nullable: true},
		{Name: "address", Type: field.TypeString, Nullable: true},
		{Name: "tax_admin", Type: field.TypeString, Nullable: true},
		{Name: "tax_no", Type: field.TypeInt, Nullable: true},
		{Name: "phone", Type: field.TypeString, Nullable: true},
		{Name: "email", Type: field.TypeString, Nullable: true},
		{Name: "note", Type: field.TypeString, Nullable: true},
		{Name: "deleted", Type: field.TypeInt, Nullable: true, Default: 0},
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
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString, Default: ""},
		{Name: "email", Type: field.TypeString, Nullable: true},
		{Name: "phone", Type: field.TypeInt, Nullable: true},
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
		JobAuthorsTable,
		JobContractorsTable,
		JobDetailsTable,
		JobLayersTable,
		JobOwnersTable,
		JobPaymentsTable,
		JobProgressesTable,
		UsersTable,
	}
)

func init() {
	CompanyDetailsTable.ForeignKeys[0].RefTable = CompanyEngineersTable
	JobDetailsTable.ForeignKeys[0].RefTable = CompanyEngineersTable
	JobDetailsTable.ForeignKeys[1].RefTable = CompanyEngineersTable
	JobDetailsTable.ForeignKeys[2].RefTable = CompanyEngineersTable
	JobDetailsTable.ForeignKeys[3].RefTable = CompanyEngineersTable
	JobDetailsTable.ForeignKeys[4].RefTable = CompanyEngineersTable
	JobDetailsTable.ForeignKeys[5].RefTable = CompanyEngineersTable
	JobDetailsTable.ForeignKeys[6].RefTable = CompanyEngineersTable
	JobDetailsTable.ForeignKeys[7].RefTable = CompanyEngineersTable
	JobDetailsTable.ForeignKeys[8].RefTable = JobAuthorsTable
	JobDetailsTable.ForeignKeys[9].RefTable = JobContractorsTable
	JobDetailsTable.ForeignKeys[10].RefTable = JobOwnersTable
	JobDetailsTable.ForeignKeys[11].RefTable = JobProgressesTable
	JobLayersTable.ForeignKeys[0].RefTable = JobDetailsTable
	JobPaymentsTable.ForeignKeys[0].RefTable = JobDetailsTable
}
