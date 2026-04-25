package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/shopspring/decimal"
)

// FinanceCurrentAccountOperation holds the schema definition for the FinanceCurrentAccountOperation entity.
type FinanceCurrentAccountOperation struct {
	ent.Schema
}

// Fields of the FinanceCurrentAccountOperation.
func (FinanceCurrentAccountOperation) Fields() []ent.Field {
	return []ent.Field{
		field.Time("Date"),
		field.
			Other("Debit", &decimal.NullDecimal{}).
			Optional().
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(18,2)",
			}),
		field.
			Other("Credit", &decimal.NullDecimal{}).
			Optional().
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(18,2)",
			}),
		field.String("Description"),
		field.Bool("Status").Default(true),
	}
}

// Edges of the FinanceCurrentAccountOperation.
func (FinanceCurrentAccountOperation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", CompanyDetail.Type).Ref("current_operations").Unique(),
		edge.From("current_account", FinanceCurrentAccount.Type).Ref("operations").Unique(),
		edge.From("operation_group", FinanceOperationGroup.Type).Ref("current_operations").Unique(),
		edge.From("account_type", FinanceCurrentAccountOperationType.Type).Ref("current_operations").Unique(),
		edge.From("bank_participant", FinanceBankOperation.Type).Ref("participant_current_operation"),
		edge.From("cash_participant", FinanceCashAccountOperation.Type).Ref("participant_current_operation"),
	}
}

// Annotations of the FinanceCurrentAccountOperation.
func (FinanceCurrentAccountOperation) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Schema: "finance",
			Table:  "current_account_operations",
		},
	}
}
