package outofgas

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

	npool "github.com/NpoolPlatform/message/npool/order/mgr/v1/outofgas"

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

var appDate = npool.OutOfGas{
	ID:      uuid.NewString(),
	OrderID: uuid.NewString(),
	Start:   10000,
	End:     10001,
}

var (
	appInfo = npool.OutOfGasReq{
		ID:      &appDate.ID,
		OrderID: &appDate.OrderID,
		Start:   &appDate.Start,
		End:     &appDate.End,
	}
)

var info *npool.OutOfGas

func createOutOfGas(t *testing.T) {
	var err error
	info, err = CreateOutOfGas(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		appDate.CreatedAt = info.CreatedAt
		appDate.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &appDate)
	}
}

func createOutOfGass(t *testing.T) {
	appDates := []npool.OutOfGas{
		{
			ID:      uuid.NewString(),
			OrderID: uuid.NewString(),
			Start:   10000,
			End:     10001,
		},
		{
			ID:      uuid.NewString(),
			OrderID: uuid.NewString(),
			Start:   10000,
			End:     10001,
		},
	}

	apps := []*npool.OutOfGasReq{}
	for key := range appDates {
		apps = append(apps, &npool.OutOfGasReq{
			ID:      &appDates[key].ID,
			OrderID: &appDates[key].OrderID,
			Start:   &appDates[key].Start,
			End:     &appDates[key].End,
		})
	}

	infos, err := CreateOutOfGass(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updateOutOfGas(t *testing.T) {
	var err error
	info, err = UpdateOutOfGas(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		appDate.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &appDate)
	}
}

func getOutOfGas(t *testing.T) {
	var err error
	info, err = GetOutOfGas(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appDate)
	}
}

func getOutOfGass(t *testing.T) {
	infos, total, err := GetOutOfGass(context.Background(),
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

func getOutOfGasOnly(t *testing.T) {
	var err error
	info, err = GetOutOfGasOnly(context.Background(),
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

func existOutOfGas(t *testing.T) {
	exist, err := ExistOutOfGas(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existOutOfGasConds(t *testing.T) {
	exist, err := ExistOutOfGasConds(context.Background(),
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

func deleteOutOfGas(t *testing.T) {
	info, err := DeleteOutOfGas(context.Background(), info.ID)
	if assert.Nil(t, err) {
		appDate.DeletedAt = info.DeletedAt
		appDate.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &appDate)
	}
}

func TestMainOutOfGas(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace(uuid.NewString(), config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createOutOfGas", createOutOfGas)
	t.Run("createOutOfGass", createOutOfGass)
	t.Run("getOutOfGas", getOutOfGas)
	t.Run("getOutOfGass", getOutOfGass)
	t.Run("getOutOfGasOnly", getOutOfGasOnly)
	t.Run("updateOutOfGas", updateOutOfGas)
	t.Run("existOutOfGas", existOutOfGas)
	t.Run("existOutOfGasConds", existOutOfGasConds)
	t.Run("delete", deleteOutOfGas)
}
