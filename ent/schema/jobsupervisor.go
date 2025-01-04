package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// JobSuperVisor holds the schema definition for the JobSuperVisor entity.
type JobSuperVisor struct {
	ent.Schema
}

// Fields of the JobSuperVisor.
func (JobSuperVisor) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name").Optional(),
		field.String("Address").Optional(),
		field.String("Phone").Optional(),
		field.String("Email").Optional(),
		field.Int("TcNo").Optional(),
		field.String("Position").Optional(),
		field.String("Career").Optional(),
		field.Int("RegisterNo").Optional(),
		field.Int("SocialSecurityNo").Optional(),
		field.String("SchoolGraduation").Optional(),
		field.Int("YDSID").Optional(),
		field.Time("CreatedAt").Default(time.Now),
		field.Time("UpdatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the JobSuperVisor.
func (JobSuperVisor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("supervisors", JobDetail.Type).StorageKey(edge.Column("supervisor_id")),
	}
}
