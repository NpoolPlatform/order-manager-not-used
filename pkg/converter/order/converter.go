package order

import (
	npool "github.com/NpoolPlatform/message/npool/order/mgr/v1/order"

	"github.com/NpoolPlatform/order-manager/pkg/db/ent"
)

func Ent2Grpc(row *ent.Order) *npool.Order {
	if row == nil {
		return nil
	}

	return &npool.Order{
		ID:                     row.ID.String(),
		GoodID:                 row.GoodID.String(),
		AppID:                  row.AppID.String(),
		UserID:                 row.UserID.String(),
		ParentOrderID:          row.ParentOrderID.String(),
		PayWithParent:          row.PayWithParent,
		Units:                  row.Units,
		PromotionID:            row.PromotionID.String(),
		DiscountCouponID:       row.DiscountCouponID.String(),
		UserSpecialReductionID: row.UserSpecialReductionID.String(),
		StartAt:                row.StartAt,
		EndAt:                  row.EndAt,
		FixAmountCouponID:      row.FixAmountCouponID.String(),
		Type:                   npool.OrderType(npool.OrderType_value[row.Type]),
		State:                  npool.OrderState(npool.OrderState_value[row.State]),
		LastBenefitAt:          row.LastBenefitAt,
		CreatedAt:              row.CreatedAt,
		UpdatedAt:              row.UpdatedAt,
		DeletedAt:              row.DeletedAt,
	}
}

func Ent2GrpcMany(rows []*ent.Order) []*npool.Order {
	infos := []*npool.Order{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
