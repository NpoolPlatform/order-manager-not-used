package state

import (
	"github.com/NpoolPlatform/message/npool/order/mgr/v1/order/state"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	state.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	state.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
