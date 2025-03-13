package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// JobLayers holds the schema definition for the JobLayers entity.
type JobLayer struct {
	ent.Schema
}

// Fields of the JobLayers.
func (JobLayer) Fields() []ent.Field {
	return []ent.Field{
		field.Int("yibfNo"),
		field.String("Name").Default(""),
		field.String("Metre").Default(""),
		field.Time("MoldDate").Optional(),
		field.Time("ConcreteDate").Optional(),
		field.Int("Samples").Optional(),
		field.String("ConcreteClass").Optional(),
		field.String("WeekResult").Optional(),
		field.String("MonthResult").Optional(),

		field.Time("CreatedAt").Default(time.Now),
		field.Time("UpdatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the JobLayers.
func (JobLayer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("layer", JobRelations.Type).Ref("layers").Unique(),
	}
}
