package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FinanceCurrentAccountOperationType holds the schema definition for the FinanceCurrentAccountOperationType entity.
type FinanceCurrentAccountOperationType struct {
	ent.Schema
}

// Fields of the FinanceCurrentAccountOperationType.
func (FinanceCurrentAccountOperationType) Fields() []ent.Field {
	return []ent.Field{
		field.String("Description"),
	}
}

// Edges of the FinanceCurrentAccountOperationType.
func (FinanceCurrentAccountOperationType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("current_operations", FinanceCurrentAccountOperation.Type).StorageKey(edge.Column("type_id")),
	}
}

// Annotations of the FinanceCurrentAccountOperationType.
func (FinanceCurrentAccountOperationType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Schema: "finance",
			Table:  "current_account_operation_types",
		},
	}
}
