package compensate

import (
	"github.com/NpoolPlatform/message/npool/order/mgr/v1/compensate"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	compensate.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	compensate.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
