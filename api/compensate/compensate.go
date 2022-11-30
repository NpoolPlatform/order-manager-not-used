//nolint:nolintlint,dupl
package compensate

import (
	"context"

	converter "github.com/NpoolPlatform/order-manager/pkg/converter/compensate"
	crud "github.com/NpoolPlatform/order-manager/pkg/crud/compensate"
	commontracer "github.com/NpoolPlatform/order-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/order-manager/pkg/tracer/compensate"

	constant "github.com/NpoolPlatform/order-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mgr/v1/compensate"

	"github.com/google/uuid"
)

func (s *Server) CreateCompensate(ctx context.Context, in *npool.CreateCompensateRequest) (*npool.CreateCompensateResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateCompensate")
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
		return &npool.CreateCompensateResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "Compensate", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create Compensate: %v", err.Error())
		return &npool.CreateCompensateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCompensateResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateCompensates(ctx context.Context, in *npool.CreateCompensatesRequest) (*npool.CreateCompensatesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateCompensates")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateCompensatesResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "Compensate", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create Compensates: %v", err)
		return &npool.CreateCompensatesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCompensatesResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateCompensate(ctx context.Context, in *npool.UpdateCompensateRequest) (*npool.UpdateCompensateResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateCompensates")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "Compensate", "crud", "CreateBulk")

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetInfo().GetID())
		return &npool.UpdateCompensateResponse{}, status.Error(codes.InvalidArgument, "ID is invalid")
	}

	rows, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create Compensates: %v", err)
		return &npool.UpdateCompensateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateCompensateResponse{
		Info: converter.Ent2Grpc(rows),
	}, nil
}

func (s *Server) GetCompensate(ctx context.Context, in *npool.GetCompensateRequest) (*npool.GetCompensateResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetCompensate")
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
		return &npool.GetCompensateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "Compensate", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get Compensate: %v", err)
		return &npool.GetCompensateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCompensateResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetCompensateOnly(ctx context.Context, in *npool.GetCompensateOnlyRequest) (*npool.GetCompensateOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetCompensateOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "Compensate", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get Compensates: %v", err)
		return &npool.GetCompensateOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCompensateOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetCompensates(ctx context.Context, in *npool.GetCompensatesRequest) (*npool.GetCompensatesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetCompensates")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "Compensate", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get Compensates: %v", err)
		return &npool.GetCompensatesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCompensatesResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistCompensate(ctx context.Context, in *npool.ExistCompensateRequest) (*npool.ExistCompensateResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistCompensate")
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
		return &npool.ExistCompensateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "Compensate", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check Compensate: %v", err)
		return &npool.ExistCompensateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistCompensateResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistCompensateConds(ctx context.Context,
	in *npool.ExistCompensateCondsRequest) (*npool.ExistCompensateCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistCompensateConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "Compensate", "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check Compensate: %v", err)
		return &npool.ExistCompensateCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistCompensateCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountCompensates(ctx context.Context, in *npool.CountCompensatesRequest) (*npool.CountCompensatesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountCompensates")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "Compensate", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count Compensates: %v", err)
		return &npool.CountCompensatesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountCompensatesResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteCompensate(ctx context.Context, in *npool.DeleteCompensateRequest) (*npool.DeleteCompensateResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateCompensates")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "Compensate", "crud", "CreateBulk")

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetID())
		return &npool.DeleteCompensateResponse{}, status.Error(codes.InvalidArgument, "ID is invalid")
	}

	rows, err := crud.Delete(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorf("fail create Compensates: %v", err)
		return &npool.DeleteCompensateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteCompensateResponse{
		Info: converter.Ent2Grpc(rows),
	}, nil
}
