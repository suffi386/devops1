package api

import (
	grpc_util "github.com/caos/zitadel/internal/api/grpc"
	"github.com/caos/zitadel/pkg/auth/api/oidc"
)

type Config struct {
	GRPC grpc_util.Config
	OIDC oidc.OPHandlerConfig
}

//
//func Start(ctx context.Context, conf Config, authZRepo *authz_repo.EsRepository, authZ auth_util.Config, defaults systemdefaults.SystemDefaults, authRepo repository.Repository) {
//	grpcServer := grpc.StartServer(conf.GRPC.ToServerConfig(), authZRepo, authZ, authRepo)
//	grpcGateway := grpc.StartGateway(conf.GRPC.ToGatewayConfig())
//	oidcHandler := oidc.NewProvider(ctx, conf.OIDC, authRepo)
//
//	server.StartServer(ctx, grpcServer, defaults)
//	server.StartGateway(ctx, grpcGateway)
//	op.Start(ctx, oidcHandler)
//}
