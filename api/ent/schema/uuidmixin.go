package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

type UUIDMixin struct {
	// We embed the `mixin.Schema` to avoid
	// implementing the rest of the methods.
	mixin.Schema
}

func (UUIDMixin) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),
	}
}
