package schema

import (
	"entgo.io/ent"
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
	return []ent.Field{}
}

func (ProductAttributeValue) Edges() []ent.Edge {
	return []ent.Edge{}
}
