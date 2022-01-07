package schema

import (
	"entgo.io/ent"
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
	return []ent.Field{}
}

func (Product) Edges() []ent.Edge {
	return []ent.Edge{}
}
