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
		field.Time("DistributionDate").Optional(),
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
		field.String("BKSReferenceNo").Optional(),
		field.String("Coordinates").Optional(),
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
		edge.To("relations", JobRelations.Type).Unique().StorageKey(edge.Column("job_id")),
	}
}
