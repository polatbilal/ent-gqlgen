// Code generated by ent, DO NOT EDIT.

package companytoken

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/polatbilal/gqlgen-ent/api-core/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldLTE(FieldID, id))
}

// YDKUsername applies equality check predicate on the "YDKUsername" field. It's identical to YDKUsernameEQ.
func YDKUsername(v string) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldEQ(FieldYDKUsername, v))
}

// YDKPassword applies equality check predicate on the "YDKPassword" field. It's identical to YDKPasswordEQ.
func YDKPassword(v string) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldEQ(FieldYDKPassword, v))
}

// Token applies equality check predicate on the "Token" field. It's identical to TokenEQ.
func Token(v string) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldEQ(FieldToken, v))
}

// DepartmentId applies equality check predicate on the "DepartmentId" field. It's identical to DepartmentIdEQ.
func DepartmentId(v int) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldEQ(FieldDepartmentId, v))
}

// CreatedAt applies equality check predicate on the "createdAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updatedAt" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldEQ(FieldUpdatedAt, v))
}

// YDKUsernameEQ applies the EQ predicate on the "YDKUsername" field.
func YDKUsernameEQ(v string) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldEQ(FieldYDKUsername, v))
}

// YDKUsernameNEQ applies the NEQ predicate on the "YDKUsername" field.
func YDKUsernameNEQ(v string) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldNEQ(FieldYDKUsername, v))
}

// YDKUsernameIn applies the In predicate on the "YDKUsername" field.
func YDKUsernameIn(vs ...string) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldIn(FieldYDKUsername, vs...))
}

// YDKUsernameNotIn applies the NotIn predicate on the "YDKUsername" field.
func YDKUsernameNotIn(vs ...string) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldNotIn(FieldYDKUsername, vs...))
}

// YDKUsernameGT applies the GT predicate on the "YDKUsername" field.
func YDKUsernameGT(v string) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldGT(FieldYDKUsername, v))
}

// YDKUsernameGTE applies the GTE predicate on the "YDKUsername" field.
func YDKUsernameGTE(v string) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldGTE(FieldYDKUsername, v))
}

// YDKUsernameLT applies the LT predicate on the "YDKUsername" field.
func YDKUsernameLT(v string) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldLT(FieldYDKUsername, v))
}

// YDKUsernameLTE applies the LTE predicate on the "YDKUsername" field.
func YDKUsernameLTE(v string) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldLTE(FieldYDKUsername, v))
}

// YDKUsernameContains applies the Contains predicate on the "YDKUsername" field.
func YDKUsernameContains(v string) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldContains(FieldYDKUsername, v))
}

// YDKUsernameHasPrefix applies the HasPrefix predicate on the "YDKUsername" field.
func YDKUsernameHasPrefix(v string) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldHasPrefix(FieldYDKUsername, v))
}

// YDKUsernameHasSuffix applies the HasSuffix predicate on the "YDKUsername" field.
func YDKUsernameHasSuffix(v string) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldHasSuffix(FieldYDKUsername, v))
}

// YDKUsernameIsNil applies the IsNil predicate on the "YDKUsername" field.
func YDKUsernameIsNil() predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldIsNull(FieldYDKUsername))
}

// YDKUsernameNotNil applies the NotNil predicate on the "YDKUsername" field.
func YDKUsernameNotNil() predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldNotNull(FieldYDKUsername))
}

// YDKUsernameEqualFold applies the EqualFold predicate on the "YDKUsername" field.
func YDKUsernameEqualFold(v string) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldEqualFold(FieldYDKUsername, v))
}

// YDKUsernameContainsFold applies the ContainsFold predicate on the "YDKUsername" field.
func YDKUsernameContainsFold(v string) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldContainsFold(FieldYDKUsername, v))
}

// YDKPasswordEQ applies the EQ predicate on the "YDKPassword" field.
func YDKPasswordEQ(v string) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldEQ(FieldYDKPassword, v))
}

// YDKPasswordNEQ applies the NEQ predicate on the "YDKPassword" field.
func YDKPasswordNEQ(v string) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldNEQ(FieldYDKPassword, v))
}

// YDKPasswordIn applies the In predicate on the "YDKPassword" field.
func YDKPasswordIn(vs ...string) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldIn(FieldYDKPassword, vs...))
}

// YDKPasswordNotIn applies the NotIn predicate on the "YDKPassword" field.
func YDKPasswordNotIn(vs ...string) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldNotIn(FieldYDKPassword, vs...))
}

// YDKPasswordGT applies the GT predicate on the "YDKPassword" field.
func YDKPasswordGT(v string) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldGT(FieldYDKPassword, v))
}

// YDKPasswordGTE applies the GTE predicate on the "YDKPassword" field.
func YDKPasswordGTE(v string) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldGTE(FieldYDKPassword, v))
}

// YDKPasswordLT applies the LT predicate on the "YDKPassword" field.
func YDKPasswordLT(v string) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldLT(FieldYDKPassword, v))
}

// YDKPasswordLTE applies the LTE predicate on the "YDKPassword" field.
func YDKPasswordLTE(v string) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldLTE(FieldYDKPassword, v))
}

// YDKPasswordContains applies the Contains predicate on the "YDKPassword" field.
func YDKPasswordContains(v string) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldContains(FieldYDKPassword, v))
}

// YDKPasswordHasPrefix applies the HasPrefix predicate on the "YDKPassword" field.
func YDKPasswordHasPrefix(v string) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldHasPrefix(FieldYDKPassword, v))
}

// YDKPasswordHasSuffix applies the HasSuffix predicate on the "YDKPassword" field.
func YDKPasswordHasSuffix(v string) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldHasSuffix(FieldYDKPassword, v))
}

// YDKPasswordIsNil applies the IsNil predicate on the "YDKPassword" field.
func YDKPasswordIsNil() predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldIsNull(FieldYDKPassword))
}

// YDKPasswordNotNil applies the NotNil predicate on the "YDKPassword" field.
func YDKPasswordNotNil() predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldNotNull(FieldYDKPassword))
}

// YDKPasswordEqualFold applies the EqualFold predicate on the "YDKPassword" field.
func YDKPasswordEqualFold(v string) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldEqualFold(FieldYDKPassword, v))
}

// YDKPasswordContainsFold applies the ContainsFold predicate on the "YDKPassword" field.
func YDKPasswordContainsFold(v string) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldContainsFold(FieldYDKPassword, v))
}

// TokenEQ applies the EQ predicate on the "Token" field.
func TokenEQ(v string) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldEQ(FieldToken, v))
}

// TokenNEQ applies the NEQ predicate on the "Token" field.
func TokenNEQ(v string) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldNEQ(FieldToken, v))
}

// TokenIn applies the In predicate on the "Token" field.
func TokenIn(vs ...string) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldIn(FieldToken, vs...))
}

// TokenNotIn applies the NotIn predicate on the "Token" field.
func TokenNotIn(vs ...string) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldNotIn(FieldToken, vs...))
}

// TokenGT applies the GT predicate on the "Token" field.
func TokenGT(v string) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldGT(FieldToken, v))
}

// TokenGTE applies the GTE predicate on the "Token" field.
func TokenGTE(v string) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldGTE(FieldToken, v))
}

// TokenLT applies the LT predicate on the "Token" field.
func TokenLT(v string) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldLT(FieldToken, v))
}

// TokenLTE applies the LTE predicate on the "Token" field.
func TokenLTE(v string) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldLTE(FieldToken, v))
}

// TokenContains applies the Contains predicate on the "Token" field.
func TokenContains(v string) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldContains(FieldToken, v))
}

// TokenHasPrefix applies the HasPrefix predicate on the "Token" field.
func TokenHasPrefix(v string) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldHasPrefix(FieldToken, v))
}

// TokenHasSuffix applies the HasSuffix predicate on the "Token" field.
func TokenHasSuffix(v string) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldHasSuffix(FieldToken, v))
}

// TokenIsNil applies the IsNil predicate on the "Token" field.
func TokenIsNil() predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldIsNull(FieldToken))
}

// TokenNotNil applies the NotNil predicate on the "Token" field.
func TokenNotNil() predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldNotNull(FieldToken))
}

// TokenEqualFold applies the EqualFold predicate on the "Token" field.
func TokenEqualFold(v string) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldEqualFold(FieldToken, v))
}

// TokenContainsFold applies the ContainsFold predicate on the "Token" field.
func TokenContainsFold(v string) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldContainsFold(FieldToken, v))
}

// DepartmentIdEQ applies the EQ predicate on the "DepartmentId" field.
func DepartmentIdEQ(v int) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldEQ(FieldDepartmentId, v))
}

// DepartmentIdNEQ applies the NEQ predicate on the "DepartmentId" field.
func DepartmentIdNEQ(v int) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldNEQ(FieldDepartmentId, v))
}

// DepartmentIdIn applies the In predicate on the "DepartmentId" field.
func DepartmentIdIn(vs ...int) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldIn(FieldDepartmentId, vs...))
}

// DepartmentIdNotIn applies the NotIn predicate on the "DepartmentId" field.
func DepartmentIdNotIn(vs ...int) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldNotIn(FieldDepartmentId, vs...))
}

// DepartmentIdGT applies the GT predicate on the "DepartmentId" field.
func DepartmentIdGT(v int) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldGT(FieldDepartmentId, v))
}

// DepartmentIdGTE applies the GTE predicate on the "DepartmentId" field.
func DepartmentIdGTE(v int) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldGTE(FieldDepartmentId, v))
}

// DepartmentIdLT applies the LT predicate on the "DepartmentId" field.
func DepartmentIdLT(v int) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldLT(FieldDepartmentId, v))
}

// DepartmentIdLTE applies the LTE predicate on the "DepartmentId" field.
func DepartmentIdLTE(v int) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldLTE(FieldDepartmentId, v))
}

// DepartmentIdIsNil applies the IsNil predicate on the "DepartmentId" field.
func DepartmentIdIsNil() predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldIsNull(FieldDepartmentId))
}

// DepartmentIdNotNil applies the NotNil predicate on the "DepartmentId" field.
func DepartmentIdNotNil() predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldNotNull(FieldDepartmentId))
}

// CreatedAtEQ applies the EQ predicate on the "createdAt" field.
func CreatedAtEQ(v time.Time) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "createdAt" field.
func CreatedAtNEQ(v time.Time) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "createdAt" field.
func CreatedAtIn(vs ...time.Time) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "createdAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "createdAt" field.
func CreatedAtGT(v time.Time) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "createdAt" field.
func CreatedAtGTE(v time.Time) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "createdAt" field.
func CreatedAtLT(v time.Time) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "createdAt" field.
func CreatedAtLTE(v time.Time) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updatedAt" field.
func UpdatedAtEQ(v time.Time) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updatedAt" field.
func UpdatedAtNEQ(v time.Time) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updatedAt" field.
func UpdatedAtIn(vs ...time.Time) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updatedAt" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updatedAt" field.
func UpdatedAtGT(v time.Time) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updatedAt" field.
func UpdatedAtGTE(v time.Time) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updatedAt" field.
func UpdatedAtLT(v time.Time) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updatedAt" field.
func UpdatedAtLTE(v time.Time) predicate.CompanyToken {
	return predicate.CompanyToken(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasCompany applies the HasEdge predicate on the "company" edge.
func HasCompany() predicate.CompanyToken {
	return predicate.CompanyToken(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CompanyTable, CompanyColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCompanyWith applies the HasEdge predicate on the "company" edge with a given conditions (other predicates).
func HasCompanyWith(preds ...predicate.CompanyDetail) predicate.CompanyToken {
	return predicate.CompanyToken(func(s *sql.Selector) {
		step := newCompanyStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CompanyToken) predicate.CompanyToken {
	return predicate.CompanyToken(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CompanyToken) predicate.CompanyToken {
	return predicate.CompanyToken(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CompanyToken) predicate.CompanyToken {
	return predicate.CompanyToken(sql.NotPredicates(p))
}
