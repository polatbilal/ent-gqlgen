package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CompanyCareer holds the schema definition for the CompanyCareer entity.
type CompanyCareer struct {
	ent.Schema
}

// Fields of the CompanyCareer.
func (CompanyCareer) Fields() []ent.Field {
	return []ent.Field{
		field.String("Career").Optional(),
	}
}

// Edges of the CompanyCareer.
func (CompanyCareer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("engineerCareers", CompanyEngineer.Type).StorageKey(edge.Column("career_id")),
		edge.To("companyOwnerCareers", CompanyOwner.Type).StorageKey(edge.Column("career_id")),
	}
}
