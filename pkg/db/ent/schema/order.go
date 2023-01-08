package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/message/npool/order/mgr/v1/order"
	"github.com/NpoolPlatform/order-manager/pkg/db/mixin"

	"github.com/google/uuid"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

func (Order) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			UUID("good_id", uuid.UUID{}),
		field.
			UUID("app_id", uuid.UUID{}),
		field.
			UUID("user_id", uuid.UUID{}),
		field.
			UUID("parent_order_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			Bool("pay_with_parent").
			Optional().
			Default(false),
		field.
			Uint32("units"),
		field.
			UUID("promotion_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("discount_coupon_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("user_special_reduction_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			Uint32("start_at").
			Optional().
			Default(0),
		field.
			Uint32("end_at").
			Optional().
			Default(0),
		field.
			UUID("fix_amount_coupon_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			String("type").
			Optional().
			Default(order.OrderType_DefaultOrderType.String()),
		field.
			String("state").
			Optional().
			Default(order.OrderState_DefaultState.String()),
		field.
			JSON("coupon_ids", []uuid.UUID{}).
			Optional().
			Default(func() []uuid.UUID {
				return []uuid.UUID{}
			}),
		field.
			Uint32("last_benefit_at").
			Optional().
			Default(0),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return nil
}
