// Code generated by ent, DO NOT EDIT.

package jobprogress

import (
	"gqlgen-ent/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldLTE(FieldID, id))
}

// One applies equality check predicate on the "One" field. It's identical to OneEQ.
func One(v int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldEQ(FieldOne, v))
}

// Two applies equality check predicate on the "Two" field. It's identical to TwoEQ.
func Two(v int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldEQ(FieldTwo, v))
}

// Three applies equality check predicate on the "Three" field. It's identical to ThreeEQ.
func Three(v int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldEQ(FieldThree, v))
}

// Four applies equality check predicate on the "Four" field. It's identical to FourEQ.
func Four(v int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldEQ(FieldFour, v))
}

// Five applies equality check predicate on the "Five" field. It's identical to FiveEQ.
func Five(v int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldEQ(FieldFive, v))
}

// Six applies equality check predicate on the "Six" field. It's identical to SixEQ.
func Six(v int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldEQ(FieldSix, v))
}

// CreatedAt applies equality check predicate on the "CreatedAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "UpdatedAt" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldEQ(FieldUpdatedAt, v))
}

// OneEQ applies the EQ predicate on the "One" field.
func OneEQ(v int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldEQ(FieldOne, v))
}

// OneNEQ applies the NEQ predicate on the "One" field.
func OneNEQ(v int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldNEQ(FieldOne, v))
}

// OneIn applies the In predicate on the "One" field.
func OneIn(vs ...int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldIn(FieldOne, vs...))
}

// OneNotIn applies the NotIn predicate on the "One" field.
func OneNotIn(vs ...int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldNotIn(FieldOne, vs...))
}

// OneGT applies the GT predicate on the "One" field.
func OneGT(v int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldGT(FieldOne, v))
}

// OneGTE applies the GTE predicate on the "One" field.
func OneGTE(v int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldGTE(FieldOne, v))
}

// OneLT applies the LT predicate on the "One" field.
func OneLT(v int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldLT(FieldOne, v))
}

// OneLTE applies the LTE predicate on the "One" field.
func OneLTE(v int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldLTE(FieldOne, v))
}

// OneIsNil applies the IsNil predicate on the "One" field.
func OneIsNil() predicate.JobProgress {
	return predicate.JobProgress(sql.FieldIsNull(FieldOne))
}

// OneNotNil applies the NotNil predicate on the "One" field.
func OneNotNil() predicate.JobProgress {
	return predicate.JobProgress(sql.FieldNotNull(FieldOne))
}

// TwoEQ applies the EQ predicate on the "Two" field.
func TwoEQ(v int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldEQ(FieldTwo, v))
}

// TwoNEQ applies the NEQ predicate on the "Two" field.
func TwoNEQ(v int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldNEQ(FieldTwo, v))
}

// TwoIn applies the In predicate on the "Two" field.
func TwoIn(vs ...int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldIn(FieldTwo, vs...))
}

// TwoNotIn applies the NotIn predicate on the "Two" field.
func TwoNotIn(vs ...int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldNotIn(FieldTwo, vs...))
}

// TwoGT applies the GT predicate on the "Two" field.
func TwoGT(v int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldGT(FieldTwo, v))
}

// TwoGTE applies the GTE predicate on the "Two" field.
func TwoGTE(v int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldGTE(FieldTwo, v))
}

// TwoLT applies the LT predicate on the "Two" field.
func TwoLT(v int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldLT(FieldTwo, v))
}

// TwoLTE applies the LTE predicate on the "Two" field.
func TwoLTE(v int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldLTE(FieldTwo, v))
}

// TwoIsNil applies the IsNil predicate on the "Two" field.
func TwoIsNil() predicate.JobProgress {
	return predicate.JobProgress(sql.FieldIsNull(FieldTwo))
}

// TwoNotNil applies the NotNil predicate on the "Two" field.
func TwoNotNil() predicate.JobProgress {
	return predicate.JobProgress(sql.FieldNotNull(FieldTwo))
}

// ThreeEQ applies the EQ predicate on the "Three" field.
func ThreeEQ(v int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldEQ(FieldThree, v))
}

// ThreeNEQ applies the NEQ predicate on the "Three" field.
func ThreeNEQ(v int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldNEQ(FieldThree, v))
}

// ThreeIn applies the In predicate on the "Three" field.
func ThreeIn(vs ...int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldIn(FieldThree, vs...))
}

// ThreeNotIn applies the NotIn predicate on the "Three" field.
func ThreeNotIn(vs ...int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldNotIn(FieldThree, vs...))
}

// ThreeGT applies the GT predicate on the "Three" field.
func ThreeGT(v int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldGT(FieldThree, v))
}

// ThreeGTE applies the GTE predicate on the "Three" field.
func ThreeGTE(v int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldGTE(FieldThree, v))
}

// ThreeLT applies the LT predicate on the "Three" field.
func ThreeLT(v int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldLT(FieldThree, v))
}

// ThreeLTE applies the LTE predicate on the "Three" field.
func ThreeLTE(v int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldLTE(FieldThree, v))
}

// ThreeIsNil applies the IsNil predicate on the "Three" field.
func ThreeIsNil() predicate.JobProgress {
	return predicate.JobProgress(sql.FieldIsNull(FieldThree))
}

// ThreeNotNil applies the NotNil predicate on the "Three" field.
func ThreeNotNil() predicate.JobProgress {
	return predicate.JobProgress(sql.FieldNotNull(FieldThree))
}

// FourEQ applies the EQ predicate on the "Four" field.
func FourEQ(v int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldEQ(FieldFour, v))
}

// FourNEQ applies the NEQ predicate on the "Four" field.
func FourNEQ(v int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldNEQ(FieldFour, v))
}

// FourIn applies the In predicate on the "Four" field.
func FourIn(vs ...int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldIn(FieldFour, vs...))
}

// FourNotIn applies the NotIn predicate on the "Four" field.
func FourNotIn(vs ...int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldNotIn(FieldFour, vs...))
}

// FourGT applies the GT predicate on the "Four" field.
func FourGT(v int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldGT(FieldFour, v))
}

// FourGTE applies the GTE predicate on the "Four" field.
func FourGTE(v int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldGTE(FieldFour, v))
}

// FourLT applies the LT predicate on the "Four" field.
func FourLT(v int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldLT(FieldFour, v))
}

// FourLTE applies the LTE predicate on the "Four" field.
func FourLTE(v int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldLTE(FieldFour, v))
}

// FourIsNil applies the IsNil predicate on the "Four" field.
func FourIsNil() predicate.JobProgress {
	return predicate.JobProgress(sql.FieldIsNull(FieldFour))
}

// FourNotNil applies the NotNil predicate on the "Four" field.
func FourNotNil() predicate.JobProgress {
	return predicate.JobProgress(sql.FieldNotNull(FieldFour))
}

// FiveEQ applies the EQ predicate on the "Five" field.
func FiveEQ(v int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldEQ(FieldFive, v))
}

// FiveNEQ applies the NEQ predicate on the "Five" field.
func FiveNEQ(v int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldNEQ(FieldFive, v))
}

// FiveIn applies the In predicate on the "Five" field.
func FiveIn(vs ...int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldIn(FieldFive, vs...))
}

// FiveNotIn applies the NotIn predicate on the "Five" field.
func FiveNotIn(vs ...int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldNotIn(FieldFive, vs...))
}

// FiveGT applies the GT predicate on the "Five" field.
func FiveGT(v int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldGT(FieldFive, v))
}

// FiveGTE applies the GTE predicate on the "Five" field.
func FiveGTE(v int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldGTE(FieldFive, v))
}

// FiveLT applies the LT predicate on the "Five" field.
func FiveLT(v int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldLT(FieldFive, v))
}

// FiveLTE applies the LTE predicate on the "Five" field.
func FiveLTE(v int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldLTE(FieldFive, v))
}

// FiveIsNil applies the IsNil predicate on the "Five" field.
func FiveIsNil() predicate.JobProgress {
	return predicate.JobProgress(sql.FieldIsNull(FieldFive))
}

// FiveNotNil applies the NotNil predicate on the "Five" field.
func FiveNotNil() predicate.JobProgress {
	return predicate.JobProgress(sql.FieldNotNull(FieldFive))
}

// SixEQ applies the EQ predicate on the "Six" field.
func SixEQ(v int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldEQ(FieldSix, v))
}

// SixNEQ applies the NEQ predicate on the "Six" field.
func SixNEQ(v int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldNEQ(FieldSix, v))
}

// SixIn applies the In predicate on the "Six" field.
func SixIn(vs ...int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldIn(FieldSix, vs...))
}

// SixNotIn applies the NotIn predicate on the "Six" field.
func SixNotIn(vs ...int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldNotIn(FieldSix, vs...))
}

// SixGT applies the GT predicate on the "Six" field.
func SixGT(v int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldGT(FieldSix, v))
}

// SixGTE applies the GTE predicate on the "Six" field.
func SixGTE(v int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldGTE(FieldSix, v))
}

// SixLT applies the LT predicate on the "Six" field.
func SixLT(v int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldLT(FieldSix, v))
}

// SixLTE applies the LTE predicate on the "Six" field.
func SixLTE(v int) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldLTE(FieldSix, v))
}

// SixIsNil applies the IsNil predicate on the "Six" field.
func SixIsNil() predicate.JobProgress {
	return predicate.JobProgress(sql.FieldIsNull(FieldSix))
}

// SixNotNil applies the NotNil predicate on the "Six" field.
func SixNotNil() predicate.JobProgress {
	return predicate.JobProgress(sql.FieldNotNull(FieldSix))
}

// CreatedAtEQ applies the EQ predicate on the "CreatedAt" field.
func CreatedAtEQ(v time.Time) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "CreatedAt" field.
func CreatedAtNEQ(v time.Time) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "CreatedAt" field.
func CreatedAtIn(vs ...time.Time) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "CreatedAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "CreatedAt" field.
func CreatedAtGT(v time.Time) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "CreatedAt" field.
func CreatedAtGTE(v time.Time) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "CreatedAt" field.
func CreatedAtLT(v time.Time) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "CreatedAt" field.
func CreatedAtLTE(v time.Time) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "UpdatedAt" field.
func UpdatedAtEQ(v time.Time) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "UpdatedAt" field.
func UpdatedAtNEQ(v time.Time) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "UpdatedAt" field.
func UpdatedAtIn(vs ...time.Time) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "UpdatedAt" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "UpdatedAt" field.
func UpdatedAtGT(v time.Time) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "UpdatedAt" field.
func UpdatedAtGTE(v time.Time) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "UpdatedAt" field.
func UpdatedAtLT(v time.Time) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "UpdatedAt" field.
func UpdatedAtLTE(v time.Time) predicate.JobProgress {
	return predicate.JobProgress(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasProgress applies the HasEdge predicate on the "progress" edge.
func HasProgress() predicate.JobProgress {
	return predicate.JobProgress(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProgressTable, ProgressColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProgressWith applies the HasEdge predicate on the "progress" edge with a given conditions (other predicates).
func HasProgressWith(preds ...predicate.JobDetail) predicate.JobProgress {
	return predicate.JobProgress(func(s *sql.Selector) {
		step := newProgressStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.JobProgress) predicate.JobProgress {
	return predicate.JobProgress(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.JobProgress) predicate.JobProgress {
	return predicate.JobProgress(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.JobProgress) predicate.JobProgress {
	return predicate.JobProgress(sql.NotPredicates(p))
}
