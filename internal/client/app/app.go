package app

import (
	"github.com/dlc-01/GophKeeper/internal/client/config"
	"github.com/dlc-01/GophKeeper/internal/client/gui"
	"github.com/dlc-01/GophKeeper/internal/client/handlers"
	"github.com/dlc-01/GophKeeper/internal/general/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type App struct {
	config   *config.Config
	logger   *logger.Logger
	conn     *grpc.ClientConn
	handlers *handlers.Handlers
	view     *gui.View
}

func New(cfg *config.Config) *App {
	loger, err := logger.Initialize(cfg.Logger)
	if err != nil {
		log.Fatal("error while starting client %w", err)
	}

	a := &App{
		config: cfg,
		logger: loger,
	}

	target := cfg.GRPCClient.Address + ":" + cfg.GRPCClient.Port
	a.conn, err = grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		a.logger.Fatal("error while starting client %w", err)
	}

	a.handlers = handlers.New(a.conn)
	a.view = gui.New(a.handlers, cfg)

	return a
}

func (a *App) Run() {
	defer a.conn.Close()
	a.view.Run()
}
