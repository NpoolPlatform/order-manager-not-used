package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/message/npool/order/mgr/v1/payment"
)

// Payment holds the schema definition for the Payment entity.
type Payment struct {
	ent.Schema
}

// Fields of the Payment.
func (Payment) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			UUID("app_id", uuid.UUID{}),
		field.
			UUID("user_id", uuid.UUID{}),
		field.
			UUID("good_id", uuid.UUID{}),
		field.
			UUID("order_id", uuid.UUID{}),
		field.
			UUID("account_id", uuid.UUID{}),
		field.
			Uint64("start_amount"),
		field.
			Other("start_amount", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Other("amount", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Other("pay_with_balance_amount", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37, 18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Other("finish_amount", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37, 18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Other("coin_usd_currency", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37, 18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Other("local_coin_usd_currency", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37, 18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Other("live_coin_usd_currency", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37, 18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			UUID("coin_info_id", uuid.UUID{}),
		field.String("state").
			Optional().
			Default(payment.PaymentState_DefaultState.String()),
		field.
			String("chain_transaction_id").
			Optional().
			Default(""),
		field.Bool("user_set_paid").
			Optional().
			Default(false),
		field.Bool("user_set_canceled").
			Optional().
			Default(false),
		field.Bool("fake_payment").
			Optional().
			Default(false),
	}
}

// Edges of the Payment.
func (Payment) Edges() []ent.Edge {
	return nil
}
