package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type ProductAttributeValue struct {
	ent.Schema
}

func (ProductAttributeValue) Config() ent.Config {
	return ent.Config{
		Table: "product_attribute_value",
	}
}

func (ProductAttributeValue) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (ProductAttributeValue) Fields() []ent.Field {
	return []ent.Field{
		field.Int("product_id"),
		field.String("value").
			Optional().
			Comment("键值"),
		field.Int("sort"),
	}
}

func (ProductAttributeValue) Edges() []ent.Edge {
	return []ent.Edge{}
}
