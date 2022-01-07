package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"orse/ent/property"
)

type Product struct {
	ent.Schema
}

func (Product) Config() ent.Config {
	return ent.Config{
		Table: "product",
	}
}

func (Product) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.Int("cate_id"),
		field.String("name").
			Optional().
			Comment("名称"),
		field.Enum("status").
			GoType((property.Status(""))),
		field.Int("create_id").
			Optional().
			Default(0),
	}
}

func (Product) Edges() []ent.Edge {
	return []ent.Edge{}
}
