package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FinanceCurrentAccount holds the schema definition for the FinanceCurrentAccount entity.
type FinanceCurrentAccount struct {
	ent.Schema
}

// Fields of the FinanceCurrentAccount.
func (FinanceCurrentAccount) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name").Default("İSİM YOK"),
		field.String("DeletedName").Optional(),
		field.String("NationalIdentificationNumber").Optional(),
		field.String("Address").Optional(),
		field.String("TaxOffice").Optional(),
		field.String("TaxNumber").Optional(),
		field.String("BusinessPhone").Optional(),
		field.String("CellPhone").Optional(),
		field.String("Fax").Optional(),
		field.String("Email").Optional(),
		field.String("WebAddress").Optional(),
		field.String("Note").Optional(),
		field.Bool("Status").Default(true),

		field.Time("createdAt").Default(time.Now),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the FinanceCurrentAccount.
func (FinanceCurrentAccount) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", CompanyDetail.Type).Ref("current_accounts").Unique(),
		edge.From("group", FinanceCurrentAccountGroup.Type).Ref("current_accounts").Unique(),
		edge.To("operations", FinanceCurrentAccountOperation.Type).StorageKey(edge.Column("account_id")),
		edge.From("account_type", FinanceCurrentAccountType.Type).Ref("current_accounts").Unique(),
	}
}

// Annotations of the FinanceCurrentAccount.
func (FinanceCurrentAccount) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Schema: "finance",
			Table:  "current_accounts",
		},
	}
}
