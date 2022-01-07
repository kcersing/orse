package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type ProductSpecsItem struct {
	ent.Schema
}

func (ProductSpecsItem) Config() ent.Config {
	return ent.Config{
		Table: "product_specs_item",
	}
}

func (ProductSpecsItem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (ProductSpecsItem) Fields() []ent.Field {
	return []ent.Field{
		field.Int("product_id"),
		field.Int("product_specs_id"),
		field.Int("key_id"),
		field.Int("value_id"),
		field.Int("sort"),
	}
}

func (ProductSpecsItem) Edges() []ent.Edge {
	return []ent.Edge{}
}
