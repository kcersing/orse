package schema

import (
	"entgo.io/ent"
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
	return []ent.Field{}
}

func (ProductAttributeKey) Edges() []ent.Edge {
	return []ent.Edge{}
}
