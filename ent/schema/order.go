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
		field.Int64("user_id"),
		field.String("sn").
			Unique(),
		field.Enum("status").
			Values("0", "10", "20", "30", "40", "50", "60", "70", "80", "90").
			Optional().
			Default("0").
			Comment("[0:待付款; 10:已付款; 20:卖家已发货; 30:交易成功; 40:待评价; 50:待退款; 60:售后维权; 70:; 80:; 90:; ]"),
	}
}

// Edges of the Order.

func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("pays", OrderPay.Type),
	}
}
func (Order) Index() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("sn").Unique(),
	}
}
