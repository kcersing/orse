package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type OrderAmounts struct {
	ent.Schema
}

func (OrderAmounts) Config() ent.Config {
	return ent.Config{
		Table: "order_amounts",
	}
}

// Fields of the Order.
func (OrderAmounts) Fields() []ent.Field {
	return []ent.Field{
		field.Int("order_id").
			Optional(),
		field.Float("total_amount").
			Default(0).
			Comment("订单总金额"),
		field.Float("pay_amount").
			Default(0).
			Comment("应付金额（实际支付金额）"),

		field.Float("freight_amount").
			Default(0).
			Comment("运费金额"),
		field.Float("promotion_amount").
			Default(0).
			Comment("促销金额"),
		field.Float("integration_amount").
			Default(0).
			Comment("积分抵扣金额"),
		field.Int("coupon_id").
			Optional().
			Comment("优惠券id"),
		field.Float("coupon_amount").
			Default(0).
			Comment("优惠券抵扣金额"),
		field.Float("discount_amount").
			Default(0).
			Comment("操作人减免的订单金额"),


	}
}

func (OrderAmounts) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("order", Order.Type).
			Ref("amounts").
			Unique().
			Field("order_id"),
	}
}

func (OrderAmounts) Index() []ent.Index {
	return []ent.Index{
		index.Fields("order_id"),
	}
}

func (OrderAmounts) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}