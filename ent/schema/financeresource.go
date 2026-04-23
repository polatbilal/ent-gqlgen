package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FinanceResource holds the schema definition for the FinanceResource entity.
type FinanceResource struct {
	ent.Schema
}

// Fields of the FinanceResource.
func (FinanceResource) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name"),

		field.Time("createdAt").Default(time.Now),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the FinanceResource.
func (FinanceResource) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", CompanyDetail.Type).Ref("resources").Unique(),
		edge.To("resources", FinanceOperation.Type).StorageKey(edge.Column("resource_id")),
	}
}

// Annotations of the FinanceResource.
func (FinanceResource) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Schema: "finance",
			Table:  "resources",
		},
	}
}
