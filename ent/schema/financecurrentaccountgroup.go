package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FinanceCurrentAccountGroup holds the schema definition for the FinanceCurrentAccountGroup entity.
type FinanceCurrentAccountGroup struct {
	ent.Schema
}

// Fields of the FinanceCurrentAccountGroup.
func (FinanceCurrentAccountGroup) Fields() []ent.Field {
	return []ent.Field{
		field.String("Description"),
	}
}

// Edges of the FinanceCurrentAccountGroup.
func (FinanceCurrentAccountGroup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("current_accounts", FinanceCurrentAccount.Type).StorageKey(edge.Column("group_id")),
	}
}

// Annotations of the FinanceCurrentAccountGroup.
func (FinanceCurrentAccountGroup) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Schema: "finance",
			Table:  "current_account_groups",
		},
	}
}
