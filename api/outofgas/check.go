package outofgas

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mgr/v1/outofgas"

	"github.com/google/uuid"
)

func validate(info *npool.OutOfGasReq) error {
	if info.OrderID == nil {
		logger.Sugar().Errorw("validate", "OrderID", info.GetOrderID())
		return status.Error(codes.InvalidArgument, "OrderID is empty")
	}

	if _, err := uuid.Parse(info.GetOrderID()); err != nil {
		logger.Sugar().Errorw("validate", "OrderID", info.GetOrderID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("OrderID is invalid: %v", err))
	}

	if info.GetStart() == 0 {
		logger.Sugar().Errorw("validate", "Start", info.GetStart())
		return status.Error(codes.InvalidArgument, "Start is zero or empty")
	}

	if info.GetEnd() == 0 {
		logger.Sugar().Errorw("validate", "End", info.GetEnd())
		return status.Error(codes.InvalidArgument, "End is zero or empty")
	}
	return nil
}
