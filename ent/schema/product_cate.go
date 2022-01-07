package schema

import (
	"entgo.io/ent"
)

type ProductCate struct {
	ent.Schema
}

func (ProductCate) Config() ent.Config {
	return ent.Config{
		Table: "product_cate",
	}
}

func (ProductCate) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (ProductCate) Fields() []ent.Field {
	return []ent.Field{}
}

func (ProductCate) Edges() []ent.Edge {
	return []ent.Edge{}
}
