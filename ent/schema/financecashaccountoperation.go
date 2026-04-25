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

// FinanceCashAccountOperation holds the schema definition for the FinanceCashAccountOperation entity.
type FinanceCashAccountOperation struct {
	ent.Schema
}

// Fields of the FinanceCashAccountOperation.
func (FinanceCashAccountOperation) Fields() []ent.Field {
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

// Edges of the FinanceCashAccountOperation.
func (FinanceCashAccountOperation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", CompanyDetail.Type).Ref("cash_operations").Unique(),
		edge.From("cash_account", FinanceCashAccount.Type).Ref("operations").Unique(),
		edge.From("operation_group", FinanceOperationGroup.Type).Ref("cash_operations").Unique(),
		edge.From("payment_class", FinancePaymentClass.Type).Ref("cash_payments").Unique(),
		edge.To("participant_current_operation", FinanceCurrentAccountOperation.Type).Unique().StorageKey(edge.Column("participant_current_account_operation_id")),
		edge.To("participant_bank_operation", FinanceBankOperation.Type).Unique().StorageKey(edge.Column("participant_bank_account_operation_id")),
		edge.From("bank_participant", FinanceBankOperation.Type).Ref("participant_cash_operation"),
	}
}

// Annotations of the FinanceCashAccountOperation.
func (FinanceCashAccountOperation) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Schema: "finance",
			Table:  "cash_account_operations",
		},
	}
}
