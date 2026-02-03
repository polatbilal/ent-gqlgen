package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FinanceGroup holds the schema definition for the FinanceGroup entity.
type FinanceGroup struct {
	ent.Schema
}

// Fields of the FinanceGroup.
func (FinanceGroup) Fields() []ent.Field {
	return []ent.Field{
		field.String("Type"),
		field.String("Description"),
	}
}

// Edges of the FinanceGroup.
func (FinanceGroup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("groups", FinanceOperation.Type).StorageKey(edge.Column("group_id")),
	}
}
