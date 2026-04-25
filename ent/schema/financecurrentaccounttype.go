package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FinanceCurrentAccountType holds the schema definition for the FinanceCurrentAccountType entity.
type FinanceCurrentAccountType struct {
	ent.Schema
}

// Fields of the FinanceCurrentAccountType.
func (FinanceCurrentAccountType) Fields() []ent.Field {
	return []ent.Field{
		field.String("Description").Default("AÇIKLAMA YOK"),
		field.String("DeletedDescription").Optional(),
		field.Time("DeletedDate").Optional(),
		field.Bool("Status").Default(true),
	}
}

// Edges of the FinanceCurrentAccountType.
func (FinanceCurrentAccountType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", CompanyDetail.Type).Ref("current_account_types").Unique(),
		edge.To("current_accounts", FinanceCurrentAccount.Type).StorageKey(edge.Column("type_id")),
	}
}

// Annotations of the FinanceCurrentAccountType.
func (FinanceCurrentAccountType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Schema: "finance",
			Table:  "current_account_types",
		},
	}
}
