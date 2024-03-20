package main

import (
	"github.com/dlc-01/GophKeeper/internal/client/app"
	"github.com/dlc-01/GophKeeper/internal/client/config"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	term := make(chan os.Signal, 1)
	signal.Notify(term, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	cfg, err := config.New()
	if err != nil {
		log.Fatalf("client config error: %s", err)
	}

	client := app.New(cfg)
	go client.Run()

	<-term
	log.Println("client has been stopped")
}
