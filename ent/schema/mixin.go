package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"time"
)

type PidMixin struct{
	// We embed the `mixin.Schema` to avoid
	// implementing the rest of the methods.
	mixin.Schema
}

func (PidMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int("parent_id").
			Optional().
			Comment("上级id"),
		field.String("tree").
			Optional().
			Comment("树"),
	}
}

// TimeMixin implements the ent.Mixin for sharing
// time fields with package schemas.
type TimeMixin struct{
	// We embed the `mixin.Schema` to avoid
	// implementing the rest of the methods.
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
