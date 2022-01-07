package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ProductSpecs 商品规格表
type ProductSpecs struct {
	ent.Schema
}

func (ProductSpecs) Config() ent.Config {
	return ent.Config{
		Table: "product_specs",
	}
}

func (ProductSpecs) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (ProductSpecs) Fields() []ent.Field {
	return []ent.Field{
		field.Int("product_id"),
		field.String("specs"),
		field.Int("stock"),
		field.Float("price"),
		field.Int("create_id").
			Optional().
			Default(0),
	}
}

func (ProductSpecs) Edges() []ent.Edge {
	return []ent.Edge{}
}
