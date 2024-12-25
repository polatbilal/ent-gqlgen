// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"gqlgen-ent/ent/jobdetail"
	"gqlgen-ent/ent/joblayer"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// JobLayer is the model entity for the JobLayer schema.
type JobLayer struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "Name" field.
	Name string `json:"Name,omitempty"`
	// Metre holds the value of the "Metre" field.
	Metre string `json:"Metre,omitempty"`
	// MoldDate holds the value of the "MoldDate" field.
	MoldDate time.Time `json:"MoldDate,omitempty"`
	// ConcreteDate holds the value of the "ConcreteDate" field.
	ConcreteDate time.Time `json:"ConcreteDate,omitempty"`
	// Samples holds the value of the "Samples" field.
	Samples int `json:"Samples,omitempty"`
	// ConcreteClass holds the value of the "ConcreteClass" field.
	ConcreteClass string `json:"ConcreteClass,omitempty"`
	// WeekResult holds the value of the "WeekResult" field.
	WeekResult string `json:"WeekResult,omitempty"`
	// MonthResult holds the value of the "MonthResult" field.
	MonthResult string `json:"MonthResult,omitempty"`
	// CreatedAt holds the value of the "CreatedAt" field.
	CreatedAt time.Time `json:"CreatedAt,omitempty"`
	// UpdatedAt holds the value of the "UpdatedAt" field.
	UpdatedAt time.Time `json:"UpdatedAt,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the JobLayerQuery when eager-loading is set.
	Edges        JobLayerEdges `json:"edges"`
	job_id       *int
	selectValues sql.SelectValues
}

// JobLayerEdges holds the relations/edges for other nodes in the graph.
type JobLayerEdges struct {
	// Layer holds the value of the layer edge.
	Layer *JobDetail `json:"layer,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// LayerOrErr returns the Layer value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e JobLayerEdges) LayerOrErr() (*JobDetail, error) {
	if e.Layer != nil {
		return e.Layer, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: jobdetail.Label}
	}
	return nil, &NotLoadedError{edge: "layer"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*JobLayer) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case joblayer.FieldID, joblayer.FieldSamples:
			values[i] = new(sql.NullInt64)
		case joblayer.FieldName, joblayer.FieldMetre, joblayer.FieldConcreteClass, joblayer.FieldWeekResult, joblayer.FieldMonthResult:
			values[i] = new(sql.NullString)
		case joblayer.FieldMoldDate, joblayer.FieldConcreteDate, joblayer.FieldCreatedAt, joblayer.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case joblayer.ForeignKeys[0]: // job_id
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the JobLayer fields.
func (jl *JobLayer) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case joblayer.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			jl.ID = int(value.Int64)
		case joblayer.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Name", values[i])
			} else if value.Valid {
				jl.Name = value.String
			}
		case joblayer.FieldMetre:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Metre", values[i])
			} else if value.Valid {
				jl.Metre = value.String
			}
		case joblayer.FieldMoldDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field MoldDate", values[i])
			} else if value.Valid {
				jl.MoldDate = value.Time
			}
		case joblayer.FieldConcreteDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field ConcreteDate", values[i])
			} else if value.Valid {
				jl.ConcreteDate = value.Time
			}
		case joblayer.FieldSamples:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field Samples", values[i])
			} else if value.Valid {
				jl.Samples = int(value.Int64)
			}
		case joblayer.FieldConcreteClass:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ConcreteClass", values[i])
			} else if value.Valid {
				jl.ConcreteClass = value.String
			}
		case joblayer.FieldWeekResult:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field WeekResult", values[i])
			} else if value.Valid {
				jl.WeekResult = value.String
			}
		case joblayer.FieldMonthResult:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field MonthResult", values[i])
			} else if value.Valid {
				jl.MonthResult = value.String
			}
		case joblayer.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field CreatedAt", values[i])
			} else if value.Valid {
				jl.CreatedAt = value.Time
			}
		case joblayer.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field UpdatedAt", values[i])
			} else if value.Valid {
				jl.UpdatedAt = value.Time
			}
		case joblayer.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field job_id", value)
			} else if value.Valid {
				jl.job_id = new(int)
				*jl.job_id = int(value.Int64)
			}
		default:
			jl.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the JobLayer.
// This includes values selected through modifiers, order, etc.
func (jl *JobLayer) Value(name string) (ent.Value, error) {
	return jl.selectValues.Get(name)
}

// QueryLayer queries the "layer" edge of the JobLayer entity.
func (jl *JobLayer) QueryLayer() *JobDetailQuery {
	return NewJobLayerClient(jl.config).QueryLayer(jl)
}

// Update returns a builder for updating this JobLayer.
// Note that you need to call JobLayer.Unwrap() before calling this method if this JobLayer
// was returned from a transaction, and the transaction was committed or rolled back.
func (jl *JobLayer) Update() *JobLayerUpdateOne {
	return NewJobLayerClient(jl.config).UpdateOne(jl)
}

// Unwrap unwraps the JobLayer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (jl *JobLayer) Unwrap() *JobLayer {
	_tx, ok := jl.config.driver.(*txDriver)
	if !ok {
		panic("ent: JobLayer is not a transactional entity")
	}
	jl.config.driver = _tx.drv
	return jl
}

// String implements the fmt.Stringer.
func (jl *JobLayer) String() string {
	var builder strings.Builder
	builder.WriteString("JobLayer(")
	builder.WriteString(fmt.Sprintf("id=%v, ", jl.ID))
	builder.WriteString("Name=")
	builder.WriteString(jl.Name)
	builder.WriteString(", ")
	builder.WriteString("Metre=")
	builder.WriteString(jl.Metre)
	builder.WriteString(", ")
	builder.WriteString("MoldDate=")
	builder.WriteString(jl.MoldDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("ConcreteDate=")
	builder.WriteString(jl.ConcreteDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("Samples=")
	builder.WriteString(fmt.Sprintf("%v", jl.Samples))
	builder.WriteString(", ")
	builder.WriteString("ConcreteClass=")
	builder.WriteString(jl.ConcreteClass)
	builder.WriteString(", ")
	builder.WriteString("WeekResult=")
	builder.WriteString(jl.WeekResult)
	builder.WriteString(", ")
	builder.WriteString("MonthResult=")
	builder.WriteString(jl.MonthResult)
	builder.WriteString(", ")
	builder.WriteString("CreatedAt=")
	builder.WriteString(jl.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("UpdatedAt=")
	builder.WriteString(jl.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// JobLayers is a parsable slice of JobLayer.
type JobLayers []*JobLayer
