// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"gqlgen-ent/ent/companycareer"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// CompanyCareer is the model entity for the CompanyCareer schema.
type CompanyCareer struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Career holds the value of the "Career" field.
	Career string `json:"Career,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CompanyCareerQuery when eager-loading is set.
	Edges        CompanyCareerEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CompanyCareerEdges holds the relations/edges for other nodes in the graph.
type CompanyCareerEdges struct {
	// EngineerCareers holds the value of the engineerCareers edge.
	EngineerCareers []*CompanyEngineer `json:"engineerCareers,omitempty"`
	// CompanyOwnerCareers holds the value of the companyOwnerCareers edge.
	CompanyOwnerCareers []*CompanyOwner `json:"companyOwnerCareers,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedEngineerCareers     map[string][]*CompanyEngineer
	namedCompanyOwnerCareers map[string][]*CompanyOwner
}

// EngineerCareersOrErr returns the EngineerCareers value or an error if the edge
// was not loaded in eager-loading.
func (e CompanyCareerEdges) EngineerCareersOrErr() ([]*CompanyEngineer, error) {
	if e.loadedTypes[0] {
		return e.EngineerCareers, nil
	}
	return nil, &NotLoadedError{edge: "engineerCareers"}
}

// CompanyOwnerCareersOrErr returns the CompanyOwnerCareers value or an error if the edge
// was not loaded in eager-loading.
func (e CompanyCareerEdges) CompanyOwnerCareersOrErr() ([]*CompanyOwner, error) {
	if e.loadedTypes[1] {
		return e.CompanyOwnerCareers, nil
	}
	return nil, &NotLoadedError{edge: "companyOwnerCareers"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CompanyCareer) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case companycareer.FieldID:
			values[i] = new(sql.NullInt64)
		case companycareer.FieldCareer:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CompanyCareer fields.
func (cc *CompanyCareer) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case companycareer.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cc.ID = int(value.Int64)
		case companycareer.FieldCareer:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Career", values[i])
			} else if value.Valid {
				cc.Career = value.String
			}
		default:
			cc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the CompanyCareer.
// This includes values selected through modifiers, order, etc.
func (cc *CompanyCareer) Value(name string) (ent.Value, error) {
	return cc.selectValues.Get(name)
}

// QueryEngineerCareers queries the "engineerCareers" edge of the CompanyCareer entity.
func (cc *CompanyCareer) QueryEngineerCareers() *CompanyEngineerQuery {
	return NewCompanyCareerClient(cc.config).QueryEngineerCareers(cc)
}

// QueryCompanyOwnerCareers queries the "companyOwnerCareers" edge of the CompanyCareer entity.
func (cc *CompanyCareer) QueryCompanyOwnerCareers() *CompanyOwnerQuery {
	return NewCompanyCareerClient(cc.config).QueryCompanyOwnerCareers(cc)
}

// Update returns a builder for updating this CompanyCareer.
// Note that you need to call CompanyCareer.Unwrap() before calling this method if this CompanyCareer
// was returned from a transaction, and the transaction was committed or rolled back.
func (cc *CompanyCareer) Update() *CompanyCareerUpdateOne {
	return NewCompanyCareerClient(cc.config).UpdateOne(cc)
}

// Unwrap unwraps the CompanyCareer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cc *CompanyCareer) Unwrap() *CompanyCareer {
	_tx, ok := cc.config.driver.(*txDriver)
	if !ok {
		panic("ent: CompanyCareer is not a transactional entity")
	}
	cc.config.driver = _tx.drv
	return cc
}

// String implements the fmt.Stringer.
func (cc *CompanyCareer) String() string {
	var builder strings.Builder
	builder.WriteString("CompanyCareer(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cc.ID))
	builder.WriteString("Career=")
	builder.WriteString(cc.Career)
	builder.WriteByte(')')
	return builder.String()
}

// NamedEngineerCareers returns the EngineerCareers named value or an error if the edge was not
// loaded in eager-loading with this name.
func (cc *CompanyCareer) NamedEngineerCareers(name string) ([]*CompanyEngineer, error) {
	if cc.Edges.namedEngineerCareers == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := cc.Edges.namedEngineerCareers[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (cc *CompanyCareer) appendNamedEngineerCareers(name string, edges ...*CompanyEngineer) {
	if cc.Edges.namedEngineerCareers == nil {
		cc.Edges.namedEngineerCareers = make(map[string][]*CompanyEngineer)
	}
	if len(edges) == 0 {
		cc.Edges.namedEngineerCareers[name] = []*CompanyEngineer{}
	} else {
		cc.Edges.namedEngineerCareers[name] = append(cc.Edges.namedEngineerCareers[name], edges...)
	}
}

// NamedCompanyOwnerCareers returns the CompanyOwnerCareers named value or an error if the edge was not
// loaded in eager-loading with this name.
func (cc *CompanyCareer) NamedCompanyOwnerCareers(name string) ([]*CompanyOwner, error) {
	if cc.Edges.namedCompanyOwnerCareers == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := cc.Edges.namedCompanyOwnerCareers[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (cc *CompanyCareer) appendNamedCompanyOwnerCareers(name string, edges ...*CompanyOwner) {
	if cc.Edges.namedCompanyOwnerCareers == nil {
		cc.Edges.namedCompanyOwnerCareers = make(map[string][]*CompanyOwner)
	}
	if len(edges) == 0 {
		cc.Edges.namedCompanyOwnerCareers[name] = []*CompanyOwner{}
	} else {
		cc.Edges.namedCompanyOwnerCareers[name] = append(cc.Edges.namedCompanyOwnerCareers[name], edges...)
	}
}

// CompanyCareers is a parsable slice of CompanyCareer.
type CompanyCareers []*CompanyCareer
