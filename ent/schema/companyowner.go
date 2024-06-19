package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CompanyOwner holds the schema definition for the CompanyOwner entity.
type CompanyOwner struct {
	ent.Schema
}

// Fields of the CompanyOwner.
func (CompanyOwner) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name").Default(""),
		field.Int("RegisterNo").Optional(),
		field.Int("CertNo").Optional(),
		field.Int("Deleted").Default(0),

		field.Time("CreatedAt").Default(time.Now),
		field.Time("UpdatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the CompanyOwner.
func (CompanyOwner) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("companyOwners", CompanyDetail.Type).StorageKey(edge.Column("owner_id")),
		edge.From("companyOwnerCareer", CompanyCareer.Type).Ref("companyOwnerCareers").Unique(),
		edge.From("companyOwnerPosition", CompanyPosition.Type).Ref("companyOwnerPositions").Unique(),
	}
}
