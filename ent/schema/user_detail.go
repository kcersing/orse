package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Menu holds the schema definition for the Menu entity.
type UserDetail struct {
	ent.Schema
}

func (UserDetail) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user_detail"},
	}
}

// Fields of the Menu.
func (UserDetail) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").
			Optional(),
		field.String("name").
			Optional().
			Comment("名称"),
		field.Int("age").
			Optional().
			Positive(),
		field.Float("rank").
			Optional(),
		field.String("pic").
			Optional().
			Comment("头像"),

	}
}

func (UserDetail) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("detail").
			Unique().
			Field("user_id"),

	}
}

func (UserDetail) Indexes() []ent.Index {
	return []ent.Index{
		// 唯一约束索引
		index.Fields("id").
			Unique(),
	}
}

