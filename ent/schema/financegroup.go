package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FinanceGroup holds the schema definition for the FinanceGroup entity.
type FinanceGroup struct {
	ent.Schema
}

// Fields of the FinanceGroup.
func (FinanceGroup) Fields() []ent.Field {
	return []ent.Field{
		field.String("Category"),
		field.String("Name"),
		field.String("DeletedName").Optional(),
		field.Time("DeletedDate").Optional(),
		field.Bool("Status").Default(true),

		field.Time("createdAt").Default(time.Now),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the FinanceGroup.
func (FinanceGroup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("groups", FinanceOperation.Type).StorageKey(edge.Column("group_id")),
		edge.To("finance_accounts", FinanceAccount.Type).StorageKey(edge.Column("group_id")),
	}
}

// Annotations of the FinanceGroup.
func (FinanceGroup) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Schema: "finance",
			Table:  "groups",
		},
	}
}
