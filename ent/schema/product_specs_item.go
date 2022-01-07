package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
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
		field.Int("product_specs_id").
			Optional(),
		field.Int("value_id").Optional(),
		field.Int("sort"),
	}
}

func (ProductSpecsItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("specs", ProductSpecs.Type).
			Ref("items").
			Unique().
			Field("product_specs_id"),
		edge.To("values", ProductAttributeValue.Type),
	}
}
