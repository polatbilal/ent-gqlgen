// Code generated by ent, DO NOT EDIT.

package jobauthor

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/polatbilal/gqlgen-ent/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldLTE(FieldID, id))
}

// Architect applies equality check predicate on the "Architect" field. It's identical to ArchitectEQ.
func Architect(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldEQ(FieldArchitect, v))
}

// Static applies equality check predicate on the "Static" field. It's identical to StaticEQ.
func Static(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldEQ(FieldStatic, v))
}

// Mechanic applies equality check predicate on the "Mechanic" field. It's identical to MechanicEQ.
func Mechanic(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldEQ(FieldMechanic, v))
}

// Electric applies equality check predicate on the "Electric" field. It's identical to ElectricEQ.
func Electric(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldEQ(FieldElectric, v))
}

// Floor applies equality check predicate on the "Floor" field. It's identical to FloorEQ.
func Floor(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldEQ(FieldFloor, v))
}

// CreatedAt applies equality check predicate on the "CreatedAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "UpdatedAt" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldEQ(FieldUpdatedAt, v))
}

// ArchitectEQ applies the EQ predicate on the "Architect" field.
func ArchitectEQ(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldEQ(FieldArchitect, v))
}

// ArchitectNEQ applies the NEQ predicate on the "Architect" field.
func ArchitectNEQ(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldNEQ(FieldArchitect, v))
}

// ArchitectIn applies the In predicate on the "Architect" field.
func ArchitectIn(vs ...string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldIn(FieldArchitect, vs...))
}

// ArchitectNotIn applies the NotIn predicate on the "Architect" field.
func ArchitectNotIn(vs ...string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldNotIn(FieldArchitect, vs...))
}

// ArchitectGT applies the GT predicate on the "Architect" field.
func ArchitectGT(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldGT(FieldArchitect, v))
}

// ArchitectGTE applies the GTE predicate on the "Architect" field.
func ArchitectGTE(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldGTE(FieldArchitect, v))
}

// ArchitectLT applies the LT predicate on the "Architect" field.
func ArchitectLT(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldLT(FieldArchitect, v))
}

// ArchitectLTE applies the LTE predicate on the "Architect" field.
func ArchitectLTE(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldLTE(FieldArchitect, v))
}

// ArchitectContains applies the Contains predicate on the "Architect" field.
func ArchitectContains(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldContains(FieldArchitect, v))
}

// ArchitectHasPrefix applies the HasPrefix predicate on the "Architect" field.
func ArchitectHasPrefix(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldHasPrefix(FieldArchitect, v))
}

// ArchitectHasSuffix applies the HasSuffix predicate on the "Architect" field.
func ArchitectHasSuffix(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldHasSuffix(FieldArchitect, v))
}

// ArchitectIsNil applies the IsNil predicate on the "Architect" field.
func ArchitectIsNil() predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldIsNull(FieldArchitect))
}

// ArchitectNotNil applies the NotNil predicate on the "Architect" field.
func ArchitectNotNil() predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldNotNull(FieldArchitect))
}

// ArchitectEqualFold applies the EqualFold predicate on the "Architect" field.
func ArchitectEqualFold(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldEqualFold(FieldArchitect, v))
}

// ArchitectContainsFold applies the ContainsFold predicate on the "Architect" field.
func ArchitectContainsFold(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldContainsFold(FieldArchitect, v))
}

// StaticEQ applies the EQ predicate on the "Static" field.
func StaticEQ(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldEQ(FieldStatic, v))
}

// StaticNEQ applies the NEQ predicate on the "Static" field.
func StaticNEQ(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldNEQ(FieldStatic, v))
}

// StaticIn applies the In predicate on the "Static" field.
func StaticIn(vs ...string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldIn(FieldStatic, vs...))
}

// StaticNotIn applies the NotIn predicate on the "Static" field.
func StaticNotIn(vs ...string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldNotIn(FieldStatic, vs...))
}

// StaticGT applies the GT predicate on the "Static" field.
func StaticGT(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldGT(FieldStatic, v))
}

// StaticGTE applies the GTE predicate on the "Static" field.
func StaticGTE(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldGTE(FieldStatic, v))
}

// StaticLT applies the LT predicate on the "Static" field.
func StaticLT(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldLT(FieldStatic, v))
}

// StaticLTE applies the LTE predicate on the "Static" field.
func StaticLTE(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldLTE(FieldStatic, v))
}

// StaticContains applies the Contains predicate on the "Static" field.
func StaticContains(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldContains(FieldStatic, v))
}

// StaticHasPrefix applies the HasPrefix predicate on the "Static" field.
func StaticHasPrefix(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldHasPrefix(FieldStatic, v))
}

// StaticHasSuffix applies the HasSuffix predicate on the "Static" field.
func StaticHasSuffix(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldHasSuffix(FieldStatic, v))
}

// StaticIsNil applies the IsNil predicate on the "Static" field.
func StaticIsNil() predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldIsNull(FieldStatic))
}

// StaticNotNil applies the NotNil predicate on the "Static" field.
func StaticNotNil() predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldNotNull(FieldStatic))
}

// StaticEqualFold applies the EqualFold predicate on the "Static" field.
func StaticEqualFold(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldEqualFold(FieldStatic, v))
}

// StaticContainsFold applies the ContainsFold predicate on the "Static" field.
func StaticContainsFold(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldContainsFold(FieldStatic, v))
}

// MechanicEQ applies the EQ predicate on the "Mechanic" field.
func MechanicEQ(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldEQ(FieldMechanic, v))
}

// MechanicNEQ applies the NEQ predicate on the "Mechanic" field.
func MechanicNEQ(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldNEQ(FieldMechanic, v))
}

// MechanicIn applies the In predicate on the "Mechanic" field.
func MechanicIn(vs ...string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldIn(FieldMechanic, vs...))
}

// MechanicNotIn applies the NotIn predicate on the "Mechanic" field.
func MechanicNotIn(vs ...string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldNotIn(FieldMechanic, vs...))
}

// MechanicGT applies the GT predicate on the "Mechanic" field.
func MechanicGT(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldGT(FieldMechanic, v))
}

// MechanicGTE applies the GTE predicate on the "Mechanic" field.
func MechanicGTE(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldGTE(FieldMechanic, v))
}

// MechanicLT applies the LT predicate on the "Mechanic" field.
func MechanicLT(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldLT(FieldMechanic, v))
}

// MechanicLTE applies the LTE predicate on the "Mechanic" field.
func MechanicLTE(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldLTE(FieldMechanic, v))
}

// MechanicContains applies the Contains predicate on the "Mechanic" field.
func MechanicContains(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldContains(FieldMechanic, v))
}

// MechanicHasPrefix applies the HasPrefix predicate on the "Mechanic" field.
func MechanicHasPrefix(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldHasPrefix(FieldMechanic, v))
}

// MechanicHasSuffix applies the HasSuffix predicate on the "Mechanic" field.
func MechanicHasSuffix(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldHasSuffix(FieldMechanic, v))
}

// MechanicIsNil applies the IsNil predicate on the "Mechanic" field.
func MechanicIsNil() predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldIsNull(FieldMechanic))
}

// MechanicNotNil applies the NotNil predicate on the "Mechanic" field.
func MechanicNotNil() predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldNotNull(FieldMechanic))
}

// MechanicEqualFold applies the EqualFold predicate on the "Mechanic" field.
func MechanicEqualFold(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldEqualFold(FieldMechanic, v))
}

// MechanicContainsFold applies the ContainsFold predicate on the "Mechanic" field.
func MechanicContainsFold(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldContainsFold(FieldMechanic, v))
}

// ElectricEQ applies the EQ predicate on the "Electric" field.
func ElectricEQ(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldEQ(FieldElectric, v))
}

// ElectricNEQ applies the NEQ predicate on the "Electric" field.
func ElectricNEQ(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldNEQ(FieldElectric, v))
}

// ElectricIn applies the In predicate on the "Electric" field.
func ElectricIn(vs ...string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldIn(FieldElectric, vs...))
}

// ElectricNotIn applies the NotIn predicate on the "Electric" field.
func ElectricNotIn(vs ...string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldNotIn(FieldElectric, vs...))
}

// ElectricGT applies the GT predicate on the "Electric" field.
func ElectricGT(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldGT(FieldElectric, v))
}

// ElectricGTE applies the GTE predicate on the "Electric" field.
func ElectricGTE(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldGTE(FieldElectric, v))
}

// ElectricLT applies the LT predicate on the "Electric" field.
func ElectricLT(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldLT(FieldElectric, v))
}

// ElectricLTE applies the LTE predicate on the "Electric" field.
func ElectricLTE(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldLTE(FieldElectric, v))
}

// ElectricContains applies the Contains predicate on the "Electric" field.
func ElectricContains(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldContains(FieldElectric, v))
}

// ElectricHasPrefix applies the HasPrefix predicate on the "Electric" field.
func ElectricHasPrefix(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldHasPrefix(FieldElectric, v))
}

// ElectricHasSuffix applies the HasSuffix predicate on the "Electric" field.
func ElectricHasSuffix(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldHasSuffix(FieldElectric, v))
}

// ElectricIsNil applies the IsNil predicate on the "Electric" field.
func ElectricIsNil() predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldIsNull(FieldElectric))
}

// ElectricNotNil applies the NotNil predicate on the "Electric" field.
func ElectricNotNil() predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldNotNull(FieldElectric))
}

// ElectricEqualFold applies the EqualFold predicate on the "Electric" field.
func ElectricEqualFold(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldEqualFold(FieldElectric, v))
}

// ElectricContainsFold applies the ContainsFold predicate on the "Electric" field.
func ElectricContainsFold(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldContainsFold(FieldElectric, v))
}

// FloorEQ applies the EQ predicate on the "Floor" field.
func FloorEQ(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldEQ(FieldFloor, v))
}

// FloorNEQ applies the NEQ predicate on the "Floor" field.
func FloorNEQ(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldNEQ(FieldFloor, v))
}

// FloorIn applies the In predicate on the "Floor" field.
func FloorIn(vs ...string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldIn(FieldFloor, vs...))
}

// FloorNotIn applies the NotIn predicate on the "Floor" field.
func FloorNotIn(vs ...string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldNotIn(FieldFloor, vs...))
}

// FloorGT applies the GT predicate on the "Floor" field.
func FloorGT(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldGT(FieldFloor, v))
}

// FloorGTE applies the GTE predicate on the "Floor" field.
func FloorGTE(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldGTE(FieldFloor, v))
}

// FloorLT applies the LT predicate on the "Floor" field.
func FloorLT(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldLT(FieldFloor, v))
}

// FloorLTE applies the LTE predicate on the "Floor" field.
func FloorLTE(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldLTE(FieldFloor, v))
}

// FloorContains applies the Contains predicate on the "Floor" field.
func FloorContains(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldContains(FieldFloor, v))
}

// FloorHasPrefix applies the HasPrefix predicate on the "Floor" field.
func FloorHasPrefix(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldHasPrefix(FieldFloor, v))
}

// FloorHasSuffix applies the HasSuffix predicate on the "Floor" field.
func FloorHasSuffix(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldHasSuffix(FieldFloor, v))
}

// FloorIsNil applies the IsNil predicate on the "Floor" field.
func FloorIsNil() predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldIsNull(FieldFloor))
}

// FloorNotNil applies the NotNil predicate on the "Floor" field.
func FloorNotNil() predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldNotNull(FieldFloor))
}

// FloorEqualFold applies the EqualFold predicate on the "Floor" field.
func FloorEqualFold(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldEqualFold(FieldFloor, v))
}

// FloorContainsFold applies the ContainsFold predicate on the "Floor" field.
func FloorContainsFold(v string) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldContainsFold(FieldFloor, v))
}

// CreatedAtEQ applies the EQ predicate on the "CreatedAt" field.
func CreatedAtEQ(v time.Time) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "CreatedAt" field.
func CreatedAtNEQ(v time.Time) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "CreatedAt" field.
func CreatedAtIn(vs ...time.Time) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "CreatedAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "CreatedAt" field.
func CreatedAtGT(v time.Time) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "CreatedAt" field.
func CreatedAtGTE(v time.Time) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "CreatedAt" field.
func CreatedAtLT(v time.Time) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "CreatedAt" field.
func CreatedAtLTE(v time.Time) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "UpdatedAt" field.
func UpdatedAtEQ(v time.Time) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "UpdatedAt" field.
func UpdatedAtNEQ(v time.Time) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "UpdatedAt" field.
func UpdatedAtIn(vs ...time.Time) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "UpdatedAt" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "UpdatedAt" field.
func UpdatedAtGT(v time.Time) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "UpdatedAt" field.
func UpdatedAtGTE(v time.Time) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "UpdatedAt" field.
func UpdatedAtLT(v time.Time) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "UpdatedAt" field.
func UpdatedAtLTE(v time.Time) predicate.JobAuthor {
	return predicate.JobAuthor(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasAuthors applies the HasEdge predicate on the "authors" edge.
func HasAuthors() predicate.JobAuthor {
	return predicate.JobAuthor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AuthorsTable, AuthorsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAuthorsWith applies the HasEdge predicate on the "authors" edge with a given conditions (other predicates).
func HasAuthorsWith(preds ...predicate.JobDetail) predicate.JobAuthor {
	return predicate.JobAuthor(func(s *sql.Selector) {
		step := newAuthorsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.JobAuthor) predicate.JobAuthor {
	return predicate.JobAuthor(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.JobAuthor) predicate.JobAuthor {
	return predicate.JobAuthor(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.JobAuthor) predicate.JobAuthor {
	return predicate.JobAuthor(sql.NotPredicates(p))
}
