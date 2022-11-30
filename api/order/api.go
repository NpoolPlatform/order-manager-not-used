package order

import (
	"github.com/NpoolPlatform/message/npool/order/mgr/v1/order"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	order.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	order.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
