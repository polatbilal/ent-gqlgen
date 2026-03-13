package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CompanyPersonnel holds the schema definition for the CompanyPersonnel entity.
type CompanyPersonnel struct {
	ent.Schema
}

// Fields of the CompanyPersonnel.
func (CompanyPersonnel) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name").Default("İSİM YOK"),
		field.String("TcNo").Optional(),
		field.String("Phone").Optional(),
		field.String("Email").Optional(),
		field.String("Address").Optional(),
		field.Time("Employment").Optional(),
		field.Bool("Status").Default(true),
		field.String("Note").Optional(),

		field.Time("CreatedAt").Default(time.Now),
		field.Time("UpdatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the CompanyPersonnel.
func (CompanyPersonnel) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", CompanyDetail.Type).Ref("personnels").Unique(),

		edge.To("finance_relations", FinanceRelations.Type).StorageKey(edge.Column("company_personnel_id")),
	}
}
