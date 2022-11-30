package payment

import (
	npool "github.com/NpoolPlatform/message/npool/order/mgr/v1/payment"

	"github.com/NpoolPlatform/order-manager/pkg/db/ent"
)

func Ent2Grpc(row *ent.Payment) *npool.Payment {
	if row == nil {
		return nil
	}

	return &npool.Payment{
		ID:                   row.ID.String(),
		AppID:                row.AppID.String(),
		UserID:               row.UserID.String(),
		GoodID:               row.GoodID.String(),
		OrderID:              row.OrderID.String(),
		AccountID:            row.AccountID.String(),
		StartAmount:          row.StartAmount.String(),
		Amount:               row.Amount.String(),
		PayWithBalanceAmount: row.PayWithBalanceAmount.String(),
		FinishAmount:         row.FinishAmount.String(),
		CoinUsdCurrency:      row.CoinUsdCurrency.String(),
		LocalCoinUsdCurrency: row.LocalCoinUsdCurrency.String(),
		LiveCoinUsdCurrency:  row.LiveCoinUsdCurrency.String(),
		CoinInfoID:           row.CoinInfoID.String(),
		State:                npool.PaymentState(npool.PaymentState_value[row.State]),
		ChainTransactionID:   row.ChainTransactionID,
		UserSetPaid:          row.UserSetPaid,
		UserSetCanceled:      row.UserSetCanceled,
		FakePayment:          row.FakePayment,
		CreatedAt:            row.CreatedAt,
		UpdatedAt:            row.UpdatedAt,
		DeletedAt:            row.DeletedAt,
	}
}

func Ent2GrpcMany(rows []*ent.Payment) []*npool.Payment {
	infos := []*npool.Payment{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
