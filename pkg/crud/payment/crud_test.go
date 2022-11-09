package payment

import (
	"context"
	"fmt"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/order-manager/pkg/db/ent"
	"github.com/shopspring/decimal"
	"os"
	"strconv"
	"testing"

	valuedef "github.com/NpoolPlatform/message/npool"
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

var appGood = ent.Payment{
	ID:                   uuid.New(),
	AppID:                uuid.New(),
	UserID:               uuid.New(),
	GoodID:               uuid.New(),
	OrderID:              uuid.New(),
	AccountID:            uuid.New(),
	StartAmount:          decimal.NewFromInt(1000),
	Amount:               decimal.NewFromInt(1000),
	PayWithBalanceAmount: decimal.NewFromInt(1000),
	FinishAmount:         decimal.NewFromInt(1000),
	CoinUsdCurrency:      decimal.NewFromInt(1000),
	LocalCoinUsdCurrency: decimal.NewFromInt(1000),
	LiveCoinUsdCurrency:  decimal.NewFromInt(1000),
	CoinInfoID:           uuid.New(),
	State:                npool.PaymentState_Done.String(),
	ChainTransactionID:   uuid.NewString(),
	UserSetPaid:          true,
	UserSetCanceled:      true,
	FakePayment:          true,
}

var (
	id                   = appGood.ID.String()
	appID                = appGood.AppID.String()
	goodID               = appGood.GoodID.String()
	orderID              = appGood.OrderID.String()
	userID               = appGood.UserID.String()
	accountID            = appGood.AccountID.String()
	startAmount          = appGood.StartAmount.String()
	amount               = appGood.Amount.String()
	payWithBalanceAmount = appGood.PayWithBalanceAmount.String()
	finishAmount         = appGood.FinishAmount.String()
	coinUsdCurrency      = appGood.CoinUsdCurrency.String()
	localCoinUsdCurrency = appGood.LocalCoinUsdCurrency.String()
	liveCoinUsdCurrency  = appGood.LiveCoinUsdCurrency.String()
	coinInfoID           = appGood.CoinInfoID.String()
	state                = npool.PaymentState_Done

	req = npool.PaymentReq{
		ID:                   &id,
		AppID:                &appID,
		UserID:               &userID,
		GoodID:               &goodID,
		OrderID:              &orderID,
		AccountID:            &accountID,
		StartAmount:          &startAmount,
		Amount:               &amount,
		PayWithBalanceAmount: &payWithBalanceAmount,
		FinishAmount:         &finishAmount,
		CoinUsdCurrency:      &coinUsdCurrency,
		LocalCoinUsdCurrency: &localCoinUsdCurrency,
		LiveCoinUsdCurrency:  &liveCoinUsdCurrency,
		CoinInfoID:           &coinInfoID,
		State:                &state,
		ChainTransactionID:   &appGood.ChainTransactionID,
		UserSetPaid:          &appGood.UserSetPaid,
		UserSetCanceled:      &appGood.UserSetCanceled,
		FakePayment:          &appGood.FakePayment,
	}
)

var info *ent.Payment

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &req)
	if assert.Nil(t, err) {
		appGood.UpdatedAt = info.UpdatedAt
		appGood.CreatedAt = info.CreatedAt
		assert.Equal(t, info.String(), appGood.String())
	}
}

func createBulk(t *testing.T) {
	entities := []*ent.Payment{
		{
			ID:                   uuid.New(),
			AppID:                uuid.New(),
			UserID:               uuid.New(),
			GoodID:               uuid.New(),
			OrderID:              uuid.New(),
			AccountID:            uuid.New(),
			StartAmount:          decimal.NewFromInt(1000),
			Amount:               decimal.NewFromInt(1000),
			PayWithBalanceAmount: decimal.NewFromInt(1000),
			FinishAmount:         decimal.NewFromInt(1000),
			CoinUsdCurrency:      decimal.NewFromInt(1000),
			LocalCoinUsdCurrency: decimal.NewFromInt(1000),
			LiveCoinUsdCurrency:  decimal.NewFromInt(1000),
			CoinInfoID:           uuid.New(),
			State:                npool.PaymentState_Done.String(),
			ChainTransactionID:   uuid.NewString(),
			UserSetPaid:          true,
			UserSetCanceled:      true,
			FakePayment:          true,
		},
		{
			ID:                   uuid.New(),
			AppID:                uuid.New(),
			UserID:               uuid.New(),
			GoodID:               uuid.New(),
			OrderID:              uuid.New(),
			AccountID:            uuid.New(),
			StartAmount:          decimal.NewFromInt(1000),
			Amount:               decimal.NewFromInt(1000),
			PayWithBalanceAmount: decimal.NewFromInt(1000),
			FinishAmount:         decimal.NewFromInt(1000),
			CoinUsdCurrency:      decimal.NewFromInt(1000),
			LocalCoinUsdCurrency: decimal.NewFromInt(1000),
			LiveCoinUsdCurrency:  decimal.NewFromInt(1000),
			CoinInfoID:           uuid.New(),
			State:                npool.PaymentState_Done.String(),
			ChainTransactionID:   uuid.NewString(),
			UserSetPaid:          true,
			UserSetCanceled:      true,
			FakePayment:          true,
		},
	}

	reqs := []*npool.PaymentReq{}
	for _, _appGood := range entities {
		_id := _appGood.ID.String()
		_appID := _appGood.AppID.String()
		_goodID := _appGood.GoodID.String()
		_orderID := _appGood.OrderID.String()
		_userID := _appGood.UserID.String()
		_accountID := _appGood.AccountID.String()
		_startAmount := _appGood.StartAmount.String()
		_amount := _appGood.Amount.String()
		_payWithBalanceAmount := _appGood.PayWithBalanceAmount.String()
		_finishAmount := _appGood.FinishAmount.String()
		_coinUsdCurrency := _appGood.CoinUsdCurrency.String()
		_localCoinUsdCurrency := _appGood.LocalCoinUsdCurrency.String()
		_liveCoinUsdCurrency := _appGood.LiveCoinUsdCurrency.String()
		_coinInfoID := _appGood.CoinInfoID.String()
		_state := npool.PaymentState_Done
		reqs = append(reqs, &npool.PaymentReq{
			ID:                   &_id,
			AppID:                &_appID,
			UserID:               &_userID,
			GoodID:               &_goodID,
			OrderID:              &_orderID,
			AccountID:            &_accountID,
			StartAmount:          &_startAmount,
			Amount:               &_amount,
			PayWithBalanceAmount: &_payWithBalanceAmount,
			FinishAmount:         &_finishAmount,
			CoinUsdCurrency:      &_coinUsdCurrency,
			LocalCoinUsdCurrency: &_localCoinUsdCurrency,
			LiveCoinUsdCurrency:  &_liveCoinUsdCurrency,
			CoinInfoID:           &_coinInfoID,
			State:                &_state,
			ChainTransactionID:   &_appGood.ChainTransactionID,
			UserSetPaid:          &_appGood.UserSetPaid,
			UserSetCanceled:      &_appGood.UserSetCanceled,
			FakePayment:          &_appGood.FakePayment,
		})
	}
	infos, err := CreateBulk(context.Background(), reqs)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func update(t *testing.T) {
	var err error
	info, err = Update(context.Background(), &req)
	if assert.Nil(t, err) {
		appGood.UpdatedAt = info.UpdatedAt
		appGood.CreatedAt = info.CreatedAt
		assert.Equal(t, info.String(), appGood.String())
	}
}
func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), appGood.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), appGood.String())
	}
}

func rows(t *testing.T) {
	infos, total, err := Rows(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		}, 0, 0)
	if assert.Nil(t, err) {
		if assert.Equal(t, total, 1) {
			assert.Equal(t, infos[0].String(), appGood.String())
		}
	}
}

func rowOnly(t *testing.T) {
	var err error
	info, err = RowOnly(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), appGood.String())
	}
}

func count(t *testing.T) {
	count, err := Count(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, count, uint32(1))
	}
}

func exist(t *testing.T) {
	exist, err := Exist(context.Background(), appGood.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existConds(t *testing.T) {
	exist, err := ExistConds(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteA(t *testing.T) {
	info, err := Delete(context.Background(), appGood.ID.String())
	if assert.Nil(t, err) {
		appGood.DeletedAt = info.DeletedAt
		appGood.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info.String(), appGood.String())
	}
}

func TestDetail(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("create", create)
	t.Run("createBulk", createBulk)
	t.Run("update", update)
	t.Run("row", row)
	t.Run("rows", rows)
	t.Run("rowOnly", rowOnly)
	t.Run("exist", exist)
	t.Run("existConds", existConds)
	t.Run("count", count)
	t.Run("delete", deleteA)
}
