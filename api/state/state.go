//nolint:nolintlint,dupl
package state

import (
	"context"

	converter "github.com/NpoolPlatform/order-manager/pkg/converter/state"
	crud "github.com/NpoolPlatform/order-manager/pkg/crud/state"
	commontracer "github.com/NpoolPlatform/order-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/order-manager/pkg/tracer/state"

	constant "github.com/NpoolPlatform/order-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/ordermgr/state"

	"github.com/google/uuid"
)

func (s *Server) CreateState(ctx context.Context, in *npool.CreateStateRequest) (*npool.CreateStateResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateState")
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
		return &npool.CreateStateResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "state", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create state: %v", err.Error())
		return &npool.CreateStateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateStateResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateStates(ctx context.Context, in *npool.CreateStatesRequest) (*npool.CreateStatesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateStates")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateStatesResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	err = duplicate(in.GetInfos())
	if err != nil {
		return &npool.CreateStatesResponse{}, err
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "state", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create states: %v", err)
		return &npool.CreateStatesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateStatesResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) GetState(ctx context.Context, in *npool.GetStateRequest) (*npool.GetStateResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetState")
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
		return &npool.GetStateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "state", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get state: %v", err)
		return &npool.GetStateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetStateResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetStateOnly(ctx context.Context, in *npool.GetStateOnlyRequest) (*npool.GetStateOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetStateOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "state", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get states: %v", err)
		return &npool.GetStateOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetStateOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetStates(ctx context.Context, in *npool.GetStatesRequest) (*npool.GetStatesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetStates")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "state", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get states: %v", err)
		return &npool.GetStatesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetStatesResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistState(ctx context.Context, in *npool.ExistStateRequest) (*npool.ExistStateResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistState")
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
		return &npool.ExistStateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "state", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check state: %v", err)
		return &npool.ExistStateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistStateResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistStateConds(ctx context.Context,
	in *npool.ExistStateCondsRequest) (*npool.ExistStateCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistStateConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "state", "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check state: %v", err)
		return &npool.ExistStateCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistStateCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountStates(ctx context.Context, in *npool.CountStatesRequest) (*npool.CountStatesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountStates")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "state", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count states: %v", err)
		return &npool.CountStatesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountStatesResponse{
		Info: total,
	}, nil
}
