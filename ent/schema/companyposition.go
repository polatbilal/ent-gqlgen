package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CompanyPosition holds the schema definition for the CompanyPosition entity.
type CompanyPosition struct {
	ent.Schema
}

// Fields of the CompanyPosition.
func (CompanyPosition) Fields() []ent.Field {
	return []ent.Field{
		field.String("Position").Optional(),
	}
}

// Edges of the CompanyPosition.
func (CompanyPosition) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("engineerPositions", CompanyEngineer.Type).StorageKey(edge.Column("position_id")),
		edge.To("companyOwnerPositions", CompanyOwner.Type).StorageKey(edge.Column("position_id")),
	}
}
