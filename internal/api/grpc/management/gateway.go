package management

import (
	"strings"

	grpc_util "github.com/caos/zitadel/internal/api/grpc"
	"github.com/caos/zitadel/internal/api/grpc/server"
	"github.com/caos/zitadel/pkg/management/grpc"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

type Gateway struct {
	grpcEndpoint string
	port         string
	cutomHeaders []string
}

func StartGateway(conf grpc_util.GatewayConfig) *Gateway {
	return &Gateway{
		grpcEndpoint: conf.GRPCEndpoint,
		port:         conf.Port,
		cutomHeaders: conf.CustomHeaders,
	}
}

func (gw *Gateway) Gateway() server.GatewayFunc {
	return grpc.RegisterManagementServiceHandlerFromEndpoint
}

func (gw *Gateway) GRPCEndpoint() string {
	return ":" + gw.grpcEndpoint
}

func (gw *Gateway) GatewayPort() string {
	return gw.port
}

func (gw *Gateway) GatewayServeMuxOptions() []runtime.ServeMuxOption {
	return []runtime.ServeMuxOption{
		runtime.WithIncomingHeaderMatcher(func(header string) (string, bool) {
			for _, customHeader := range gw.cutomHeaders {
				if strings.HasPrefix(strings.ToLower(header), customHeader) {
					return header, true
				}
			}
			return runtime.DefaultHeaderMatcher(header)
		}),
	}
}
