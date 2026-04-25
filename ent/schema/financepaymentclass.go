package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FinancePaymentClass holds the schema definition for the FinancePaymentClass entity.
type FinancePaymentClass struct {
	ent.Schema
}

// Fields of the FinancePaymentClass.
func (FinancePaymentClass) Fields() []ent.Field {
	return []ent.Field{
		field.String("Description"),
		field.String("DeletedDescription").Optional(),
		field.Time("DeletedDate").Optional(),
		field.Bool("Status").Default(true),
	}
}

// Edges of the FinancePaymentClass.
func (FinancePaymentClass) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", CompanyDetail.Type).Ref("cash_payment_classes").Unique(),
		edge.To("bank_payments", FinanceBankOperation.Type).StorageKey(edge.Column("payment_class_id")),
		edge.To("cash_payments", FinanceCashAccountOperation.Type).StorageKey(edge.Column("payment_class_id")),
	}
}

// Annotations of the FinancePaymentClass.
func (FinancePaymentClass) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Schema: "finance",
			Table:  "payment_classes",
		},
	}
}
