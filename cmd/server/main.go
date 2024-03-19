package main

import (
	"fmt"
	"github.com/dlc-01/GophKeeper/internal/general/logger"

	"github.com/dlc-01/GophKeeper/internal/server/adapter/conf"

	"github.com/dlc-01/GophKeeper/internal/server/app"

	"log"
	"os"
)

var (
	lgr *logger.Logger
	cfg *conf.Config
	err error
)

func main() {
	if cfg, err = conf.InitConf(); err != nil {
		log.Fatal("err while parsing conf :%w", err)
	}
	if lgr, err = logger.Initialize(logger.ConfigLog{AppMode: cfg.AppMod, LoggerDirectory: cfg.Logger.File.Directory, LoggerFileMaxSize: cfg.Logger.File.MaxSize, LoggerFileCompress: cfg.Logger.File.Compress, LoggerFileMaxBackups: cfg.Logger.File.MaxBackups, LoggerFileMaxAge: cfg.Logger.File.MaxAge}); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	appCfg, err := app.New(cfg, lgr)
	if err != nil {
		lgr.Fatalf("error while starting server: %s", err)
	}

	lgr.Infof("gRPC run: %s", cfg.GRPCServer.Address)

	if err = appCfg.GrpcSrv.Serve(appCfg.Listener); err != nil {
		lgr.Fatal(fmt.Errorf("gRPC - Serve: %w", err))
	}
}
