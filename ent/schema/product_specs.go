package schema

import (
	"entgo.io/ent"
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
	return []ent.Field{}
}

func (ProductSpecs) Edges() []ent.Edge {
	return []ent.Edge{}
}
