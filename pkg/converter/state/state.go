package state

import (
	npool "github.com/NpoolPlatform/message/npool/ordermgr/state"
	"github.com/NpoolPlatform/order-manager/pkg/db/ent"
)

func Ent2Grpc(row *ent.State) *npool.State {
	if row == nil {
		return nil
	}

	return &npool.State{
		ID:      row.ID.String(),
		OrderID: row.OrderID.String(),
		State:   npool.EState(npool.EState_value[row.State]),
	}
}

func Ent2GrpcMany(rows []*ent.State) []*npool.State {
	infos := []*npool.State{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
