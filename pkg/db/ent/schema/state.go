package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/order-manager/pkg/db/mixin"
	"github.com/google/uuid"

	"github.com/NpoolPlatform/message/npool/order/mgr/v1/state"
)

// State holds the schema definition for the State entity.
type State struct {
	ent.Schema
}

func (State) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the State.
func (State) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			UUID("order_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			String("state").
			Optional().
			Default(state.EState_DefaultState.String()),
	}
}

// Edges of the State.
func (State) Edges() []ent.Edge {
	return nil
}
