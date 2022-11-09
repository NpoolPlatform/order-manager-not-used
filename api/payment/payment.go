//nolint:nolintlint,dupl
package payment

import (
	"context"

	converter "github.com/NpoolPlatform/order-manager/pkg/converter/payment"
	crud "github.com/NpoolPlatform/order-manager/pkg/crud/payment"
	commontracer "github.com/NpoolPlatform/order-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/order-manager/pkg/tracer/payment"

	constant "github.com/NpoolPlatform/order-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mgr/v1/payment"

	"github.com/google/uuid"
)

func (s *Server) CreatePayment(ctx context.Context, in *npool.CreatePaymentRequest) (*npool.CreatePaymentResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreatePayment")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	err = validate(in.GetInfo())
	if err != nil {
		return &npool.CreatePaymentResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "Payment", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create Payment: %v", err.Error())
		return &npool.CreatePaymentResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreatePaymentResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreatePayments(ctx context.Context, in *npool.CreatePaymentsRequest) (*npool.CreatePaymentsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreatePayments")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreatePaymentsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "Payment", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create Payments: %v", err)
		return &npool.CreatePaymentsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreatePaymentsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdatePayment(ctx context.Context, in *npool.UpdatePaymentRequest) (*npool.UpdatePaymentResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreatePayments")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "Payment", "crud", "CreateBulk")

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetInfo().GetID())
		return &npool.UpdatePaymentResponse{}, status.Error(codes.InvalidArgument, "ID is invalid")
	}

	rows, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create Payments: %v", err)
		return &npool.UpdatePaymentResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdatePaymentResponse{
		Info: converter.Ent2Grpc(rows),
	}, nil
}

func (s *Server) GetPayment(ctx context.Context, in *npool.GetPaymentRequest) (*npool.GetPaymentResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetPayment")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, in.GetID())

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.GetPaymentResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "Payment", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get Payment: %v", err)
		return &npool.GetPaymentResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetPaymentResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetPaymentOnly(ctx context.Context, in *npool.GetPaymentOnlyRequest) (*npool.GetPaymentOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetPaymentOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "Payment", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get Payments: %v", err)
		return &npool.GetPaymentOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetPaymentOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetPayments(ctx context.Context, in *npool.GetPaymentsRequest) (*npool.GetPaymentsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetPayments")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "Payment", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get Payments: %v", err)
		return &npool.GetPaymentsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetPaymentsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistPayment(ctx context.Context, in *npool.ExistPaymentRequest) (*npool.ExistPaymentResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistPayment")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, in.GetID())

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.ExistPaymentResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "Payment", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check Payment: %v", err)
		return &npool.ExistPaymentResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistPaymentResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistPaymentConds(ctx context.Context,
	in *npool.ExistPaymentCondsRequest) (*npool.ExistPaymentCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistPaymentConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "Payment", "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check Payment: %v", err)
		return &npool.ExistPaymentCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistPaymentCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountPayments(ctx context.Context, in *npool.CountPaymentsRequest) (*npool.CountPaymentsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountPayments")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "Payment", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count Payments: %v", err)
		return &npool.CountPaymentsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountPaymentsResponse{
		Info: total,
	}, nil
}

func (s *Server) DeletePayment(ctx context.Context, in *npool.DeletePaymentRequest) (*npool.DeletePaymentResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreatePayments")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "Payment", "crud", "CreateBulk")

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetID())
		return &npool.DeletePaymentResponse{}, status.Error(codes.InvalidArgument, "ID is invalid")
	}

	rows, err := crud.Delete(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorf("fail create Payments: %v", err)
		return &npool.DeletePaymentResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeletePaymentResponse{
		Info: converter.Ent2Grpc(rows),
	}, nil
}
