//nolint:dupl
package order

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/order-manager/pkg/db/ent"

	valuedef "github.com/NpoolPlatform/message/npool"
	npool "github.com/NpoolPlatform/message/npool/order/mgr/v1/order"

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

var appGood = ent.Order{
	ID:                     uuid.New(),
	GoodID:                 uuid.New(),
	AppID:                  uuid.New(),
	UserID:                 uuid.New(),
	ParentOrderID:          uuid.New(),
	PayWithParent:          true,
	UnitsV1:                decimal.NewFromInt(100),
	Units:                  0,
	PromotionID:            uuid.New(),
	DiscountCouponID:       uuid.New(),
	UserSpecialReductionID: uuid.New(),
	StartAt:                uint32(time.Now().Unix()),
	EndAt:                  uint32(time.Now().Add(100 * time.Hour).Unix()),
	FixAmountCouponID:      uuid.New(),
	Type:                   npool.OrderType_Airdrop.String(),
	State:                  npool.OrderState_InService.String(),
}

var (
	id                     = appGood.ID.String()
	appID                  = appGood.AppID.String()
	orderID                = appGood.GoodID.String()
	userID                 = appGood.UserID.String()
	parentOrderID          = appGood.ParentOrderID.String()
	promotionID            = appGood.PromotionID.String()
	discountCouponID       = appGood.DiscountCouponID.String()
	userSpecialReductionID = appGood.UserSpecialReductionID.String()
	couponID               = appGood.FixAmountCouponID.String()
	tp                     = npool.OrderType_Airdrop
	state                  = npool.OrderState_InService
	units                  = appGood.UnitsV1.String()
	req                    = npool.OrderReq{
		ID:                     &id,
		GoodID:                 &orderID,
		AppID:                  &appID,
		UserID:                 &userID,
		ParentOrderID:          &parentOrderID,
		PayWithParent:          &appGood.PayWithParent,
		Units:                  &units,
		PromotionID:            &promotionID,
		DiscountCouponID:       &discountCouponID,
		UserSpecialReductionID: &userSpecialReductionID,
		StartAt:                &appGood.StartAt,
		EndAt:                  &appGood.EndAt,
		FixAmountCouponID:      &couponID,
		Type:                   &tp,
		State:                  &state,
	}
)

var info *ent.Order

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
	entities := []*ent.Order{
		{
			ID:                     uuid.New(),
			GoodID:                 uuid.New(),
			AppID:                  uuid.New(),
			UserID:                 uuid.New(),
			ParentOrderID:          uuid.New(),
			PayWithParent:          true,
			UnitsV1:                decimal.NewFromInt(100),
			Units:                  0,
			PromotionID:            uuid.New(),
			DiscountCouponID:       uuid.New(),
			UserSpecialReductionID: uuid.New(),
			StartAt:                uint32(time.Now().Unix()),
			EndAt:                  uint32(time.Now().Add(100 * time.Hour).Unix()),
			FixAmountCouponID:      uuid.New(),
			Type:                   npool.OrderType_Airdrop.String(),
			State:                  npool.OrderState_InService.String(),
		},
		{
			ID:                     uuid.New(),
			GoodID:                 uuid.New(),
			AppID:                  uuid.New(),
			UserID:                 uuid.New(),
			ParentOrderID:          uuid.New(),
			PayWithParent:          true,
			UnitsV1:                decimal.NewFromInt(100),
			Units:                  0,
			PromotionID:            uuid.New(),
			DiscountCouponID:       uuid.New(),
			UserSpecialReductionID: uuid.New(),
			StartAt:                uint32(time.Now().Unix()),
			EndAt:                  uint32(time.Now().Add(100 * time.Hour).Unix()),
			FixAmountCouponID:      uuid.New(),
			Type:                   npool.OrderType_Airdrop.String(),
			State:                  npool.OrderState_InService.String(),
		},
	}

	reqs := []*npool.OrderReq{}
	for _, _appGood := range entities {
		_id := _appGood.ID.String()
		_appID := _appGood.AppID.String()
		_orderID := _appGood.GoodID.String()
		_userID := _appGood.UserID.String()
		_parentOrderID := _appGood.ParentOrderID.String()
		_promotionID := _appGood.PromotionID.String()
		_discountCouponID := _appGood.DiscountCouponID.String()
		_userSpecialReductionID := _appGood.UserSpecialReductionID.String()
		_couponID := _appGood.FixAmountCouponID.String()
		_tp := npool.OrderType_Airdrop
		_state := npool.OrderState_InService
		_units := appGood.UnitsV1.String()
		reqs = append(reqs, &npool.OrderReq{
			ID:                     &_id,
			GoodID:                 &_orderID,
			AppID:                  &_appID,
			UserID:                 &_userID,
			ParentOrderID:          &_parentOrderID,
			PayWithParent:          &_appGood.PayWithParent,
			Units:                  &_units,
			PromotionID:            &_promotionID,
			DiscountCouponID:       &_discountCouponID,
			UserSpecialReductionID: &_userSpecialReductionID,
			StartAt:                &_appGood.StartAt,
			EndAt:                  &_appGood.EndAt,
			FixAmountCouponID:      &_couponID,
			Type:                   &_tp,
			State:                  &_state,
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
