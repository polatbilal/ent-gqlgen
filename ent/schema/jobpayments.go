package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/shopspring/decimal"
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
		field.
			Other("Amount", &decimal.NullDecimal{}).
			Optional().
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(12,2)",
			}),
		field.Bool("AtMunicipality").Optional().Default(false),
		field.Time("MunicipalityDeliveryDate").Optional(),
		field.Bool("InvoiceIssued").Optional().Default(false),
		field.Time("InvoiceIssuedDate").Optional(),
		field.
			Other("InvoiceIssuedAmount", &decimal.NullDecimal{}).
			Optional().
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(12,2)",
			}),
		field.Bool("InvoiceReceived").Optional().Default(false),
		field.Time("InvoiceReceivedDate").Optional(),
		field.
			Other("InvoiceReceivedAmount", &decimal.NullDecimal{}).
			Optional().
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(12,2)",
			}),

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
