package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

func (Order) Config() ent.Config {
	return ent.Config{
		Table: "order",
	}
}

func (Order) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.String("sn").
			Unique(),
		field.Enum("source").
			Values("0", "1").
			Optional().
			Default("0").
			Comment("[0:PC订单; 1:app订单;]"),
		field.Enum("status").
			Values("0", "10", "20", "30", "40").
			Optional().
			Default("0").
			Comment("[0:待付款; 10:已付款; 20:已发货; 30:交易成功; 40:已关闭; ]"),

		field.Int("integration").
			Comment("可以获得的积分"),

		field.Time("payment_time").
			Optional().
			Comment("支付时间"),

		field.String("note").
			Unique().
			Comment("订单备注"),
		field.Time("comment_time").
			Optional().
			Comment("评价时间"),
		field.Enum("delete").
			Values("0", "1").
			Optional().
			Default("0").
			Comment("[0:未删除; 1:已删除;]"),
	}
}

// Edges of the Order.

func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("items", OrderItem.Type),
		edge.To("amounts", OrderAmounts.Type),
		edge.To("pays", OrderPay.Type),
		edge.To("deliverys", OrderDelivery.Type),
	}
}
func (Order) Index() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("sn").Unique(),
	}
}
