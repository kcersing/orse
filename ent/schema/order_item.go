package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type OrderItem struct {
	ent.Schema
}

func (OrderItem) Config() ent.Config {
	return ent.Config{
		Table: "order_pay",
	}
}

// Fields of the Order.
func (OrderItem) Fields() []ent.Field {
	return []ent.Field{
		field.Int("order_id").Optional(),
		field.String("sn").
			Unique(),
		field.Int("product_cate_id").Comment("商品分类id"),
		field.Int("product_id").Comment("商品id"),
		field.String("product_pic").Comment("商品名称"),
		field.String("product_name").Comment("商品名称"),
		field.Int("product_sn").Comment("商品条码"),
		field.Float("product_price").
			Default(0).
			Comment("销售价格（单个）"),
		field.Int("quantity").
			Comment("购买数量"),
		field.Int("product_specs_id").Comment("商品规格id"),
		field.String("product_specs_sn").Comment("商品规格条码"),
		field.String("product_specs_attr").Comment("商品规格属性:[{\"key\":\"大小\",\"value\":\"大\"},{\"key\":\"颜色\",\"value\":\"红\"}]"),

	}
}

// Edges of the OrderPay.
func (OrderItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("order", Order.Type).
			Ref("items").
			Unique().
			Field("order_id"),
	}
}

func (OrderItem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
func (OrderItem) Index() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("sn").Unique(),
	}
}
