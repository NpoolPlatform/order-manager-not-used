package payment

import (
	"context"
	"fmt"
	"time"

	"github.com/shopspring/decimal"

	constant "github.com/NpoolPlatform/order-manager/pkg/message/const"
	commontracer "github.com/NpoolPlatform/order-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/order-manager/pkg/tracer/payment"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/order/mgr/v1/payment"
	"github.com/NpoolPlatform/order-manager/pkg/db"
	"github.com/NpoolPlatform/order-manager/pkg/db/ent"
	"github.com/NpoolPlatform/order-manager/pkg/db/ent/payment"

	"github.com/google/uuid"
)

//nolint
func CreateSet(c *ent.PaymentCreate, in *npool.PaymentReq) (*ent.PaymentCreate, error) {
	if in.ID != nil {
		c.SetID(uuid.MustParse(in.GetID()))
	}
	if in.AppID != nil {
		c.SetAppID(uuid.MustParse(in.GetAppID()))
	}
	if in.GoodID != nil {
		c.SetGoodID(uuid.MustParse(in.GetGoodID()))
	}
	if in.UserID != nil {
		c.SetUserID(uuid.MustParse(in.GetUserID()))
	}
	if in.OrderID != nil {
		c.SetOrderID(uuid.MustParse(in.GetOrderID()))
	}
	if in.AccountID != nil {
		c.SetAccountID(uuid.MustParse(in.GetAccountID()))
	}
	if in.StartAmount != nil {
		d, err := decimal.NewFromString(in.GetStartAmount())
		if err != nil {
			return nil, err
		}
		c.SetStartAmount(d)
	}
	if in.Amount != nil {
		d, err := decimal.NewFromString(in.GetAmount())
		if err != nil {
			return nil, err
		}
		c.SetAmount(d)
	}
	if in.PayWithBalanceAmount != nil {
		d, err := decimal.NewFromString(in.GetPayWithBalanceAmount())
		if err != nil {
			return nil, err
		}
		c.SetPayWithBalanceAmount(d)
	}
	if in.FinishAmount != nil {
		d, err := decimal.NewFromString(in.GetFinishAmount())
		if err != nil {
			return nil, err
		}
		c.SetFinishAmount(d)
	}
	if in.CoinUsdCurrency != nil {
		d, err := decimal.NewFromString(in.GetCoinUsdCurrency())
		if err != nil {
			return nil, err
		}
		c.SetCoinUsdCurrency(d)
	}
	if in.LocalCoinUsdCurrency != nil {
		d, err := decimal.NewFromString(in.GetLocalCoinUsdCurrency())
		if err != nil {
			return nil, err
		}
		c.SetLocalCoinUsdCurrency(d)
	}
	if in.LiveCoinUsdCurrency != nil {
		d, err := decimal.NewFromString(in.GetLiveCoinUsdCurrency())
		if err != nil {
			return nil, err
		}
		c.SetLiveCoinUsdCurrency(d)
	}
	if in.CoinInfoID != nil {
		c.SetCoinInfoID(uuid.MustParse(in.GetCoinInfoID()))
	}
	if in.State != nil {
		c.SetState(in.GetState().String())
	}
	if in.ChainTransactionID != nil {
		c.SetChainTransactionID(in.GetChainTransactionID())
	}
	if in.UserSetPaid != nil {
		c.SetUserSetPaid(in.GetUserSetPaid())
	}
	if in.UserSetCanceled != nil {
		c.SetUserSetCanceled(in.GetUserSetCanceled())
	}
	if in.FakePayment != nil {
		c.SetFakePayment(in.GetFakePayment())
	}
	if in.CreatedAt != nil {
		c.SetCreatedAt(in.GetCreatedAt())
	}

	return c, nil
}

func Create(ctx context.Context, in *npool.PaymentReq) (*ent.Payment, error) {
	var info *ent.Payment
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Create")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := cli.Payment.Create()
		stm, err := CreateSet(c, in)
		if err != nil {
			return err
		}
		info, err = stm.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateBulk(ctx context.Context, in []*npool.PaymentReq) ([]*ent.Payment, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateBulk")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceMany(span, in)

	rows := []*ent.Payment{}
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.PaymentCreate, len(in))
		for i, info := range in {
			bulk[i] = tx.Payment.Create()
			bulk[i], err = CreateSet(bulk[i], info)
			if err != nil {
				return err
			}
		}
		rows, err = tx.Payment.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func UpdateSet(u *ent.PaymentUpdateOne, in *npool.PaymentReq) (*ent.PaymentUpdateOne, error) {
	if in.State != nil {
		u.SetState(in.GetState().String())
	}
	if in.UserSetPaid != nil {
		u.SetUserSetPaid(in.GetUserSetPaid())
	}
	if in.UserSetCanceled != nil {
		u.SetUserSetCanceled(in.GetUserSetCanceled())
	}
	if in.FakePayment != nil {
		u.SetFakePayment(in.GetFakePayment())
	}
	if in.FinishAmount != nil {
		finishAmount, err := decimal.NewFromString(in.GetFinishAmount())
		if err != nil {
			return nil, err
		}
		u.SetFinishAmount(finishAmount)
	}
	return u, nil
}

func Update(ctx context.Context, in *npool.PaymentReq) (*ent.Payment, error) {
	var info *ent.Payment
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Update")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in)

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		info, err = tx.Payment.Query().Where(payment.ID(uuid.MustParse(in.GetID()))).ForUpdate().Only(_ctx)
		if err != nil {
			return err
		}

		stm, err := UpdateSet(info.Update(), in)
		if err != nil {
			return err
		}

		info, err = stm.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func Row(ctx context.Context, id uuid.UUID) (*ent.Payment, error) {
	var info *ent.Payment
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Row")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.Payment.Query().Where(payment.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

//nolint
func setQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.PaymentQuery, error) {
	stm := cli.Payment.Query()
	if conds == nil {
		return stm, nil
	}
	if conds.ID != nil {
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(payment.ID(uuid.MustParse(conds.GetID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid payment field")
		}
	}
	if conds.AppID != nil {
		switch conds.GetAppID().GetOp() {
		case cruder.EQ:
			stm.Where(payment.AppID(uuid.MustParse(conds.GetAppID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid payment field")
		}
	}
	if conds.UserID != nil {
		switch conds.GetUserID().GetOp() {
		case cruder.EQ:
			stm.Where(payment.UserID(uuid.MustParse(conds.GetUserID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid payment field")
		}
	}
	if conds.GoodID != nil {
		switch conds.GetGoodID().GetOp() {
		case cruder.EQ:
			stm.Where(payment.GoodID(uuid.MustParse(conds.GetGoodID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid payment field")
		}
	}
	if conds.OrderID != nil {
		switch conds.GetOrderID().GetOp() {
		case cruder.EQ:
			stm.Where(payment.OrderID(uuid.MustParse(conds.GetOrderID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid payment field")
		}
	}
	if conds.AccountID != nil {
		switch conds.GetAccountID().GetOp() {
		case cruder.EQ:
			stm.Where(payment.AccountID(uuid.MustParse(conds.GetAccountID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid payment field")
		}
	}
	if conds.CoinInfoID != nil {
		switch conds.GetCoinInfoID().GetOp() {
		case cruder.EQ:
			stm.Where(payment.CoinInfoID(uuid.MustParse(conds.GetCoinInfoID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid payment field")
		}
	}
	if conds.State != nil {
		switch conds.GetState().GetOp() {
		case cruder.EQ:
			stm.Where(payment.State(npool.PaymentState(conds.GetState().GetValue()).String()))
		default:
			return nil, fmt.Errorf("invalid payment field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.Payment, int, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Rows")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)
	span = commontracer.TraceOffsetLimit(span, offset, limit)

	rows := []*ent.Payment{}
	var total int
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}

		rows, err = stm.
			Offset(offset).
			Order(ent.Desc(payment.FieldUpdatedAt)).
			Limit(limit).
			All(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.Payment, error) {
	var info *ent.Payment
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "RowOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}

		info, err = stm.Only(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func Count(ctx context.Context, conds *npool.Conds) (uint32, error) {
	var err error
	var total int

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Count")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}

	return uint32(total), nil
}

func Exist(ctx context.Context, id uuid.UUID) (bool, error) {
	var err error
	exist := false

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Exist")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.Payment.Query().Where(payment.ID(id)).Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func ExistConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	var err error
	exist := false

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}

		exist, err = stm.Exist(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func Delete(ctx context.Context, id string) (*ent.Payment, error) {
	var info *ent.Payment
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Delete")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.Payment.UpdateOneID(uuid.MustParse(id)).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
