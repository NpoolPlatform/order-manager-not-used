//nolint:dupl
package state

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/order/mgr/v1/order/state"

	constant "github.com/NpoolPlatform/order-manager/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get state connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func CreateState(ctx context.Context, in *npool.StateReq) (*npool.State, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateState(ctx, &npool.CreateStateRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create state: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create state: %v", err)
	}
	return info.(*npool.State), nil
}

func CreateStates(ctx context.Context, in []*npool.StateReq) ([]*npool.State, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateStates(ctx, &npool.CreateStatesRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create states: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create states: %v", err)
	}
	return infos.([]*npool.State), nil
}

func UpdateState(ctx context.Context, in *npool.StateReq) (*npool.State, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.UpdateState(ctx, &npool.UpdateStateRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail update state: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update state: %v", err)
	}
	return info.(*npool.State), nil
}

func GetState(ctx context.Context, id string) (*npool.State, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetState(ctx, &npool.GetStateRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get state: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get state: %v", err)
	}
	return info.(*npool.State), nil
}

func GetStateOnly(ctx context.Context, conds *npool.Conds) (*npool.State, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetStateOnly(ctx, &npool.GetStateOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get state: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get state: %v", err)
	}
	return info.(*npool.State), nil
}

func GetStates(ctx context.Context, conds *npool.Conds, limit, offset int32) ([]*npool.State, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetStates(ctx, &npool.GetStatesRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get states: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get states: %v", err)
	}
	return infos.([]*npool.State), total, nil
}

func ExistState(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistState(ctx, &npool.ExistStateRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get state: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get state: %v", err)
	}
	return infos.(bool), nil
}

func ExistStateConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistStateConds(ctx, &npool.ExistStateCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get state: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get state: %v", err)
	}
	return infos.(bool), nil
}

func CountStates(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CountStates(ctx, &npool.CountStatesRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count state: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count state: %v", err)
	}
	return infos.(uint32), nil
}
