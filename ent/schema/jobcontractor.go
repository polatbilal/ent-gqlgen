package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// JobContractor holds the schema definition for the JobContractor entity.
type JobContractor struct {
	ent.Schema
}

// Fields of the JobContractor.
func (JobContractor) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name").Default(""),
		field.Int("TcNo").Optional(),
		field.String("Address").Optional(),
		field.Int("RegisterNo").Optional(),
		field.String("TaxAdmin").Optional(),
		field.Int("TaxNo").Optional(),
		field.String("Phone").Optional(),
		field.String("Email").Optional(),
		field.Int("yds_id").Optional(),
		field.String("Note").Optional(),

		field.Time("CreatedAt").Default(time.Now),
		field.Time("UpdatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the JobContractor.
func (JobContractor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("contractors", JobDetail.Type).StorageKey(edge.Column("contractor_id")),
	}
}
