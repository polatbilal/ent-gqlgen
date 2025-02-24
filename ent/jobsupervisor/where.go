// Code generated by ent, DO NOT EDIT.

package jobsupervisor

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/polatbilal/gqlgen-ent/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "Name" field. It's identical to NameEQ.
func Name(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldEQ(FieldName, v))
}

// Address applies equality check predicate on the "Address" field. It's identical to AddressEQ.
func Address(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldEQ(FieldAddress, v))
}

// Phone applies equality check predicate on the "Phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldEQ(FieldPhone, v))
}

// Email applies equality check predicate on the "Email" field. It's identical to EmailEQ.
func Email(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldEQ(FieldEmail, v))
}

// TCNO applies equality check predicate on the "TCNO" field. It's identical to TCNOEQ.
func TCNO(v int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldEQ(FieldTCNO, v))
}

// Position applies equality check predicate on the "Position" field. It's identical to PositionEQ.
func Position(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldEQ(FieldPosition, v))
}

// Career applies equality check predicate on the "Career" field. It's identical to CareerEQ.
func Career(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldEQ(FieldCareer, v))
}

// RegNo applies equality check predicate on the "RegNo" field. It's identical to RegNoEQ.
func RegNo(v int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldEQ(FieldRegNo, v))
}

// SocialSecurityNo applies equality check predicate on the "SocialSecurityNo" field. It's identical to SocialSecurityNoEQ.
func SocialSecurityNo(v int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldEQ(FieldSocialSecurityNo, v))
}

// SchoolGraduation applies equality check predicate on the "SchoolGraduation" field. It's identical to SchoolGraduationEQ.
func SchoolGraduation(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldEQ(FieldSchoolGraduation, v))
}

// YDSID applies equality check predicate on the "YDSID" field. It's identical to YDSIDEQ.
func YDSID(v int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldEQ(FieldYDSID, v))
}

// CreatedAt applies equality check predicate on the "CreatedAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "UpdatedAt" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldEQ(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "Name" field.
func NameEQ(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "Name" field.
func NameNEQ(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "Name" field.
func NameIn(vs ...string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "Name" field.
func NameNotIn(vs ...string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "Name" field.
func NameGT(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "Name" field.
func NameGTE(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "Name" field.
func NameLT(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "Name" field.
func NameLTE(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "Name" field.
func NameContains(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "Name" field.
func NameHasPrefix(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "Name" field.
func NameHasSuffix(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldHasSuffix(FieldName, v))
}

// NameIsNil applies the IsNil predicate on the "Name" field.
func NameIsNil() predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldIsNull(FieldName))
}

// NameNotNil applies the NotNil predicate on the "Name" field.
func NameNotNil() predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldNotNull(FieldName))
}

// NameEqualFold applies the EqualFold predicate on the "Name" field.
func NameEqualFold(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "Name" field.
func NameContainsFold(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldContainsFold(FieldName, v))
}

// AddressEQ applies the EQ predicate on the "Address" field.
func AddressEQ(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldEQ(FieldAddress, v))
}

// AddressNEQ applies the NEQ predicate on the "Address" field.
func AddressNEQ(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldNEQ(FieldAddress, v))
}

// AddressIn applies the In predicate on the "Address" field.
func AddressIn(vs ...string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldIn(FieldAddress, vs...))
}

// AddressNotIn applies the NotIn predicate on the "Address" field.
func AddressNotIn(vs ...string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldNotIn(FieldAddress, vs...))
}

// AddressGT applies the GT predicate on the "Address" field.
func AddressGT(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldGT(FieldAddress, v))
}

// AddressGTE applies the GTE predicate on the "Address" field.
func AddressGTE(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldGTE(FieldAddress, v))
}

// AddressLT applies the LT predicate on the "Address" field.
func AddressLT(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldLT(FieldAddress, v))
}

// AddressLTE applies the LTE predicate on the "Address" field.
func AddressLTE(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldLTE(FieldAddress, v))
}

// AddressContains applies the Contains predicate on the "Address" field.
func AddressContains(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldContains(FieldAddress, v))
}

// AddressHasPrefix applies the HasPrefix predicate on the "Address" field.
func AddressHasPrefix(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldHasPrefix(FieldAddress, v))
}

// AddressHasSuffix applies the HasSuffix predicate on the "Address" field.
func AddressHasSuffix(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldHasSuffix(FieldAddress, v))
}

// AddressIsNil applies the IsNil predicate on the "Address" field.
func AddressIsNil() predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldIsNull(FieldAddress))
}

// AddressNotNil applies the NotNil predicate on the "Address" field.
func AddressNotNil() predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldNotNull(FieldAddress))
}

// AddressEqualFold applies the EqualFold predicate on the "Address" field.
func AddressEqualFold(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldEqualFold(FieldAddress, v))
}

// AddressContainsFold applies the ContainsFold predicate on the "Address" field.
func AddressContainsFold(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldContainsFold(FieldAddress, v))
}

// PhoneEQ applies the EQ predicate on the "Phone" field.
func PhoneEQ(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldEQ(FieldPhone, v))
}

// PhoneNEQ applies the NEQ predicate on the "Phone" field.
func PhoneNEQ(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldNEQ(FieldPhone, v))
}

// PhoneIn applies the In predicate on the "Phone" field.
func PhoneIn(vs ...string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldIn(FieldPhone, vs...))
}

// PhoneNotIn applies the NotIn predicate on the "Phone" field.
func PhoneNotIn(vs ...string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldNotIn(FieldPhone, vs...))
}

// PhoneGT applies the GT predicate on the "Phone" field.
func PhoneGT(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldGT(FieldPhone, v))
}

// PhoneGTE applies the GTE predicate on the "Phone" field.
func PhoneGTE(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldGTE(FieldPhone, v))
}

// PhoneLT applies the LT predicate on the "Phone" field.
func PhoneLT(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldLT(FieldPhone, v))
}

// PhoneLTE applies the LTE predicate on the "Phone" field.
func PhoneLTE(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldLTE(FieldPhone, v))
}

// PhoneContains applies the Contains predicate on the "Phone" field.
func PhoneContains(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldContains(FieldPhone, v))
}

// PhoneHasPrefix applies the HasPrefix predicate on the "Phone" field.
func PhoneHasPrefix(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldHasPrefix(FieldPhone, v))
}

// PhoneHasSuffix applies the HasSuffix predicate on the "Phone" field.
func PhoneHasSuffix(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldHasSuffix(FieldPhone, v))
}

// PhoneIsNil applies the IsNil predicate on the "Phone" field.
func PhoneIsNil() predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldIsNull(FieldPhone))
}

// PhoneNotNil applies the NotNil predicate on the "Phone" field.
func PhoneNotNil() predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldNotNull(FieldPhone))
}

// PhoneEqualFold applies the EqualFold predicate on the "Phone" field.
func PhoneEqualFold(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldEqualFold(FieldPhone, v))
}

// PhoneContainsFold applies the ContainsFold predicate on the "Phone" field.
func PhoneContainsFold(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldContainsFold(FieldPhone, v))
}

// EmailEQ applies the EQ predicate on the "Email" field.
func EmailEQ(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "Email" field.
func EmailNEQ(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "Email" field.
func EmailIn(vs ...string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "Email" field.
func EmailNotIn(vs ...string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "Email" field.
func EmailGT(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "Email" field.
func EmailGTE(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "Email" field.
func EmailLT(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "Email" field.
func EmailLTE(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "Email" field.
func EmailContains(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "Email" field.
func EmailHasPrefix(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "Email" field.
func EmailHasSuffix(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailIsNil applies the IsNil predicate on the "Email" field.
func EmailIsNil() predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldIsNull(FieldEmail))
}

// EmailNotNil applies the NotNil predicate on the "Email" field.
func EmailNotNil() predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldNotNull(FieldEmail))
}

// EmailEqualFold applies the EqualFold predicate on the "Email" field.
func EmailEqualFold(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "Email" field.
func EmailContainsFold(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldContainsFold(FieldEmail, v))
}

// TCNOEQ applies the EQ predicate on the "TCNO" field.
func TCNOEQ(v int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldEQ(FieldTCNO, v))
}

// TCNONEQ applies the NEQ predicate on the "TCNO" field.
func TCNONEQ(v int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldNEQ(FieldTCNO, v))
}

// TCNOIn applies the In predicate on the "TCNO" field.
func TCNOIn(vs ...int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldIn(FieldTCNO, vs...))
}

// TCNONotIn applies the NotIn predicate on the "TCNO" field.
func TCNONotIn(vs ...int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldNotIn(FieldTCNO, vs...))
}

// TCNOGT applies the GT predicate on the "TCNO" field.
func TCNOGT(v int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldGT(FieldTCNO, v))
}

// TCNOGTE applies the GTE predicate on the "TCNO" field.
func TCNOGTE(v int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldGTE(FieldTCNO, v))
}

// TCNOLT applies the LT predicate on the "TCNO" field.
func TCNOLT(v int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldLT(FieldTCNO, v))
}

// TCNOLTE applies the LTE predicate on the "TCNO" field.
func TCNOLTE(v int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldLTE(FieldTCNO, v))
}

// TCNOIsNil applies the IsNil predicate on the "TCNO" field.
func TCNOIsNil() predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldIsNull(FieldTCNO))
}

// TCNONotNil applies the NotNil predicate on the "TCNO" field.
func TCNONotNil() predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldNotNull(FieldTCNO))
}

// PositionEQ applies the EQ predicate on the "Position" field.
func PositionEQ(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldEQ(FieldPosition, v))
}

// PositionNEQ applies the NEQ predicate on the "Position" field.
func PositionNEQ(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldNEQ(FieldPosition, v))
}

// PositionIn applies the In predicate on the "Position" field.
func PositionIn(vs ...string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldIn(FieldPosition, vs...))
}

// PositionNotIn applies the NotIn predicate on the "Position" field.
func PositionNotIn(vs ...string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldNotIn(FieldPosition, vs...))
}

// PositionGT applies the GT predicate on the "Position" field.
func PositionGT(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldGT(FieldPosition, v))
}

// PositionGTE applies the GTE predicate on the "Position" field.
func PositionGTE(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldGTE(FieldPosition, v))
}

// PositionLT applies the LT predicate on the "Position" field.
func PositionLT(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldLT(FieldPosition, v))
}

// PositionLTE applies the LTE predicate on the "Position" field.
func PositionLTE(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldLTE(FieldPosition, v))
}

// PositionContains applies the Contains predicate on the "Position" field.
func PositionContains(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldContains(FieldPosition, v))
}

// PositionHasPrefix applies the HasPrefix predicate on the "Position" field.
func PositionHasPrefix(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldHasPrefix(FieldPosition, v))
}

// PositionHasSuffix applies the HasSuffix predicate on the "Position" field.
func PositionHasSuffix(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldHasSuffix(FieldPosition, v))
}

// PositionIsNil applies the IsNil predicate on the "Position" field.
func PositionIsNil() predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldIsNull(FieldPosition))
}

// PositionNotNil applies the NotNil predicate on the "Position" field.
func PositionNotNil() predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldNotNull(FieldPosition))
}

// PositionEqualFold applies the EqualFold predicate on the "Position" field.
func PositionEqualFold(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldEqualFold(FieldPosition, v))
}

// PositionContainsFold applies the ContainsFold predicate on the "Position" field.
func PositionContainsFold(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldContainsFold(FieldPosition, v))
}

// CareerEQ applies the EQ predicate on the "Career" field.
func CareerEQ(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldEQ(FieldCareer, v))
}

// CareerNEQ applies the NEQ predicate on the "Career" field.
func CareerNEQ(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldNEQ(FieldCareer, v))
}

// CareerIn applies the In predicate on the "Career" field.
func CareerIn(vs ...string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldIn(FieldCareer, vs...))
}

// CareerNotIn applies the NotIn predicate on the "Career" field.
func CareerNotIn(vs ...string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldNotIn(FieldCareer, vs...))
}

// CareerGT applies the GT predicate on the "Career" field.
func CareerGT(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldGT(FieldCareer, v))
}

// CareerGTE applies the GTE predicate on the "Career" field.
func CareerGTE(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldGTE(FieldCareer, v))
}

// CareerLT applies the LT predicate on the "Career" field.
func CareerLT(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldLT(FieldCareer, v))
}

// CareerLTE applies the LTE predicate on the "Career" field.
func CareerLTE(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldLTE(FieldCareer, v))
}

// CareerContains applies the Contains predicate on the "Career" field.
func CareerContains(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldContains(FieldCareer, v))
}

// CareerHasPrefix applies the HasPrefix predicate on the "Career" field.
func CareerHasPrefix(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldHasPrefix(FieldCareer, v))
}

// CareerHasSuffix applies the HasSuffix predicate on the "Career" field.
func CareerHasSuffix(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldHasSuffix(FieldCareer, v))
}

// CareerIsNil applies the IsNil predicate on the "Career" field.
func CareerIsNil() predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldIsNull(FieldCareer))
}

// CareerNotNil applies the NotNil predicate on the "Career" field.
func CareerNotNil() predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldNotNull(FieldCareer))
}

// CareerEqualFold applies the EqualFold predicate on the "Career" field.
func CareerEqualFold(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldEqualFold(FieldCareer, v))
}

// CareerContainsFold applies the ContainsFold predicate on the "Career" field.
func CareerContainsFold(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldContainsFold(FieldCareer, v))
}

// RegNoEQ applies the EQ predicate on the "RegNo" field.
func RegNoEQ(v int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldEQ(FieldRegNo, v))
}

// RegNoNEQ applies the NEQ predicate on the "RegNo" field.
func RegNoNEQ(v int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldNEQ(FieldRegNo, v))
}

// RegNoIn applies the In predicate on the "RegNo" field.
func RegNoIn(vs ...int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldIn(FieldRegNo, vs...))
}

// RegNoNotIn applies the NotIn predicate on the "RegNo" field.
func RegNoNotIn(vs ...int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldNotIn(FieldRegNo, vs...))
}

// RegNoGT applies the GT predicate on the "RegNo" field.
func RegNoGT(v int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldGT(FieldRegNo, v))
}

// RegNoGTE applies the GTE predicate on the "RegNo" field.
func RegNoGTE(v int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldGTE(FieldRegNo, v))
}

// RegNoLT applies the LT predicate on the "RegNo" field.
func RegNoLT(v int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldLT(FieldRegNo, v))
}

// RegNoLTE applies the LTE predicate on the "RegNo" field.
func RegNoLTE(v int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldLTE(FieldRegNo, v))
}

// RegNoIsNil applies the IsNil predicate on the "RegNo" field.
func RegNoIsNil() predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldIsNull(FieldRegNo))
}

// RegNoNotNil applies the NotNil predicate on the "RegNo" field.
func RegNoNotNil() predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldNotNull(FieldRegNo))
}

// SocialSecurityNoEQ applies the EQ predicate on the "SocialSecurityNo" field.
func SocialSecurityNoEQ(v int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldEQ(FieldSocialSecurityNo, v))
}

// SocialSecurityNoNEQ applies the NEQ predicate on the "SocialSecurityNo" field.
func SocialSecurityNoNEQ(v int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldNEQ(FieldSocialSecurityNo, v))
}

// SocialSecurityNoIn applies the In predicate on the "SocialSecurityNo" field.
func SocialSecurityNoIn(vs ...int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldIn(FieldSocialSecurityNo, vs...))
}

// SocialSecurityNoNotIn applies the NotIn predicate on the "SocialSecurityNo" field.
func SocialSecurityNoNotIn(vs ...int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldNotIn(FieldSocialSecurityNo, vs...))
}

// SocialSecurityNoGT applies the GT predicate on the "SocialSecurityNo" field.
func SocialSecurityNoGT(v int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldGT(FieldSocialSecurityNo, v))
}

// SocialSecurityNoGTE applies the GTE predicate on the "SocialSecurityNo" field.
func SocialSecurityNoGTE(v int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldGTE(FieldSocialSecurityNo, v))
}

// SocialSecurityNoLT applies the LT predicate on the "SocialSecurityNo" field.
func SocialSecurityNoLT(v int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldLT(FieldSocialSecurityNo, v))
}

// SocialSecurityNoLTE applies the LTE predicate on the "SocialSecurityNo" field.
func SocialSecurityNoLTE(v int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldLTE(FieldSocialSecurityNo, v))
}

// SocialSecurityNoIsNil applies the IsNil predicate on the "SocialSecurityNo" field.
func SocialSecurityNoIsNil() predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldIsNull(FieldSocialSecurityNo))
}

// SocialSecurityNoNotNil applies the NotNil predicate on the "SocialSecurityNo" field.
func SocialSecurityNoNotNil() predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldNotNull(FieldSocialSecurityNo))
}

// SchoolGraduationEQ applies the EQ predicate on the "SchoolGraduation" field.
func SchoolGraduationEQ(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldEQ(FieldSchoolGraduation, v))
}

// SchoolGraduationNEQ applies the NEQ predicate on the "SchoolGraduation" field.
func SchoolGraduationNEQ(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldNEQ(FieldSchoolGraduation, v))
}

// SchoolGraduationIn applies the In predicate on the "SchoolGraduation" field.
func SchoolGraduationIn(vs ...string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldIn(FieldSchoolGraduation, vs...))
}

// SchoolGraduationNotIn applies the NotIn predicate on the "SchoolGraduation" field.
func SchoolGraduationNotIn(vs ...string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldNotIn(FieldSchoolGraduation, vs...))
}

// SchoolGraduationGT applies the GT predicate on the "SchoolGraduation" field.
func SchoolGraduationGT(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldGT(FieldSchoolGraduation, v))
}

// SchoolGraduationGTE applies the GTE predicate on the "SchoolGraduation" field.
func SchoolGraduationGTE(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldGTE(FieldSchoolGraduation, v))
}

// SchoolGraduationLT applies the LT predicate on the "SchoolGraduation" field.
func SchoolGraduationLT(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldLT(FieldSchoolGraduation, v))
}

// SchoolGraduationLTE applies the LTE predicate on the "SchoolGraduation" field.
func SchoolGraduationLTE(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldLTE(FieldSchoolGraduation, v))
}

// SchoolGraduationContains applies the Contains predicate on the "SchoolGraduation" field.
func SchoolGraduationContains(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldContains(FieldSchoolGraduation, v))
}

// SchoolGraduationHasPrefix applies the HasPrefix predicate on the "SchoolGraduation" field.
func SchoolGraduationHasPrefix(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldHasPrefix(FieldSchoolGraduation, v))
}

// SchoolGraduationHasSuffix applies the HasSuffix predicate on the "SchoolGraduation" field.
func SchoolGraduationHasSuffix(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldHasSuffix(FieldSchoolGraduation, v))
}

// SchoolGraduationIsNil applies the IsNil predicate on the "SchoolGraduation" field.
func SchoolGraduationIsNil() predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldIsNull(FieldSchoolGraduation))
}

// SchoolGraduationNotNil applies the NotNil predicate on the "SchoolGraduation" field.
func SchoolGraduationNotNil() predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldNotNull(FieldSchoolGraduation))
}

// SchoolGraduationEqualFold applies the EqualFold predicate on the "SchoolGraduation" field.
func SchoolGraduationEqualFold(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldEqualFold(FieldSchoolGraduation, v))
}

// SchoolGraduationContainsFold applies the ContainsFold predicate on the "SchoolGraduation" field.
func SchoolGraduationContainsFold(v string) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldContainsFold(FieldSchoolGraduation, v))
}

// YDSIDEQ applies the EQ predicate on the "YDSID" field.
func YDSIDEQ(v int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldEQ(FieldYDSID, v))
}

// YDSIDNEQ applies the NEQ predicate on the "YDSID" field.
func YDSIDNEQ(v int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldNEQ(FieldYDSID, v))
}

// YDSIDIn applies the In predicate on the "YDSID" field.
func YDSIDIn(vs ...int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldIn(FieldYDSID, vs...))
}

// YDSIDNotIn applies the NotIn predicate on the "YDSID" field.
func YDSIDNotIn(vs ...int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldNotIn(FieldYDSID, vs...))
}

// YDSIDGT applies the GT predicate on the "YDSID" field.
func YDSIDGT(v int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldGT(FieldYDSID, v))
}

// YDSIDGTE applies the GTE predicate on the "YDSID" field.
func YDSIDGTE(v int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldGTE(FieldYDSID, v))
}

// YDSIDLT applies the LT predicate on the "YDSID" field.
func YDSIDLT(v int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldLT(FieldYDSID, v))
}

// YDSIDLTE applies the LTE predicate on the "YDSID" field.
func YDSIDLTE(v int) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldLTE(FieldYDSID, v))
}

// YDSIDIsNil applies the IsNil predicate on the "YDSID" field.
func YDSIDIsNil() predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldIsNull(FieldYDSID))
}

// YDSIDNotNil applies the NotNil predicate on the "YDSID" field.
func YDSIDNotNil() predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldNotNull(FieldYDSID))
}

// CreatedAtEQ applies the EQ predicate on the "CreatedAt" field.
func CreatedAtEQ(v time.Time) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "CreatedAt" field.
func CreatedAtNEQ(v time.Time) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "CreatedAt" field.
func CreatedAtIn(vs ...time.Time) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "CreatedAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "CreatedAt" field.
func CreatedAtGT(v time.Time) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "CreatedAt" field.
func CreatedAtGTE(v time.Time) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "CreatedAt" field.
func CreatedAtLT(v time.Time) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "CreatedAt" field.
func CreatedAtLTE(v time.Time) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "UpdatedAt" field.
func UpdatedAtEQ(v time.Time) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "UpdatedAt" field.
func UpdatedAtNEQ(v time.Time) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "UpdatedAt" field.
func UpdatedAtIn(vs ...time.Time) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "UpdatedAt" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "UpdatedAt" field.
func UpdatedAtGT(v time.Time) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "UpdatedAt" field.
func UpdatedAtGTE(v time.Time) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "UpdatedAt" field.
func UpdatedAtLT(v time.Time) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "UpdatedAt" field.
func UpdatedAtLTE(v time.Time) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasSupervisors applies the HasEdge predicate on the "supervisors" edge.
func HasSupervisors() predicate.JobSuperVisor {
	return predicate.JobSuperVisor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SupervisorsTable, SupervisorsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSupervisorsWith applies the HasEdge predicate on the "supervisors" edge with a given conditions (other predicates).
func HasSupervisorsWith(preds ...predicate.JobDetail) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(func(s *sql.Selector) {
		step := newSupervisorsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.JobSuperVisor) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.JobSuperVisor) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.JobSuperVisor) predicate.JobSuperVisor {
	return predicate.JobSuperVisor(sql.NotPredicates(p))
}
