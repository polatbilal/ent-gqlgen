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
		field.Int("CompanyCode"),
		field.String("Name").Default(""),
		field.String("Address").Optional(),
		field.String("City").Optional(),
		field.String("State").Optional(),
		field.String("Phone").Optional(),
		field.String("Fax").Optional(),
		field.String("Mobile").Optional(),
		field.String("Email").Optional(),
		field.String("Website").Optional(),
		field.String("TaxAdmin").Optional(),
		field.Int("TaxNo").Default(0).Optional(),
		field.String("Commerce").Optional(),
		field.String("CommerceReg").Optional(),
		field.Time("VisaDate").Optional(),

		field.Time("CreatedAt").Default(time.Now),
		field.Time("UpdatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the CompanyDetail.
func (CompanyDetail) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("companyOwner", CompanyEngineer.Type).Ref("companyOwners").Unique(),
		edge.To("engineers", CompanyEngineer.Type).StorageKey(edge.Column("company_id")),
		edge.To("users", CompanyUser.Type).StorageKey(edge.Column("company_id")),
		edge.To("jobs", JobDetail.Type).StorageKey(edge.Column("company_id")),
	}
}
