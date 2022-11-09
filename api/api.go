package api

import (
	"context"

	ordermgr "github.com/NpoolPlatform/message/npool/order/mgr/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	ordermgr.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	ordermgr.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := ordermgr.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
