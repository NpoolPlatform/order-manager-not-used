package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/order-manager/pkg/db/mixin"

	"github.com/google/uuid"
)

// Compensate holds the schema definition for the Compensate entity.
type Compensate struct {
	ent.Schema
}

func (Compensate) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the Compensate.
func (Compensate) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			UUID("order_id", uuid.UUID{}),
		field.
			Uint32("start").
			Optional().
			Default(0),
		field.
			Uint32("end").
			Optional().
			Default(0),
		field.
			String("message").
			Optional().
			Default(""),
	}
}

// Edges of the Compensate.
func (Compensate) Edges() []ent.Edge {
	return nil
}
