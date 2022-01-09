package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type OrderPay struct {
	ent.Schema
}

func (OrderPay) Config() ent.Config {
	return ent.Config{
		Table: "order_pay",
	}
}

// Fields of the Order.
func (OrderPay) Fields() []ent.Field {
	return []ent.Field{
		field.Int("order_id").Optional(),
		field.Int("user_id"),
		field.String("sn").
			Unique(),
		field.Float("price").
			Default(0),
		field.String("pay_mode").
			Optional(),
	}
}

// Edges of the OrderPay.
func (OrderPay) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("order", Order.Type).
			Ref("pays").
			Unique().
			Field("order_id"),
	}
}

func (OrderPay) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
func (OrderPay) Index() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("sn").Unique(),
	}
}
