package outofgas

import (
	npool "github.com/NpoolPlatform/message/npool/order/mgr/v1/outofgas"

	"github.com/NpoolPlatform/order-manager/pkg/db/ent"
)

func Ent2Grpc(row *ent.OutOfGas) *npool.OutOfGas {
	if row == nil {
		return nil
	}

	return &npool.OutOfGas{
		ID:        row.ID.String(),
		OrderID:   row.OrderID.String(),
		Start:     row.Start,
		End:       row.End,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
		DeletedAt: row.DeletedAt,
	}
}

func Ent2GrpcMany(rows []*ent.OutOfGas) []*npool.OutOfGas {
	infos := []*npool.OutOfGas{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
