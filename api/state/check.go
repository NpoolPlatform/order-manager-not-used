package state

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mgr/v1/order/state"
)

func validate(info *npool.StateReq) error {
	if info.OrderID == nil {
		logger.Sugar().Errorw("validate", "OrderID", info.OrderID)
		return status.Error(codes.InvalidArgument, "OrderID is empty")
	}

	if info.State == nil {
		logger.Sugar().Errorw("validate", "State", info.State)
		return status.Error(codes.InvalidArgument, "State is empty")
	}

	switch info.GetState() {
	case npool.EState_WaitPayment:
	case npool.EState_Paid:
	case npool.EState_PaymentTimeout:
	case npool.EState_Canceled:
	case npool.EState_InService:
	case npool.EState_Expired:
	default:
		logger.Sugar().Errorw("validate", "State", info.GetState())
		return status.Error(codes.InvalidArgument, "State is invalid")
	}

	return nil
}

func duplicate(infos []*npool.StateReq) error {
	keys := map[string]struct{}{}

	for _, info := range infos {
		if err := validate(info); err != nil {
			return status.Error(codes.InvalidArgument, fmt.Sprintf("Infos has invalid element %v", err))
		}

		key := fmt.Sprintf("%v", info.OrderID)
		if _, ok := keys[key]; ok {
			return status.Error(codes.InvalidArgument, "Infos has duplicate OrderID")
		}

		keys[key] = struct{}{}
	}

	return nil
}
