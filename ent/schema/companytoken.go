package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CompanyToken holds the schema definition for the CompanyToken entity.
type CompanyToken struct {
	ent.Schema
}

// Fields of the CompanyToken.
func (CompanyToken) Fields() []ent.Field {
	return []ent.Field{
		field.Text("token").Optional(),
	}
}

// Edges of the CompanyToken.
func (CompanyToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", CompanyDetail.Type).Ref("tokens").Unique(),
	}
}
