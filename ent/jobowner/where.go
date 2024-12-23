// Code generated by ent, DO NOT EDIT.

package jobowner

import (
	"gqlgen-ent/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "Name" field. It's identical to NameEQ.
func Name(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldEQ(FieldName, v))
}

// TcNo applies equality check predicate on the "TcNo" field. It's identical to TcNoEQ.
func TcNo(v int) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldEQ(FieldTcNo, v))
}

// Address applies equality check predicate on the "Address" field. It's identical to AddressEQ.
func Address(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldEQ(FieldAddress, v))
}

// TaxAdmin applies equality check predicate on the "TaxAdmin" field. It's identical to TaxAdminEQ.
func TaxAdmin(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldEQ(FieldTaxAdmin, v))
}

// TaxNo applies equality check predicate on the "TaxNo" field. It's identical to TaxNoEQ.
func TaxNo(v int) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldEQ(FieldTaxNo, v))
}

// Phone applies equality check predicate on the "Phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldEQ(FieldPhone, v))
}

// Email applies equality check predicate on the "Email" field. It's identical to EmailEQ.
func Email(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldEQ(FieldEmail, v))
}

// Note applies equality check predicate on the "Note" field. It's identical to NoteEQ.
func Note(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldEQ(FieldNote, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldEQ(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "Name" field.
func NameEQ(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "Name" field.
func NameNEQ(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "Name" field.
func NameIn(vs ...string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "Name" field.
func NameNotIn(vs ...string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "Name" field.
func NameGT(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "Name" field.
func NameGTE(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "Name" field.
func NameLT(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "Name" field.
func NameLTE(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "Name" field.
func NameContains(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "Name" field.
func NameHasPrefix(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "Name" field.
func NameHasSuffix(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "Name" field.
func NameEqualFold(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "Name" field.
func NameContainsFold(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldContainsFold(FieldName, v))
}

// TcNoEQ applies the EQ predicate on the "TcNo" field.
func TcNoEQ(v int) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldEQ(FieldTcNo, v))
}

// TcNoNEQ applies the NEQ predicate on the "TcNo" field.
func TcNoNEQ(v int) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldNEQ(FieldTcNo, v))
}

// TcNoIn applies the In predicate on the "TcNo" field.
func TcNoIn(vs ...int) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldIn(FieldTcNo, vs...))
}

// TcNoNotIn applies the NotIn predicate on the "TcNo" field.
func TcNoNotIn(vs ...int) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldNotIn(FieldTcNo, vs...))
}

// TcNoGT applies the GT predicate on the "TcNo" field.
func TcNoGT(v int) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldGT(FieldTcNo, v))
}

// TcNoGTE applies the GTE predicate on the "TcNo" field.
func TcNoGTE(v int) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldGTE(FieldTcNo, v))
}

// TcNoLT applies the LT predicate on the "TcNo" field.
func TcNoLT(v int) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldLT(FieldTcNo, v))
}

// TcNoLTE applies the LTE predicate on the "TcNo" field.
func TcNoLTE(v int) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldLTE(FieldTcNo, v))
}

// TcNoIsNil applies the IsNil predicate on the "TcNo" field.
func TcNoIsNil() predicate.JobOwner {
	return predicate.JobOwner(sql.FieldIsNull(FieldTcNo))
}

// TcNoNotNil applies the NotNil predicate on the "TcNo" field.
func TcNoNotNil() predicate.JobOwner {
	return predicate.JobOwner(sql.FieldNotNull(FieldTcNo))
}

// AddressEQ applies the EQ predicate on the "Address" field.
func AddressEQ(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldEQ(FieldAddress, v))
}

// AddressNEQ applies the NEQ predicate on the "Address" field.
func AddressNEQ(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldNEQ(FieldAddress, v))
}

// AddressIn applies the In predicate on the "Address" field.
func AddressIn(vs ...string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldIn(FieldAddress, vs...))
}

// AddressNotIn applies the NotIn predicate on the "Address" field.
func AddressNotIn(vs ...string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldNotIn(FieldAddress, vs...))
}

// AddressGT applies the GT predicate on the "Address" field.
func AddressGT(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldGT(FieldAddress, v))
}

// AddressGTE applies the GTE predicate on the "Address" field.
func AddressGTE(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldGTE(FieldAddress, v))
}

// AddressLT applies the LT predicate on the "Address" field.
func AddressLT(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldLT(FieldAddress, v))
}

// AddressLTE applies the LTE predicate on the "Address" field.
func AddressLTE(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldLTE(FieldAddress, v))
}

// AddressContains applies the Contains predicate on the "Address" field.
func AddressContains(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldContains(FieldAddress, v))
}

// AddressHasPrefix applies the HasPrefix predicate on the "Address" field.
func AddressHasPrefix(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldHasPrefix(FieldAddress, v))
}

// AddressHasSuffix applies the HasSuffix predicate on the "Address" field.
func AddressHasSuffix(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldHasSuffix(FieldAddress, v))
}

// AddressIsNil applies the IsNil predicate on the "Address" field.
func AddressIsNil() predicate.JobOwner {
	return predicate.JobOwner(sql.FieldIsNull(FieldAddress))
}

// AddressNotNil applies the NotNil predicate on the "Address" field.
func AddressNotNil() predicate.JobOwner {
	return predicate.JobOwner(sql.FieldNotNull(FieldAddress))
}

// AddressEqualFold applies the EqualFold predicate on the "Address" field.
func AddressEqualFold(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldEqualFold(FieldAddress, v))
}

// AddressContainsFold applies the ContainsFold predicate on the "Address" field.
func AddressContainsFold(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldContainsFold(FieldAddress, v))
}

// TaxAdminEQ applies the EQ predicate on the "TaxAdmin" field.
func TaxAdminEQ(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldEQ(FieldTaxAdmin, v))
}

// TaxAdminNEQ applies the NEQ predicate on the "TaxAdmin" field.
func TaxAdminNEQ(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldNEQ(FieldTaxAdmin, v))
}

// TaxAdminIn applies the In predicate on the "TaxAdmin" field.
func TaxAdminIn(vs ...string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldIn(FieldTaxAdmin, vs...))
}

// TaxAdminNotIn applies the NotIn predicate on the "TaxAdmin" field.
func TaxAdminNotIn(vs ...string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldNotIn(FieldTaxAdmin, vs...))
}

// TaxAdminGT applies the GT predicate on the "TaxAdmin" field.
func TaxAdminGT(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldGT(FieldTaxAdmin, v))
}

// TaxAdminGTE applies the GTE predicate on the "TaxAdmin" field.
func TaxAdminGTE(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldGTE(FieldTaxAdmin, v))
}

// TaxAdminLT applies the LT predicate on the "TaxAdmin" field.
func TaxAdminLT(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldLT(FieldTaxAdmin, v))
}

// TaxAdminLTE applies the LTE predicate on the "TaxAdmin" field.
func TaxAdminLTE(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldLTE(FieldTaxAdmin, v))
}

// TaxAdminContains applies the Contains predicate on the "TaxAdmin" field.
func TaxAdminContains(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldContains(FieldTaxAdmin, v))
}

// TaxAdminHasPrefix applies the HasPrefix predicate on the "TaxAdmin" field.
func TaxAdminHasPrefix(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldHasPrefix(FieldTaxAdmin, v))
}

// TaxAdminHasSuffix applies the HasSuffix predicate on the "TaxAdmin" field.
func TaxAdminHasSuffix(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldHasSuffix(FieldTaxAdmin, v))
}

// TaxAdminIsNil applies the IsNil predicate on the "TaxAdmin" field.
func TaxAdminIsNil() predicate.JobOwner {
	return predicate.JobOwner(sql.FieldIsNull(FieldTaxAdmin))
}

// TaxAdminNotNil applies the NotNil predicate on the "TaxAdmin" field.
func TaxAdminNotNil() predicate.JobOwner {
	return predicate.JobOwner(sql.FieldNotNull(FieldTaxAdmin))
}

// TaxAdminEqualFold applies the EqualFold predicate on the "TaxAdmin" field.
func TaxAdminEqualFold(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldEqualFold(FieldTaxAdmin, v))
}

// TaxAdminContainsFold applies the ContainsFold predicate on the "TaxAdmin" field.
func TaxAdminContainsFold(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldContainsFold(FieldTaxAdmin, v))
}

// TaxNoEQ applies the EQ predicate on the "TaxNo" field.
func TaxNoEQ(v int) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldEQ(FieldTaxNo, v))
}

// TaxNoNEQ applies the NEQ predicate on the "TaxNo" field.
func TaxNoNEQ(v int) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldNEQ(FieldTaxNo, v))
}

// TaxNoIn applies the In predicate on the "TaxNo" field.
func TaxNoIn(vs ...int) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldIn(FieldTaxNo, vs...))
}

// TaxNoNotIn applies the NotIn predicate on the "TaxNo" field.
func TaxNoNotIn(vs ...int) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldNotIn(FieldTaxNo, vs...))
}

// TaxNoGT applies the GT predicate on the "TaxNo" field.
func TaxNoGT(v int) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldGT(FieldTaxNo, v))
}

// TaxNoGTE applies the GTE predicate on the "TaxNo" field.
func TaxNoGTE(v int) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldGTE(FieldTaxNo, v))
}

// TaxNoLT applies the LT predicate on the "TaxNo" field.
func TaxNoLT(v int) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldLT(FieldTaxNo, v))
}

// TaxNoLTE applies the LTE predicate on the "TaxNo" field.
func TaxNoLTE(v int) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldLTE(FieldTaxNo, v))
}

// TaxNoIsNil applies the IsNil predicate on the "TaxNo" field.
func TaxNoIsNil() predicate.JobOwner {
	return predicate.JobOwner(sql.FieldIsNull(FieldTaxNo))
}

// TaxNoNotNil applies the NotNil predicate on the "TaxNo" field.
func TaxNoNotNil() predicate.JobOwner {
	return predicate.JobOwner(sql.FieldNotNull(FieldTaxNo))
}

// PhoneEQ applies the EQ predicate on the "Phone" field.
func PhoneEQ(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldEQ(FieldPhone, v))
}

// PhoneNEQ applies the NEQ predicate on the "Phone" field.
func PhoneNEQ(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldNEQ(FieldPhone, v))
}

// PhoneIn applies the In predicate on the "Phone" field.
func PhoneIn(vs ...string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldIn(FieldPhone, vs...))
}

// PhoneNotIn applies the NotIn predicate on the "Phone" field.
func PhoneNotIn(vs ...string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldNotIn(FieldPhone, vs...))
}

// PhoneGT applies the GT predicate on the "Phone" field.
func PhoneGT(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldGT(FieldPhone, v))
}

// PhoneGTE applies the GTE predicate on the "Phone" field.
func PhoneGTE(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldGTE(FieldPhone, v))
}

// PhoneLT applies the LT predicate on the "Phone" field.
func PhoneLT(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldLT(FieldPhone, v))
}

// PhoneLTE applies the LTE predicate on the "Phone" field.
func PhoneLTE(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldLTE(FieldPhone, v))
}

// PhoneContains applies the Contains predicate on the "Phone" field.
func PhoneContains(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldContains(FieldPhone, v))
}

// PhoneHasPrefix applies the HasPrefix predicate on the "Phone" field.
func PhoneHasPrefix(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldHasPrefix(FieldPhone, v))
}

// PhoneHasSuffix applies the HasSuffix predicate on the "Phone" field.
func PhoneHasSuffix(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldHasSuffix(FieldPhone, v))
}

// PhoneIsNil applies the IsNil predicate on the "Phone" field.
func PhoneIsNil() predicate.JobOwner {
	return predicate.JobOwner(sql.FieldIsNull(FieldPhone))
}

// PhoneNotNil applies the NotNil predicate on the "Phone" field.
func PhoneNotNil() predicate.JobOwner {
	return predicate.JobOwner(sql.FieldNotNull(FieldPhone))
}

// PhoneEqualFold applies the EqualFold predicate on the "Phone" field.
func PhoneEqualFold(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldEqualFold(FieldPhone, v))
}

// PhoneContainsFold applies the ContainsFold predicate on the "Phone" field.
func PhoneContainsFold(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldContainsFold(FieldPhone, v))
}

// EmailEQ applies the EQ predicate on the "Email" field.
func EmailEQ(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "Email" field.
func EmailNEQ(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "Email" field.
func EmailIn(vs ...string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "Email" field.
func EmailNotIn(vs ...string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "Email" field.
func EmailGT(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "Email" field.
func EmailGTE(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "Email" field.
func EmailLT(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "Email" field.
func EmailLTE(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "Email" field.
func EmailContains(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "Email" field.
func EmailHasPrefix(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "Email" field.
func EmailHasSuffix(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailIsNil applies the IsNil predicate on the "Email" field.
func EmailIsNil() predicate.JobOwner {
	return predicate.JobOwner(sql.FieldIsNull(FieldEmail))
}

// EmailNotNil applies the NotNil predicate on the "Email" field.
func EmailNotNil() predicate.JobOwner {
	return predicate.JobOwner(sql.FieldNotNull(FieldEmail))
}

// EmailEqualFold applies the EqualFold predicate on the "Email" field.
func EmailEqualFold(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "Email" field.
func EmailContainsFold(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldContainsFold(FieldEmail, v))
}

// NoteEQ applies the EQ predicate on the "Note" field.
func NoteEQ(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldEQ(FieldNote, v))
}

// NoteNEQ applies the NEQ predicate on the "Note" field.
func NoteNEQ(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldNEQ(FieldNote, v))
}

// NoteIn applies the In predicate on the "Note" field.
func NoteIn(vs ...string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldIn(FieldNote, vs...))
}

// NoteNotIn applies the NotIn predicate on the "Note" field.
func NoteNotIn(vs ...string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldNotIn(FieldNote, vs...))
}

// NoteGT applies the GT predicate on the "Note" field.
func NoteGT(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldGT(FieldNote, v))
}

// NoteGTE applies the GTE predicate on the "Note" field.
func NoteGTE(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldGTE(FieldNote, v))
}

// NoteLT applies the LT predicate on the "Note" field.
func NoteLT(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldLT(FieldNote, v))
}

// NoteLTE applies the LTE predicate on the "Note" field.
func NoteLTE(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldLTE(FieldNote, v))
}

// NoteContains applies the Contains predicate on the "Note" field.
func NoteContains(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldContains(FieldNote, v))
}

// NoteHasPrefix applies the HasPrefix predicate on the "Note" field.
func NoteHasPrefix(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldHasPrefix(FieldNote, v))
}

// NoteHasSuffix applies the HasSuffix predicate on the "Note" field.
func NoteHasSuffix(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldHasSuffix(FieldNote, v))
}

// NoteIsNil applies the IsNil predicate on the "Note" field.
func NoteIsNil() predicate.JobOwner {
	return predicate.JobOwner(sql.FieldIsNull(FieldNote))
}

// NoteNotNil applies the NotNil predicate on the "Note" field.
func NoteNotNil() predicate.JobOwner {
	return predicate.JobOwner(sql.FieldNotNull(FieldNote))
}

// NoteEqualFold applies the EqualFold predicate on the "Note" field.
func NoteEqualFold(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldEqualFold(FieldNote, v))
}

// NoteContainsFold applies the ContainsFold predicate on the "Note" field.
func NoteContainsFold(v string) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldContainsFold(FieldNote, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.JobOwner {
	return predicate.JobOwner(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasOwners applies the HasEdge predicate on the "owners" edge.
func HasOwners() predicate.JobOwner {
	return predicate.JobOwner(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OwnersTable, OwnersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnersWith applies the HasEdge predicate on the "owners" edge with a given conditions (other predicates).
func HasOwnersWith(preds ...predicate.JobDetail) predicate.JobOwner {
	return predicate.JobOwner(func(s *sql.Selector) {
		step := newOwnersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.JobOwner) predicate.JobOwner {
	return predicate.JobOwner(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.JobOwner) predicate.JobOwner {
	return predicate.JobOwner(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.JobOwner) predicate.JobOwner {
	return predicate.JobOwner(sql.NotPredicates(p))
}
