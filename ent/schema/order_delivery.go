package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type OrderDelivery struct {
	ent.Schema
}

func (OrderDelivery) Config() ent.Config {
	return ent.Config{
		Table: "order_delivery",
	}
}

// Fields of the Order.
func (OrderDelivery) Fields() []ent.Field {
	return []ent.Field{
		field.Int("order_id").
			Optional(),
		field.String("sn").
			Unique().
			Comment("物流单号"),
		field.String("delivery_company").
			Optional().
			Comment("配送方式"),
		field.Time("delivery_time").
			Optional().
			Comment("发货时间"),

		field.String("receiver_name").
			Optional().
			Comment("收货人姓名"),
		field.String("receiver_phone").
			Optional().
			Comment("收货人电话"),
		field.String("receiver_post_code").
			Optional().
			Comment("收货人邮编"),
		field.String("receiver_province").
			Optional().
			Comment("省份/直辖市"),

		field.String("receiver_city").
			Optional().
			Comment("城市"),
		field.String("receiver_region").
			Optional().
			Comment("区"),
		field.String("receiver_detail_address").
			Optional().
			Comment("详细地址"),

		field.Time("receive_time").
			Optional().
			Comment("确认收货时间"),
	}
}

// Edges of the OrderPay.
func (OrderDelivery) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("order", Order.Type).
			Ref("deliverys").
			Unique().
			Field("order_id"),
	}
}

func (OrderDelivery) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
func (OrderDelivery) Index() []ent.Index {
	return []ent.Index{
		index.Fields("order_id"),
		index.Fields("sn").Unique(),
	}
}
