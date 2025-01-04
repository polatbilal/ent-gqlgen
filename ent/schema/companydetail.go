package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CompanyDetail holds the schema definition for the CompanyDetail entity.
type CompanyDetail struct {
	ent.Schema
}

// Fields of the CompanyDetail.
func (CompanyDetail) Fields() []ent.Field {
	return []ent.Field{
		field.Int("CompanyCode").Unique(),
		field.String("Name").Default(""),
		field.String("Address").Optional(),
		field.String("Phone").Optional(),
		field.String("Email").Optional(),
		field.String("Website").Optional(),
		field.String("TaxAdmin").Optional(),
		field.Int("TaxNo").Default(0).Optional(),
		field.String("ChamberInfo").Optional(),
		field.String("ChamberRegisterNo").Optional(),
		field.Time("VisaDate").Optional(),
		field.Time("VisaEndDate").Optional(),
		field.Bool("VisaFinishedFor90Days").Default(false).Optional(),
		field.Bool("CorePersonAbsent90Days").Default(false).Optional(),
		field.Bool("IsClosed").Default(false),

		field.String("OwnerName").Optional(),
		field.Int("OwnerTcNo").Optional(),
		field.String("OwnerAddress").Optional(),
		field.String("OwnerPhone").Optional(),
		field.String("OwnerEmail").Optional(),
		field.Int("OwnerRegisterNo").Optional(),
		field.String("OwnerCareer").Optional(),

		field.Time("CreatedAt").Default(time.Now),
		field.Time("UpdatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the CompanyDetail.
func (CompanyDetail) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("engineers", CompanyEngineer.Type).StorageKey(edge.Column("company_id")),
		edge.To("users", CompanyUser.Type).StorageKey(edge.Column("company_id")),
		edge.To("jobs", JobDetail.Type).StorageKey(edge.Column("company_id")),
	}
}
