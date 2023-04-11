package system

import (
	"google.golang.org/grpc"

	"github.com/zitadel/zitadel/internal/admin/repository"

	"github.com/zitadel/zitadel/internal/admin/repository/eventsourcing"
	"github.com/zitadel/zitadel/internal/api/authz"
	"github.com/zitadel/zitadel/internal/api/grpc/server"
	"github.com/zitadel/zitadel/internal/command"
	"github.com/zitadel/zitadel/internal/query"
	"github.com/zitadel/zitadel/pkg/grpc/system"
)

const (
	systemAPI = "System-API"
)

var _ system.SystemServiceServer = (*Server)(nil)

type Server struct {
	system.UnimplementedSystemServiceServer
	database        string
	command         *command.Commands
	query           *query.Queries
	administrator   repository.AdministratorRepository
	defaultInstance command.InstanceSetup
	externalDomain  string
}

type Config struct {
	Repository eventsourcing.Config
}

func CreateServer(
	command *command.Commands,
	query *query.Queries,
	repo repository.Repository,
	database string,
	defaultInstance command.InstanceSetup,
	externalDomain string,
) *Server {
	return &Server{
		command:         command,
		query:           query,
		administrator:   repo,
		database:        database,
		defaultInstance: defaultInstance,
		externalDomain:  externalDomain,
	}
}

func (s *Server) RegisterServer(grpcServer *grpc.Server) {
	system.RegisterSystemServiceServer(grpcServer, s)
}

func (s *Server) AppName() string {
	return systemAPI
}

func (s *Server) MethodPrefix() string {
	return system.SystemService_MethodPrefix
}

func (s *Server) AuthMethods() authz.MethodMapping {
	return system.SystemService_AuthMethods
}

func (s *Server) RegisterGateway() server.RegisterGatewayFunc {
	return system.RegisterSystemServiceHandler
}

func (s *Server) GatewayPathPrefix() string {
	return "/system/v1"
}
