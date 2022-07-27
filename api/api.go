package api

import (
	"context"

	"github.com/NpoolPlatform/message/npool/ordermgr"

	"github.com/NpoolPlatform/order-manager/api/state"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	ordermgr.UnimplementedOrderManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	ordermgr.RegisterOrderManagerServer(server, &Server{})
	state.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := ordermgr.RegisterOrderManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	if err := state.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
