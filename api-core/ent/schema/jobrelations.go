package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// JobRelations holds the schema definition for the JobRelations entity.
type JobRelations struct {
	ent.Schema
}

// Fields of the JobRelations.
func (JobRelations) Fields() []ent.Field {
	return []ent.Field{
		field.Int("yibfNo"),

		field.Time("CreatedAt").Default(time.Now),
		field.Time("UpdatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the JobRelations.
func (JobRelations) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("job", JobDetail.Type).Ref("relations").Unique(),

		edge.From("owner", JobOwner.Type).Ref("owners").Unique(),
		edge.From("author", JobAuthor.Type).Ref("authors").Unique(),
		edge.From("company", CompanyDetail.Type).Ref("jobs").Unique(),
		edge.From("progress", JobProgress.Type).Ref("progress").Unique(),
		edge.From("contractor", JobContractor.Type).Ref("contractors").Unique(),
		edge.From("supervisor", JobSupervisor.Type).Ref("supervisors").Unique(),

		edge.From("static", CompanyEngineer.Type).Ref("statics").Unique(),
		edge.From("mechanic", CompanyEngineer.Type).Ref("mechanics").Unique(),
		edge.From("electric", CompanyEngineer.Type).Ref("electrics").Unique(),
		edge.From("inspector", CompanyEngineer.Type).Ref("inspectors").Unique(),
		edge.From("architect", CompanyEngineer.Type).Ref("architects").Unique(),

		edge.From("controller", CompanyEngineer.Type).Ref("controllers").Unique(),
		edge.From("mechaniccontroller", CompanyEngineer.Type).Ref("mechaniccontrollers").Unique(),
		edge.From("electriccontroller", CompanyEngineer.Type).Ref("electriccontrollers").Unique(),

		edge.To("layers", JobLayer.Type).StorageKey(edge.Column("job_id")),
		edge.To("payments", JobPayments.Type).StorageKey(edge.Column("job_id")),
	}
}
