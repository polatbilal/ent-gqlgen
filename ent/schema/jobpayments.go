package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// JobProgress holds the schema definition for the JobProgress entity.
type JobPayments struct {
	ent.Schema
}

// Fields of the JobProgress.
func (JobPayments) Fields() []ent.Field {
	return []ent.Field{
		field.Time("Date").Default(time.Now),
		field.Int("Amount").Optional().Default(0),
		field.String("Description").Optional().Default(""),
		field.String("Status").Optional().Default(""),
		field.Float("Percentage").Optional().Default(0.5),

		field.Time("CreatedAt").Default(time.Now),
		field.Time("UpdatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the JobProgress.
func (JobPayments) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("payments", JobDetail.Type).Ref("payments").Unique(),
	}
}
