package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// CompanyOwner holds the schema definition for the CompanyOwner entity.
type CompanyOwner struct {
	ent.Schema
}

// Fields of the CompanyOwner.
func (CompanyOwner) Fields() []ent.Field {
	return nil
}

// Edges of the CompanyOwner.
func (CompanyOwner) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", CompanyDetail.Type).Ref("owners").Unique(),
		edge.From("engineer", CompanyEngineer.Type).Ref("owners").Unique(),
	}
}
