package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Job_Author holds the schema definition for the Job_Author entity.
type JobAuthor struct {
	ent.Schema
}

// Fields of the Job_Author.
func (JobAuthor) Fields() []ent.Field {
	return []ent.Field{
		field.String("Architect").Default("").Optional(),
		field.String("Static").Default("").Optional(),
		field.String("Mechanic").Default("").Optional(),
		field.String("Electric").Default("").Optional(),
		field.String("Floor").Default("").Optional(),

		field.Time("CreatedAt").Default(time.Now),
		field.Time("UpdatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Job_Author.
func (JobAuthor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("authors", JobDetail.Type).StorageKey(edge.Column("author_id")),
	}
}
