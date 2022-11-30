package payment

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/order/mgr/v1/payment"
)

func trace(span trace1.Span, in *npool.PaymentReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("AppID.%v", index), in.GetAppID()),
		attribute.String(fmt.Sprintf("GoodID.%v", index), in.GetGoodID()),
		attribute.String(fmt.Sprintf("UserID.%v", index), in.GetUserID()),
		attribute.String(fmt.Sprintf("OrderID.%v", index), in.GetOrderID()),
		attribute.String(fmt.Sprintf("StartAmount.%v", index), in.GetStartAmount()),
		attribute.String(fmt.Sprintf("Amount.%v", index), in.GetAmount()),
		attribute.String(fmt.Sprintf("PayWithBalanceAmount.%v", index), in.GetPayWithBalanceAmount()),
		attribute.String(fmt.Sprintf("FinishAmount.%v", index), in.GetFinishAmount()),
		attribute.String(fmt.Sprintf("CoinUsdCurrency.%v", index), in.GetCoinUsdCurrency()),
		attribute.String(fmt.Sprintf("LocalCoinUsdCurrency.%v", index), in.GetLocalCoinUsdCurrency()),
		attribute.String(fmt.Sprintf("LiveCoinUsdCurrency.%v", index), in.GetLiveCoinUsdCurrency()),
		attribute.String(fmt.Sprintf("CoinInfoID.%v", index), in.GetCoinInfoID()),
		attribute.String(fmt.Sprintf("State.%v", index), in.GetState().String()),
		attribute.String(fmt.Sprintf("ChainTransactionID.%v", index), in.GetChainTransactionID()),
		attribute.Bool(fmt.Sprintf("UserSetPaid.%v", index), in.GetUserSetPaid()),
		attribute.Bool(fmt.Sprintf("UserSetCanceled.%v", index), in.GetUserSetCanceled()),
		attribute.Bool(fmt.Sprintf("FakePayment.%v", index), in.GetFakePayment()),
		attribute.Int(fmt.Sprintf("CreateAt.%v", index), int(in.GetCreatedAt())),
	)
	return span
}

func Trace(span trace1.Span, in *npool.PaymentReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	span.SetAttributes(
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Value", in.GetID().GetValue()),
		attribute.String("AppID.Op", in.GetAppID().GetOp()),
		attribute.String("AppID.Value", in.GetAppID().GetValue()),
		attribute.String("UserID.Op", in.GetUserID().GetOp()),
		attribute.String("UserID.Value", in.GetUserID().GetValue()),
		attribute.String("GoodID.Op", in.GetGoodID().GetOp()),
		attribute.String("GoodID.Value", in.GetGoodID().GetValue()),
		attribute.String("OrderID.Op", in.GetOrderID().GetOp()),
		attribute.String("OrderID.Value", in.GetOrderID().GetValue()),
		attribute.String("AccountID.Op", in.GetAccountID().GetOp()),
		attribute.String("AccountID.Value", in.GetAccountID().GetValue()),
		attribute.String("CoinInfoID.Op", in.GetCoinInfoID().GetOp()),
		attribute.String("CoinInfoID.Value", in.GetCoinInfoID().GetValue()),
		attribute.String("State.Op", in.GetState().GetOp()),
		attribute.Int("State.Value", int(in.GetState().GetValue())),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.PaymentReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
