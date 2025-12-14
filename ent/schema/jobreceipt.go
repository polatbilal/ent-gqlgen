package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// JobReceipt holds the schema definition for the JobReceipt entity.
type JobReceipt struct {
	ent.Schema
}

// Fields of the JobReceipt.
func (JobReceipt) Fields() []ent.Field {
	return []ent.Field{
		field.Int("yibfNo").Default(0),
		field.Time("receiptDate").Default(time.Now),
		field.Float("Amount").Optional().Default(0.0),
		field.String("Note").Optional(),

		field.Time("CreatedAt").Default(time.Now),
		field.Time("UpdatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the JobReceipt.
func (JobReceipt) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("receipt", JobRelations.Type).Ref("receipts").Unique(),
	}
}
