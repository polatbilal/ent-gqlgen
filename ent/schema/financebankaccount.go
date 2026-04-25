package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FinanceBankAccount holds the schema definition for the FinanceBankAccount entity.
type FinanceBankAccount struct {
	ent.Schema
}

// Fields of the FinanceBankAccount.
func (FinanceBankAccount) Fields() []ent.Field {
	return []ent.Field{
		field.String("Description"),
		field.String("DeletedDescription").Optional(),
		field.Time("DeletedDate").Optional(),
		field.Bool("Status").Default(true),
	}
}

// Edges of the FinanceBankAccount.
func (FinanceBankAccount) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", CompanyDetail.Type).Ref("bank_accounts").Unique(),
		edge.To("operations", FinanceBankOperation.Type).StorageKey(edge.Column("bank_account_id")),
	}
}

// Annotations of the FinanceBankAccount.
func (FinanceBankAccount) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Schema: "finance",
			Table:  "bank_accounts",
		},
	}
}
