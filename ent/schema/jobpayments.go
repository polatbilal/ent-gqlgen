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
		field.Int("yibfNo").Default(0),
		field.Int("PaymentNo").Optional().Default(0),
		field.Time("PaymentDate").Default(time.Now),
		field.String("PaymentType").Optional(),
		field.String("State").Optional(),
		field.Float("LevelRequest").Optional(),
		field.Float("LevelApprove").Optional(),
		field.Float("Amount").Optional(),
		field.Bool("AtMunicipality").Optional().Default(false),
		field.Time("MunicipalityDeliveryDate").Optional(),
		field.Bool("InvoiceIssued").Optional().Default(false),
		field.Time("InvoiceIssuedDate").Optional(),
		field.Bool("InvoiceReceived").Optional().Default(false),
		field.Time("InvoiceReceivedDate").Optional(),

		field.Time("CreatedAt").Default(time.Now),
		field.Time("UpdatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the JobProgress.
func (JobPayments) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("payments", JobRelations.Type).Ref("payments").Unique(),
	}
}
