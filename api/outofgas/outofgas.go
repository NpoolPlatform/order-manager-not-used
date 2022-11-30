//nolint:nolintlint,dupl
package outofgas

import (
	"context"

	converter "github.com/NpoolPlatform/order-manager/pkg/converter/outofgas"
	crud "github.com/NpoolPlatform/order-manager/pkg/crud/outofgas"
	commontracer "github.com/NpoolPlatform/order-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/order-manager/pkg/tracer/outofgas"

	constant "github.com/NpoolPlatform/order-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mgr/v1/outofgas"

	"github.com/google/uuid"
)

func (s *Server) CreateOutOfGas(ctx context.Context, in *npool.CreateOutOfGasRequest) (*npool.CreateOutOfGasResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateOutOfGas")
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
		return &npool.CreateOutOfGasResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "OutOfGas", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create OutOfGas: %v", err.Error())
		return &npool.CreateOutOfGasResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateOutOfGasResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateOutOfGass(ctx context.Context, in *npool.CreateOutOfGassRequest) (*npool.CreateOutOfGassResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateOutOfGass")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateOutOfGassResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "OutOfGas", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create OutOfGass: %v", err)
		return &npool.CreateOutOfGassResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateOutOfGassResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateOutOfGas(ctx context.Context, in *npool.UpdateOutOfGasRequest) (*npool.UpdateOutOfGasResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateOutOfGass")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "OutOfGas", "crud", "CreateBulk")

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetInfo().GetID())
		return &npool.UpdateOutOfGasResponse{}, status.Error(codes.InvalidArgument, "ID is invalid")
	}

	rows, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create OutOfGass: %v", err)
		return &npool.UpdateOutOfGasResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateOutOfGasResponse{
		Info: converter.Ent2Grpc(rows),
	}, nil
}

func (s *Server) GetOutOfGas(ctx context.Context, in *npool.GetOutOfGasRequest) (*npool.GetOutOfGasResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetOutOfGas")
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
		return &npool.GetOutOfGasResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "OutOfGas", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get OutOfGas: %v", err)
		return &npool.GetOutOfGasResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetOutOfGasResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetOutOfGasOnly(ctx context.Context, in *npool.GetOutOfGasOnlyRequest) (*npool.GetOutOfGasOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetOutOfGasOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "OutOfGas", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get OutOfGass: %v", err)
		return &npool.GetOutOfGasOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetOutOfGasOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetOutOfGass(ctx context.Context, in *npool.GetOutOfGassRequest) (*npool.GetOutOfGassResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetOutOfGass")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "OutOfGas", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get OutOfGass: %v", err)
		return &npool.GetOutOfGassResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetOutOfGassResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistOutOfGas(ctx context.Context, in *npool.ExistOutOfGasRequest) (*npool.ExistOutOfGasResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistOutOfGas")
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
		return &npool.ExistOutOfGasResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "OutOfGas", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check OutOfGas: %v", err)
		return &npool.ExistOutOfGasResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistOutOfGasResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistOutOfGasConds(ctx context.Context,
	in *npool.ExistOutOfGasCondsRequest) (*npool.ExistOutOfGasCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistOutOfGasConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "OutOfGas", "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check OutOfGas: %v", err)
		return &npool.ExistOutOfGasCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistOutOfGasCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountOutOfGass(ctx context.Context, in *npool.CountOutOfGassRequest) (*npool.CountOutOfGassResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountOutOfGass")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "OutOfGas", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count OutOfGass: %v", err)
		return &npool.CountOutOfGassResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountOutOfGassResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteOutOfGas(ctx context.Context, in *npool.DeleteOutOfGasRequest) (*npool.DeleteOutOfGasResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateOutOfGass")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "OutOfGas", "crud", "CreateBulk")

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetID())
		return &npool.DeleteOutOfGasResponse{}, status.Error(codes.InvalidArgument, "ID is invalid")
	}

	rows, err := crud.Delete(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorf("fail create OutOfGass: %v", err)
		return &npool.DeleteOutOfGasResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteOutOfGasResponse{
		Info: converter.Ent2Grpc(rows),
	}, nil
}
