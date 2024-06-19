package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// JobOwner holds the schema definition for the JobOwner entity.
type JobOwner struct {
	ent.Schema
}

// Fields of the JobOwner.
func (JobOwner) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name").Default(""),
		field.Int("TcNo").Optional(),
		field.String("Address").Optional(),
		field.String("TaxAdmin").Optional(),
		field.Int("TaxNo").Optional(),
		field.String("Phone").Optional(),
		field.String("Email").Optional(),
		field.String("Note").Optional(),
		field.Int("Deleted").Default(0),

		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the JobOwner.
func (JobOwner) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("owners", JobDetail.Type).StorageKey(edge.Column("owner_id")),
	}
}
