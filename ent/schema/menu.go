package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"saas/ent/property"
)

// Menu holds the schema definition for the Menu entity.
type Menu struct {
	ent.Schema
}

//func (Menu) Config() ent.Config {
//	return ent.Config{
//		Table: "menu",
//	}
//}

func (Menu) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "menu"},
	}
}

func (Menu) Mixin() []ent.Mixin {
	return []ent.Mixin{
		PidMixin{},
		TimeMixin{},
	}
}

// Fields of the Menu.
func (Menu) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			Comment("标题"),

		field.String("name").
			Optional().
			Comment("名称"),

		field.String("url").
			Optional().
			Comment("路由"),

		field.Int("level").
			Default(1).
			Comment("级别"),

		field.Int("sort").
			Default(1).
			Comment("排序"),

		field.Enum("status").
			GoType((property.Status(""))),
		field.String("icon").
			Optional().
			Comment("样式"),

		field.Enum("hidden").
			GoType((property.Status(""))).
			Optional().
			Comment("是否隐藏:[0:显示;1:隐藏]"),

		field.String("desc").
			Optional().
			Comment("菜单描述"),
	}
}


// Edges of the Node.
func (Menu) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("prev", Menu.Type).
			Unique().
			Field("pid").
			From("next").
			Unique(),

	}
}

func (Menu) Indexes() []ent.Index {
	return []ent.Index{
		// 非唯一约束索引
		index.Fields("url"),
		// 唯一约束索引
		index.Fields("id").
			Unique(),
	}
}