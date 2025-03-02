// Code generated by ent, DO NOT EDIT.

package jobpayments

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/polatbilal/gqlgen-ent/api-core/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldLTE(FieldID, id))
}

// YibfNo applies equality check predicate on the "yibfNo" field. It's identical to YibfNoEQ.
func YibfNo(v int) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldEQ(FieldYibfNo, v))
}

// PaymentNo applies equality check predicate on the "PaymentNo" field. It's identical to PaymentNoEQ.
func PaymentNo(v int) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldEQ(FieldPaymentNo, v))
}

// PaymentDate applies equality check predicate on the "PaymentDate" field. It's identical to PaymentDateEQ.
func PaymentDate(v time.Time) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldEQ(FieldPaymentDate, v))
}

// PaymentType applies equality check predicate on the "PaymentType" field. It's identical to PaymentTypeEQ.
func PaymentType(v string) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldEQ(FieldPaymentType, v))
}

// State applies equality check predicate on the "State" field. It's identical to StateEQ.
func State(v string) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldEQ(FieldState, v))
}

// TotalPayment applies equality check predicate on the "TotalPayment" field. It's identical to TotalPaymentEQ.
func TotalPayment(v float64) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldEQ(FieldTotalPayment, v))
}

// LevelRequest applies equality check predicate on the "LevelRequest" field. It's identical to LevelRequestEQ.
func LevelRequest(v float64) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldEQ(FieldLevelRequest, v))
}

// LevelApprove applies equality check predicate on the "LevelApprove" field. It's identical to LevelApproveEQ.
func LevelApprove(v float64) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldEQ(FieldLevelApprove, v))
}

// Amount applies equality check predicate on the "Amount" field. It's identical to AmountEQ.
func Amount(v float64) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldEQ(FieldAmount, v))
}

// CreatedAt applies equality check predicate on the "CreatedAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "UpdatedAt" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldEQ(FieldUpdatedAt, v))
}

// YibfNoEQ applies the EQ predicate on the "yibfNo" field.
func YibfNoEQ(v int) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldEQ(FieldYibfNo, v))
}

// YibfNoNEQ applies the NEQ predicate on the "yibfNo" field.
func YibfNoNEQ(v int) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldNEQ(FieldYibfNo, v))
}

// YibfNoIn applies the In predicate on the "yibfNo" field.
func YibfNoIn(vs ...int) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldIn(FieldYibfNo, vs...))
}

// YibfNoNotIn applies the NotIn predicate on the "yibfNo" field.
func YibfNoNotIn(vs ...int) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldNotIn(FieldYibfNo, vs...))
}

// YibfNoGT applies the GT predicate on the "yibfNo" field.
func YibfNoGT(v int) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldGT(FieldYibfNo, v))
}

// YibfNoGTE applies the GTE predicate on the "yibfNo" field.
func YibfNoGTE(v int) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldGTE(FieldYibfNo, v))
}

// YibfNoLT applies the LT predicate on the "yibfNo" field.
func YibfNoLT(v int) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldLT(FieldYibfNo, v))
}

// YibfNoLTE applies the LTE predicate on the "yibfNo" field.
func YibfNoLTE(v int) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldLTE(FieldYibfNo, v))
}

// PaymentNoEQ applies the EQ predicate on the "PaymentNo" field.
func PaymentNoEQ(v int) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldEQ(FieldPaymentNo, v))
}

// PaymentNoNEQ applies the NEQ predicate on the "PaymentNo" field.
func PaymentNoNEQ(v int) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldNEQ(FieldPaymentNo, v))
}

// PaymentNoIn applies the In predicate on the "PaymentNo" field.
func PaymentNoIn(vs ...int) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldIn(FieldPaymentNo, vs...))
}

// PaymentNoNotIn applies the NotIn predicate on the "PaymentNo" field.
func PaymentNoNotIn(vs ...int) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldNotIn(FieldPaymentNo, vs...))
}

// PaymentNoGT applies the GT predicate on the "PaymentNo" field.
func PaymentNoGT(v int) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldGT(FieldPaymentNo, v))
}

// PaymentNoGTE applies the GTE predicate on the "PaymentNo" field.
func PaymentNoGTE(v int) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldGTE(FieldPaymentNo, v))
}

// PaymentNoLT applies the LT predicate on the "PaymentNo" field.
func PaymentNoLT(v int) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldLT(FieldPaymentNo, v))
}

// PaymentNoLTE applies the LTE predicate on the "PaymentNo" field.
func PaymentNoLTE(v int) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldLTE(FieldPaymentNo, v))
}

// PaymentNoIsNil applies the IsNil predicate on the "PaymentNo" field.
func PaymentNoIsNil() predicate.JobPayments {
	return predicate.JobPayments(sql.FieldIsNull(FieldPaymentNo))
}

// PaymentNoNotNil applies the NotNil predicate on the "PaymentNo" field.
func PaymentNoNotNil() predicate.JobPayments {
	return predicate.JobPayments(sql.FieldNotNull(FieldPaymentNo))
}

// PaymentDateEQ applies the EQ predicate on the "PaymentDate" field.
func PaymentDateEQ(v time.Time) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldEQ(FieldPaymentDate, v))
}

// PaymentDateNEQ applies the NEQ predicate on the "PaymentDate" field.
func PaymentDateNEQ(v time.Time) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldNEQ(FieldPaymentDate, v))
}

// PaymentDateIn applies the In predicate on the "PaymentDate" field.
func PaymentDateIn(vs ...time.Time) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldIn(FieldPaymentDate, vs...))
}

// PaymentDateNotIn applies the NotIn predicate on the "PaymentDate" field.
func PaymentDateNotIn(vs ...time.Time) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldNotIn(FieldPaymentDate, vs...))
}

// PaymentDateGT applies the GT predicate on the "PaymentDate" field.
func PaymentDateGT(v time.Time) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldGT(FieldPaymentDate, v))
}

// PaymentDateGTE applies the GTE predicate on the "PaymentDate" field.
func PaymentDateGTE(v time.Time) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldGTE(FieldPaymentDate, v))
}

// PaymentDateLT applies the LT predicate on the "PaymentDate" field.
func PaymentDateLT(v time.Time) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldLT(FieldPaymentDate, v))
}

// PaymentDateLTE applies the LTE predicate on the "PaymentDate" field.
func PaymentDateLTE(v time.Time) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldLTE(FieldPaymentDate, v))
}

// PaymentTypeEQ applies the EQ predicate on the "PaymentType" field.
func PaymentTypeEQ(v string) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldEQ(FieldPaymentType, v))
}

// PaymentTypeNEQ applies the NEQ predicate on the "PaymentType" field.
func PaymentTypeNEQ(v string) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldNEQ(FieldPaymentType, v))
}

// PaymentTypeIn applies the In predicate on the "PaymentType" field.
func PaymentTypeIn(vs ...string) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldIn(FieldPaymentType, vs...))
}

// PaymentTypeNotIn applies the NotIn predicate on the "PaymentType" field.
func PaymentTypeNotIn(vs ...string) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldNotIn(FieldPaymentType, vs...))
}

// PaymentTypeGT applies the GT predicate on the "PaymentType" field.
func PaymentTypeGT(v string) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldGT(FieldPaymentType, v))
}

// PaymentTypeGTE applies the GTE predicate on the "PaymentType" field.
func PaymentTypeGTE(v string) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldGTE(FieldPaymentType, v))
}

// PaymentTypeLT applies the LT predicate on the "PaymentType" field.
func PaymentTypeLT(v string) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldLT(FieldPaymentType, v))
}

// PaymentTypeLTE applies the LTE predicate on the "PaymentType" field.
func PaymentTypeLTE(v string) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldLTE(FieldPaymentType, v))
}

// PaymentTypeContains applies the Contains predicate on the "PaymentType" field.
func PaymentTypeContains(v string) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldContains(FieldPaymentType, v))
}

// PaymentTypeHasPrefix applies the HasPrefix predicate on the "PaymentType" field.
func PaymentTypeHasPrefix(v string) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldHasPrefix(FieldPaymentType, v))
}

// PaymentTypeHasSuffix applies the HasSuffix predicate on the "PaymentType" field.
func PaymentTypeHasSuffix(v string) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldHasSuffix(FieldPaymentType, v))
}

// PaymentTypeIsNil applies the IsNil predicate on the "PaymentType" field.
func PaymentTypeIsNil() predicate.JobPayments {
	return predicate.JobPayments(sql.FieldIsNull(FieldPaymentType))
}

// PaymentTypeNotNil applies the NotNil predicate on the "PaymentType" field.
func PaymentTypeNotNil() predicate.JobPayments {
	return predicate.JobPayments(sql.FieldNotNull(FieldPaymentType))
}

// PaymentTypeEqualFold applies the EqualFold predicate on the "PaymentType" field.
func PaymentTypeEqualFold(v string) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldEqualFold(FieldPaymentType, v))
}

// PaymentTypeContainsFold applies the ContainsFold predicate on the "PaymentType" field.
func PaymentTypeContainsFold(v string) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldContainsFold(FieldPaymentType, v))
}

// StateEQ applies the EQ predicate on the "State" field.
func StateEQ(v string) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldEQ(FieldState, v))
}

// StateNEQ applies the NEQ predicate on the "State" field.
func StateNEQ(v string) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldNEQ(FieldState, v))
}

// StateIn applies the In predicate on the "State" field.
func StateIn(vs ...string) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldIn(FieldState, vs...))
}

// StateNotIn applies the NotIn predicate on the "State" field.
func StateNotIn(vs ...string) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldNotIn(FieldState, vs...))
}

// StateGT applies the GT predicate on the "State" field.
func StateGT(v string) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldGT(FieldState, v))
}

// StateGTE applies the GTE predicate on the "State" field.
func StateGTE(v string) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldGTE(FieldState, v))
}

// StateLT applies the LT predicate on the "State" field.
func StateLT(v string) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldLT(FieldState, v))
}

// StateLTE applies the LTE predicate on the "State" field.
func StateLTE(v string) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldLTE(FieldState, v))
}

// StateContains applies the Contains predicate on the "State" field.
func StateContains(v string) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldContains(FieldState, v))
}

// StateHasPrefix applies the HasPrefix predicate on the "State" field.
func StateHasPrefix(v string) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldHasPrefix(FieldState, v))
}

// StateHasSuffix applies the HasSuffix predicate on the "State" field.
func StateHasSuffix(v string) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldHasSuffix(FieldState, v))
}

// StateIsNil applies the IsNil predicate on the "State" field.
func StateIsNil() predicate.JobPayments {
	return predicate.JobPayments(sql.FieldIsNull(FieldState))
}

// StateNotNil applies the NotNil predicate on the "State" field.
func StateNotNil() predicate.JobPayments {
	return predicate.JobPayments(sql.FieldNotNull(FieldState))
}

// StateEqualFold applies the EqualFold predicate on the "State" field.
func StateEqualFold(v string) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldEqualFold(FieldState, v))
}

// StateContainsFold applies the ContainsFold predicate on the "State" field.
func StateContainsFold(v string) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldContainsFold(FieldState, v))
}

// TotalPaymentEQ applies the EQ predicate on the "TotalPayment" field.
func TotalPaymentEQ(v float64) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldEQ(FieldTotalPayment, v))
}

// TotalPaymentNEQ applies the NEQ predicate on the "TotalPayment" field.
func TotalPaymentNEQ(v float64) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldNEQ(FieldTotalPayment, v))
}

// TotalPaymentIn applies the In predicate on the "TotalPayment" field.
func TotalPaymentIn(vs ...float64) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldIn(FieldTotalPayment, vs...))
}

// TotalPaymentNotIn applies the NotIn predicate on the "TotalPayment" field.
func TotalPaymentNotIn(vs ...float64) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldNotIn(FieldTotalPayment, vs...))
}

// TotalPaymentGT applies the GT predicate on the "TotalPayment" field.
func TotalPaymentGT(v float64) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldGT(FieldTotalPayment, v))
}

// TotalPaymentGTE applies the GTE predicate on the "TotalPayment" field.
func TotalPaymentGTE(v float64) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldGTE(FieldTotalPayment, v))
}

// TotalPaymentLT applies the LT predicate on the "TotalPayment" field.
func TotalPaymentLT(v float64) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldLT(FieldTotalPayment, v))
}

// TotalPaymentLTE applies the LTE predicate on the "TotalPayment" field.
func TotalPaymentLTE(v float64) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldLTE(FieldTotalPayment, v))
}

// TotalPaymentIsNil applies the IsNil predicate on the "TotalPayment" field.
func TotalPaymentIsNil() predicate.JobPayments {
	return predicate.JobPayments(sql.FieldIsNull(FieldTotalPayment))
}

// TotalPaymentNotNil applies the NotNil predicate on the "TotalPayment" field.
func TotalPaymentNotNil() predicate.JobPayments {
	return predicate.JobPayments(sql.FieldNotNull(FieldTotalPayment))
}

// LevelRequestEQ applies the EQ predicate on the "LevelRequest" field.
func LevelRequestEQ(v float64) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldEQ(FieldLevelRequest, v))
}

// LevelRequestNEQ applies the NEQ predicate on the "LevelRequest" field.
func LevelRequestNEQ(v float64) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldNEQ(FieldLevelRequest, v))
}

// LevelRequestIn applies the In predicate on the "LevelRequest" field.
func LevelRequestIn(vs ...float64) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldIn(FieldLevelRequest, vs...))
}

// LevelRequestNotIn applies the NotIn predicate on the "LevelRequest" field.
func LevelRequestNotIn(vs ...float64) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldNotIn(FieldLevelRequest, vs...))
}

// LevelRequestGT applies the GT predicate on the "LevelRequest" field.
func LevelRequestGT(v float64) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldGT(FieldLevelRequest, v))
}

// LevelRequestGTE applies the GTE predicate on the "LevelRequest" field.
func LevelRequestGTE(v float64) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldGTE(FieldLevelRequest, v))
}

// LevelRequestLT applies the LT predicate on the "LevelRequest" field.
func LevelRequestLT(v float64) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldLT(FieldLevelRequest, v))
}

// LevelRequestLTE applies the LTE predicate on the "LevelRequest" field.
func LevelRequestLTE(v float64) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldLTE(FieldLevelRequest, v))
}

// LevelRequestIsNil applies the IsNil predicate on the "LevelRequest" field.
func LevelRequestIsNil() predicate.JobPayments {
	return predicate.JobPayments(sql.FieldIsNull(FieldLevelRequest))
}

// LevelRequestNotNil applies the NotNil predicate on the "LevelRequest" field.
func LevelRequestNotNil() predicate.JobPayments {
	return predicate.JobPayments(sql.FieldNotNull(FieldLevelRequest))
}

// LevelApproveEQ applies the EQ predicate on the "LevelApprove" field.
func LevelApproveEQ(v float64) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldEQ(FieldLevelApprove, v))
}

// LevelApproveNEQ applies the NEQ predicate on the "LevelApprove" field.
func LevelApproveNEQ(v float64) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldNEQ(FieldLevelApprove, v))
}

// LevelApproveIn applies the In predicate on the "LevelApprove" field.
func LevelApproveIn(vs ...float64) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldIn(FieldLevelApprove, vs...))
}

// LevelApproveNotIn applies the NotIn predicate on the "LevelApprove" field.
func LevelApproveNotIn(vs ...float64) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldNotIn(FieldLevelApprove, vs...))
}

// LevelApproveGT applies the GT predicate on the "LevelApprove" field.
func LevelApproveGT(v float64) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldGT(FieldLevelApprove, v))
}

// LevelApproveGTE applies the GTE predicate on the "LevelApprove" field.
func LevelApproveGTE(v float64) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldGTE(FieldLevelApprove, v))
}

// LevelApproveLT applies the LT predicate on the "LevelApprove" field.
func LevelApproveLT(v float64) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldLT(FieldLevelApprove, v))
}

// LevelApproveLTE applies the LTE predicate on the "LevelApprove" field.
func LevelApproveLTE(v float64) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldLTE(FieldLevelApprove, v))
}

// LevelApproveIsNil applies the IsNil predicate on the "LevelApprove" field.
func LevelApproveIsNil() predicate.JobPayments {
	return predicate.JobPayments(sql.FieldIsNull(FieldLevelApprove))
}

// LevelApproveNotNil applies the NotNil predicate on the "LevelApprove" field.
func LevelApproveNotNil() predicate.JobPayments {
	return predicate.JobPayments(sql.FieldNotNull(FieldLevelApprove))
}

// AmountEQ applies the EQ predicate on the "Amount" field.
func AmountEQ(v float64) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldEQ(FieldAmount, v))
}

// AmountNEQ applies the NEQ predicate on the "Amount" field.
func AmountNEQ(v float64) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldNEQ(FieldAmount, v))
}

// AmountIn applies the In predicate on the "Amount" field.
func AmountIn(vs ...float64) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldIn(FieldAmount, vs...))
}

// AmountNotIn applies the NotIn predicate on the "Amount" field.
func AmountNotIn(vs ...float64) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldNotIn(FieldAmount, vs...))
}

// AmountGT applies the GT predicate on the "Amount" field.
func AmountGT(v float64) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldGT(FieldAmount, v))
}

// AmountGTE applies the GTE predicate on the "Amount" field.
func AmountGTE(v float64) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldGTE(FieldAmount, v))
}

// AmountLT applies the LT predicate on the "Amount" field.
func AmountLT(v float64) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldLT(FieldAmount, v))
}

// AmountLTE applies the LTE predicate on the "Amount" field.
func AmountLTE(v float64) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldLTE(FieldAmount, v))
}

// AmountIsNil applies the IsNil predicate on the "Amount" field.
func AmountIsNil() predicate.JobPayments {
	return predicate.JobPayments(sql.FieldIsNull(FieldAmount))
}

// AmountNotNil applies the NotNil predicate on the "Amount" field.
func AmountNotNil() predicate.JobPayments {
	return predicate.JobPayments(sql.FieldNotNull(FieldAmount))
}

// CreatedAtEQ applies the EQ predicate on the "CreatedAt" field.
func CreatedAtEQ(v time.Time) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "CreatedAt" field.
func CreatedAtNEQ(v time.Time) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "CreatedAt" field.
func CreatedAtIn(vs ...time.Time) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "CreatedAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "CreatedAt" field.
func CreatedAtGT(v time.Time) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "CreatedAt" field.
func CreatedAtGTE(v time.Time) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "CreatedAt" field.
func CreatedAtLT(v time.Time) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "CreatedAt" field.
func CreatedAtLTE(v time.Time) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "UpdatedAt" field.
func UpdatedAtEQ(v time.Time) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "UpdatedAt" field.
func UpdatedAtNEQ(v time.Time) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "UpdatedAt" field.
func UpdatedAtIn(vs ...time.Time) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "UpdatedAt" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "UpdatedAt" field.
func UpdatedAtGT(v time.Time) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "UpdatedAt" field.
func UpdatedAtGTE(v time.Time) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "UpdatedAt" field.
func UpdatedAtLT(v time.Time) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "UpdatedAt" field.
func UpdatedAtLTE(v time.Time) predicate.JobPayments {
	return predicate.JobPayments(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasPayments applies the HasEdge predicate on the "payments" edge.
func HasPayments() predicate.JobPayments {
	return predicate.JobPayments(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PaymentsTable, PaymentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPaymentsWith applies the HasEdge predicate on the "payments" edge with a given conditions (other predicates).
func HasPaymentsWith(preds ...predicate.JobRelations) predicate.JobPayments {
	return predicate.JobPayments(func(s *sql.Selector) {
		step := newPaymentsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.JobPayments) predicate.JobPayments {
	return predicate.JobPayments(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.JobPayments) predicate.JobPayments {
	return predicate.JobPayments(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.JobPayments) predicate.JobPayments {
	return predicate.JobPayments(sql.NotPredicates(p))
}
