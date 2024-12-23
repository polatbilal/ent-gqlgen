// Code generated by ent, DO NOT EDIT.

package jobcontractor

import (
	"gqlgen-ent/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "Name" field. It's identical to NameEQ.
func Name(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldEQ(FieldName, v))
}

// TcNo applies equality check predicate on the "TcNo" field. It's identical to TcNoEQ.
func TcNo(v int) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldEQ(FieldTcNo, v))
}

// Address applies equality check predicate on the "Address" field. It's identical to AddressEQ.
func Address(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldEQ(FieldAddress, v))
}

// RegisterNo applies equality check predicate on the "RegisterNo" field. It's identical to RegisterNoEQ.
func RegisterNo(v int) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldEQ(FieldRegisterNo, v))
}

// TaxAdmin applies equality check predicate on the "TaxAdmin" field. It's identical to TaxAdminEQ.
func TaxAdmin(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldEQ(FieldTaxAdmin, v))
}

// TaxNo applies equality check predicate on the "TaxNo" field. It's identical to TaxNoEQ.
func TaxNo(v int) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldEQ(FieldTaxNo, v))
}

// Phone applies equality check predicate on the "Phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldEQ(FieldPhone, v))
}

// Email applies equality check predicate on the "Email" field. It's identical to EmailEQ.
func Email(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldEQ(FieldEmail, v))
}

// Note applies equality check predicate on the "Note" field. It's identical to NoteEQ.
func Note(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldEQ(FieldNote, v))
}

// CreatedAt applies equality check predicate on the "CreatedAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "UpdatedAt" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldEQ(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "Name" field.
func NameEQ(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "Name" field.
func NameNEQ(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "Name" field.
func NameIn(vs ...string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "Name" field.
func NameNotIn(vs ...string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "Name" field.
func NameGT(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "Name" field.
func NameGTE(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "Name" field.
func NameLT(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "Name" field.
func NameLTE(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "Name" field.
func NameContains(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "Name" field.
func NameHasPrefix(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "Name" field.
func NameHasSuffix(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "Name" field.
func NameEqualFold(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "Name" field.
func NameContainsFold(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldContainsFold(FieldName, v))
}

// TcNoEQ applies the EQ predicate on the "TcNo" field.
func TcNoEQ(v int) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldEQ(FieldTcNo, v))
}

// TcNoNEQ applies the NEQ predicate on the "TcNo" field.
func TcNoNEQ(v int) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldNEQ(FieldTcNo, v))
}

// TcNoIn applies the In predicate on the "TcNo" field.
func TcNoIn(vs ...int) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldIn(FieldTcNo, vs...))
}

// TcNoNotIn applies the NotIn predicate on the "TcNo" field.
func TcNoNotIn(vs ...int) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldNotIn(FieldTcNo, vs...))
}

// TcNoGT applies the GT predicate on the "TcNo" field.
func TcNoGT(v int) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldGT(FieldTcNo, v))
}

// TcNoGTE applies the GTE predicate on the "TcNo" field.
func TcNoGTE(v int) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldGTE(FieldTcNo, v))
}

// TcNoLT applies the LT predicate on the "TcNo" field.
func TcNoLT(v int) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldLT(FieldTcNo, v))
}

// TcNoLTE applies the LTE predicate on the "TcNo" field.
func TcNoLTE(v int) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldLTE(FieldTcNo, v))
}

// TcNoIsNil applies the IsNil predicate on the "TcNo" field.
func TcNoIsNil() predicate.JobContractor {
	return predicate.JobContractor(sql.FieldIsNull(FieldTcNo))
}

// TcNoNotNil applies the NotNil predicate on the "TcNo" field.
func TcNoNotNil() predicate.JobContractor {
	return predicate.JobContractor(sql.FieldNotNull(FieldTcNo))
}

// AddressEQ applies the EQ predicate on the "Address" field.
func AddressEQ(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldEQ(FieldAddress, v))
}

// AddressNEQ applies the NEQ predicate on the "Address" field.
func AddressNEQ(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldNEQ(FieldAddress, v))
}

// AddressIn applies the In predicate on the "Address" field.
func AddressIn(vs ...string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldIn(FieldAddress, vs...))
}

// AddressNotIn applies the NotIn predicate on the "Address" field.
func AddressNotIn(vs ...string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldNotIn(FieldAddress, vs...))
}

// AddressGT applies the GT predicate on the "Address" field.
func AddressGT(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldGT(FieldAddress, v))
}

// AddressGTE applies the GTE predicate on the "Address" field.
func AddressGTE(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldGTE(FieldAddress, v))
}

// AddressLT applies the LT predicate on the "Address" field.
func AddressLT(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldLT(FieldAddress, v))
}

// AddressLTE applies the LTE predicate on the "Address" field.
func AddressLTE(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldLTE(FieldAddress, v))
}

// AddressContains applies the Contains predicate on the "Address" field.
func AddressContains(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldContains(FieldAddress, v))
}

// AddressHasPrefix applies the HasPrefix predicate on the "Address" field.
func AddressHasPrefix(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldHasPrefix(FieldAddress, v))
}

// AddressHasSuffix applies the HasSuffix predicate on the "Address" field.
func AddressHasSuffix(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldHasSuffix(FieldAddress, v))
}

// AddressIsNil applies the IsNil predicate on the "Address" field.
func AddressIsNil() predicate.JobContractor {
	return predicate.JobContractor(sql.FieldIsNull(FieldAddress))
}

// AddressNotNil applies the NotNil predicate on the "Address" field.
func AddressNotNil() predicate.JobContractor {
	return predicate.JobContractor(sql.FieldNotNull(FieldAddress))
}

// AddressEqualFold applies the EqualFold predicate on the "Address" field.
func AddressEqualFold(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldEqualFold(FieldAddress, v))
}

// AddressContainsFold applies the ContainsFold predicate on the "Address" field.
func AddressContainsFold(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldContainsFold(FieldAddress, v))
}

// RegisterNoEQ applies the EQ predicate on the "RegisterNo" field.
func RegisterNoEQ(v int) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldEQ(FieldRegisterNo, v))
}

// RegisterNoNEQ applies the NEQ predicate on the "RegisterNo" field.
func RegisterNoNEQ(v int) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldNEQ(FieldRegisterNo, v))
}

// RegisterNoIn applies the In predicate on the "RegisterNo" field.
func RegisterNoIn(vs ...int) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldIn(FieldRegisterNo, vs...))
}

// RegisterNoNotIn applies the NotIn predicate on the "RegisterNo" field.
func RegisterNoNotIn(vs ...int) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldNotIn(FieldRegisterNo, vs...))
}

// RegisterNoGT applies the GT predicate on the "RegisterNo" field.
func RegisterNoGT(v int) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldGT(FieldRegisterNo, v))
}

// RegisterNoGTE applies the GTE predicate on the "RegisterNo" field.
func RegisterNoGTE(v int) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldGTE(FieldRegisterNo, v))
}

// RegisterNoLT applies the LT predicate on the "RegisterNo" field.
func RegisterNoLT(v int) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldLT(FieldRegisterNo, v))
}

// RegisterNoLTE applies the LTE predicate on the "RegisterNo" field.
func RegisterNoLTE(v int) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldLTE(FieldRegisterNo, v))
}

// RegisterNoIsNil applies the IsNil predicate on the "RegisterNo" field.
func RegisterNoIsNil() predicate.JobContractor {
	return predicate.JobContractor(sql.FieldIsNull(FieldRegisterNo))
}

// RegisterNoNotNil applies the NotNil predicate on the "RegisterNo" field.
func RegisterNoNotNil() predicate.JobContractor {
	return predicate.JobContractor(sql.FieldNotNull(FieldRegisterNo))
}

// TaxAdminEQ applies the EQ predicate on the "TaxAdmin" field.
func TaxAdminEQ(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldEQ(FieldTaxAdmin, v))
}

// TaxAdminNEQ applies the NEQ predicate on the "TaxAdmin" field.
func TaxAdminNEQ(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldNEQ(FieldTaxAdmin, v))
}

// TaxAdminIn applies the In predicate on the "TaxAdmin" field.
func TaxAdminIn(vs ...string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldIn(FieldTaxAdmin, vs...))
}

// TaxAdminNotIn applies the NotIn predicate on the "TaxAdmin" field.
func TaxAdminNotIn(vs ...string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldNotIn(FieldTaxAdmin, vs...))
}

// TaxAdminGT applies the GT predicate on the "TaxAdmin" field.
func TaxAdminGT(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldGT(FieldTaxAdmin, v))
}

// TaxAdminGTE applies the GTE predicate on the "TaxAdmin" field.
func TaxAdminGTE(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldGTE(FieldTaxAdmin, v))
}

// TaxAdminLT applies the LT predicate on the "TaxAdmin" field.
func TaxAdminLT(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldLT(FieldTaxAdmin, v))
}

// TaxAdminLTE applies the LTE predicate on the "TaxAdmin" field.
func TaxAdminLTE(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldLTE(FieldTaxAdmin, v))
}

// TaxAdminContains applies the Contains predicate on the "TaxAdmin" field.
func TaxAdminContains(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldContains(FieldTaxAdmin, v))
}

// TaxAdminHasPrefix applies the HasPrefix predicate on the "TaxAdmin" field.
func TaxAdminHasPrefix(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldHasPrefix(FieldTaxAdmin, v))
}

// TaxAdminHasSuffix applies the HasSuffix predicate on the "TaxAdmin" field.
func TaxAdminHasSuffix(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldHasSuffix(FieldTaxAdmin, v))
}

// TaxAdminIsNil applies the IsNil predicate on the "TaxAdmin" field.
func TaxAdminIsNil() predicate.JobContractor {
	return predicate.JobContractor(sql.FieldIsNull(FieldTaxAdmin))
}

// TaxAdminNotNil applies the NotNil predicate on the "TaxAdmin" field.
func TaxAdminNotNil() predicate.JobContractor {
	return predicate.JobContractor(sql.FieldNotNull(FieldTaxAdmin))
}

// TaxAdminEqualFold applies the EqualFold predicate on the "TaxAdmin" field.
func TaxAdminEqualFold(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldEqualFold(FieldTaxAdmin, v))
}

// TaxAdminContainsFold applies the ContainsFold predicate on the "TaxAdmin" field.
func TaxAdminContainsFold(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldContainsFold(FieldTaxAdmin, v))
}

// TaxNoEQ applies the EQ predicate on the "TaxNo" field.
func TaxNoEQ(v int) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldEQ(FieldTaxNo, v))
}

// TaxNoNEQ applies the NEQ predicate on the "TaxNo" field.
func TaxNoNEQ(v int) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldNEQ(FieldTaxNo, v))
}

// TaxNoIn applies the In predicate on the "TaxNo" field.
func TaxNoIn(vs ...int) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldIn(FieldTaxNo, vs...))
}

// TaxNoNotIn applies the NotIn predicate on the "TaxNo" field.
func TaxNoNotIn(vs ...int) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldNotIn(FieldTaxNo, vs...))
}

// TaxNoGT applies the GT predicate on the "TaxNo" field.
func TaxNoGT(v int) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldGT(FieldTaxNo, v))
}

// TaxNoGTE applies the GTE predicate on the "TaxNo" field.
func TaxNoGTE(v int) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldGTE(FieldTaxNo, v))
}

// TaxNoLT applies the LT predicate on the "TaxNo" field.
func TaxNoLT(v int) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldLT(FieldTaxNo, v))
}

// TaxNoLTE applies the LTE predicate on the "TaxNo" field.
func TaxNoLTE(v int) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldLTE(FieldTaxNo, v))
}

// TaxNoIsNil applies the IsNil predicate on the "TaxNo" field.
func TaxNoIsNil() predicate.JobContractor {
	return predicate.JobContractor(sql.FieldIsNull(FieldTaxNo))
}

// TaxNoNotNil applies the NotNil predicate on the "TaxNo" field.
func TaxNoNotNil() predicate.JobContractor {
	return predicate.JobContractor(sql.FieldNotNull(FieldTaxNo))
}

// PhoneEQ applies the EQ predicate on the "Phone" field.
func PhoneEQ(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldEQ(FieldPhone, v))
}

// PhoneNEQ applies the NEQ predicate on the "Phone" field.
func PhoneNEQ(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldNEQ(FieldPhone, v))
}

// PhoneIn applies the In predicate on the "Phone" field.
func PhoneIn(vs ...string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldIn(FieldPhone, vs...))
}

// PhoneNotIn applies the NotIn predicate on the "Phone" field.
func PhoneNotIn(vs ...string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldNotIn(FieldPhone, vs...))
}

// PhoneGT applies the GT predicate on the "Phone" field.
func PhoneGT(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldGT(FieldPhone, v))
}

// PhoneGTE applies the GTE predicate on the "Phone" field.
func PhoneGTE(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldGTE(FieldPhone, v))
}

// PhoneLT applies the LT predicate on the "Phone" field.
func PhoneLT(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldLT(FieldPhone, v))
}

// PhoneLTE applies the LTE predicate on the "Phone" field.
func PhoneLTE(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldLTE(FieldPhone, v))
}

// PhoneContains applies the Contains predicate on the "Phone" field.
func PhoneContains(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldContains(FieldPhone, v))
}

// PhoneHasPrefix applies the HasPrefix predicate on the "Phone" field.
func PhoneHasPrefix(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldHasPrefix(FieldPhone, v))
}

// PhoneHasSuffix applies the HasSuffix predicate on the "Phone" field.
func PhoneHasSuffix(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldHasSuffix(FieldPhone, v))
}

// PhoneIsNil applies the IsNil predicate on the "Phone" field.
func PhoneIsNil() predicate.JobContractor {
	return predicate.JobContractor(sql.FieldIsNull(FieldPhone))
}

// PhoneNotNil applies the NotNil predicate on the "Phone" field.
func PhoneNotNil() predicate.JobContractor {
	return predicate.JobContractor(sql.FieldNotNull(FieldPhone))
}

// PhoneEqualFold applies the EqualFold predicate on the "Phone" field.
func PhoneEqualFold(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldEqualFold(FieldPhone, v))
}

// PhoneContainsFold applies the ContainsFold predicate on the "Phone" field.
func PhoneContainsFold(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldContainsFold(FieldPhone, v))
}

// EmailEQ applies the EQ predicate on the "Email" field.
func EmailEQ(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "Email" field.
func EmailNEQ(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "Email" field.
func EmailIn(vs ...string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "Email" field.
func EmailNotIn(vs ...string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "Email" field.
func EmailGT(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "Email" field.
func EmailGTE(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "Email" field.
func EmailLT(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "Email" field.
func EmailLTE(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "Email" field.
func EmailContains(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "Email" field.
func EmailHasPrefix(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "Email" field.
func EmailHasSuffix(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailIsNil applies the IsNil predicate on the "Email" field.
func EmailIsNil() predicate.JobContractor {
	return predicate.JobContractor(sql.FieldIsNull(FieldEmail))
}

// EmailNotNil applies the NotNil predicate on the "Email" field.
func EmailNotNil() predicate.JobContractor {
	return predicate.JobContractor(sql.FieldNotNull(FieldEmail))
}

// EmailEqualFold applies the EqualFold predicate on the "Email" field.
func EmailEqualFold(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "Email" field.
func EmailContainsFold(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldContainsFold(FieldEmail, v))
}

// NoteEQ applies the EQ predicate on the "Note" field.
func NoteEQ(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldEQ(FieldNote, v))
}

// NoteNEQ applies the NEQ predicate on the "Note" field.
func NoteNEQ(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldNEQ(FieldNote, v))
}

// NoteIn applies the In predicate on the "Note" field.
func NoteIn(vs ...string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldIn(FieldNote, vs...))
}

// NoteNotIn applies the NotIn predicate on the "Note" field.
func NoteNotIn(vs ...string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldNotIn(FieldNote, vs...))
}

// NoteGT applies the GT predicate on the "Note" field.
func NoteGT(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldGT(FieldNote, v))
}

// NoteGTE applies the GTE predicate on the "Note" field.
func NoteGTE(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldGTE(FieldNote, v))
}

// NoteLT applies the LT predicate on the "Note" field.
func NoteLT(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldLT(FieldNote, v))
}

// NoteLTE applies the LTE predicate on the "Note" field.
func NoteLTE(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldLTE(FieldNote, v))
}

// NoteContains applies the Contains predicate on the "Note" field.
func NoteContains(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldContains(FieldNote, v))
}

// NoteHasPrefix applies the HasPrefix predicate on the "Note" field.
func NoteHasPrefix(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldHasPrefix(FieldNote, v))
}

// NoteHasSuffix applies the HasSuffix predicate on the "Note" field.
func NoteHasSuffix(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldHasSuffix(FieldNote, v))
}

// NoteIsNil applies the IsNil predicate on the "Note" field.
func NoteIsNil() predicate.JobContractor {
	return predicate.JobContractor(sql.FieldIsNull(FieldNote))
}

// NoteNotNil applies the NotNil predicate on the "Note" field.
func NoteNotNil() predicate.JobContractor {
	return predicate.JobContractor(sql.FieldNotNull(FieldNote))
}

// NoteEqualFold applies the EqualFold predicate on the "Note" field.
func NoteEqualFold(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldEqualFold(FieldNote, v))
}

// NoteContainsFold applies the ContainsFold predicate on the "Note" field.
func NoteContainsFold(v string) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldContainsFold(FieldNote, v))
}

// CreatedAtEQ applies the EQ predicate on the "CreatedAt" field.
func CreatedAtEQ(v time.Time) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "CreatedAt" field.
func CreatedAtNEQ(v time.Time) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "CreatedAt" field.
func CreatedAtIn(vs ...time.Time) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "CreatedAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "CreatedAt" field.
func CreatedAtGT(v time.Time) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "CreatedAt" field.
func CreatedAtGTE(v time.Time) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "CreatedAt" field.
func CreatedAtLT(v time.Time) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "CreatedAt" field.
func CreatedAtLTE(v time.Time) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "UpdatedAt" field.
func UpdatedAtEQ(v time.Time) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "UpdatedAt" field.
func UpdatedAtNEQ(v time.Time) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "UpdatedAt" field.
func UpdatedAtIn(vs ...time.Time) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "UpdatedAt" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "UpdatedAt" field.
func UpdatedAtGT(v time.Time) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "UpdatedAt" field.
func UpdatedAtGTE(v time.Time) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "UpdatedAt" field.
func UpdatedAtLT(v time.Time) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "UpdatedAt" field.
func UpdatedAtLTE(v time.Time) predicate.JobContractor {
	return predicate.JobContractor(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasContractors applies the HasEdge predicate on the "contractors" edge.
func HasContractors() predicate.JobContractor {
	return predicate.JobContractor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ContractorsTable, ContractorsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasContractorsWith applies the HasEdge predicate on the "contractors" edge with a given conditions (other predicates).
func HasContractorsWith(preds ...predicate.JobDetail) predicate.JobContractor {
	return predicate.JobContractor(func(s *sql.Selector) {
		step := newContractorsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.JobContractor) predicate.JobContractor {
	return predicate.JobContractor(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.JobContractor) predicate.JobContractor {
	return predicate.JobContractor(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.JobContractor) predicate.JobContractor {
	return predicate.JobContractor(sql.NotPredicates(p))
}
