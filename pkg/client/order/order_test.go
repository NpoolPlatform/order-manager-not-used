package order

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	val "github.com/NpoolPlatform/message/npool"

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

var appDate = npool.Order{
	ID:                     uuid.NewString(),
	GoodID:                 uuid.NewString(),
	AppID:                  uuid.NewString(),
	UserID:                 uuid.NewString(),
	ParentOrderID:          uuid.NewString(),
	PayWithParent:          true,
	Units:                  1001,
	PromotionID:            uuid.NewString(),
	DiscountCouponID:       uuid.NewString(),
	UserSpecialReductionID: uuid.NewString(),
	StartAt:                1002,
	EndAt:                  1003,
	FixAmountCouponID:      uuid.NewString(),
	Type:                   npool.OrderType_Airdrop,
	State:                  npool.OrderState_WaitPayment,
}

var (
	appInfo = npool.OrderReq{
		ID:                     &appDate.ID,
		GoodID:                 &appDate.GoodID,
		AppID:                  &appDate.AppID,
		UserID:                 &appDate.UserID,
		ParentOrderID:          &appDate.ParentOrderID,
		PayWithParent:          &appDate.PayWithParent,
		Units:                  &appDate.Units,
		PromotionID:            &appDate.PromotionID,
		DiscountCouponID:       &appDate.DiscountCouponID,
		UserSpecialReductionID: &appDate.UserSpecialReductionID,
		StartAt:                &appDate.StartAt,
		EndAt:                  &appDate.EndAt,
		FixAmountCouponID:      &appDate.FixAmountCouponID,
		Type:                   &appDate.Type,
		State:                  &appDate.State,
	}
)

var info *npool.Order

func createOrder(t *testing.T) {
	var err error
	info, err = CreateOrder(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		appDate.CreatedAt = info.CreatedAt
		appDate.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &appDate)
	}
}

func createOrders(t *testing.T) {
	appDates := []npool.Order{
		{
			ID:                     uuid.NewString(),
			GoodID:                 uuid.NewString(),
			AppID:                  uuid.NewString(),
			UserID:                 uuid.NewString(),
			ParentOrderID:          uuid.NewString(),
			PayWithParent:          true,
			Units:                  1001,
			PromotionID:            uuid.NewString(),
			DiscountCouponID:       uuid.NewString(),
			UserSpecialReductionID: uuid.NewString(),
			StartAt:                1002,
			EndAt:                  1003,
			FixAmountCouponID:      uuid.NewString(),
			Type:                   npool.OrderType_Airdrop,
			State:                  npool.OrderState_WaitPayment,
		},
		{
			ID:                     uuid.NewString(),
			GoodID:                 uuid.NewString(),
			AppID:                  uuid.NewString(),
			UserID:                 uuid.NewString(),
			ParentOrderID:          uuid.NewString(),
			PayWithParent:          true,
			Units:                  1001,
			PromotionID:            uuid.NewString(),
			DiscountCouponID:       uuid.NewString(),
			UserSpecialReductionID: uuid.NewString(),
			StartAt:                1002,
			EndAt:                  1003,
			FixAmountCouponID:      uuid.NewString(),
			Type:                   npool.OrderType_Airdrop,
			State:                  npool.OrderState_WaitPayment,
		},
	}

	apps := []*npool.OrderReq{}
	for key := range appDates {
		apps = append(apps, &npool.OrderReq{
			ID:                     &appDates[key].ID,
			GoodID:                 &appDates[key].GoodID,
			AppID:                  &appDates[key].AppID,
			UserID:                 &appDates[key].UserID,
			ParentOrderID:          &appDates[key].ParentOrderID,
			PayWithParent:          &appDates[key].PayWithParent,
			Units:                  &appDates[key].Units,
			PromotionID:            &appDates[key].PromotionID,
			DiscountCouponID:       &appDates[key].DiscountCouponID,
			UserSpecialReductionID: &appDates[key].UserSpecialReductionID,
			StartAt:                &appDates[key].StartAt,
			EndAt:                  &appDates[key].EndAt,
			FixAmountCouponID:      &appDates[key].FixAmountCouponID,
			Type:                   &appDates[key].Type,
			State:                  &appDates[key].State,
		})
	}

	infos, err := CreateOrders(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updateOrder(t *testing.T) {
	var err error
	info, err = UpdateOrder(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		appDate.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &appDate)
	}
}

func getOrder(t *testing.T) {
	var err error
	info, err = GetOrder(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appDate)
	}
}

func getOrders(t *testing.T) {
	infos, total, err := GetOrders(context.Background(),
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

func getOrderOnly(t *testing.T) {
	var err error
	info, err = GetOrderOnly(context.Background(),
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

func existOrder(t *testing.T) {
	exist, err := ExistOrder(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existOrderConds(t *testing.T) {
	exist, err := ExistOrderConds(context.Background(),
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

func deleteOrder(t *testing.T) {
	info, err := DeleteOrder(context.Background(), info.ID)
	if assert.Nil(t, err) {
		appDate.DeletedAt = info.DeletedAt
		assert.Equal(t, info, &appDate)
	}
}

func TestMainOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace(uuid.NewString(), config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createOrder", createOrder)
	t.Run("createOrders", createOrders)
	t.Run("getOrder", getOrder)
	t.Run("getOrders", getOrders)
	t.Run("getOrderOnly", getOrderOnly)
	t.Run("updateOrder", updateOrder)
	t.Run("existOrder", existOrder)
	t.Run("existOrderConds", existOrderConds)
	t.Run("delete", deleteOrder)
}
