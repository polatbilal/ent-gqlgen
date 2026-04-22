package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FinanceAccount holds the schema definition for the FinanceAccount entity.
type FinanceAccount struct {
	ent.Schema
}

// Fields of the FinanceAccount.
func (FinanceAccount) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name").Default("İSİM YOK"),
		field.String("TcNo").Optional(),
		field.String("TaxNo").Optional(),
		field.String("TaxAdmin").Optional(),
		field.String("Phone").Optional(),
		field.String("Email").Optional(),
		field.String("Address").Optional(),
		field.String("Note").Optional(),

		field.Time("createdAt").Default(time.Now),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the FinanceAccount.
func (FinanceAccount) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", CompanyDetail.Type).Ref("accounts").Unique(),
		edge.From("group", FinanceGroup.Type).Ref("finance_accounts").Unique(),
		edge.From("type", FinanceClass.Type).Ref("types").Unique(),
		edge.To("operations", FinanceOperation.Type).StorageKey(edge.Column("account_id")),
	}
}

// Annotations of the FinanceAccount.
func (FinanceAccount) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Schema: "finance",
			Table:  "accounts",
		},
	}
}
