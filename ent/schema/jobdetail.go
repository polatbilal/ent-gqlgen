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
		field.Int("YibfNo").Unique(),
		field.String("Title").Optional(),
		field.String("Administration").Optional(),
		field.String("State").Optional(),
		field.String("Island").Optional(),
		field.String("Parcel").Optional(),
		field.String("Sheet").Optional(),
		field.Time("ContractDate").Optional(),
		field.Time("StartDate").Optional(),
		field.Time("LicenseDate").Optional(),
		field.String("LicenseNo").Optional(),
		field.Time("CompletionDate").Optional(),
		field.Float("LandArea").Optional(),
		field.Float("TotalArea").Optional(),
		field.Float("ConstructionArea").Optional(),
		field.Float("LeftArea").Optional(),
		field.String("YDSAddress").Optional(),
		field.String("Address").Optional(),
		field.String("BuildingClass").Optional(),
		field.String("BuildingType").Optional(),
		field.Float("Level").Optional(),
		field.Float("UnitPrice").Optional(),
		field.Int("FloorCount").Optional(),
		field.Int("BKSReferenceNo").Optional(),
		field.Text("Coordinates").Optional(),
		field.String("FolderNo").Optional(),
		field.Bool("UploadedFile").Default(false).Optional(),
		field.Bool("IndustryArea").Default(false).Optional(),
		field.Bool("ClusterStructure").Default(false).Optional(),
		field.Bool("IsLicenseExpired").Default(false).Optional(),
		field.Bool("IsCompleted").Default(false).Optional(),
		field.Text("Note").Optional(),

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
		edge.From("supervisor", JobSupervisor.Type).Ref("supervisors").Unique(),

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
