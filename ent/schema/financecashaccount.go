package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FinanceCashAccount holds the schema definition for the FinanceCashAccount entity.
type FinanceCashAccount struct {
	ent.Schema
}

// Fields of the FinanceCashAccount.
func (FinanceCashAccount) Fields() []ent.Field {
	return []ent.Field{
		field.String("Description"),
		field.String("DeletedDescription").Optional(),
		field.Time("DeletedDate").Optional(),
		field.Bool("Status").Default(true),
	}
}

// Edges of the FinanceCashAccount.
func (FinanceCashAccount) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", CompanyDetail.Type).Ref("cash_accounts").Unique(),
		edge.To("operations", FinanceCashAccountOperation.Type).StorageKey(edge.Column("cash_account_id")),
	}
}

// Annotations of the FinanceCashAccount.
func (FinanceCashAccount) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Schema: "finance",
			Table:  "cash_accounts",
		},
	}
}
