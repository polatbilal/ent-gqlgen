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

// FinanceBankAccountOperation holds the schema definition for the FinanceBankAccountOperation entity.
type FinanceBankOperation struct {
	ent.Schema
}

// Fields of the FinanceBankAccountOperation.
func (FinanceBankOperation) Fields() []ent.Field {
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

// Edges of the FinanceBankAccountOperation.
func (FinanceBankOperation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", CompanyDetail.Type).Ref("bank_operations").Unique(),
		edge.From("bank_account", FinanceBankAccount.Type).Ref("operations").Unique(),
		edge.From("operation_group", FinanceOperationGroup.Type).Ref("bank_operations").Unique(),
		edge.From("payment_class", FinancePaymentClass.Type).Ref("bank_payments").Unique(),
		edge.To("participant_current_operation", FinanceCurrentAccountOperation.Type).Unique().StorageKey(edge.Column("participant_current_account_operation_id")),
		edge.To("participant_cash_operation", FinanceCashAccountOperation.Type).Unique().StorageKey(edge.Column("participant_cash_account_operation_id")),
		edge.From("cash_participant", FinanceCashAccountOperation.Type).Ref("participant_bank_operation"),
	}
}

// Annotations of the FinanceBankAccountOperation.
func (FinanceBankOperation) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Schema: "finance",
			Table:  "bank_account_operations",
		},
	}
}
