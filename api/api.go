package api

import (
	"context"

	compensate "github.com/NpoolPlatform/order-manager/api/compensate"
	"github.com/NpoolPlatform/order-manager/api/order"
	"github.com/NpoolPlatform/order-manager/api/outofgas"
	"github.com/NpoolPlatform/order-manager/api/payment"

	v1 "github.com/NpoolPlatform/message/npool/order/mgr/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	v1.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	v1.RegisterManagerServer(server, &Server{})
	order.Register(server)
	compensate.Register(server)
	outofgas.Register(server)
	payment.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := v1.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
