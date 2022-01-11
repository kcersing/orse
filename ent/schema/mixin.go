package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"time"
)

type PidMixin struct{
	mixin.Schema
}
func (PidMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int("parent_id").
			Optional().
			Comment("上级id"),
	}
}

type TimeMixin struct{
	mixin.Schema
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Optional().
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			Optional().
			UpdateDefault(time.Now),
	}
}
