package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// JobFloors holds the schema definition for the JobFloors entity.
type JobFloor struct {
	ent.Schema
}

// Fields of the JobFloors.
func (JobFloor) Fields() []ent.Field {
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

// Edges of the JobFloors.
func (JobFloor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("floor", JobRelations.Type).Ref("floors").Unique(),
	}
}
