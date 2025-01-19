// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/polatbilal/gqlgen-ent/ent/companydetail"
	"github.com/polatbilal/gqlgen-ent/ent/companyengineer"
	"github.com/polatbilal/gqlgen-ent/ent/jobauthor"
	"github.com/polatbilal/gqlgen-ent/ent/jobcontractor"
	"github.com/polatbilal/gqlgen-ent/ent/jobdetail"
	"github.com/polatbilal/gqlgen-ent/ent/jobowner"
	"github.com/polatbilal/gqlgen-ent/ent/jobprogress"
	"github.com/polatbilal/gqlgen-ent/ent/jobsupervisor"
)

// JobDetail is the model entity for the JobDetail schema.
type JobDetail struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// YibfNo holds the value of the "YibfNo" field.
	YibfNo int `json:"YibfNo,omitempty"`
	// Title holds the value of the "Title" field.
	Title string `json:"Title,omitempty"`
	// Administration holds the value of the "Administration" field.
	Administration string `json:"Administration,omitempty"`
	// State holds the value of the "State" field.
	State string `json:"State,omitempty"`
	// Island holds the value of the "Island" field.
	Island string `json:"Island,omitempty"`
	// Parcel holds the value of the "Parcel" field.
	Parcel string `json:"Parcel,omitempty"`
	// Sheet holds the value of the "Sheet" field.
	Sheet string `json:"Sheet,omitempty"`
	// ContractDate holds the value of the "ContractDate" field.
	ContractDate time.Time `json:"ContractDate,omitempty"`
	// StartDate holds the value of the "StartDate" field.
	StartDate time.Time `json:"StartDate,omitempty"`
	// LicenseDate holds the value of the "LicenseDate" field.
	LicenseDate time.Time `json:"LicenseDate,omitempty"`
	// LicenseNo holds the value of the "LicenseNo" field.
	LicenseNo string `json:"LicenseNo,omitempty"`
	// DistributionDate holds the value of the "DistributionDate" field.
	DistributionDate time.Time `json:"DistributionDate,omitempty"`
	// CompletionDate holds the value of the "CompletionDate" field.
	CompletionDate time.Time `json:"CompletionDate,omitempty"`
	// LandArea holds the value of the "LandArea" field.
	LandArea float64 `json:"LandArea,omitempty"`
	// TotalArea holds the value of the "TotalArea" field.
	TotalArea float64 `json:"TotalArea,omitempty"`
	// ConstructionArea holds the value of the "ConstructionArea" field.
	ConstructionArea float64 `json:"ConstructionArea,omitempty"`
	// LeftArea holds the value of the "LeftArea" field.
	LeftArea float64 `json:"LeftArea,omitempty"`
	// YDSAddress holds the value of the "YDSAddress" field.
	YDSAddress string `json:"YDSAddress,omitempty"`
	// Address holds the value of the "Address" field.
	Address string `json:"Address,omitempty"`
	// BuildingClass holds the value of the "BuildingClass" field.
	BuildingClass string `json:"BuildingClass,omitempty"`
	// BuildingType holds the value of the "BuildingType" field.
	BuildingType string `json:"BuildingType,omitempty"`
	// Level holds the value of the "Level" field.
	Level float64 `json:"Level,omitempty"`
	// UnitPrice holds the value of the "UnitPrice" field.
	UnitPrice float64 `json:"UnitPrice,omitempty"`
	// FloorCount holds the value of the "FloorCount" field.
	FloorCount int `json:"FloorCount,omitempty"`
	// BKSReferenceNo holds the value of the "BKSReferenceNo" field.
	BKSReferenceNo int `json:"BKSReferenceNo,omitempty"`
	// Coordinates holds the value of the "Coordinates" field.
	Coordinates string `json:"Coordinates,omitempty"`
	// FolderNo holds the value of the "FolderNo" field.
	FolderNo string `json:"FolderNo,omitempty"`
	// UploadedFile holds the value of the "UploadedFile" field.
	UploadedFile bool `json:"UploadedFile,omitempty"`
	// IndustryArea holds the value of the "IndustryArea" field.
	IndustryArea bool `json:"IndustryArea,omitempty"`
	// ClusterStructure holds the value of the "ClusterStructure" field.
	ClusterStructure bool `json:"ClusterStructure,omitempty"`
	// IsLicenseExpired holds the value of the "IsLicenseExpired" field.
	IsLicenseExpired bool `json:"IsLicenseExpired,omitempty"`
	// IsCompleted holds the value of the "IsCompleted" field.
	IsCompleted bool `json:"IsCompleted,omitempty"`
	// Note holds the value of the "Note" field.
	Note string `json:"Note,omitempty"`
	// CreatedAt holds the value of the "CreatedAt" field.
	CreatedAt time.Time `json:"CreatedAt,omitempty"`
	// UpdatedAt holds the value of the "UpdatedAt" field.
	UpdatedAt time.Time `json:"UpdatedAt,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the JobDetailQuery when eager-loading is set.
	Edges                 JobDetailEdges `json:"edges"`
	company_id            *int
	static_id             *int
	mechanic_id           *int
	electric_id           *int
	inspector_id          *int
	architect_id          *int
	controller_id         *int
	mechaniccontroller_id *int
	electriccontroller_id *int
	author_id             *int
	contractor_id         *int
	owner_id              *int
	progress_id           *int
	supervisor_id         *int
	selectValues          sql.SelectValues
}

// JobDetailEdges holds the relations/edges for other nodes in the graph.
type JobDetailEdges struct {
	// Owner holds the value of the owner edge.
	Owner *JobOwner `json:"owner,omitempty"`
	// Author holds the value of the author edge.
	Author *JobAuthor `json:"author,omitempty"`
	// Company holds the value of the company edge.
	Company *CompanyDetail `json:"company,omitempty"`
	// Progress holds the value of the progress edge.
	Progress *JobProgress `json:"progress,omitempty"`
	// Contractor holds the value of the contractor edge.
	Contractor *JobContractor `json:"contractor,omitempty"`
	// Supervisor holds the value of the supervisor edge.
	Supervisor *JobSupervisor `json:"supervisor,omitempty"`
	// Static holds the value of the static edge.
	Static *CompanyEngineer `json:"static,omitempty"`
	// Mechanic holds the value of the mechanic edge.
	Mechanic *CompanyEngineer `json:"mechanic,omitempty"`
	// Electric holds the value of the electric edge.
	Electric *CompanyEngineer `json:"electric,omitempty"`
	// Inspector holds the value of the inspector edge.
	Inspector *CompanyEngineer `json:"inspector,omitempty"`
	// Architect holds the value of the architect edge.
	Architect *CompanyEngineer `json:"architect,omitempty"`
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
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [16]bool
	// totalCount holds the count of the edges above.
	totalCount [16]map[string]int

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

// AuthorOrErr returns the Author value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e JobDetailEdges) AuthorOrErr() (*JobAuthor, error) {
	if e.Author != nil {
		return e.Author, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: jobauthor.Label}
	}
	return nil, &NotLoadedError{edge: "author"}
}

// CompanyOrErr returns the Company value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e JobDetailEdges) CompanyOrErr() (*CompanyDetail, error) {
	if e.Company != nil {
		return e.Company, nil
	} else if e.loadedTypes[2] {
		return nil, &NotFoundError{label: companydetail.Label}
	}
	return nil, &NotLoadedError{edge: "company"}
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

// ContractorOrErr returns the Contractor value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e JobDetailEdges) ContractorOrErr() (*JobContractor, error) {
	if e.Contractor != nil {
		return e.Contractor, nil
	} else if e.loadedTypes[4] {
		return nil, &NotFoundError{label: jobcontractor.Label}
	}
	return nil, &NotLoadedError{edge: "contractor"}
}

// SupervisorOrErr returns the Supervisor value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e JobDetailEdges) SupervisorOrErr() (*JobSupervisor, error) {
	if e.Supervisor != nil {
		return e.Supervisor, nil
	} else if e.loadedTypes[5] {
		return nil, &NotFoundError{label: jobsupervisor.Label}
	}
	return nil, &NotLoadedError{edge: "supervisor"}
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

// InspectorOrErr returns the Inspector value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e JobDetailEdges) InspectorOrErr() (*CompanyEngineer, error) {
	if e.Inspector != nil {
		return e.Inspector, nil
	} else if e.loadedTypes[9] {
		return nil, &NotFoundError{label: companyengineer.Label}
	}
	return nil, &NotLoadedError{edge: "inspector"}
}

// ArchitectOrErr returns the Architect value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e JobDetailEdges) ArchitectOrErr() (*CompanyEngineer, error) {
	if e.Architect != nil {
		return e.Architect, nil
	} else if e.loadedTypes[10] {
		return nil, &NotFoundError{label: companyengineer.Label}
	}
	return nil, &NotLoadedError{edge: "architect"}
}

// ControllerOrErr returns the Controller value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e JobDetailEdges) ControllerOrErr() (*CompanyEngineer, error) {
	if e.Controller != nil {
		return e.Controller, nil
	} else if e.loadedTypes[11] {
		return nil, &NotFoundError{label: companyengineer.Label}
	}
	return nil, &NotLoadedError{edge: "controller"}
}

// MechaniccontrollerOrErr returns the Mechaniccontroller value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e JobDetailEdges) MechaniccontrollerOrErr() (*CompanyEngineer, error) {
	if e.Mechaniccontroller != nil {
		return e.Mechaniccontroller, nil
	} else if e.loadedTypes[12] {
		return nil, &NotFoundError{label: companyengineer.Label}
	}
	return nil, &NotLoadedError{edge: "mechaniccontroller"}
}

// ElectriccontrollerOrErr returns the Electriccontroller value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e JobDetailEdges) ElectriccontrollerOrErr() (*CompanyEngineer, error) {
	if e.Electriccontroller != nil {
		return e.Electriccontroller, nil
	} else if e.loadedTypes[13] {
		return nil, &NotFoundError{label: companyengineer.Label}
	}
	return nil, &NotLoadedError{edge: "electriccontroller"}
}

// LayersOrErr returns the Layers value or an error if the edge
// was not loaded in eager-loading.
func (e JobDetailEdges) LayersOrErr() ([]*JobLayer, error) {
	if e.loadedTypes[14] {
		return e.Layers, nil
	}
	return nil, &NotLoadedError{edge: "layers"}
}

// PaymentsOrErr returns the Payments value or an error if the edge
// was not loaded in eager-loading.
func (e JobDetailEdges) PaymentsOrErr() ([]*JobPayments, error) {
	if e.loadedTypes[15] {
		return e.Payments, nil
	}
	return nil, &NotLoadedError{edge: "payments"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*JobDetail) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case jobdetail.FieldUploadedFile, jobdetail.FieldIndustryArea, jobdetail.FieldClusterStructure, jobdetail.FieldIsLicenseExpired, jobdetail.FieldIsCompleted:
			values[i] = new(sql.NullBool)
		case jobdetail.FieldLandArea, jobdetail.FieldTotalArea, jobdetail.FieldConstructionArea, jobdetail.FieldLeftArea, jobdetail.FieldLevel, jobdetail.FieldUnitPrice:
			values[i] = new(sql.NullFloat64)
		case jobdetail.FieldID, jobdetail.FieldYibfNo, jobdetail.FieldFloorCount, jobdetail.FieldBKSReferenceNo:
			values[i] = new(sql.NullInt64)
		case jobdetail.FieldTitle, jobdetail.FieldAdministration, jobdetail.FieldState, jobdetail.FieldIsland, jobdetail.FieldParcel, jobdetail.FieldSheet, jobdetail.FieldLicenseNo, jobdetail.FieldYDSAddress, jobdetail.FieldAddress, jobdetail.FieldBuildingClass, jobdetail.FieldBuildingType, jobdetail.FieldCoordinates, jobdetail.FieldFolderNo, jobdetail.FieldNote:
			values[i] = new(sql.NullString)
		case jobdetail.FieldContractDate, jobdetail.FieldStartDate, jobdetail.FieldLicenseDate, jobdetail.FieldDistributionDate, jobdetail.FieldCompletionDate, jobdetail.FieldCreatedAt, jobdetail.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case jobdetail.ForeignKeys[0]: // company_id
			values[i] = new(sql.NullInt64)
		case jobdetail.ForeignKeys[1]: // static_id
			values[i] = new(sql.NullInt64)
		case jobdetail.ForeignKeys[2]: // mechanic_id
			values[i] = new(sql.NullInt64)
		case jobdetail.ForeignKeys[3]: // electric_id
			values[i] = new(sql.NullInt64)
		case jobdetail.ForeignKeys[4]: // inspector_id
			values[i] = new(sql.NullInt64)
		case jobdetail.ForeignKeys[5]: // architect_id
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
		case jobdetail.ForeignKeys[13]: // supervisor_id
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
		case jobdetail.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Title", values[i])
			} else if value.Valid {
				jd.Title = value.String
			}
		case jobdetail.FieldAdministration:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Administration", values[i])
			} else if value.Valid {
				jd.Administration = value.String
			}
		case jobdetail.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field State", values[i])
			} else if value.Valid {
				jd.State = value.String
			}
		case jobdetail.FieldIsland:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Island", values[i])
			} else if value.Valid {
				jd.Island = value.String
			}
		case jobdetail.FieldParcel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Parcel", values[i])
			} else if value.Valid {
				jd.Parcel = value.String
			}
		case jobdetail.FieldSheet:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Sheet", values[i])
			} else if value.Valid {
				jd.Sheet = value.String
			}
		case jobdetail.FieldContractDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field ContractDate", values[i])
			} else if value.Valid {
				jd.ContractDate = value.Time
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
		case jobdetail.FieldDistributionDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field DistributionDate", values[i])
			} else if value.Valid {
				jd.DistributionDate = value.Time
			}
		case jobdetail.FieldCompletionDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field CompletionDate", values[i])
			} else if value.Valid {
				jd.CompletionDate = value.Time
			}
		case jobdetail.FieldLandArea:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field LandArea", values[i])
			} else if value.Valid {
				jd.LandArea = value.Float64
			}
		case jobdetail.FieldTotalArea:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field TotalArea", values[i])
			} else if value.Valid {
				jd.TotalArea = value.Float64
			}
		case jobdetail.FieldConstructionArea:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field ConstructionArea", values[i])
			} else if value.Valid {
				jd.ConstructionArea = value.Float64
			}
		case jobdetail.FieldLeftArea:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field LeftArea", values[i])
			} else if value.Valid {
				jd.LeftArea = value.Float64
			}
		case jobdetail.FieldYDSAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field YDSAddress", values[i])
			} else if value.Valid {
				jd.YDSAddress = value.String
			}
		case jobdetail.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Address", values[i])
			} else if value.Valid {
				jd.Address = value.String
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
		case jobdetail.FieldLevel:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field Level", values[i])
			} else if value.Valid {
				jd.Level = value.Float64
			}
		case jobdetail.FieldUnitPrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field UnitPrice", values[i])
			} else if value.Valid {
				jd.UnitPrice = value.Float64
			}
		case jobdetail.FieldFloorCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field FloorCount", values[i])
			} else if value.Valid {
				jd.FloorCount = int(value.Int64)
			}
		case jobdetail.FieldBKSReferenceNo:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field BKSReferenceNo", values[i])
			} else if value.Valid {
				jd.BKSReferenceNo = int(value.Int64)
			}
		case jobdetail.FieldCoordinates:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Coordinates", values[i])
			} else if value.Valid {
				jd.Coordinates = value.String
			}
		case jobdetail.FieldFolderNo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field FolderNo", values[i])
			} else if value.Valid {
				jd.FolderNo = value.String
			}
		case jobdetail.FieldUploadedFile:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field UploadedFile", values[i])
			} else if value.Valid {
				jd.UploadedFile = value.Bool
			}
		case jobdetail.FieldIndustryArea:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field IndustryArea", values[i])
			} else if value.Valid {
				jd.IndustryArea = value.Bool
			}
		case jobdetail.FieldClusterStructure:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field ClusterStructure", values[i])
			} else if value.Valid {
				jd.ClusterStructure = value.Bool
			}
		case jobdetail.FieldIsLicenseExpired:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field IsLicenseExpired", values[i])
			} else if value.Valid {
				jd.IsLicenseExpired = value.Bool
			}
		case jobdetail.FieldIsCompleted:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field IsCompleted", values[i])
			} else if value.Valid {
				jd.IsCompleted = value.Bool
			}
		case jobdetail.FieldNote:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Note", values[i])
			} else if value.Valid {
				jd.Note = value.String
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
				return fmt.Errorf("unexpected type %T for edge-field static_id", value)
			} else if value.Valid {
				jd.static_id = new(int)
				*jd.static_id = int(value.Int64)
			}
		case jobdetail.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field mechanic_id", value)
			} else if value.Valid {
				jd.mechanic_id = new(int)
				*jd.mechanic_id = int(value.Int64)
			}
		case jobdetail.ForeignKeys[3]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field electric_id", value)
			} else if value.Valid {
				jd.electric_id = new(int)
				*jd.electric_id = int(value.Int64)
			}
		case jobdetail.ForeignKeys[4]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field inspector_id", value)
			} else if value.Valid {
				jd.inspector_id = new(int)
				*jd.inspector_id = int(value.Int64)
			}
		case jobdetail.ForeignKeys[5]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field architect_id", value)
			} else if value.Valid {
				jd.architect_id = new(int)
				*jd.architect_id = int(value.Int64)
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
		case jobdetail.ForeignKeys[13]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field supervisor_id", value)
			} else if value.Valid {
				jd.supervisor_id = new(int)
				*jd.supervisor_id = int(value.Int64)
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

// QueryAuthor queries the "author" edge of the JobDetail entity.
func (jd *JobDetail) QueryAuthor() *JobAuthorQuery {
	return NewJobDetailClient(jd.config).QueryAuthor(jd)
}

// QueryCompany queries the "company" edge of the JobDetail entity.
func (jd *JobDetail) QueryCompany() *CompanyDetailQuery {
	return NewJobDetailClient(jd.config).QueryCompany(jd)
}

// QueryProgress queries the "progress" edge of the JobDetail entity.
func (jd *JobDetail) QueryProgress() *JobProgressQuery {
	return NewJobDetailClient(jd.config).QueryProgress(jd)
}

// QueryContractor queries the "contractor" edge of the JobDetail entity.
func (jd *JobDetail) QueryContractor() *JobContractorQuery {
	return NewJobDetailClient(jd.config).QueryContractor(jd)
}

// QuerySupervisor queries the "supervisor" edge of the JobDetail entity.
func (jd *JobDetail) QuerySupervisor() *JobSupervisorQuery {
	return NewJobDetailClient(jd.config).QuerySupervisor(jd)
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

// QueryInspector queries the "inspector" edge of the JobDetail entity.
func (jd *JobDetail) QueryInspector() *CompanyEngineerQuery {
	return NewJobDetailClient(jd.config).QueryInspector(jd)
}

// QueryArchitect queries the "architect" edge of the JobDetail entity.
func (jd *JobDetail) QueryArchitect() *CompanyEngineerQuery {
	return NewJobDetailClient(jd.config).QueryArchitect(jd)
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
	builder.WriteString("Title=")
	builder.WriteString(jd.Title)
	builder.WriteString(", ")
	builder.WriteString("Administration=")
	builder.WriteString(jd.Administration)
	builder.WriteString(", ")
	builder.WriteString("State=")
	builder.WriteString(jd.State)
	builder.WriteString(", ")
	builder.WriteString("Island=")
	builder.WriteString(jd.Island)
	builder.WriteString(", ")
	builder.WriteString("Parcel=")
	builder.WriteString(jd.Parcel)
	builder.WriteString(", ")
	builder.WriteString("Sheet=")
	builder.WriteString(jd.Sheet)
	builder.WriteString(", ")
	builder.WriteString("ContractDate=")
	builder.WriteString(jd.ContractDate.Format(time.ANSIC))
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
	builder.WriteString("DistributionDate=")
	builder.WriteString(jd.DistributionDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("CompletionDate=")
	builder.WriteString(jd.CompletionDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("LandArea=")
	builder.WriteString(fmt.Sprintf("%v", jd.LandArea))
	builder.WriteString(", ")
	builder.WriteString("TotalArea=")
	builder.WriteString(fmt.Sprintf("%v", jd.TotalArea))
	builder.WriteString(", ")
	builder.WriteString("ConstructionArea=")
	builder.WriteString(fmt.Sprintf("%v", jd.ConstructionArea))
	builder.WriteString(", ")
	builder.WriteString("LeftArea=")
	builder.WriteString(fmt.Sprintf("%v", jd.LeftArea))
	builder.WriteString(", ")
	builder.WriteString("YDSAddress=")
	builder.WriteString(jd.YDSAddress)
	builder.WriteString(", ")
	builder.WriteString("Address=")
	builder.WriteString(jd.Address)
	builder.WriteString(", ")
	builder.WriteString("BuildingClass=")
	builder.WriteString(jd.BuildingClass)
	builder.WriteString(", ")
	builder.WriteString("BuildingType=")
	builder.WriteString(jd.BuildingType)
	builder.WriteString(", ")
	builder.WriteString("Level=")
	builder.WriteString(fmt.Sprintf("%v", jd.Level))
	builder.WriteString(", ")
	builder.WriteString("UnitPrice=")
	builder.WriteString(fmt.Sprintf("%v", jd.UnitPrice))
	builder.WriteString(", ")
	builder.WriteString("FloorCount=")
	builder.WriteString(fmt.Sprintf("%v", jd.FloorCount))
	builder.WriteString(", ")
	builder.WriteString("BKSReferenceNo=")
	builder.WriteString(fmt.Sprintf("%v", jd.BKSReferenceNo))
	builder.WriteString(", ")
	builder.WriteString("Coordinates=")
	builder.WriteString(jd.Coordinates)
	builder.WriteString(", ")
	builder.WriteString("FolderNo=")
	builder.WriteString(jd.FolderNo)
	builder.WriteString(", ")
	builder.WriteString("UploadedFile=")
	builder.WriteString(fmt.Sprintf("%v", jd.UploadedFile))
	builder.WriteString(", ")
	builder.WriteString("IndustryArea=")
	builder.WriteString(fmt.Sprintf("%v", jd.IndustryArea))
	builder.WriteString(", ")
	builder.WriteString("ClusterStructure=")
	builder.WriteString(fmt.Sprintf("%v", jd.ClusterStructure))
	builder.WriteString(", ")
	builder.WriteString("IsLicenseExpired=")
	builder.WriteString(fmt.Sprintf("%v", jd.IsLicenseExpired))
	builder.WriteString(", ")
	builder.WriteString("IsCompleted=")
	builder.WriteString(fmt.Sprintf("%v", jd.IsCompleted))
	builder.WriteString(", ")
	builder.WriteString("Note=")
	builder.WriteString(jd.Note)
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
