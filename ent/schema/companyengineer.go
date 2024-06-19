package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CompanyEngineer holds the schema definition for the CompanyEngineer entity.
type CompanyEngineer struct {
	ent.Schema
}

// Fields of the CompanyEngineer.
func (CompanyEngineer) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name").Default(""),
		field.String("Address").Optional(),
		field.String("Email").Optional(),
		field.Int("TcNo").Optional(),
		field.String("Phone").Optional(),
		field.Int("RegNo").Optional(),
		field.Int("CertNo").Optional(),
		field.String("Note").Optional(),
		field.Int("Status").Default(0),
		field.Int("Deleted").Default(0),
		field.Time("Employment").Optional(),
		field.Time("Dismissal").Optional(),

		field.Time("CreatedAt").Default(time.Now),
		field.Time("UpdatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the CompanyEngineer.
func (CompanyEngineer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("engineerCareer", CompanyCareer.Type).Ref("engineerCareers").Unique(),
		edge.From("engineerPosition", CompanyPosition.Type).Ref("engineerPositions").Unique(),

		edge.To("inspectors", JobDetail.Type).StorageKey(edge.Column("inspector_id")),
		edge.To("architects", JobDetail.Type).StorageKey(edge.Column("architect_id")),
		edge.To("statics", JobDetail.Type).StorageKey(edge.Column("static_id")),
		edge.To("mechanics", JobDetail.Type).StorageKey(edge.Column("mechanic_id")),
		edge.To("electrics", JobDetail.Type).StorageKey(edge.Column("electric_id")),

		edge.To("controllers", JobDetail.Type).StorageKey(edge.Column("controller_id")),
		edge.To("mechaniccontrollers", JobDetail.Type).StorageKey(edge.Column("mechaniccontroller_id")),
		edge.To("electriccontrollers", JobDetail.Type).StorageKey(edge.Column("electriccontroller_id")),
	}
}
