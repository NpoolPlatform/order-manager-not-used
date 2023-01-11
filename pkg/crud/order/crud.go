package order

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"

	constant "github.com/NpoolPlatform/order-manager/pkg/message/const"
	commontracer "github.com/NpoolPlatform/order-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/order-manager/pkg/tracer/order"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/order/mgr/v1/order"
	"github.com/NpoolPlatform/order-manager/pkg/db"
	"github.com/NpoolPlatform/order-manager/pkg/db/ent"
	"github.com/NpoolPlatform/order-manager/pkg/db/ent/order"

	"github.com/google/uuid"
)

//nolint
func CreateSet(c *ent.OrderCreate, in *npool.OrderReq) (*ent.OrderCreate, error) {
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
	if in.ParentOrderID != nil {
		c.SetParentOrderID(uuid.MustParse(in.GetParentOrderID()))
	}
	if in.PayWithParent != nil {
		c.SetPayWithParent(in.GetPayWithParent())
	}
	if in.Units != nil {
		c.SetUnits(in.GetUnits())
	}
	if in.PromotionID != nil {
		c.SetPromotionID(uuid.MustParse(in.GetPromotionID()))
	}
	if in.DiscountCouponID != nil {
		c.SetDiscountCouponID(uuid.MustParse(in.GetDiscountCouponID()))
	}
	if in.UserSpecialReductionID != nil {
		c.SetUserSpecialReductionID(uuid.MustParse(in.GetUserSpecialReductionID()))
	}
	if in.StartAt != nil {
		c.SetStartAt(in.GetStartAt())
	}
	if in.EndAt != nil {
		c.SetEndAt(in.GetEndAt())
	}
	if in.FixAmountCouponID != nil {
		c.SetFixAmountCouponID(uuid.MustParse(in.GetFixAmountCouponID()))
	}
	if in.Type != nil {
		c.SetType(in.GetType().String())
	}
	if in.State != nil {
		c.SetState(in.GetState().String())
	}
	if in.CouponIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range in.GetCouponIDs() {
			ids = append(ids, uuid.MustParse(id))
		}
		c.SetCouponIds(ids)
	}
	if in.CreatedAt != nil {
		c.SetCreatedAt(in.GetCreatedAt())
	}
	c.SetLastBenefitAt(0)

	return c, nil
}

func Create(ctx context.Context, in *npool.OrderReq) (*ent.Order, error) {
	var info *ent.Order
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
		c := cli.Order.Create()
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

func CreateBulk(ctx context.Context, in []*npool.OrderReq) ([]*ent.Order, error) {
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

	rows := []*ent.Order{}
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.OrderCreate, len(in))
		for i, info := range in {
			bulk[i] = tx.Order.Create()
			bulk[i], err = CreateSet(bulk[i], info)
			if err != nil {
				return err
			}
		}
		rows, err = tx.Order.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func UpdateSet(u *ent.OrderUpdateOne, in *npool.OrderReq) (*ent.OrderUpdateOne, error) {
	if in.State != nil {
		u.SetState(in.GetState().String())
	}
	if in.StartAt != nil {
		u.SetStartAt(in.GetStartAt())
	}
	if in.EndAt != nil {
		u.SetEndAt(in.GetEndAt())
	}
	if in.LastBenefitAt != nil {
		u.SetLastBenefitAt(in.GetLastBenefitAt())
	}
	return u, nil
}

func Update(ctx context.Context, in *npool.OrderReq) (*ent.Order, error) {
	var info *ent.Order
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
		info, err = tx.Order.Query().Where(order.ID(uuid.MustParse(in.GetID()))).ForUpdate().Only(_ctx)
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

func Row(ctx context.Context, id uuid.UUID) (*ent.Order, error) {
	var info *ent.Order
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
		info, err = cli.Order.Query().Where(order.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

//nolint
func SetQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.OrderQuery, error) {
	stm := cli.Order.Query()
	if conds == nil {
		return stm, nil
	}
	if conds.ID != nil {
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(order.ID(uuid.MustParse(conds.GetID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	if conds.IDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetIDs().GetValue() {
			ids = append(ids, uuid.MustParse(id))
		}

		switch conds.GetIDs().GetOp() {
		case cruder.IN:
			stm.Where(order.IDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	if conds.GoodID != nil {
		switch conds.GetGoodID().GetOp() {
		case cruder.EQ:
			stm.Where(order.GoodID(uuid.MustParse(conds.GetGoodID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	if conds.AppID != nil {
		switch conds.GetAppID().GetOp() {
		case cruder.EQ:
			stm.Where(order.AppID(uuid.MustParse(conds.GetAppID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	if conds.UserID != nil {
		switch conds.GetUserID().GetOp() {
		case cruder.EQ:
			stm.Where(order.UserID(uuid.MustParse(conds.GetUserID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	if conds.Type != nil {
		switch conds.GetType().GetOp() {
		case cruder.EQ:
			stm.Where(order.Type(npool.OrderType(conds.GetType().GetValue()).String()))
		default:
			return nil, fmt.Errorf("invalid payment field")
		}
	}
	if conds.State != nil {
		switch conds.GetState().GetOp() {
		case cruder.EQ:
			stm.Where(order.State(npool.OrderState(conds.GetState().GetValue()).String()))
		default:
			return nil, fmt.Errorf("invalid payment field")
		}
	}
	if conds.FixAmountCouponID != nil {
		switch conds.GetFixAmountCouponID().GetOp() {
		case cruder.EQ:
			stm.Where(order.FixAmountCouponID(uuid.MustParse(conds.GetFixAmountCouponID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid payment field")
		}
	}
	if conds.DiscountCouponID != nil {
		switch conds.GetDiscountCouponID().GetOp() {
		case cruder.EQ:
			stm.Where(order.DiscountCouponID(uuid.MustParse(conds.GetDiscountCouponID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid payment field")
		}
	}
	if conds.UserSpecialReductionID != nil {
		switch conds.GetUserSpecialReductionID().GetOp() {
		case cruder.EQ:
			stm.Where(order.UserSpecialReductionID(uuid.MustParse(conds.GetUserSpecialReductionID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid payment field")
		}
	}
	if conds.LastBenefitAt != nil {
		switch conds.GetLastBenefitAt().GetOp() {
		case cruder.EQ:
			stm.Where(order.LastBenefitAt(conds.GetLastBenefitAt().GetValue()))
		default:
			return nil, fmt.Errorf("invalid payment field")
		}
	}
	if conds.CouponID != nil {
		switch conds.GetCouponID().GetOp() {
		case cruder.LIKE:
			stm.Where(func(selector *sql.Selector) {
				selector.Where(sqljson.ValueContains(order.FieldCouponIds, conds.GetCouponID().GetValue()))
			})
		default:
			return nil, fmt.Errorf("invalid payment field")
		}
	}
	if len(conds.GetCouponIDs().GetValue()) > 0 {
		stm.Where(func(selector *sql.Selector) {
			for _, val := range conds.GetCouponIDs().GetValue() {
				selector.Or().Where(sqljson.ValueContains(order.FieldCouponIds, val))
			}
		})
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.Order, int, error) {
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

	rows := []*ent.Order{}
	var total int
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		cli = cli.Debug()
		stm, err := SetQueryConds(conds, cli)
		if err != nil {
			return err
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}

		rows, err = stm.
			Offset(offset).
			Order(ent.Desc(order.FieldUpdatedAt)).
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

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.Order, error) {
	var info *ent.Order
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
		stm, err := SetQueryConds(conds, cli)
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
		stm, err := SetQueryConds(conds, cli)
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
		exist, err = cli.Order.Query().Where(order.ID(id)).Exist(_ctx)
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
		stm, err := SetQueryConds(conds, cli)
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

func Delete(ctx context.Context, id string) (*ent.Order, error) {
	var info *ent.Order
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
		info, err = cli.Order.UpdateOneID(uuid.MustParse(id)).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
