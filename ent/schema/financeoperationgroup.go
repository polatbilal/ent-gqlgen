package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FinanceOperationGroup holds the schema definition for the FinanceOperationGroup entity.
type FinanceOperationGroup struct {
	ent.Schema
}

// Fields of the FinanceOperationGroup.
func (FinanceOperationGroup) Fields() []ent.Field {
	return []ent.Field{
		field.String("Description"),
	}
}

// Edges of the FinanceOperationGroup.
func (FinanceOperationGroup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("bank_operations", FinanceBankOperation.Type).StorageKey(edge.Column("operation_group_id")),
		edge.To("cash_operations", FinanceCashAccountOperation.Type).StorageKey(edge.Column("operation_group_id")),
		edge.To("current_operations", FinanceCurrentAccountOperation.Type).StorageKey(edge.Column("operation_group_id")),
	}
}

// Annotations of the FinanceOperationGroup.
func (FinanceOperationGroup) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Schema: "finance",
			Table:  "operation_groups",
		},
	}
}
