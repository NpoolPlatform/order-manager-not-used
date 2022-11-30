package order

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/order/mgr/v1/order"
)

func trace(span trace1.Span, in *npool.OrderReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("AppID.%v", index), in.GetAppID()),
		attribute.String(fmt.Sprintf("GoodID.%v", index), in.GetGoodID()),
		attribute.String(fmt.Sprintf("UserID.%v", index), in.GetUserID()),
		attribute.String(fmt.Sprintf("ParentOrderID.%v", index), in.GetParentOrderID()),
		attribute.Bool(fmt.Sprintf("PayWithParent.%v", index), in.GetPayWithParent()),
		attribute.Int(fmt.Sprintf("Units.%v", index), int(in.GetUnits())),
		attribute.String(fmt.Sprintf("PromotionID.%v", index), in.GetPromotionID()),
		attribute.String(fmt.Sprintf("DiscountCouponID.%v", index), in.GetDiscountCouponID()),
		attribute.String(fmt.Sprintf("UserSpecialReductionID.%v", index), in.GetUserSpecialReductionID()),
		attribute.Int(fmt.Sprintf("StartAt.%v", index), int(in.GetStartAt())),
		attribute.String(fmt.Sprintf("FixAmountCouponID.%v", index), in.GetFixAmountCouponID()),
		attribute.String(fmt.Sprintf("Type.%v", index), in.GetType().String()),
		attribute.String(fmt.Sprintf("State.%v", index), in.GetState().String()),
		attribute.Int(fmt.Sprintf("CreatedAt.%v", index), int(in.GetCreatedAt())),
	)
	return span
}

func Trace(span trace1.Span, in *npool.OrderReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	span.SetAttributes(
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Value", in.GetID().GetValue()),
		attribute.String("AppID.Op", in.GetAppID().GetOp()),
		attribute.String("AppID.Value", in.GetAppID().GetValue()),
		attribute.String("GoodID.Op", in.GetGoodID().GetOp()),
		attribute.String("GoodID.Value", in.GetGoodID().GetValue()),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.OrderReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
