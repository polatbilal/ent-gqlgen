package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FinanceClass holds the schema definition for the FinanceClass entity.
type FinanceClass struct {
	ent.Schema
}

// Fields of the FinanceClass.
func (FinanceClass) Fields() []ent.Field {
	return []ent.Field{
		field.String("Category").Optional(),
		field.String("Name"),
		field.String("DeletedName").Optional(),
		field.Time("DeletedDate").Optional(),
		field.Bool("Status").Default(true),

		field.Time("createdAt").Default(time.Now),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the FinanceClass.
func (FinanceClass) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", CompanyDetail.Type).Ref("methods").Unique(),
		edge.To("classes", FinanceOperation.Type).StorageKey(edge.Column("class_id")),
	}
}
