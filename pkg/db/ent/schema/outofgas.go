package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/order-manager/pkg/db/mixin"

	"github.com/google/uuid"
)

// OutOfGas holds the schema definition for the OutOfGas entity.
type OutOfGas struct {
	ent.Schema
}

func (OutOfGas) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the OutOfGas.
func (OutOfGas) Fields() []ent.Field {
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
	}
}

// Edges of the OutOfGas.
func (OutOfGas) Edges() []ent.Edge {
	return nil
}
