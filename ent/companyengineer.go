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
)

// CompanyEngineer is the model entity for the CompanyEngineer schema.
type CompanyEngineer struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "Name" field.
	Name string `json:"Name,omitempty"`
	// TcNo holds the value of the "TcNo" field.
	TcNo int `json:"TcNo,omitempty"`
	// Phone holds the value of the "Phone" field.
	Phone string `json:"Phone,omitempty"`
	// Email holds the value of the "Email" field.
	Email string `json:"Email,omitempty"`
	// Address holds the value of the "Address" field.
	Address string `json:"Address,omitempty"`
	// Career holds the value of the "Career" field.
	Career string `json:"Career,omitempty"`
	// Position holds the value of the "Position" field.
	Position string `json:"Position,omitempty"`
	// RegisterNo holds the value of the "RegisterNo" field.
	RegisterNo int `json:"RegisterNo,omitempty"`
	// CertNo holds the value of the "CertNo" field.
	CertNo int `json:"CertNo,omitempty"`
	// YDSID holds the value of the "YDSID" field.
	YDSID int `json:"YDSID,omitempty"`
	// Employment holds the value of the "Employment" field.
	Employment time.Time `json:"Employment,omitempty"`
	// Status holds the value of the "Status" field.
	Status int `json:"Status,omitempty"`
	// Note holds the value of the "Note" field.
	Note string `json:"Note,omitempty"`
	// CreatedAt holds the value of the "CreatedAt" field.
	CreatedAt time.Time `json:"CreatedAt,omitempty"`
	// UpdatedAt holds the value of the "UpdatedAt" field.
	UpdatedAt time.Time `json:"UpdatedAt,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CompanyEngineerQuery when eager-loading is set.
	Edges        CompanyEngineerEdges `json:"edges"`
	company_id   *int
	selectValues sql.SelectValues
}

// CompanyEngineerEdges holds the relations/edges for other nodes in the graph.
type CompanyEngineerEdges struct {
	// Company holds the value of the company edge.
	Company *CompanyDetail `json:"company,omitempty"`
	// Statics holds the value of the statics edge.
	Statics []*JobDetail `json:"statics,omitempty"`
	// Mechanics holds the value of the mechanics edge.
	Mechanics []*JobDetail `json:"mechanics,omitempty"`
	// Electrics holds the value of the electrics edge.
	Electrics []*JobDetail `json:"electrics,omitempty"`
	// Inspectors holds the value of the inspectors edge.
	Inspectors []*JobDetail `json:"inspectors,omitempty"`
	// Architects holds the value of the architects edge.
	Architects []*JobDetail `json:"architects,omitempty"`
	// Controllers holds the value of the controllers edge.
	Controllers []*JobDetail `json:"controllers,omitempty"`
	// Mechaniccontrollers holds the value of the mechaniccontrollers edge.
	Mechaniccontrollers []*JobDetail `json:"mechaniccontrollers,omitempty"`
	// Electriccontrollers holds the value of the electriccontrollers edge.
	Electriccontrollers []*JobDetail `json:"electriccontrollers,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [9]bool
	// totalCount holds the count of the edges above.
	totalCount [9]map[string]int

	namedStatics             map[string][]*JobDetail
	namedMechanics           map[string][]*JobDetail
	namedElectrics           map[string][]*JobDetail
	namedInspectors          map[string][]*JobDetail
	namedArchitects          map[string][]*JobDetail
	namedControllers         map[string][]*JobDetail
	namedMechaniccontrollers map[string][]*JobDetail
	namedElectriccontrollers map[string][]*JobDetail
}

// CompanyOrErr returns the Company value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CompanyEngineerEdges) CompanyOrErr() (*CompanyDetail, error) {
	if e.Company != nil {
		return e.Company, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: companydetail.Label}
	}
	return nil, &NotLoadedError{edge: "company"}
}

// StaticsOrErr returns the Statics value or an error if the edge
// was not loaded in eager-loading.
func (e CompanyEngineerEdges) StaticsOrErr() ([]*JobDetail, error) {
	if e.loadedTypes[1] {
		return e.Statics, nil
	}
	return nil, &NotLoadedError{edge: "statics"}
}

// MechanicsOrErr returns the Mechanics value or an error if the edge
// was not loaded in eager-loading.
func (e CompanyEngineerEdges) MechanicsOrErr() ([]*JobDetail, error) {
	if e.loadedTypes[2] {
		return e.Mechanics, nil
	}
	return nil, &NotLoadedError{edge: "mechanics"}
}

// ElectricsOrErr returns the Electrics value or an error if the edge
// was not loaded in eager-loading.
func (e CompanyEngineerEdges) ElectricsOrErr() ([]*JobDetail, error) {
	if e.loadedTypes[3] {
		return e.Electrics, nil
	}
	return nil, &NotLoadedError{edge: "electrics"}
}

// InspectorsOrErr returns the Inspectors value or an error if the edge
// was not loaded in eager-loading.
func (e CompanyEngineerEdges) InspectorsOrErr() ([]*JobDetail, error) {
	if e.loadedTypes[4] {
		return e.Inspectors, nil
	}
	return nil, &NotLoadedError{edge: "inspectors"}
}

// ArchitectsOrErr returns the Architects value or an error if the edge
// was not loaded in eager-loading.
func (e CompanyEngineerEdges) ArchitectsOrErr() ([]*JobDetail, error) {
	if e.loadedTypes[5] {
		return e.Architects, nil
	}
	return nil, &NotLoadedError{edge: "architects"}
}

// ControllersOrErr returns the Controllers value or an error if the edge
// was not loaded in eager-loading.
func (e CompanyEngineerEdges) ControllersOrErr() ([]*JobDetail, error) {
	if e.loadedTypes[6] {
		return e.Controllers, nil
	}
	return nil, &NotLoadedError{edge: "controllers"}
}

// MechaniccontrollersOrErr returns the Mechaniccontrollers value or an error if the edge
// was not loaded in eager-loading.
func (e CompanyEngineerEdges) MechaniccontrollersOrErr() ([]*JobDetail, error) {
	if e.loadedTypes[7] {
		return e.Mechaniccontrollers, nil
	}
	return nil, &NotLoadedError{edge: "mechaniccontrollers"}
}

// ElectriccontrollersOrErr returns the Electriccontrollers value or an error if the edge
// was not loaded in eager-loading.
func (e CompanyEngineerEdges) ElectriccontrollersOrErr() ([]*JobDetail, error) {
	if e.loadedTypes[8] {
		return e.Electriccontrollers, nil
	}
	return nil, &NotLoadedError{edge: "electriccontrollers"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CompanyEngineer) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case companyengineer.FieldID, companyengineer.FieldTcNo, companyengineer.FieldRegisterNo, companyengineer.FieldCertNo, companyengineer.FieldYDSID, companyengineer.FieldStatus:
			values[i] = new(sql.NullInt64)
		case companyengineer.FieldName, companyengineer.FieldPhone, companyengineer.FieldEmail, companyengineer.FieldAddress, companyengineer.FieldCareer, companyengineer.FieldPosition, companyengineer.FieldNote:
			values[i] = new(sql.NullString)
		case companyengineer.FieldEmployment, companyengineer.FieldCreatedAt, companyengineer.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case companyengineer.ForeignKeys[0]: // company_id
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CompanyEngineer fields.
func (ce *CompanyEngineer) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case companyengineer.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ce.ID = int(value.Int64)
		case companyengineer.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Name", values[i])
			} else if value.Valid {
				ce.Name = value.String
			}
		case companyengineer.FieldTcNo:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field TcNo", values[i])
			} else if value.Valid {
				ce.TcNo = int(value.Int64)
			}
		case companyengineer.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Phone", values[i])
			} else if value.Valid {
				ce.Phone = value.String
			}
		case companyengineer.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Email", values[i])
			} else if value.Valid {
				ce.Email = value.String
			}
		case companyengineer.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Address", values[i])
			} else if value.Valid {
				ce.Address = value.String
			}
		case companyengineer.FieldCareer:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Career", values[i])
			} else if value.Valid {
				ce.Career = value.String
			}
		case companyengineer.FieldPosition:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Position", values[i])
			} else if value.Valid {
				ce.Position = value.String
			}
		case companyengineer.FieldRegisterNo:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field RegisterNo", values[i])
			} else if value.Valid {
				ce.RegisterNo = int(value.Int64)
			}
		case companyengineer.FieldCertNo:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field CertNo", values[i])
			} else if value.Valid {
				ce.CertNo = int(value.Int64)
			}
		case companyengineer.FieldYDSID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field YDSID", values[i])
			} else if value.Valid {
				ce.YDSID = int(value.Int64)
			}
		case companyengineer.FieldEmployment:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field Employment", values[i])
			} else if value.Valid {
				ce.Employment = value.Time
			}
		case companyengineer.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field Status", values[i])
			} else if value.Valid {
				ce.Status = int(value.Int64)
			}
		case companyengineer.FieldNote:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Note", values[i])
			} else if value.Valid {
				ce.Note = value.String
			}
		case companyengineer.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field CreatedAt", values[i])
			} else if value.Valid {
				ce.CreatedAt = value.Time
			}
		case companyengineer.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field UpdatedAt", values[i])
			} else if value.Valid {
				ce.UpdatedAt = value.Time
			}
		case companyengineer.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field company_id", value)
			} else if value.Valid {
				ce.company_id = new(int)
				*ce.company_id = int(value.Int64)
			}
		default:
			ce.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the CompanyEngineer.
// This includes values selected through modifiers, order, etc.
func (ce *CompanyEngineer) Value(name string) (ent.Value, error) {
	return ce.selectValues.Get(name)
}

// QueryCompany queries the "company" edge of the CompanyEngineer entity.
func (ce *CompanyEngineer) QueryCompany() *CompanyDetailQuery {
	return NewCompanyEngineerClient(ce.config).QueryCompany(ce)
}

// QueryStatics queries the "statics" edge of the CompanyEngineer entity.
func (ce *CompanyEngineer) QueryStatics() *JobDetailQuery {
	return NewCompanyEngineerClient(ce.config).QueryStatics(ce)
}

// QueryMechanics queries the "mechanics" edge of the CompanyEngineer entity.
func (ce *CompanyEngineer) QueryMechanics() *JobDetailQuery {
	return NewCompanyEngineerClient(ce.config).QueryMechanics(ce)
}

// QueryElectrics queries the "electrics" edge of the CompanyEngineer entity.
func (ce *CompanyEngineer) QueryElectrics() *JobDetailQuery {
	return NewCompanyEngineerClient(ce.config).QueryElectrics(ce)
}

// QueryInspectors queries the "inspectors" edge of the CompanyEngineer entity.
func (ce *CompanyEngineer) QueryInspectors() *JobDetailQuery {
	return NewCompanyEngineerClient(ce.config).QueryInspectors(ce)
}

// QueryArchitects queries the "architects" edge of the CompanyEngineer entity.
func (ce *CompanyEngineer) QueryArchitects() *JobDetailQuery {
	return NewCompanyEngineerClient(ce.config).QueryArchitects(ce)
}

// QueryControllers queries the "controllers" edge of the CompanyEngineer entity.
func (ce *CompanyEngineer) QueryControllers() *JobDetailQuery {
	return NewCompanyEngineerClient(ce.config).QueryControllers(ce)
}

// QueryMechaniccontrollers queries the "mechaniccontrollers" edge of the CompanyEngineer entity.
func (ce *CompanyEngineer) QueryMechaniccontrollers() *JobDetailQuery {
	return NewCompanyEngineerClient(ce.config).QueryMechaniccontrollers(ce)
}

// QueryElectriccontrollers queries the "electriccontrollers" edge of the CompanyEngineer entity.
func (ce *CompanyEngineer) QueryElectriccontrollers() *JobDetailQuery {
	return NewCompanyEngineerClient(ce.config).QueryElectriccontrollers(ce)
}

// Update returns a builder for updating this CompanyEngineer.
// Note that you need to call CompanyEngineer.Unwrap() before calling this method if this CompanyEngineer
// was returned from a transaction, and the transaction was committed or rolled back.
func (ce *CompanyEngineer) Update() *CompanyEngineerUpdateOne {
	return NewCompanyEngineerClient(ce.config).UpdateOne(ce)
}

// Unwrap unwraps the CompanyEngineer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ce *CompanyEngineer) Unwrap() *CompanyEngineer {
	_tx, ok := ce.config.driver.(*txDriver)
	if !ok {
		panic("ent: CompanyEngineer is not a transactional entity")
	}
	ce.config.driver = _tx.drv
	return ce
}

// String implements the fmt.Stringer.
func (ce *CompanyEngineer) String() string {
	var builder strings.Builder
	builder.WriteString("CompanyEngineer(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ce.ID))
	builder.WriteString("Name=")
	builder.WriteString(ce.Name)
	builder.WriteString(", ")
	builder.WriteString("TcNo=")
	builder.WriteString(fmt.Sprintf("%v", ce.TcNo))
	builder.WriteString(", ")
	builder.WriteString("Phone=")
	builder.WriteString(ce.Phone)
	builder.WriteString(", ")
	builder.WriteString("Email=")
	builder.WriteString(ce.Email)
	builder.WriteString(", ")
	builder.WriteString("Address=")
	builder.WriteString(ce.Address)
	builder.WriteString(", ")
	builder.WriteString("Career=")
	builder.WriteString(ce.Career)
	builder.WriteString(", ")
	builder.WriteString("Position=")
	builder.WriteString(ce.Position)
	builder.WriteString(", ")
	builder.WriteString("RegisterNo=")
	builder.WriteString(fmt.Sprintf("%v", ce.RegisterNo))
	builder.WriteString(", ")
	builder.WriteString("CertNo=")
	builder.WriteString(fmt.Sprintf("%v", ce.CertNo))
	builder.WriteString(", ")
	builder.WriteString("YDSID=")
	builder.WriteString(fmt.Sprintf("%v", ce.YDSID))
	builder.WriteString(", ")
	builder.WriteString("Employment=")
	builder.WriteString(ce.Employment.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("Status=")
	builder.WriteString(fmt.Sprintf("%v", ce.Status))
	builder.WriteString(", ")
	builder.WriteString("Note=")
	builder.WriteString(ce.Note)
	builder.WriteString(", ")
	builder.WriteString("CreatedAt=")
	builder.WriteString(ce.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("UpdatedAt=")
	builder.WriteString(ce.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// NamedStatics returns the Statics named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ce *CompanyEngineer) NamedStatics(name string) ([]*JobDetail, error) {
	if ce.Edges.namedStatics == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ce.Edges.namedStatics[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ce *CompanyEngineer) appendNamedStatics(name string, edges ...*JobDetail) {
	if ce.Edges.namedStatics == nil {
		ce.Edges.namedStatics = make(map[string][]*JobDetail)
	}
	if len(edges) == 0 {
		ce.Edges.namedStatics[name] = []*JobDetail{}
	} else {
		ce.Edges.namedStatics[name] = append(ce.Edges.namedStatics[name], edges...)
	}
}

// NamedMechanics returns the Mechanics named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ce *CompanyEngineer) NamedMechanics(name string) ([]*JobDetail, error) {
	if ce.Edges.namedMechanics == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ce.Edges.namedMechanics[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ce *CompanyEngineer) appendNamedMechanics(name string, edges ...*JobDetail) {
	if ce.Edges.namedMechanics == nil {
		ce.Edges.namedMechanics = make(map[string][]*JobDetail)
	}
	if len(edges) == 0 {
		ce.Edges.namedMechanics[name] = []*JobDetail{}
	} else {
		ce.Edges.namedMechanics[name] = append(ce.Edges.namedMechanics[name], edges...)
	}
}

// NamedElectrics returns the Electrics named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ce *CompanyEngineer) NamedElectrics(name string) ([]*JobDetail, error) {
	if ce.Edges.namedElectrics == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ce.Edges.namedElectrics[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ce *CompanyEngineer) appendNamedElectrics(name string, edges ...*JobDetail) {
	if ce.Edges.namedElectrics == nil {
		ce.Edges.namedElectrics = make(map[string][]*JobDetail)
	}
	if len(edges) == 0 {
		ce.Edges.namedElectrics[name] = []*JobDetail{}
	} else {
		ce.Edges.namedElectrics[name] = append(ce.Edges.namedElectrics[name], edges...)
	}
}

// NamedInspectors returns the Inspectors named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ce *CompanyEngineer) NamedInspectors(name string) ([]*JobDetail, error) {
	if ce.Edges.namedInspectors == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ce.Edges.namedInspectors[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ce *CompanyEngineer) appendNamedInspectors(name string, edges ...*JobDetail) {
	if ce.Edges.namedInspectors == nil {
		ce.Edges.namedInspectors = make(map[string][]*JobDetail)
	}
	if len(edges) == 0 {
		ce.Edges.namedInspectors[name] = []*JobDetail{}
	} else {
		ce.Edges.namedInspectors[name] = append(ce.Edges.namedInspectors[name], edges...)
	}
}

// NamedArchitects returns the Architects named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ce *CompanyEngineer) NamedArchitects(name string) ([]*JobDetail, error) {
	if ce.Edges.namedArchitects == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ce.Edges.namedArchitects[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ce *CompanyEngineer) appendNamedArchitects(name string, edges ...*JobDetail) {
	if ce.Edges.namedArchitects == nil {
		ce.Edges.namedArchitects = make(map[string][]*JobDetail)
	}
	if len(edges) == 0 {
		ce.Edges.namedArchitects[name] = []*JobDetail{}
	} else {
		ce.Edges.namedArchitects[name] = append(ce.Edges.namedArchitects[name], edges...)
	}
}

// NamedControllers returns the Controllers named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ce *CompanyEngineer) NamedControllers(name string) ([]*JobDetail, error) {
	if ce.Edges.namedControllers == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ce.Edges.namedControllers[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ce *CompanyEngineer) appendNamedControllers(name string, edges ...*JobDetail) {
	if ce.Edges.namedControllers == nil {
		ce.Edges.namedControllers = make(map[string][]*JobDetail)
	}
	if len(edges) == 0 {
		ce.Edges.namedControllers[name] = []*JobDetail{}
	} else {
		ce.Edges.namedControllers[name] = append(ce.Edges.namedControllers[name], edges...)
	}
}

// NamedMechaniccontrollers returns the Mechaniccontrollers named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ce *CompanyEngineer) NamedMechaniccontrollers(name string) ([]*JobDetail, error) {
	if ce.Edges.namedMechaniccontrollers == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ce.Edges.namedMechaniccontrollers[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ce *CompanyEngineer) appendNamedMechaniccontrollers(name string, edges ...*JobDetail) {
	if ce.Edges.namedMechaniccontrollers == nil {
		ce.Edges.namedMechaniccontrollers = make(map[string][]*JobDetail)
	}
	if len(edges) == 0 {
		ce.Edges.namedMechaniccontrollers[name] = []*JobDetail{}
	} else {
		ce.Edges.namedMechaniccontrollers[name] = append(ce.Edges.namedMechaniccontrollers[name], edges...)
	}
}

// NamedElectriccontrollers returns the Electriccontrollers named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ce *CompanyEngineer) NamedElectriccontrollers(name string) ([]*JobDetail, error) {
	if ce.Edges.namedElectriccontrollers == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ce.Edges.namedElectriccontrollers[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ce *CompanyEngineer) appendNamedElectriccontrollers(name string, edges ...*JobDetail) {
	if ce.Edges.namedElectriccontrollers == nil {
		ce.Edges.namedElectriccontrollers = make(map[string][]*JobDetail)
	}
	if len(edges) == 0 {
		ce.Edges.namedElectriccontrollers[name] = []*JobDetail{}
	} else {
		ce.Edges.namedElectriccontrollers[name] = append(ce.Edges.namedElectriccontrollers[name], edges...)
	}
}

// CompanyEngineers is a parsable slice of CompanyEngineer.
type CompanyEngineers []*CompanyEngineer
