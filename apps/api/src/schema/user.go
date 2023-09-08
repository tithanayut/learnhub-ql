package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).StorageKey("user_id"),
		field.String("username").Unique(),
		field.String("password"),
		field.String("name"),
		field.Time("registered_at").
			Default(time.Now),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("videos", Video.Type).StorageKey(edge.Column("posted_by")),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("username", "password"),
	}
}
