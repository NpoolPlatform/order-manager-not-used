package state

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/order/mgr/v1/state"
)

func trace(span trace1.Span, in *npool.StateReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("OrderID.%v", index), in.GetOrderID()),
		attribute.String(fmt.Sprintf("State.%v", index), in.GetState().String()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.StateReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	span.SetAttributes(
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Value", in.GetID().GetValue()),
		attribute.String("OrderID.Op", in.GetOrderID().GetOp()),
		attribute.String("OrderID.Value", in.GetOrderID().GetValue()),
		attribute.String("State.Op", in.GetState().GetOp()),
		attribute.String("State.Value", npool.EState(in.GetState().GetValue()).String()),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.StateReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
