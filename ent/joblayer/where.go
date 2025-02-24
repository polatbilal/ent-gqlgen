// Code generated by ent, DO NOT EDIT.

package joblayer

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/polatbilal/gqlgen-ent/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "Name" field. It's identical to NameEQ.
func Name(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldEQ(FieldName, v))
}

// Metre applies equality check predicate on the "Metre" field. It's identical to MetreEQ.
func Metre(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldEQ(FieldMetre, v))
}

// MoldDate applies equality check predicate on the "MoldDate" field. It's identical to MoldDateEQ.
func MoldDate(v time.Time) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldEQ(FieldMoldDate, v))
}

// ConcreteDate applies equality check predicate on the "ConcreteDate" field. It's identical to ConcreteDateEQ.
func ConcreteDate(v time.Time) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldEQ(FieldConcreteDate, v))
}

// Samples applies equality check predicate on the "Samples" field. It's identical to SamplesEQ.
func Samples(v int) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldEQ(FieldSamples, v))
}

// ConcreteClass applies equality check predicate on the "ConcreteClass" field. It's identical to ConcreteClassEQ.
func ConcreteClass(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldEQ(FieldConcreteClass, v))
}

// WeekResult applies equality check predicate on the "WeekResult" field. It's identical to WeekResultEQ.
func WeekResult(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldEQ(FieldWeekResult, v))
}

// MonthResult applies equality check predicate on the "MonthResult" field. It's identical to MonthResultEQ.
func MonthResult(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldEQ(FieldMonthResult, v))
}

// CreatedAt applies equality check predicate on the "CreatedAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "UpdatedAt" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldEQ(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "Name" field.
func NameEQ(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "Name" field.
func NameNEQ(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "Name" field.
func NameIn(vs ...string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "Name" field.
func NameNotIn(vs ...string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "Name" field.
func NameGT(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "Name" field.
func NameGTE(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "Name" field.
func NameLT(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "Name" field.
func NameLTE(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "Name" field.
func NameContains(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "Name" field.
func NameHasPrefix(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "Name" field.
func NameHasSuffix(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "Name" field.
func NameEqualFold(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "Name" field.
func NameContainsFold(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldContainsFold(FieldName, v))
}

// MetreEQ applies the EQ predicate on the "Metre" field.
func MetreEQ(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldEQ(FieldMetre, v))
}

// MetreNEQ applies the NEQ predicate on the "Metre" field.
func MetreNEQ(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldNEQ(FieldMetre, v))
}

// MetreIn applies the In predicate on the "Metre" field.
func MetreIn(vs ...string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldIn(FieldMetre, vs...))
}

// MetreNotIn applies the NotIn predicate on the "Metre" field.
func MetreNotIn(vs ...string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldNotIn(FieldMetre, vs...))
}

// MetreGT applies the GT predicate on the "Metre" field.
func MetreGT(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldGT(FieldMetre, v))
}

// MetreGTE applies the GTE predicate on the "Metre" field.
func MetreGTE(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldGTE(FieldMetre, v))
}

// MetreLT applies the LT predicate on the "Metre" field.
func MetreLT(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldLT(FieldMetre, v))
}

// MetreLTE applies the LTE predicate on the "Metre" field.
func MetreLTE(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldLTE(FieldMetre, v))
}

// MetreContains applies the Contains predicate on the "Metre" field.
func MetreContains(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldContains(FieldMetre, v))
}

// MetreHasPrefix applies the HasPrefix predicate on the "Metre" field.
func MetreHasPrefix(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldHasPrefix(FieldMetre, v))
}

// MetreHasSuffix applies the HasSuffix predicate on the "Metre" field.
func MetreHasSuffix(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldHasSuffix(FieldMetre, v))
}

// MetreEqualFold applies the EqualFold predicate on the "Metre" field.
func MetreEqualFold(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldEqualFold(FieldMetre, v))
}

// MetreContainsFold applies the ContainsFold predicate on the "Metre" field.
func MetreContainsFold(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldContainsFold(FieldMetre, v))
}

// MoldDateEQ applies the EQ predicate on the "MoldDate" field.
func MoldDateEQ(v time.Time) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldEQ(FieldMoldDate, v))
}

// MoldDateNEQ applies the NEQ predicate on the "MoldDate" field.
func MoldDateNEQ(v time.Time) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldNEQ(FieldMoldDate, v))
}

// MoldDateIn applies the In predicate on the "MoldDate" field.
func MoldDateIn(vs ...time.Time) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldIn(FieldMoldDate, vs...))
}

// MoldDateNotIn applies the NotIn predicate on the "MoldDate" field.
func MoldDateNotIn(vs ...time.Time) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldNotIn(FieldMoldDate, vs...))
}

// MoldDateGT applies the GT predicate on the "MoldDate" field.
func MoldDateGT(v time.Time) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldGT(FieldMoldDate, v))
}

// MoldDateGTE applies the GTE predicate on the "MoldDate" field.
func MoldDateGTE(v time.Time) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldGTE(FieldMoldDate, v))
}

// MoldDateLT applies the LT predicate on the "MoldDate" field.
func MoldDateLT(v time.Time) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldLT(FieldMoldDate, v))
}

// MoldDateLTE applies the LTE predicate on the "MoldDate" field.
func MoldDateLTE(v time.Time) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldLTE(FieldMoldDate, v))
}

// MoldDateIsNil applies the IsNil predicate on the "MoldDate" field.
func MoldDateIsNil() predicate.JobLayer {
	return predicate.JobLayer(sql.FieldIsNull(FieldMoldDate))
}

// MoldDateNotNil applies the NotNil predicate on the "MoldDate" field.
func MoldDateNotNil() predicate.JobLayer {
	return predicate.JobLayer(sql.FieldNotNull(FieldMoldDate))
}

// ConcreteDateEQ applies the EQ predicate on the "ConcreteDate" field.
func ConcreteDateEQ(v time.Time) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldEQ(FieldConcreteDate, v))
}

// ConcreteDateNEQ applies the NEQ predicate on the "ConcreteDate" field.
func ConcreteDateNEQ(v time.Time) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldNEQ(FieldConcreteDate, v))
}

// ConcreteDateIn applies the In predicate on the "ConcreteDate" field.
func ConcreteDateIn(vs ...time.Time) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldIn(FieldConcreteDate, vs...))
}

// ConcreteDateNotIn applies the NotIn predicate on the "ConcreteDate" field.
func ConcreteDateNotIn(vs ...time.Time) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldNotIn(FieldConcreteDate, vs...))
}

// ConcreteDateGT applies the GT predicate on the "ConcreteDate" field.
func ConcreteDateGT(v time.Time) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldGT(FieldConcreteDate, v))
}

// ConcreteDateGTE applies the GTE predicate on the "ConcreteDate" field.
func ConcreteDateGTE(v time.Time) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldGTE(FieldConcreteDate, v))
}

// ConcreteDateLT applies the LT predicate on the "ConcreteDate" field.
func ConcreteDateLT(v time.Time) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldLT(FieldConcreteDate, v))
}

// ConcreteDateLTE applies the LTE predicate on the "ConcreteDate" field.
func ConcreteDateLTE(v time.Time) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldLTE(FieldConcreteDate, v))
}

// ConcreteDateIsNil applies the IsNil predicate on the "ConcreteDate" field.
func ConcreteDateIsNil() predicate.JobLayer {
	return predicate.JobLayer(sql.FieldIsNull(FieldConcreteDate))
}

// ConcreteDateNotNil applies the NotNil predicate on the "ConcreteDate" field.
func ConcreteDateNotNil() predicate.JobLayer {
	return predicate.JobLayer(sql.FieldNotNull(FieldConcreteDate))
}

// SamplesEQ applies the EQ predicate on the "Samples" field.
func SamplesEQ(v int) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldEQ(FieldSamples, v))
}

// SamplesNEQ applies the NEQ predicate on the "Samples" field.
func SamplesNEQ(v int) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldNEQ(FieldSamples, v))
}

// SamplesIn applies the In predicate on the "Samples" field.
func SamplesIn(vs ...int) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldIn(FieldSamples, vs...))
}

// SamplesNotIn applies the NotIn predicate on the "Samples" field.
func SamplesNotIn(vs ...int) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldNotIn(FieldSamples, vs...))
}

// SamplesGT applies the GT predicate on the "Samples" field.
func SamplesGT(v int) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldGT(FieldSamples, v))
}

// SamplesGTE applies the GTE predicate on the "Samples" field.
func SamplesGTE(v int) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldGTE(FieldSamples, v))
}

// SamplesLT applies the LT predicate on the "Samples" field.
func SamplesLT(v int) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldLT(FieldSamples, v))
}

// SamplesLTE applies the LTE predicate on the "Samples" field.
func SamplesLTE(v int) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldLTE(FieldSamples, v))
}

// SamplesIsNil applies the IsNil predicate on the "Samples" field.
func SamplesIsNil() predicate.JobLayer {
	return predicate.JobLayer(sql.FieldIsNull(FieldSamples))
}

// SamplesNotNil applies the NotNil predicate on the "Samples" field.
func SamplesNotNil() predicate.JobLayer {
	return predicate.JobLayer(sql.FieldNotNull(FieldSamples))
}

// ConcreteClassEQ applies the EQ predicate on the "ConcreteClass" field.
func ConcreteClassEQ(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldEQ(FieldConcreteClass, v))
}

// ConcreteClassNEQ applies the NEQ predicate on the "ConcreteClass" field.
func ConcreteClassNEQ(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldNEQ(FieldConcreteClass, v))
}

// ConcreteClassIn applies the In predicate on the "ConcreteClass" field.
func ConcreteClassIn(vs ...string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldIn(FieldConcreteClass, vs...))
}

// ConcreteClassNotIn applies the NotIn predicate on the "ConcreteClass" field.
func ConcreteClassNotIn(vs ...string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldNotIn(FieldConcreteClass, vs...))
}

// ConcreteClassGT applies the GT predicate on the "ConcreteClass" field.
func ConcreteClassGT(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldGT(FieldConcreteClass, v))
}

// ConcreteClassGTE applies the GTE predicate on the "ConcreteClass" field.
func ConcreteClassGTE(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldGTE(FieldConcreteClass, v))
}

// ConcreteClassLT applies the LT predicate on the "ConcreteClass" field.
func ConcreteClassLT(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldLT(FieldConcreteClass, v))
}

// ConcreteClassLTE applies the LTE predicate on the "ConcreteClass" field.
func ConcreteClassLTE(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldLTE(FieldConcreteClass, v))
}

// ConcreteClassContains applies the Contains predicate on the "ConcreteClass" field.
func ConcreteClassContains(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldContains(FieldConcreteClass, v))
}

// ConcreteClassHasPrefix applies the HasPrefix predicate on the "ConcreteClass" field.
func ConcreteClassHasPrefix(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldHasPrefix(FieldConcreteClass, v))
}

// ConcreteClassHasSuffix applies the HasSuffix predicate on the "ConcreteClass" field.
func ConcreteClassHasSuffix(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldHasSuffix(FieldConcreteClass, v))
}

// ConcreteClassIsNil applies the IsNil predicate on the "ConcreteClass" field.
func ConcreteClassIsNil() predicate.JobLayer {
	return predicate.JobLayer(sql.FieldIsNull(FieldConcreteClass))
}

// ConcreteClassNotNil applies the NotNil predicate on the "ConcreteClass" field.
func ConcreteClassNotNil() predicate.JobLayer {
	return predicate.JobLayer(sql.FieldNotNull(FieldConcreteClass))
}

// ConcreteClassEqualFold applies the EqualFold predicate on the "ConcreteClass" field.
func ConcreteClassEqualFold(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldEqualFold(FieldConcreteClass, v))
}

// ConcreteClassContainsFold applies the ContainsFold predicate on the "ConcreteClass" field.
func ConcreteClassContainsFold(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldContainsFold(FieldConcreteClass, v))
}

// WeekResultEQ applies the EQ predicate on the "WeekResult" field.
func WeekResultEQ(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldEQ(FieldWeekResult, v))
}

// WeekResultNEQ applies the NEQ predicate on the "WeekResult" field.
func WeekResultNEQ(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldNEQ(FieldWeekResult, v))
}

// WeekResultIn applies the In predicate on the "WeekResult" field.
func WeekResultIn(vs ...string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldIn(FieldWeekResult, vs...))
}

// WeekResultNotIn applies the NotIn predicate on the "WeekResult" field.
func WeekResultNotIn(vs ...string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldNotIn(FieldWeekResult, vs...))
}

// WeekResultGT applies the GT predicate on the "WeekResult" field.
func WeekResultGT(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldGT(FieldWeekResult, v))
}

// WeekResultGTE applies the GTE predicate on the "WeekResult" field.
func WeekResultGTE(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldGTE(FieldWeekResult, v))
}

// WeekResultLT applies the LT predicate on the "WeekResult" field.
func WeekResultLT(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldLT(FieldWeekResult, v))
}

// WeekResultLTE applies the LTE predicate on the "WeekResult" field.
func WeekResultLTE(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldLTE(FieldWeekResult, v))
}

// WeekResultContains applies the Contains predicate on the "WeekResult" field.
func WeekResultContains(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldContains(FieldWeekResult, v))
}

// WeekResultHasPrefix applies the HasPrefix predicate on the "WeekResult" field.
func WeekResultHasPrefix(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldHasPrefix(FieldWeekResult, v))
}

// WeekResultHasSuffix applies the HasSuffix predicate on the "WeekResult" field.
func WeekResultHasSuffix(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldHasSuffix(FieldWeekResult, v))
}

// WeekResultIsNil applies the IsNil predicate on the "WeekResult" field.
func WeekResultIsNil() predicate.JobLayer {
	return predicate.JobLayer(sql.FieldIsNull(FieldWeekResult))
}

// WeekResultNotNil applies the NotNil predicate on the "WeekResult" field.
func WeekResultNotNil() predicate.JobLayer {
	return predicate.JobLayer(sql.FieldNotNull(FieldWeekResult))
}

// WeekResultEqualFold applies the EqualFold predicate on the "WeekResult" field.
func WeekResultEqualFold(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldEqualFold(FieldWeekResult, v))
}

// WeekResultContainsFold applies the ContainsFold predicate on the "WeekResult" field.
func WeekResultContainsFold(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldContainsFold(FieldWeekResult, v))
}

// MonthResultEQ applies the EQ predicate on the "MonthResult" field.
func MonthResultEQ(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldEQ(FieldMonthResult, v))
}

// MonthResultNEQ applies the NEQ predicate on the "MonthResult" field.
func MonthResultNEQ(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldNEQ(FieldMonthResult, v))
}

// MonthResultIn applies the In predicate on the "MonthResult" field.
func MonthResultIn(vs ...string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldIn(FieldMonthResult, vs...))
}

// MonthResultNotIn applies the NotIn predicate on the "MonthResult" field.
func MonthResultNotIn(vs ...string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldNotIn(FieldMonthResult, vs...))
}

// MonthResultGT applies the GT predicate on the "MonthResult" field.
func MonthResultGT(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldGT(FieldMonthResult, v))
}

// MonthResultGTE applies the GTE predicate on the "MonthResult" field.
func MonthResultGTE(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldGTE(FieldMonthResult, v))
}

// MonthResultLT applies the LT predicate on the "MonthResult" field.
func MonthResultLT(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldLT(FieldMonthResult, v))
}

// MonthResultLTE applies the LTE predicate on the "MonthResult" field.
func MonthResultLTE(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldLTE(FieldMonthResult, v))
}

// MonthResultContains applies the Contains predicate on the "MonthResult" field.
func MonthResultContains(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldContains(FieldMonthResult, v))
}

// MonthResultHasPrefix applies the HasPrefix predicate on the "MonthResult" field.
func MonthResultHasPrefix(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldHasPrefix(FieldMonthResult, v))
}

// MonthResultHasSuffix applies the HasSuffix predicate on the "MonthResult" field.
func MonthResultHasSuffix(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldHasSuffix(FieldMonthResult, v))
}

// MonthResultIsNil applies the IsNil predicate on the "MonthResult" field.
func MonthResultIsNil() predicate.JobLayer {
	return predicate.JobLayer(sql.FieldIsNull(FieldMonthResult))
}

// MonthResultNotNil applies the NotNil predicate on the "MonthResult" field.
func MonthResultNotNil() predicate.JobLayer {
	return predicate.JobLayer(sql.FieldNotNull(FieldMonthResult))
}

// MonthResultEqualFold applies the EqualFold predicate on the "MonthResult" field.
func MonthResultEqualFold(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldEqualFold(FieldMonthResult, v))
}

// MonthResultContainsFold applies the ContainsFold predicate on the "MonthResult" field.
func MonthResultContainsFold(v string) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldContainsFold(FieldMonthResult, v))
}

// CreatedAtEQ applies the EQ predicate on the "CreatedAt" field.
func CreatedAtEQ(v time.Time) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "CreatedAt" field.
func CreatedAtNEQ(v time.Time) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "CreatedAt" field.
func CreatedAtIn(vs ...time.Time) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "CreatedAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "CreatedAt" field.
func CreatedAtGT(v time.Time) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "CreatedAt" field.
func CreatedAtGTE(v time.Time) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "CreatedAt" field.
func CreatedAtLT(v time.Time) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "CreatedAt" field.
func CreatedAtLTE(v time.Time) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "UpdatedAt" field.
func UpdatedAtEQ(v time.Time) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "UpdatedAt" field.
func UpdatedAtNEQ(v time.Time) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "UpdatedAt" field.
func UpdatedAtIn(vs ...time.Time) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "UpdatedAt" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "UpdatedAt" field.
func UpdatedAtGT(v time.Time) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "UpdatedAt" field.
func UpdatedAtGTE(v time.Time) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "UpdatedAt" field.
func UpdatedAtLT(v time.Time) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "UpdatedAt" field.
func UpdatedAtLTE(v time.Time) predicate.JobLayer {
	return predicate.JobLayer(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasLayer applies the HasEdge predicate on the "layer" edge.
func HasLayer() predicate.JobLayer {
	return predicate.JobLayer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, LayerTable, LayerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLayerWith applies the HasEdge predicate on the "layer" edge with a given conditions (other predicates).
func HasLayerWith(preds ...predicate.JobDetail) predicate.JobLayer {
	return predicate.JobLayer(func(s *sql.Selector) {
		step := newLayerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.JobLayer) predicate.JobLayer {
	return predicate.JobLayer(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.JobLayer) predicate.JobLayer {
	return predicate.JobLayer(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.JobLayer) predicate.JobLayer {
	return predicate.JobLayer(sql.NotPredicates(p))
}
