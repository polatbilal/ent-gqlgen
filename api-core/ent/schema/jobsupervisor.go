package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// JobSupervisor holds the schema definition for the JobSupervisor entity.
type JobSupervisor struct {
	ent.Schema
}

// Fields of the JobSupervisor.
func (JobSupervisor) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name").Optional(),
		field.String("Address").Optional(),
		field.String("Phone").Optional(),
		field.String("Email").Optional(),
		field.String("TcNo").Optional(),
		field.String("Position").Optional(),
		field.String("Career").Optional(),
		field.String("RegisterNo").Optional(),
		field.String("SocialSecurityNo").Optional(),
		field.String("SchoolGraduation").Optional(),
		field.Int("YDSID").Optional().Unique(),

		field.Time("CreatedAt").Default(time.Now),
		field.Time("UpdatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the JobSupervisor.
func (JobSupervisor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("supervisors", JobRelations.Type).StorageKey(edge.Column("supervisor_id")),
	}
}
