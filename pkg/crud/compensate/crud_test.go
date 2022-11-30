package compensate

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/order-manager/pkg/db/ent"

	valuedef "github.com/NpoolPlatform/message/npool"
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

var appGood = ent.Compensate{
	ID:      uuid.New(),
	OrderID: uuid.New(),
	Start:   1000,
	End:     1000,
	Message: uuid.NewString(),
}

var (
	id      = appGood.ID.String()
	orderID = appGood.OrderID.String()
	req     = npool.CompensateReq{
		ID:      &id,
		OrderID: &orderID,
		Start:   &appGood.Start,
		End:     &appGood.End,
		Message: &appGood.Message,
	}
)

var info *ent.Compensate

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
	entities := []*ent.Compensate{
		{
			ID:      uuid.New(),
			OrderID: uuid.New(),
			Start:   1000,
			End:     1000,
			Message: uuid.NewString(),
		},
		{
			ID:      uuid.New(),
			OrderID: uuid.New(),
			Start:   1000,
			End:     1000,
			Message: uuid.NewString(),
		},
	}

	reqs := []*npool.CompensateReq{}
	for _, _appGood := range entities {
		_id := _appGood.ID.String()
		_orderID := _appGood.OrderID.String()
		reqs = append(reqs, &npool.CompensateReq{
			ID:      &_id,
			OrderID: &_orderID,
			Start:   &_appGood.Start,
			End:     &_appGood.End,
			Message: &_appGood.Message,
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
