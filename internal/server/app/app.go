package app

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dlc-01/GophKeeper/internal/general/logger"
	proto "github.com/dlc-01/GophKeeper/internal/general/proto/gen"
	"github.com/dlc-01/GophKeeper/internal/server/adapter/auth/jwt/manager"
	"github.com/dlc-01/GophKeeper/internal/server/adapter/conf"
	"github.com/dlc-01/GophKeeper/internal/server/adapter/handlers"
	"github.com/dlc-01/GophKeeper/internal/server/adapter/repository/postgres"
	"github.com/dlc-01/GophKeeper/internal/server/adapter/repository/postgres/repositories"
	"github.com/dlc-01/GophKeeper/internal/server/core/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net"
)

var (
	lgrS       *logger.Logger
	jwtManager *manager.TokenService
	db         *sql.DB
)

type GrpcServer struct {
	GrpcSrv  *grpc.Server
	Listener net.Listener
}

func New(cfg *conf.Config, lgr *logger.Logger) (*GrpcServer, error) {
	lgrS = lgr
	sqlClient, err := postgres.NewSQLClient(*cfg)
	if err != nil {
		return nil, err
	}

	userRepository, err := repositories.NewUserRepository(lgr, sqlClient)
	if err != nil {
		return nil, err
	}

	pairRepository, err := repositories.NewPairRepository(lgr, sqlClient)
	if err != nil {
		return nil, err
	}

	textReposiry, err := repositories.NewTextRepository(lgr, sqlClient)
	if err != nil {
		return nil, err
	}

	bankRepository, err := repositories.NewBankRepository(lgr, sqlClient)
	if err != nil {
		return nil, err
	}

	jwtManager = manager.NewTokenService(cfg.JWT.SecretKey, cfg.JWT.Expire)

	authModule := service.NewAuthService(userRepository, jwtManager, lgr)

	userSevice := service.NewUserService(userRepository, lgr)

	bankService := service.NewBankService(bankRepository, lgr)

	textServie := service.NewTextService(textReposiry, lgr)

	pairService := service.NewPairService(pairRepository, lgr)

	authHandler := handlers.NewAuthServer(authModule)

	userHandler := handlers.NewUserServer(userSevice, authModule)

	pairHandlers := handlers.NewPairServer(pairService)

	textHandlers := handlers.NewTextServer(textServie)

	bankHandler := handlers.NewBankServer(bankService)

	listen, err := net.Listen("tcp", cfg.GRPCServer.Address)
	if err != nil {
		return nil, fmt.Errorf("gRPC - net.Listen: %w", err)
	}

	grpcSrv := grpc.NewServer(grpc.UnaryInterceptor(userIdentity))

	proto.RegisterUserServer(grpcSrv, userHandler)
	proto.RegisterAuthServer(grpcSrv, authHandler)
	proto.RegisterPairServer(grpcSrv, pairHandlers)
	proto.RegisterTextServer(grpcSrv, textHandlers)
	proto.RegisterBanksServer(grpcSrv, bankHandler)
	return &GrpcServer{GrpcSrv: grpcSrv, Listener: listen}, err
}
func CLose() error {
	err := db.Close()
	if err != nil {
		return fmt.Errorf("error file closing connection : %w", err)
	}
	return nil
}

func userIdentity(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	if _, ok := info.Server.(*handlers.AuthServer); ok {
		return handler(ctx, req)
	}

	var token string
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		values := md.Get("token")
		if len(values) > 0 {
			token = values[0]
		}
	}

	if token == "" {
		lgrS.Error(fmt.Errorf("missing token in metadata"))
		return nil, status.Error(codes.FailedPrecondition, "missing token")
	}

	user, err := jwtManager.VerifyToken(token)
	if err != nil {
		lgrS.Error(fmt.Errorf("user identity: %w", err))
		return nil, status.Error(codes.Aborted, "user identity error")
	}
	ctx = context.WithValue(ctx, handlers.UserIDKey, user.ID)

	return handler(ctx, req)
}
