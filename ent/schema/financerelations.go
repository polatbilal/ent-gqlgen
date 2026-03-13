package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FinanceRelations holds the schema definition for the FinanceRelations entity.
type FinanceRelations struct {
	ent.Schema
}

// Fields of the FinanceRelations.
func (FinanceRelations) Fields() []ent.Field {
	return []ent.Field{
		field.Time("createdAt").Default(time.Now),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the FinanceRelations.
func (FinanceRelations) Edges() []ent.Edge {
	return []ent.Edge{
		// Account edges (optional - sadece 1'i dolu olacak)
		edge.From("job_owner", JobOwner.Type).Ref("finance_relations").Unique(),
		edge.From("company_personnel", CompanyPersonnel.Type).Ref("finance_relations").Unique(),
		edge.From("company_engineer", CompanyEngineer.Type).Ref("finance_relations").Unique(),
		edge.From("finance_account", FinanceAccount.Type).Ref("finance_relations").Unique(),

		// Group edge (zorunlu)
		edge.From("group", FinanceGroup.Type).Ref("finance_account_relations").Unique(),

		// Operations edge (one-to-many)
		edge.To("operations", FinanceOperation.Type).StorageKey(edge.Column("relations_id")),
	}
}
