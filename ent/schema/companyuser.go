package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// CompanyUser holds the schema definition for the CompanyUser entity.
type CompanyUser struct {
	ent.Schema
}

// Fields of the CompanyUser.
func (CompanyUser) Fields() []ent.Field {
	return nil
}

// Edges of the CompanyUser.
func (CompanyUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", CompanyDetail.Type).Ref("users").Unique(),
		edge.From("user", User.Type).Ref("companies").Unique(),
	}
}
