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
		field.String("Province").Default("").Optional(),
		field.String("Idare").Default("").Optional(),
		field.String("Pafta").Default("").Optional(),
		field.String("Ada").Default("").Optional(),
		field.String("Parsel").Default("").Optional(),
		field.String("FolderNo").Default("").Optional(),
		field.Int("Status").Default(0),
		field.Time("ContractDate").Optional(),
		field.Time("CompletionDate").Optional(),
		field.Time("StartDate").Optional(),
		field.Time("LicenseDate").Optional(),
		field.String("LicenseNo").Default("").Optional(),
		field.String("ConstructionArea").Optional(),
		field.String("District").Default("").Optional(),
		field.String("Village").Default("").Optional(),
		field.String("Street").Default("").Optional(),
		field.String("BuildingClass").Default("").Optional(),
		field.String("BuildingType").Default("").Optional(),
		field.String("BuildingBlock").Default("").Optional(),
		field.String("LandArea").Optional(),
		field.Int("Floors").Optional(),
		field.String("UsagePurpose").Default("").Optional(),
		field.Text("Note").Optional(),
		field.Int("Started").Default(0),
		field.Int("Deleted").Default(0),

		field.Time("CreatedAt").Default(time.Now),
		field.Time("UpdatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the JobDetail.
func (JobDetail) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", JobOwner.Type).Ref("owners").Unique(),
		edge.From("contractor", JobContractor.Type).Ref("contractors").Unique(),
		edge.From("author", JobAuthor.Type).Ref("authors").Unique(),
		edge.From("progress", JobProgress.Type).Ref("progress").Unique(),

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
