package server

import (
	"context"

	"github.com/caos/logging"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"golang.org/x/text/language"
	"google.golang.org/grpc"

	"github.com/caos/zitadel/internal/api/grpc/server/middleware"
	"github.com/caos/zitadel/internal/api/http"
)

const (
	defaultGrpcPort = "80"
)

type Server interface {
	Gateway
	RegisterServer(*grpc.Server)
	AuthInterceptor() grpc.UnaryServerInterceptor
	//GRPCServer(defaults systemdefaults.SystemDefaults) (*grpc.Server, error) TODO: remove
}

func CreateServer(servers []Server) *grpc.Server {
	authInterceptors := make([]grpc.UnaryServerInterceptor, len(servers))
	for i, server := range servers {
		authInterceptors[i] = server.AuthInterceptor()
	}
	return grpc.NewServer(
		middleware.TracingStatsServer(http.Healthz, http.Readiness, http.Validation),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				middleware.ErrorHandler(language.German),
				grpc_middleware.ChainUnaryServer(
					authInterceptors...,
				),
			),
		),
	)
}

func Serve(ctx context.Context, server *grpc.Server, port string) {
	go func() {
		<-ctx.Done()
		server.GracefulStop()
	}()

	go func() {
		listener := http.CreateListener(port)
		err := server.Serve(listener)
		logging.Log("SERVE-Ga3e94").OnError(err).Panic("grpc server serve failed")
	}()
	logging.LogWithFields("SERVE-bZ44QM", "port", port).Info("grpc server is listening")
}

func grpcPort(port string) string {
	if port == "" {
		return defaultGrpcPort
	}
	return port
}
