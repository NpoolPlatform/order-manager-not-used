package compensate

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

	npool "github.com/NpoolPlatform/message/npool/order/mgr/v1/compensate"

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

var appDate = npool.Compensate{
	ID:      uuid.NewString(),
	OrderID: uuid.NewString(),
	Start:   10000,
	End:     10001,
	Message: uuid.NewString(),
}

var (
	appInfo = npool.CompensateReq{
		ID:      &appDate.ID,
		OrderID: &appDate.OrderID,
		Start:   &appDate.Start,
		End:     &appDate.End,
		Message: &appDate.Message,
	}
)

var info *npool.Compensate

func createCompensate(t *testing.T) {
	var err error
	info, err = CreateCompensate(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		appDate.CreatedAt = info.CreatedAt
		appDate.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &appDate)
	}
}

func createCompensates(t *testing.T) {
	appDates := []npool.Compensate{
		{
			ID:      uuid.NewString(),
			OrderID: uuid.NewString(),
			Start:   10000,
			End:     10001,
			Message: uuid.NewString(),
		},
		{
			ID:      uuid.NewString(),
			OrderID: uuid.NewString(),
			Start:   10000,
			End:     10001,
			Message: uuid.NewString(),
		},
	}

	apps := []*npool.CompensateReq{}
	for key := range appDates {
		apps = append(apps, &npool.CompensateReq{
			ID:      &appDates[key].ID,
			OrderID: &appDates[key].OrderID,
			Start:   &appDates[key].Start,
			End:     &appDates[key].End,
			Message: &appDates[key].Message,
		})
	}

	infos, err := CreateCompensates(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updateCompensate(t *testing.T) {
	var err error
	info, err = UpdateCompensate(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		appDate.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &appDate)
	}
}

func getCompensate(t *testing.T) {
	var err error
	info, err = GetCompensate(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appDate)
	}
}

func getCompensates(t *testing.T) {
	infos, total, err := GetCompensates(context.Background(),
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

func getCompensateOnly(t *testing.T) {
	var err error
	info, err = GetCompensateOnly(context.Background(),
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

func existCompensate(t *testing.T) {
	exist, err := ExistCompensate(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existCompensateConds(t *testing.T) {
	exist, err := ExistCompensateConds(context.Background(),
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

func deleteCompensate(t *testing.T) {
	info, err := DeleteCompensate(context.Background(), info.ID)
	if assert.Nil(t, err) {
		appDate.DeletedAt = info.DeletedAt
		assert.Equal(t, info, &appDate)
	}
}

func TestMainCompensate(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace(uuid.NewString(), config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createCompensate", createCompensate)
	t.Run("createCompensates", createCompensates)
	t.Run("getCompensate", getCompensate)
	t.Run("getCompensates", getCompensates)
	t.Run("getCompensateOnly", getCompensateOnly)
	t.Run("updateCompensate", updateCompensate)
	t.Run("existCompensate", existCompensate)
	t.Run("existCompensateConds", existCompensateConds)
	t.Run("delete", deleteCompensate)
}
