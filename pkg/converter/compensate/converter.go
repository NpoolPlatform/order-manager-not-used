package compensate

import (
	npool "github.com/NpoolPlatform/message/npool/order/mgr/v1/compensate"

	"github.com/NpoolPlatform/order-manager/pkg/db/ent"
)

func Ent2Grpc(row *ent.Compensate) *npool.Compensate {
	if row == nil {
		return nil
	}

	return &npool.Compensate{
		ID:        row.ID.String(),
		OrderID:   row.OrderID.String(),
		Start:     row.Start,
		End:       row.End,
		Message:   row.Message,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
		DeletedAt: row.DeletedAt,
	}
}

func Ent2GrpcMany(rows []*ent.Compensate) []*npool.Compensate {
	infos := []*npool.Compensate{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
