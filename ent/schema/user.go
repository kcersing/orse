package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"time"
)

// Menu holds the schema definition for the Menu entity.
type User struct {
	ent.Schema
}

//func (Menu) Config() ent.Config {
//	return ent.Config{
//		Table: "menu",
//	}
//}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user"},
	}
}

// Fields of the Menu.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			Unique().
			Comment("用户名"),
		field.String("mobile").
			Unique().
			Comment("手机号"),
		field.String("pass").
			Sensitive().
			Optional().
			Comment("密码"),
		field.UUID("uuid", uuid.UUID{}).
			Default(uuid.New),
		field.Int("role").
			Optional().
			Comment("Role"),
		field.Bool("active").
			Default(false),
		field.Enum("state").
			Values("on", "off").
			Optional(),
		field.Time("login").
			Default(time.Now).
			Comment("上次登录的时间"),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		// 唯一约束索引
		index.Fields("id").
			Unique(),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("detail", UserDetail.Type),
	}
}