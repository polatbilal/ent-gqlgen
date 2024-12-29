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
		field.Int("TcNo").Optional(),
		field.String("Phone").Optional(),
		field.String("Email").Optional(),
		field.String("Address").Optional(),
		field.String("Career").Optional(),
		field.String("Position").Optional(),
		field.Int("RegNo").Optional(),
		field.Int("CertNo").Optional(),
		field.Int("yds_id").Optional(),
		field.Time("Employment").Optional(),
		field.Time("Dismissal").Optional(),
		field.Int("Status").Default(1),
		field.String("Note").Optional(),

		field.Time("CreatedAt").Default(time.Now),
		field.Time("UpdatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the CompanyEngineer.
func (CompanyEngineer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", CompanyDetail.Type).Ref("engineers").Unique(),
		edge.To("companyOwners", CompanyDetail.Type).StorageKey(edge.Column("owner_id")),

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
