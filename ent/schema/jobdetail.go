package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// JobDetail holds the schema definition for the JobDetail entity.
type JobDetail struct {
	ent.Schema
}

// Fields of the JobDetail.
func (JobDetail) Fields() []ent.Field {
	return []ent.Field{
		field.Int("YibfNo").Positive().Unique(),
		field.String("Idare").Optional(),
		field.String("Pafta").Optional(),
		field.String("Ada").Optional(),
		field.String("Parsel").Optional(),
		field.String("FolderNo").Optional(),
		field.Int("Status").Default(0),
		field.Time("ContractDate").Optional(),
		field.Time("CompletionDate").Optional(),
		field.Time("StartDate").Optional(),
		field.Time("LicenseDate").Optional(),
		field.String("LicenseNo").Optional(),
		field.String("ConstructionArea").Optional(),
		field.String("YDSAddress").Optional(),
		field.String("Address").Optional(),
		field.String("BuildingClass").Optional(),
		field.String("BuildingType").Optional(),
		field.Float("UnitPrice").Optional(),
		field.String("LandArea").Optional(),
		field.Int("Floors").Optional(),
		field.String("UsagePurpose").Optional(),
		field.Text("Note").Optional(),
		field.String("Coordinates").Optional(),
		field.Int("Started").Default(0),

		field.Time("CreatedAt").Default(time.Now),
		field.Time("UpdatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the JobDetail.
func (JobDetail) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", CompanyDetail.Type).Ref("jobs").Unique(),
		edge.From("owner", JobOwner.Type).Ref("owners").Unique(),
		edge.From("contractor", JobContractor.Type).Ref("contractors").Unique(),
		edge.From("author", JobAuthor.Type).Ref("authors").Unique(),
		edge.From("progress", JobProgress.Type).Ref("progress").Unique(),
		edge.From("supervisor", JobSuperVisor.Type).Ref("supervisors").Unique(),

		edge.From("inspector", CompanyEngineer.Type).Ref("inspectors").Unique(),
		edge.From("architect", CompanyEngineer.Type).Ref("architects").Unique(),
		edge.From("static", CompanyEngineer.Type).Ref("statics").Unique(),
		edge.From("mechanic", CompanyEngineer.Type).Ref("mechanics").Unique(),
		edge.From("electric", CompanyEngineer.Type).Ref("electrics").Unique(),

		edge.From("controller", CompanyEngineer.Type).Ref("controllers").Unique(),
		edge.From("mechaniccontroller", CompanyEngineer.Type).Ref("mechaniccontrollers").Unique(),
		edge.From("electriccontroller", CompanyEngineer.Type).Ref("electriccontrollers").Unique(),

		edge.To("layers", JobLayer.Type).StorageKey(edge.Column("job_id")),
		edge.To("payments", JobPayments.Type).StorageKey(edge.Column("payments_id")),
	}
}
