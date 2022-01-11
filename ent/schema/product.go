package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"orse/ent/property"
)

type Product struct {
	ent.Schema
}

func (Product) Config() ent.Config {
	return ent.Config{
		Table: "product",
	}
}

func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.Int("cate_id").
			Optional(),
		field.String("name").
			Optional().
			Comment("名称"),
		field.Enum("status").
			GoType(property.Status("")),
		field.Int("create_id").
			Optional().
			Default(0),
	}
}

func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("cate", ProductCate.Type).
			Ref("products").
			Unique().
			Field("cate_id"),
		edge.To("specs", ProductSpecs.Type),
	}
}

func (Product) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
