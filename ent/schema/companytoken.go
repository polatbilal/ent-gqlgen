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
		field.String("Token").Optional(),
		field.Int("DepartmentId").Optional().Unique(),
		field.Int("Expire").Optional(),
		field.String("RefreshToken").Optional(),
		field.String("SecretKey").Optional(),
		field.String("SecureSecretKey").Optional(),
		field.String("OtpUri").Optional(),

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
