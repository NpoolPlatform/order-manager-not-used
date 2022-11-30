//nolint:dupl
package payment

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/shopspring/decimal"

	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	val "github.com/NpoolPlatform/message/npool"

	npool "github.com/NpoolPlatform/message/npool/order/mgr/v1/payment"

	testinit "github.com/NpoolPlatform/order-manager/pkg/testinit"
	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var appDate = npool.Payment{
	ID:                   uuid.NewString(),
	AppID:                uuid.NewString(),
	UserID:               uuid.NewString(),
	GoodID:               uuid.NewString(),
	OrderID:              uuid.NewString(),
	AccountID:            uuid.NewString(),
	StartAmount:          decimal.NewFromInt(1000).String(),
	Amount:               decimal.NewFromInt(1000).String(),
	PayWithBalanceAmount: decimal.NewFromInt(1000).String(),
	FinishAmount:         decimal.NewFromInt(1000).String(),
	CoinUsdCurrency:      decimal.NewFromInt(1000).String(),
	LocalCoinUsdCurrency: decimal.NewFromInt(1000).String(),
	LiveCoinUsdCurrency:  decimal.NewFromInt(1000).String(),
	CoinInfoID:           uuid.NewString(),
	State:                npool.PaymentState_Wait,
	ChainTransactionID:   uuid.NewString(),
	UserSetPaid:          true,
	UserSetCanceled:      true,
	FakePayment:          true,
}

var (
	appInfo = npool.PaymentReq{
		ID:                   &appDate.ID,
		AppID:                &appDate.AppID,
		UserID:               &appDate.UserID,
		GoodID:               &appDate.GoodID,
		OrderID:              &appDate.OrderID,
		AccountID:            &appDate.AccountID,
		StartAmount:          &appDate.StartAmount,
		Amount:               &appDate.Amount,
		PayWithBalanceAmount: &appDate.PayWithBalanceAmount,
		FinishAmount:         &appDate.FinishAmount,
		CoinUsdCurrency:      &appDate.CoinUsdCurrency,
		LocalCoinUsdCurrency: &appDate.LocalCoinUsdCurrency,
		LiveCoinUsdCurrency:  &appDate.LiveCoinUsdCurrency,
		CoinInfoID:           &appDate.CoinInfoID,
		State:                &appDate.State,
		ChainTransactionID:   &appDate.ChainTransactionID,
		UserSetPaid:          &appDate.UserSetPaid,
		UserSetCanceled:      &appDate.UserSetCanceled,
		FakePayment:          &appDate.FakePayment,
	}
)

var info *npool.Payment

func createPayment(t *testing.T) {
	var err error
	info, err = CreatePayment(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		appDate.CreatedAt = info.CreatedAt
		appDate.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &appDate)
	}
}

func createPayments(t *testing.T) {
	appDates := []npool.Payment{
		{
			ID:                   uuid.NewString(),
			AppID:                uuid.NewString(),
			UserID:               uuid.NewString(),
			GoodID:               uuid.NewString(),
			OrderID:              uuid.NewString(),
			AccountID:            uuid.NewString(),
			StartAmount:          decimal.NewFromInt(1000).String(),
			Amount:               decimal.NewFromInt(1000).String(),
			PayWithBalanceAmount: decimal.NewFromInt(1000).String(),
			FinishAmount:         decimal.NewFromInt(1000).String(),
			CoinUsdCurrency:      decimal.NewFromInt(1000).String(),
			LocalCoinUsdCurrency: decimal.NewFromInt(1000).String(),
			LiveCoinUsdCurrency:  decimal.NewFromInt(1000).String(),
			CoinInfoID:           uuid.NewString(),
			State:                npool.PaymentState_Wait,
			ChainTransactionID:   uuid.NewString(),
			UserSetPaid:          true,
			UserSetCanceled:      true,
			FakePayment:          true,
		},
		{
			ID:                   uuid.NewString(),
			AppID:                uuid.NewString(),
			UserID:               uuid.NewString(),
			GoodID:               uuid.NewString(),
			OrderID:              uuid.NewString(),
			AccountID:            uuid.NewString(),
			StartAmount:          decimal.NewFromInt(1000).String(),
			Amount:               decimal.NewFromInt(1000).String(),
			PayWithBalanceAmount: decimal.NewFromInt(1000).String(),
			FinishAmount:         decimal.NewFromInt(1000).String(),
			CoinUsdCurrency:      decimal.NewFromInt(1000).String(),
			LocalCoinUsdCurrency: decimal.NewFromInt(1000).String(),
			LiveCoinUsdCurrency:  decimal.NewFromInt(1000).String(),
			CoinInfoID:           uuid.NewString(),
			State:                npool.PaymentState_Wait,
			ChainTransactionID:   uuid.NewString(),
			UserSetPaid:          true,
			UserSetCanceled:      true,
			FakePayment:          true,
		},
	}

	apps := []*npool.PaymentReq{}
	for key := range appDates {
		apps = append(apps, &npool.PaymentReq{
			ID:                   &appDates[key].ID,
			AppID:                &appDates[key].AppID,
			UserID:               &appDates[key].UserID,
			GoodID:               &appDates[key].GoodID,
			OrderID:              &appDates[key].OrderID,
			AccountID:            &appDates[key].AccountID,
			StartAmount:          &appDates[key].StartAmount,
			Amount:               &appDates[key].Amount,
			PayWithBalanceAmount: &appDates[key].PayWithBalanceAmount,
			FinishAmount:         &appDates[key].FinishAmount,
			CoinUsdCurrency:      &appDates[key].CoinUsdCurrency,
			LocalCoinUsdCurrency: &appDates[key].LocalCoinUsdCurrency,
			LiveCoinUsdCurrency:  &appDates[key].LiveCoinUsdCurrency,
			CoinInfoID:           &appDates[key].CoinInfoID,
			State:                &appDates[key].State,
			ChainTransactionID:   &appDates[key].ChainTransactionID,
			UserSetPaid:          &appDates[key].UserSetPaid,
			UserSetCanceled:      &appDates[key].UserSetCanceled,
			FakePayment:          &appDates[key].FakePayment,
		})
	}

	infos, err := CreatePayments(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updatePayment(t *testing.T) {
	var err error
	info, err = UpdatePayment(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		appDate.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &appDate)
	}
}

func getPayment(t *testing.T) {
	var err error
	info, err = GetPayment(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appDate)
	}
}

func getPayments(t *testing.T) {
	infos, total, err := GetPayments(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, infos[0], &appDate)
	}
}

func getPaymentOnly(t *testing.T) {
	var err error
	info, err = GetPaymentOnly(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appDate)
	}
}

func existPayment(t *testing.T) {
	exist, err := ExistPayment(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existPaymentConds(t *testing.T) {
	exist, err := ExistPaymentConds(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deletePayment(t *testing.T) {
	info, err := DeletePayment(context.Background(), info.ID)
	if assert.Nil(t, err) {
		appDate.DeletedAt = info.DeletedAt
		assert.Equal(t, info, &appDate)
	}
}

func TestMainPayment(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace(uuid.NewString(), config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createPayment", createPayment)
	t.Run("createPayments", createPayments)
	t.Run("getPayment", getPayment)
	t.Run("getPayments", getPayments)
	t.Run("getPaymentOnly", getPaymentOnly)
	t.Run("updatePayment", updatePayment)
	t.Run("existPayment", existPayment)
	t.Run("existPaymentConds", existPaymentConds)
	t.Run("delete", deletePayment)
}
