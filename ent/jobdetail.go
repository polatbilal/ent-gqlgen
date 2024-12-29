// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"gqlgen-ent/ent/companydetail"
	"gqlgen-ent/ent/companyengineer"
	"gqlgen-ent/ent/jobauthor"
	"gqlgen-ent/ent/jobcontractor"
	"gqlgen-ent/ent/jobdetail"
	"gqlgen-ent/ent/jobowner"
	"gqlgen-ent/ent/jobprogress"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// JobDetail is the model entity for the JobDetail schema.
type JobDetail struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// YibfNo holds the value of the "YibfNo" field.
	YibfNo int `json:"YibfNo,omitempty"`
	// Idare holds the value of the "Idare" field.
	Idare string `json:"Idare,omitempty"`
	// Pafta holds the value of the "Pafta" field.
	Pafta string `json:"Pafta,omitempty"`
	// Ada holds the value of the "Ada" field.
	Ada string `json:"Ada,omitempty"`
	// Parsel holds the value of the "Parsel" field.
	Parsel string `json:"Parsel,omitempty"`
	// FolderNo holds the value of the "FolderNo" field.
	FolderNo string `json:"FolderNo,omitempty"`
	// Status holds the value of the "Status" field.
	Status int `json:"Status,omitempty"`
	// ContractDate holds the value of the "ContractDate" field.
	ContractDate time.Time `json:"ContractDate,omitempty"`
	// CompletionDate holds the value of the "CompletionDate" field.
	CompletionDate time.Time `json:"CompletionDate,omitempty"`
	// StartDate holds the value of the "StartDate" field.
	StartDate time.Time `json:"StartDate,omitempty"`
	// LicenseDate holds the value of the "LicenseDate" field.
	LicenseDate time.Time `json:"LicenseDate,omitempty"`
	// LicenseNo holds the value of the "LicenseNo" field.
	LicenseNo string `json:"LicenseNo,omitempty"`
	// ConstructionArea holds the value of the "ConstructionArea" field.
	ConstructionArea string `json:"ConstructionArea,omitempty"`
	// City holds the value of the "City" field.
	City string `json:"City,omitempty"`
	// District holds the value of the "District" field.
	District string `json:"District,omitempty"`
	// Village holds the value of the "Village" field.
	Village string `json:"Village,omitempty"`
	// Street holds the value of the "Street" field.
	Street string `json:"Street,omitempty"`
	// BuildingClass holds the value of the "BuildingClass" field.
	BuildingClass string `json:"BuildingClass,omitempty"`
	// BuildingType holds the value of the "BuildingType" field.
	BuildingType string `json:"BuildingType,omitempty"`
	// BuildingBlock holds the value of the "BuildingBlock" field.
	BuildingBlock string `json:"BuildingBlock,omitempty"`
	// LandArea holds the value of the "LandArea" field.
	LandArea string `json:"LandArea,omitempty"`
	// Floors holds the value of the "Floors" field.
	Floors int `json:"Floors,omitempty"`
	// UsagePurpose holds the value of the "UsagePurpose" field.
	UsagePurpose string `json:"UsagePurpose,omitempty"`
	// Note holds the value of the "Note" field.
	Note string `json:"Note,omitempty"`
	// Started holds the value of the "Started" field.
	Started int `json:"Started,omitempty"`
	// Deleted holds the value of the "Deleted" field.
	Deleted int `json:"Deleted,omitempty"`
	// CreatedAt holds the value of the "CreatedAt" field.
	CreatedAt time.Time `json:"CreatedAt,omitempty"`
	// UpdatedAt holds the value of the "UpdatedAt" field.
	UpdatedAt time.Time `json:"UpdatedAt,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the JobDetailQuery when eager-loading is set.
	Edges                 JobDetailEdges `json:"edges"`
	company_id            *int
	inspector_id          *int
	architect_id          *int
	static_id             *int
	mechanic_id           *int
	electric_id           *int
	controller_id         *int
	mechaniccontroller_id *int
	electriccontroller_id *int
	author_id             *int
	contractor_id         *int
	owner_id              *int
	progress_id           *int
	selectValues          sql.SelectValues
}

// JobDetailEdges holds the relations/edges for other nodes in the graph.
type JobDetailEdges struct {
	// Owner holds the value of the owner edge.
	Owner *JobOwner `json:"owner,omitempty"`
	// Contractor holds the value of the contractor edge.
	Contractor *JobContractor `json:"contractor,omitempty"`
	// Author holds the value of the author edge.
	Author *JobAuthor `json:"author,omitempty"`
	// Progress holds the value of the progress edge.
	Progress *JobProgress `json:"progress,omitempty"`
	// Inspector holds the value of the inspector edge.
	Inspector *CompanyEngineer `json:"inspector,omitempty"`
	// Architect holds the value of the architect edge.
	Architect *CompanyEngineer `json:"architect,omitempty"`
	// Static holds the value of the static edge.
	Static *CompanyEngineer `json:"static,omitempty"`
	// Mechanic holds the value of the mechanic edge.
	Mechanic *CompanyEngineer `json:"mechanic,omitempty"`
	// Electric holds the value of the electric edge.
	Electric *CompanyEngineer `json:"electric,omitempty"`
	// Controller holds the value of the controller edge.
	Controller *CompanyEngineer `json:"controller,omitempty"`
	// Mechaniccontroller holds the value of the mechaniccontroller edge.
	Mechaniccontroller *CompanyEngineer `json:"mechaniccontroller,omitempty"`
	// Electriccontroller holds the value of the electriccontroller edge.
	Electriccontroller *CompanyEngineer `json:"electriccontroller,omitempty"`
	// Layers holds the value of the layers edge.
	Layers []*JobLayer `json:"layers,omitempty"`
	// Payments holds the value of the payments edge.
	Payments []*JobPayments `json:"payments,omitempty"`
	// Company holds the value of the company edge.
	Company *CompanyDetail `json:"company,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [15]bool
	// totalCount holds the count of the edges above.
	totalCount [15]map[string]int

	namedLayers   map[string][]*JobLayer
	namedPayments map[string][]*JobPayments
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e JobDetailEdges) OwnerOrErr() (*JobOwner, error) {
	if e.Owner != nil {
		return e.Owner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: jobowner.Label}
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// ContractorOrErr returns the Contractor value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e JobDetailEdges) ContractorOrErr() (*JobContractor, error) {
	if e.Contractor != nil {
		return e.Contractor, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: jobcontractor.Label}
	}
	return nil, &NotLoadedError{edge: "contractor"}
}

// AuthorOrErr returns the Author value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e JobDetailEdges) AuthorOrErr() (*JobAuthor, error) {
	if e.Author != nil {
		return e.Author, nil
	} else if e.loadedTypes[2] {
		return nil, &NotFoundError{label: jobauthor.Label}
	}
	return nil, &NotLoadedError{edge: "author"}
}

// ProgressOrErr returns the Progress value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e JobDetailEdges) ProgressOrErr() (*JobProgress, error) {
	if e.Progress != nil {
		return e.Progress, nil
	} else if e.loadedTypes[3] {
		return nil, &NotFoundError{label: jobprogress.Label}
	}
	return nil, &NotLoadedError{edge: "progress"}
}

// InspectorOrErr returns the Inspector value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e JobDetailEdges) InspectorOrErr() (*CompanyEngineer, error) {
	if e.Inspector != nil {
		return e.Inspector, nil
	} else if e.loadedTypes[4] {
		return nil, &NotFoundError{label: companyengineer.Label}
	}
	return nil, &NotLoadedError{edge: "inspector"}
}

// ArchitectOrErr returns the Architect value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e JobDetailEdges) ArchitectOrErr() (*CompanyEngineer, error) {
	if e.Architect != nil {
		return e.Architect, nil
	} else if e.loadedTypes[5] {
		return nil, &NotFoundError{label: companyengineer.Label}
	}
	return nil, &NotLoadedError{edge: "architect"}
}

// StaticOrErr returns the Static value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e JobDetailEdges) StaticOrErr() (*CompanyEngineer, error) {
	if e.Static != nil {
		return e.Static, nil
	} else if e.loadedTypes[6] {
		return nil, &NotFoundError{label: companyengineer.Label}
	}
	return nil, &NotLoadedError{edge: "static"}
}

// MechanicOrErr returns the Mechanic value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e JobDetailEdges) MechanicOrErr() (*CompanyEngineer, error) {
	if e.Mechanic != nil {
		return e.Mechanic, nil
	} else if e.loadedTypes[7] {
		return nil, &NotFoundError{label: companyengineer.Label}
	}
	return nil, &NotLoadedError{edge: "mechanic"}
}

// ElectricOrErr returns the Electric value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e JobDetailEdges) ElectricOrErr() (*CompanyEngineer, error) {
	if e.Electric != nil {
		return e.Electric, nil
	} else if e.loadedTypes[8] {
		return nil, &NotFoundError{label: companyengineer.Label}
	}
	return nil, &NotLoadedError{edge: "electric"}
}

// ControllerOrErr returns the Controller value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e JobDetailEdges) ControllerOrErr() (*CompanyEngineer, error) {
	if e.Controller != nil {
		return e.Controller, nil
	} else if e.loadedTypes[9] {
		return nil, &NotFoundError{label: companyengineer.Label}
	}
	return nil, &NotLoadedError{edge: "controller"}
}

// MechaniccontrollerOrErr returns the Mechaniccontroller value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e JobDetailEdges) MechaniccontrollerOrErr() (*CompanyEngineer, error) {
	if e.Mechaniccontroller != nil {
		return e.Mechaniccontroller, nil
	} else if e.loadedTypes[10] {
		return nil, &NotFoundError{label: companyengineer.Label}
	}
	return nil, &NotLoadedError{edge: "mechaniccontroller"}
}

// ElectriccontrollerOrErr returns the Electriccontroller value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e JobDetailEdges) ElectriccontrollerOrErr() (*CompanyEngineer, error) {
	if e.Electriccontroller != nil {
		return e.Electriccontroller, nil
	} else if e.loadedTypes[11] {
		return nil, &NotFoundError{label: companyengineer.Label}
	}
	return nil, &NotLoadedError{edge: "electriccontroller"}
}

// LayersOrErr returns the Layers value or an error if the edge
// was not loaded in eager-loading.
func (e JobDetailEdges) LayersOrErr() ([]*JobLayer, error) {
	if e.loadedTypes[12] {
		return e.Layers, nil
	}
	return nil, &NotLoadedError{edge: "layers"}
}

// PaymentsOrErr returns the Payments value or an error if the edge
// was not loaded in eager-loading.
func (e JobDetailEdges) PaymentsOrErr() ([]*JobPayments, error) {
	if e.loadedTypes[13] {
		return e.Payments, nil
	}
	return nil, &NotLoadedError{edge: "payments"}
}

// CompanyOrErr returns the Company value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e JobDetailEdges) CompanyOrErr() (*CompanyDetail, error) {
	if e.Company != nil {
		return e.Company, nil
	} else if e.loadedTypes[14] {
		return nil, &NotFoundError{label: companydetail.Label}
	}
	return nil, &NotLoadedError{edge: "company"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*JobDetail) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case jobdetail.FieldID, jobdetail.FieldYibfNo, jobdetail.FieldStatus, jobdetail.FieldFloors, jobdetail.FieldStarted, jobdetail.FieldDeleted:
			values[i] = new(sql.NullInt64)
		case jobdetail.FieldIdare, jobdetail.FieldPafta, jobdetail.FieldAda, jobdetail.FieldParsel, jobdetail.FieldFolderNo, jobdetail.FieldLicenseNo, jobdetail.FieldConstructionArea, jobdetail.FieldCity, jobdetail.FieldDistrict, jobdetail.FieldVillage, jobdetail.FieldStreet, jobdetail.FieldBuildingClass, jobdetail.FieldBuildingType, jobdetail.FieldBuildingBlock, jobdetail.FieldLandArea, jobdetail.FieldUsagePurpose, jobdetail.FieldNote:
			values[i] = new(sql.NullString)
		case jobdetail.FieldContractDate, jobdetail.FieldCompletionDate, jobdetail.FieldStartDate, jobdetail.FieldLicenseDate, jobdetail.FieldCreatedAt, jobdetail.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case jobdetail.ForeignKeys[0]: // company_id
			values[i] = new(sql.NullInt64)
		case jobdetail.ForeignKeys[1]: // inspector_id
			values[i] = new(sql.NullInt64)
		case jobdetail.ForeignKeys[2]: // architect_id
			values[i] = new(sql.NullInt64)
		case jobdetail.ForeignKeys[3]: // static_id
			values[i] = new(sql.NullInt64)
		case jobdetail.ForeignKeys[4]: // mechanic_id
			values[i] = new(sql.NullInt64)
		case jobdetail.ForeignKeys[5]: // electric_id
			values[i] = new(sql.NullInt64)
		case jobdetail.ForeignKeys[6]: // controller_id
			values[i] = new(sql.NullInt64)
		case jobdetail.ForeignKeys[7]: // mechaniccontroller_id
			values[i] = new(sql.NullInt64)
		case jobdetail.ForeignKeys[8]: // electriccontroller_id
			values[i] = new(sql.NullInt64)
		case jobdetail.ForeignKeys[9]: // author_id
			values[i] = new(sql.NullInt64)
		case jobdetail.ForeignKeys[10]: // contractor_id
			values[i] = new(sql.NullInt64)
		case jobdetail.ForeignKeys[11]: // owner_id
			values[i] = new(sql.NullInt64)
		case jobdetail.ForeignKeys[12]: // progress_id
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the JobDetail fields.
func (jd *JobDetail) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case jobdetail.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			jd.ID = int(value.Int64)
		case jobdetail.FieldYibfNo:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field YibfNo", values[i])
			} else if value.Valid {
				jd.YibfNo = int(value.Int64)
			}
		case jobdetail.FieldIdare:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Idare", values[i])
			} else if value.Valid {
				jd.Idare = value.String
			}
		case jobdetail.FieldPafta:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Pafta", values[i])
			} else if value.Valid {
				jd.Pafta = value.String
			}
		case jobdetail.FieldAda:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Ada", values[i])
			} else if value.Valid {
				jd.Ada = value.String
			}
		case jobdetail.FieldParsel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Parsel", values[i])
			} else if value.Valid {
				jd.Parsel = value.String
			}
		case jobdetail.FieldFolderNo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field FolderNo", values[i])
			} else if value.Valid {
				jd.FolderNo = value.String
			}
		case jobdetail.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field Status", values[i])
			} else if value.Valid {
				jd.Status = int(value.Int64)
			}
		case jobdetail.FieldContractDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field ContractDate", values[i])
			} else if value.Valid {
				jd.ContractDate = value.Time
			}
		case jobdetail.FieldCompletionDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field CompletionDate", values[i])
			} else if value.Valid {
				jd.CompletionDate = value.Time
			}
		case jobdetail.FieldStartDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field StartDate", values[i])
			} else if value.Valid {
				jd.StartDate = value.Time
			}
		case jobdetail.FieldLicenseDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field LicenseDate", values[i])
			} else if value.Valid {
				jd.LicenseDate = value.Time
			}
		case jobdetail.FieldLicenseNo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field LicenseNo", values[i])
			} else if value.Valid {
				jd.LicenseNo = value.String
			}
		case jobdetail.FieldConstructionArea:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ConstructionArea", values[i])
			} else if value.Valid {
				jd.ConstructionArea = value.String
			}
		case jobdetail.FieldCity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field City", values[i])
			} else if value.Valid {
				jd.City = value.String
			}
		case jobdetail.FieldDistrict:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field District", values[i])
			} else if value.Valid {
				jd.District = value.String
			}
		case jobdetail.FieldVillage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Village", values[i])
			} else if value.Valid {
				jd.Village = value.String
			}
		case jobdetail.FieldStreet:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Street", values[i])
			} else if value.Valid {
				jd.Street = value.String
			}
		case jobdetail.FieldBuildingClass:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field BuildingClass", values[i])
			} else if value.Valid {
				jd.BuildingClass = value.String
			}
		case jobdetail.FieldBuildingType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field BuildingType", values[i])
			} else if value.Valid {
				jd.BuildingType = value.String
			}
		case jobdetail.FieldBuildingBlock:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field BuildingBlock", values[i])
			} else if value.Valid {
				jd.BuildingBlock = value.String
			}
		case jobdetail.FieldLandArea:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field LandArea", values[i])
			} else if value.Valid {
				jd.LandArea = value.String
			}
		case jobdetail.FieldFloors:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field Floors", values[i])
			} else if value.Valid {
				jd.Floors = int(value.Int64)
			}
		case jobdetail.FieldUsagePurpose:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field UsagePurpose", values[i])
			} else if value.Valid {
				jd.UsagePurpose = value.String
			}
		case jobdetail.FieldNote:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Note", values[i])
			} else if value.Valid {
				jd.Note = value.String
			}
		case jobdetail.FieldStarted:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field Started", values[i])
			} else if value.Valid {
				jd.Started = int(value.Int64)
			}
		case jobdetail.FieldDeleted:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field Deleted", values[i])
			} else if value.Valid {
				jd.Deleted = int(value.Int64)
			}
		case jobdetail.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field CreatedAt", values[i])
			} else if value.Valid {
				jd.CreatedAt = value.Time
			}
		case jobdetail.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field UpdatedAt", values[i])
			} else if value.Valid {
				jd.UpdatedAt = value.Time
			}
		case jobdetail.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field company_id", value)
			} else if value.Valid {
				jd.company_id = new(int)
				*jd.company_id = int(value.Int64)
			}
		case jobdetail.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field inspector_id", value)
			} else if value.Valid {
				jd.inspector_id = new(int)
				*jd.inspector_id = int(value.Int64)
			}
		case jobdetail.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field architect_id", value)
			} else if value.Valid {
				jd.architect_id = new(int)
				*jd.architect_id = int(value.Int64)
			}
		case jobdetail.ForeignKeys[3]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field static_id", value)
			} else if value.Valid {
				jd.static_id = new(int)
				*jd.static_id = int(value.Int64)
			}
		case jobdetail.ForeignKeys[4]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field mechanic_id", value)
			} else if value.Valid {
				jd.mechanic_id = new(int)
				*jd.mechanic_id = int(value.Int64)
			}
		case jobdetail.ForeignKeys[5]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field electric_id", value)
			} else if value.Valid {
				jd.electric_id = new(int)
				*jd.electric_id = int(value.Int64)
			}
		case jobdetail.ForeignKeys[6]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field controller_id", value)
			} else if value.Valid {
				jd.controller_id = new(int)
				*jd.controller_id = int(value.Int64)
			}
		case jobdetail.ForeignKeys[7]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field mechaniccontroller_id", value)
			} else if value.Valid {
				jd.mechaniccontroller_id = new(int)
				*jd.mechaniccontroller_id = int(value.Int64)
			}
		case jobdetail.ForeignKeys[8]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field electriccontroller_id", value)
			} else if value.Valid {
				jd.electriccontroller_id = new(int)
				*jd.electriccontroller_id = int(value.Int64)
			}
		case jobdetail.ForeignKeys[9]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field author_id", value)
			} else if value.Valid {
				jd.author_id = new(int)
				*jd.author_id = int(value.Int64)
			}
		case jobdetail.ForeignKeys[10]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field contractor_id", value)
			} else if value.Valid {
				jd.contractor_id = new(int)
				*jd.contractor_id = int(value.Int64)
			}
		case jobdetail.ForeignKeys[11]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field owner_id", value)
			} else if value.Valid {
				jd.owner_id = new(int)
				*jd.owner_id = int(value.Int64)
			}
		case jobdetail.ForeignKeys[12]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field progress_id", value)
			} else if value.Valid {
				jd.progress_id = new(int)
				*jd.progress_id = int(value.Int64)
			}
		default:
			jd.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the JobDetail.
// This includes values selected through modifiers, order, etc.
func (jd *JobDetail) Value(name string) (ent.Value, error) {
	return jd.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the JobDetail entity.
func (jd *JobDetail) QueryOwner() *JobOwnerQuery {
	return NewJobDetailClient(jd.config).QueryOwner(jd)
}

// QueryContractor queries the "contractor" edge of the JobDetail entity.
func (jd *JobDetail) QueryContractor() *JobContractorQuery {
	return NewJobDetailClient(jd.config).QueryContractor(jd)
}

// QueryAuthor queries the "author" edge of the JobDetail entity.
func (jd *JobDetail) QueryAuthor() *JobAuthorQuery {
	return NewJobDetailClient(jd.config).QueryAuthor(jd)
}

// QueryProgress queries the "progress" edge of the JobDetail entity.
func (jd *JobDetail) QueryProgress() *JobProgressQuery {
	return NewJobDetailClient(jd.config).QueryProgress(jd)
}

// QueryInspector queries the "inspector" edge of the JobDetail entity.
func (jd *JobDetail) QueryInspector() *CompanyEngineerQuery {
	return NewJobDetailClient(jd.config).QueryInspector(jd)
}

// QueryArchitect queries the "architect" edge of the JobDetail entity.
func (jd *JobDetail) QueryArchitect() *CompanyEngineerQuery {
	return NewJobDetailClient(jd.config).QueryArchitect(jd)
}

// QueryStatic queries the "static" edge of the JobDetail entity.
func (jd *JobDetail) QueryStatic() *CompanyEngineerQuery {
	return NewJobDetailClient(jd.config).QueryStatic(jd)
}

// QueryMechanic queries the "mechanic" edge of the JobDetail entity.
func (jd *JobDetail) QueryMechanic() *CompanyEngineerQuery {
	return NewJobDetailClient(jd.config).QueryMechanic(jd)
}

// QueryElectric queries the "electric" edge of the JobDetail entity.
func (jd *JobDetail) QueryElectric() *CompanyEngineerQuery {
	return NewJobDetailClient(jd.config).QueryElectric(jd)
}

// QueryController queries the "controller" edge of the JobDetail entity.
func (jd *JobDetail) QueryController() *CompanyEngineerQuery {
	return NewJobDetailClient(jd.config).QueryController(jd)
}

// QueryMechaniccontroller queries the "mechaniccontroller" edge of the JobDetail entity.
func (jd *JobDetail) QueryMechaniccontroller() *CompanyEngineerQuery {
	return NewJobDetailClient(jd.config).QueryMechaniccontroller(jd)
}

// QueryElectriccontroller queries the "electriccontroller" edge of the JobDetail entity.
func (jd *JobDetail) QueryElectriccontroller() *CompanyEngineerQuery {
	return NewJobDetailClient(jd.config).QueryElectriccontroller(jd)
}

// QueryLayers queries the "layers" edge of the JobDetail entity.
func (jd *JobDetail) QueryLayers() *JobLayerQuery {
	return NewJobDetailClient(jd.config).QueryLayers(jd)
}

// QueryPayments queries the "payments" edge of the JobDetail entity.
func (jd *JobDetail) QueryPayments() *JobPaymentsQuery {
	return NewJobDetailClient(jd.config).QueryPayments(jd)
}

// QueryCompany queries the "company" edge of the JobDetail entity.
func (jd *JobDetail) QueryCompany() *CompanyDetailQuery {
	return NewJobDetailClient(jd.config).QueryCompany(jd)
}

// Update returns a builder for updating this JobDetail.
// Note that you need to call JobDetail.Unwrap() before calling this method if this JobDetail
// was returned from a transaction, and the transaction was committed or rolled back.
func (jd *JobDetail) Update() *JobDetailUpdateOne {
	return NewJobDetailClient(jd.config).UpdateOne(jd)
}

// Unwrap unwraps the JobDetail entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (jd *JobDetail) Unwrap() *JobDetail {
	_tx, ok := jd.config.driver.(*txDriver)
	if !ok {
		panic("ent: JobDetail is not a transactional entity")
	}
	jd.config.driver = _tx.drv
	return jd
}

// String implements the fmt.Stringer.
func (jd *JobDetail) String() string {
	var builder strings.Builder
	builder.WriteString("JobDetail(")
	builder.WriteString(fmt.Sprintf("id=%v, ", jd.ID))
	builder.WriteString("YibfNo=")
	builder.WriteString(fmt.Sprintf("%v", jd.YibfNo))
	builder.WriteString(", ")
	builder.WriteString("Idare=")
	builder.WriteString(jd.Idare)
	builder.WriteString(", ")
	builder.WriteString("Pafta=")
	builder.WriteString(jd.Pafta)
	builder.WriteString(", ")
	builder.WriteString("Ada=")
	builder.WriteString(jd.Ada)
	builder.WriteString(", ")
	builder.WriteString("Parsel=")
	builder.WriteString(jd.Parsel)
	builder.WriteString(", ")
	builder.WriteString("FolderNo=")
	builder.WriteString(jd.FolderNo)
	builder.WriteString(", ")
	builder.WriteString("Status=")
	builder.WriteString(fmt.Sprintf("%v", jd.Status))
	builder.WriteString(", ")
	builder.WriteString("ContractDate=")
	builder.WriteString(jd.ContractDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("CompletionDate=")
	builder.WriteString(jd.CompletionDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("StartDate=")
	builder.WriteString(jd.StartDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("LicenseDate=")
	builder.WriteString(jd.LicenseDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("LicenseNo=")
	builder.WriteString(jd.LicenseNo)
	builder.WriteString(", ")
	builder.WriteString("ConstructionArea=")
	builder.WriteString(jd.ConstructionArea)
	builder.WriteString(", ")
	builder.WriteString("City=")
	builder.WriteString(jd.City)
	builder.WriteString(", ")
	builder.WriteString("District=")
	builder.WriteString(jd.District)
	builder.WriteString(", ")
	builder.WriteString("Village=")
	builder.WriteString(jd.Village)
	builder.WriteString(", ")
	builder.WriteString("Street=")
	builder.WriteString(jd.Street)
	builder.WriteString(", ")
	builder.WriteString("BuildingClass=")
	builder.WriteString(jd.BuildingClass)
	builder.WriteString(", ")
	builder.WriteString("BuildingType=")
	builder.WriteString(jd.BuildingType)
	builder.WriteString(", ")
	builder.WriteString("BuildingBlock=")
	builder.WriteString(jd.BuildingBlock)
	builder.WriteString(", ")
	builder.WriteString("LandArea=")
	builder.WriteString(jd.LandArea)
	builder.WriteString(", ")
	builder.WriteString("Floors=")
	builder.WriteString(fmt.Sprintf("%v", jd.Floors))
	builder.WriteString(", ")
	builder.WriteString("UsagePurpose=")
	builder.WriteString(jd.UsagePurpose)
	builder.WriteString(", ")
	builder.WriteString("Note=")
	builder.WriteString(jd.Note)
	builder.WriteString(", ")
	builder.WriteString("Started=")
	builder.WriteString(fmt.Sprintf("%v", jd.Started))
	builder.WriteString(", ")
	builder.WriteString("Deleted=")
	builder.WriteString(fmt.Sprintf("%v", jd.Deleted))
	builder.WriteString(", ")
	builder.WriteString("CreatedAt=")
	builder.WriteString(jd.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("UpdatedAt=")
	builder.WriteString(jd.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// NamedLayers returns the Layers named value or an error if the edge was not
// loaded in eager-loading with this name.
func (jd *JobDetail) NamedLayers(name string) ([]*JobLayer, error) {
	if jd.Edges.namedLayers == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := jd.Edges.namedLayers[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (jd *JobDetail) appendNamedLayers(name string, edges ...*JobLayer) {
	if jd.Edges.namedLayers == nil {
		jd.Edges.namedLayers = make(map[string][]*JobLayer)
	}
	if len(edges) == 0 {
		jd.Edges.namedLayers[name] = []*JobLayer{}
	} else {
		jd.Edges.namedLayers[name] = append(jd.Edges.namedLayers[name], edges...)
	}
}

// NamedPayments returns the Payments named value or an error if the edge was not
// loaded in eager-loading with this name.
func (jd *JobDetail) NamedPayments(name string) ([]*JobPayments, error) {
	if jd.Edges.namedPayments == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := jd.Edges.namedPayments[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (jd *JobDetail) appendNamedPayments(name string, edges ...*JobPayments) {
	if jd.Edges.namedPayments == nil {
		jd.Edges.namedPayments = make(map[string][]*JobPayments)
	}
	if len(edges) == 0 {
		jd.Edges.namedPayments[name] = []*JobPayments{}
	} else {
		jd.Edges.namedPayments[name] = append(jd.Edges.namedPayments[name], edges...)
	}
}

// JobDetails is a parsable slice of JobDetail.
type JobDetails []*JobDetail
