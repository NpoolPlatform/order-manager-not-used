package payment

import (
	"fmt"

	"github.com/shopspring/decimal"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mgr/v1/payment"

	"github.com/google/uuid"
)

//nolint
func validate(info *npool.PaymentReq) error {
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

	if info.OrderID == nil {
		logger.Sugar().Errorw("validate", "OrderID", info.GetOrderID())
		return status.Error(codes.InvalidArgument, "OrderID is empty")
	}

	if _, err := uuid.Parse(info.GetOrderID()); err != nil {
		logger.Sugar().Errorw("validate", "OrderID", info.GetOrderID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("OrderID is invalid: %v", err))
	}

	if info.AccountID == nil {
		logger.Sugar().Errorw("validate", "AccountID", info.GetAccountID())
		return status.Error(codes.InvalidArgument, "AccountID is empty")
	}

	if _, err := uuid.Parse(info.GetAccountID()); err != nil {
		logger.Sugar().Errorw("validate", "AccountID", info.GetAccountID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("AccountID is invalid: %v", err))
	}

	if info.StartAmount == nil {
		logger.Sugar().Errorw("validate", "StartAmount", info.StartAmount)
		return status.Error(codes.InvalidArgument, "StartAmount is empty")
	}

	startAmount, err := decimal.NewFromString(info.GetStartAmount())
	if err != nil {
		logger.Sugar().Errorw("validate", "StartAmount", info.GetStartAmount(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("StartAmount is invalid: %v", err))
	}

	if startAmount.Cmp(decimal.NewFromInt(0)) <= 0 {
		logger.Sugar().Errorw("validate", "StartAmount", info.GetStartAmount(), "error", "less than 0")
		return status.Error(codes.InvalidArgument, "GetStartAmount is Less than or equal to 0")
	}

	if info.Amount == nil {
		logger.Sugar().Errorw("validate", "Amount", info.Amount)
		return status.Error(codes.InvalidArgument, "Amount is empty")
	}

	amount, err := decimal.NewFromString(info.GetAmount())
	if err != nil {
		logger.Sugar().Errorw("validate", "Amount", info.GetAmount(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("Amount is invalid: %v", err))
	}

	if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
		logger.Sugar().Errorw("validate", "Amount", info.GetAmount(), "error", "less than 0")
		return status.Error(codes.InvalidArgument, "GetAmount is Less than or equal to 0")
	}

	if info.FinishAmount == nil {
		logger.Sugar().Errorw("validate", "FinishAmount", info.FinishAmount)
		return status.Error(codes.InvalidArgument, "FinishAmount is empty")
	}

	finishAmount, err := decimal.NewFromString(info.GetFinishAmount())
	if err != nil {
		logger.Sugar().Errorw("validate", "FinishAmount", info.GetFinishAmount(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("FinishAmount is invalid: %v", err))
	}

	if finishAmount.Cmp(decimal.NewFromInt(0)) <= 0 {
		logger.Sugar().Errorw("validate", "FinishAmount", info.GetFinishAmount(), "error", "less than 0")
		return status.Error(codes.InvalidArgument, "GetFinishAmount is Less than or equal to 0")
	}

	if info.CoinUsdCurrency == nil {
		logger.Sugar().Errorw("validate", "CoinUsdCurrency", info.CoinUsdCurrency)
		return status.Error(codes.InvalidArgument, "CoinUsdCurrency is empty")
	}

	coinUsdCurrency, err := decimal.NewFromString(info.GetCoinUsdCurrency())
	if err != nil {
		logger.Sugar().Errorw("validate", "CoinUsdCurrency", info.GetCoinUsdCurrency(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("CoinUsdCurrency is invalid: %v", err))
	}

	if coinUsdCurrency.Cmp(decimal.NewFromInt(0)) <= 0 {
		logger.Sugar().Errorw("validate", "CoinUsdCurrency", info.GetCoinUsdCurrency(), "error", "less than 0")
		return status.Error(codes.InvalidArgument, "GetCoinUsdCurrency is Less than or equal to 0")
	}

	if info.LocalCoinUsdCurrency == nil {
		logger.Sugar().Errorw("validate", "LocalCoinUsdCurrency", info.LocalCoinUsdCurrency)
		return status.Error(codes.InvalidArgument, "LocalCoinUsdCurrency is empty")
	}

	localCoinUsdCurrency, err := decimal.NewFromString(info.GetLocalCoinUsdCurrency())
	if err != nil {
		logger.Sugar().Errorw("validate", "LocalCoinUsdCurrency", info.GetLocalCoinUsdCurrency(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("LocalCoinUsdCurrency is invalid: %v", err))
	}

	if localCoinUsdCurrency.Cmp(decimal.NewFromInt(0)) <= 0 {
		logger.Sugar().Errorw("validate", "LocalCoinUsdCurrency", info.GetLocalCoinUsdCurrency(), "error", "less than 0")
		return status.Error(codes.InvalidArgument, "GetLocalCoinUsdCurrency is Less than or equal to 0")
	}

	if info.LiveCoinUsdCurrency == nil {
		logger.Sugar().Errorw("validate", "LiveCoinUsdCurrency", info.LiveCoinUsdCurrency)
		return status.Error(codes.InvalidArgument, "LiveCoinUsdCurrency is empty")
	}

	liveCoinUsdCurrency, err := decimal.NewFromString(info.GetLiveCoinUsdCurrency())
	if err != nil {
		logger.Sugar().Errorw("validate", "LiveCoinUsdCurrency", info.GetLiveCoinUsdCurrency(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("LiveCoinUsdCurrency is invalid: %v", err))
	}

	if liveCoinUsdCurrency.Cmp(decimal.NewFromInt(0)) <= 0 {
		logger.Sugar().Errorw("validate", "LiveCoinUsdCurrency", info.GetLiveCoinUsdCurrency(), "error", "less than 0")
		return status.Error(codes.InvalidArgument, "GetLiveCoinUsdCurrency is Less than or equal to 0")
	}

	if info.CoinInfoID == nil {
		logger.Sugar().Errorw("validate", "CoinInfoID", info.GetCoinInfoID())
		return status.Error(codes.InvalidArgument, "CoinInfoID is empty")
	}

	if _, err := uuid.Parse(info.GetCoinInfoID()); err != nil {
		logger.Sugar().Errorw("validate", "CoinInfoIDID", info.GetCoinInfoID(), "error", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("CoinInfoID is invalid: %v", err))
	}

	switch info.GetState() {
	case npool.PaymentState_Wait:
	default:
		logger.Sugar().Errorw("validate", "PaymentState", info.GetState())
		return status.Error(codes.InvalidArgument, fmt.Sprintf("PaymentState is invalid"))
	}

	return nil
}
