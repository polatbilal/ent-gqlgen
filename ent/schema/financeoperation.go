package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FinanceOperation holds the schema definition for the FinanceOperation entity.
type FinanceOperation struct {
	ent.Schema
}

// Fields of the FinanceOperation.
func (FinanceOperation) Fields() []ent.Field {
	return []ent.Field{
		field.String("Operation"),
		field.Time("Date"),
		field.String("Debit"),
		field.String("Credit"),
		field.String("Description"),

		field.Time("createdAt").Default(time.Now),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the FinanceOperation.
func (FinanceOperation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account", FinanceAccount.Type).Ref("accounts").Unique(),
		edge.From("method", FinanceClass.Type).Ref("methods").Unique(),
		edge.From("company", CompanyDetail.Type).Ref("operations").Unique(),
		edge.From("resource", FinanceResource.Type).Ref("resources").Unique(),
		edge.From("group", FinanceGroup.Type).Ref("groups").Unique(),
	}
}
