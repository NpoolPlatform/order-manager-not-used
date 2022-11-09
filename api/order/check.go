package order

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mgr/v1/order"

	"github.com/google/uuid"
)

//nolint
func validate(info *npool.OrderReq) error {
	if info.AppID == nil {
		logger.Sugar().Errorw("validate", "AppID", info.GetAppID())
		return status.Error(codes.InvalidArgument, "AppID is empty")
	}

	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", info.GetAppID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("AppID is invalid: %v", err))
	}

	if info.GoodID == nil {
		logger.Sugar().Errorw("validate", "GoodID", info.GetGoodID())
		return status.Error(codes.InvalidArgument, "GoodID is empty")
	}

	if _, err := uuid.Parse(info.GetGoodID()); err != nil {
		logger.Sugar().Errorw("validate", "GoodID", info.GetGoodID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("GoodID is invalid: %v", err))
	}

	if info.UserID == nil {
		logger.Sugar().Errorw("validate", "UserID", info.GetUserID())
		return status.Error(codes.InvalidArgument, "UserID is empty")
	}

	if _, err := uuid.Parse(info.GetUserID()); err != nil {
		logger.Sugar().Errorw("validate", "UserID", info.GetUserID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("UserID is invalid: %v", err))
	}

	if info.ParentOrderID == nil {
		logger.Sugar().Errorw("validate", "ParentOrderID", info.GetParentOrderID())
		return status.Error(codes.InvalidArgument, "ParentOrderID is empty")
	}

	if _, err := uuid.Parse(info.GetParentOrderID()); err != nil {
		logger.Sugar().Errorw("validate", "ParentOrderIDID", info.GetParentOrderID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("ParentOrderID is invalid: %v", err))
	}

	if info.GetUnits() == 0 {
		logger.Sugar().Errorw("validate", "Units", info.GetUnits())
		return status.Error(codes.InvalidArgument, "Units is zero or empty")
	}

	if info.PromotionID == nil {
		logger.Sugar().Errorw("validate", "PromotionID", info.GetPromotionID())
		return status.Error(codes.InvalidArgument, "PromotionID is empty")
	}

	if _, err := uuid.Parse(info.GetPromotionID()); err != nil {
		logger.Sugar().Errorw("validate", "PromotionIDID", info.GetPromotionID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("PromotionID is invalid: %v", err))
	}

	if info.DiscountCouponID == nil {
		logger.Sugar().Errorw("validate", "DiscountCouponID", info.GetDiscountCouponID())
		return status.Error(codes.InvalidArgument, "DiscountCouponID is empty")
	}

	if _, err := uuid.Parse(info.GetDiscountCouponID()); err != nil {
		logger.Sugar().Errorw("validate", "DiscountCouponIDID", info.GetDiscountCouponID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("DiscountCouponID is invalid: %v", err))
	}

	if info.UserSpecialReductionID == nil {
		logger.Sugar().Errorw("validate", "UserSpecialReductionID", info.GetUserSpecialReductionID())
		return status.Error(codes.InvalidArgument, "UserSpecialReductionID is empty")
	}

	if _, err := uuid.Parse(info.GetUserSpecialReductionID()); err != nil {
		logger.Sugar().Errorw("validate", "UserSpecialReductionIDID", info.GetUserSpecialReductionID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("UserSpecialReductionID is invalid: %v", err))
	}

	if info.GetStartAt() == 0 {
		logger.Sugar().Errorw("validate", "StartAt", info.GetStartAt())
		return status.Error(codes.InvalidArgument, "Start is zero or empty")
	}

	if info.GetEndAt() == 0 {
		logger.Sugar().Errorw("validate", "End", info.GetUnits())
		return status.Error(codes.InvalidArgument, "End is zero or empty")
	}

	switch info.GetType() {
	case npool.OrderType_Normal:
	case npool.OrderType_Offline:
	case npool.OrderType_Airdrop:
	default:
		logger.Sugar().Errorw("validate", "OrderType", info.GetType())
		return status.Error(codes.InvalidArgument, "OrderType is invalid")
	}

	switch info.GetState() {
	case npool.OrderState_WaitPayment:
	default:
		logger.Sugar().Errorw("validate", "State", info.GetState())
		return status.Error(codes.InvalidArgument, "State is invalid")
	}

	return nil
}
