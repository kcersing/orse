package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type ProductAttributeKey struct {
	ent.Schema
}

func (ProductAttributeKey) Config() ent.Config {
	return ent.Config{
		Table: "product_attribute_key",
	}
}

func (ProductAttributeKey) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (ProductAttributeKey) Fields() []ent.Field {
	return []ent.Field{
		field.Int("product_id"),
		field.String("name").
			Optional().
			Comment("名称"),
		field.Int("sort"),
	}
}

func (ProductAttributeKey) Edges() []ent.Edge {
	return []ent.Edge{}
}
