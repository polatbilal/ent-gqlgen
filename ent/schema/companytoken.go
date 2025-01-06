package schema

import (
	"time"

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
		field.Int("department_id").Optional(),
		field.Time("expire_date").Optional(),

		field.Time("createdAt").Default(time.Now),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the CompanyToken.
func (CompanyToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", CompanyDetail.Type).Ref("tokens").Unique(),
	}
}
