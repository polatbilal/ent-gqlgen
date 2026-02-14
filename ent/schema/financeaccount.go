package schema

import (
	"time"

	"entgo.io/ent"
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
		field.String("Name").Default(""),

		field.Time("createdAt").Default(time.Now),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the FinanceAccount.
func (FinanceAccount) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", CompanyDetail.Type).Ref("accounts").Unique(),
		edge.To("accounts", FinanceOperation.Type).StorageKey(edge.Column("account_id")),
	}
}
