package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// Video holds the schema definition for the Video entity.
type Video struct {
	ent.Schema
}

// Fields of the Video.
func (Video) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").StorageKey("video_id"),
		field.String("video_title"),
		field.String("video_url").Unique(),
		field.String("video_thumbnail").Optional(),
		field.String("creator_name"),
		field.String("creator_url"),
		field.Time("created_at").Default(time.Now()),
	}

}

// Edges of the Video.
func (Video) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("posted_by", User.Type).Ref("videos").Unique().Required(),
	}
}

func (Video) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("video_url"),
	}
}
