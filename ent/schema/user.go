package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("Username").MaxLen(191).Unique(),
		field.String("Name").Default(""),
		field.String("Email").Default(""),
		field.String("Phone").Optional(),
		field.String("Password"),
		field.String("Role").Default("User"),
		field.String("RefreshToken").Optional(),
		field.Time("RefreshTokenExpireAt").Optional(),
		field.Time("LicenseExpireDate").Optional(),
		field.Time("LastLogin").Optional(),
		field.Bool("Active").Default(true),
		field.Time("CreatedAt").Default(time.Now),
		field.Time("UpdatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("companies", CompanyUser.Type).StorageKey(edge.Column("user_id")),
	}
}
