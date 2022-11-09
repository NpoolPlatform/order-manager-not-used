package outofgas

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/order/mgr/v1/outofgas"
)

func trace(span trace1.Span, in *npool.OutOfGasReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("OrderID.%v", index), in.GetOrderID()),
		attribute.Int(fmt.Sprintf("Start.%v", index), int(in.GetStart())),
		attribute.Int(fmt.Sprintf("End.%v", index), int(in.GetEnd())),
	)
	return span
}

func Trace(span trace1.Span, in *npool.OutOfGasReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	span.SetAttributes(
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Value", in.GetID().GetValue()),
		attribute.String("OrderID.Op", in.GetOrderID().GetOp()),
		attribute.String("OrderID.Value", in.GetOrderID().GetValue()),
		attribute.String("Start.Op", in.GetStart().GetOp()),
		attribute.Int("Start.Value", int(in.GetStart().GetValue())),
		attribute.String("End.Op", in.GetEnd().GetOp()),
		attribute.Int("End.Value", int(in.GetEnd().GetValue())),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.OutOfGasReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
