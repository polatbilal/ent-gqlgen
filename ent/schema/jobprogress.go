package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// JobProgress holds the schema definition for the JobProgress entity.
type JobProgress struct {
	ent.Schema
}

// Fields of the JobProgress.
func (JobProgress) Fields() []ent.Field {
	return []ent.Field{
		field.Int("One").Optional().Default(0),
		field.Int("Two").Optional().Default(0),
		field.Int("Three").Optional().Default(0),
		field.Int("Four").Optional().Default(0),
		field.Int("Five").Optional().Default(0),
		field.Int("Six").Optional().Default(0),

		field.Time("CreatedAt").Default(time.Now),
		field.Time("UpdatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the JobProgress.
func (JobProgress) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("progress", JobDetail.Type).StorageKey(edge.Column("progress_id")),
	}
}
