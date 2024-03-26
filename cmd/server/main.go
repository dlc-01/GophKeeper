package main

import (
	"fmt"
	"github.com/dlc-01/GophKeeper/internal/general/logger"
	"os"
	"os/signal"
	"syscall"

	"github.com/dlc-01/GophKeeper/internal/server/adapter/conf"

	"github.com/dlc-01/GophKeeper/internal/server/app"

	"log"
)

func main() {

	term := make(chan os.Signal, 1)
	signal.Notify(term, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	cfg, err := conf.InitConf()
	if err != nil {
		log.Fatal("err while parsing conf :%w", err)
	}
	lgr, err := logger.Initialize(cfg.Logger)
	if err != nil {
		lgr.Fatal(err.Error())
	}

	grpcServer, err := app.New(cfg, lgr)
	if err != nil {
		lgr.Fatalf("error while starting server: %s", err)
	}

	lgr.Infof("gRPC run: %s", cfg.GRPCServer.Address)

	if err = grpcServer.GrpcSrv.Serve(grpcServer.Listener); err != nil {
		lgr.Fatal(fmt.Errorf("gRPC - Serve: %w", err))
	}

	<-term

	if err := app.CLose(); err != nil {
		lgr.Fatalf("error while closing server: %s", err)
	}
	
	lgr.Info("server close")
}
