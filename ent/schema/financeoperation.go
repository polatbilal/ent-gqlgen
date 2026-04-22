package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/shopspring/decimal"
)

// FinanceOperation holds the schema definition for the FinanceOperation entity.
type FinanceOperation struct {
	ent.Schema
}

// Fields of the FinanceOperation.
func (FinanceOperation) Fields() []ent.Field {
	return []ent.Field{
		field.Time("Date"),
		field.
			Other("Debit", &decimal.NullDecimal{}).
			Optional().
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(12,2)",
			}),
		field.
			Other("Credit", &decimal.NullDecimal{}).
			Optional().
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(12,2)",
			}),
		field.String("Description"),

		field.Time("createdAt").Default(time.Now),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the FinanceOperation.
func (FinanceOperation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("class", FinanceClass.Type).Ref("classes").Unique(),
		edge.From("company", CompanyDetail.Type).Ref("operations").Unique(),
		edge.From("resource", FinanceResource.Type).Ref("resources").Unique(),
		edge.From("group", FinanceGroup.Type).Ref("groups").Unique(),
		edge.From("account", FinanceAccount.Type).Ref("operations").Unique(),
	}
}
