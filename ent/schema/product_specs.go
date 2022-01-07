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
		field.String("name"),
		field.String("sn"),
		field.Int("stock").
			Comment("库存"),
		field.Int("sales").
			Comment("销量"),
		field.Float("price").
			Comment("原价"),
		field.Float("sale_price").
			Comment("售卖价"),
		field.Int("create_id").
			Optional().
			Default(0),
	}
}

func (ProductSpecs) Edges() []ent.Edge {
	return []ent.Edge{}
}
