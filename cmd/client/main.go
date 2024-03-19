package main

import (
	"github.com/dlc-01/GophKeeper/internal/client/app"
	"github.com/dlc-01/GophKeeper/internal/client/config"
	"log"
)

func main() {

	cfg, err := config.New()
	if err != nil {
		log.Fatalf("client config error: %s", err)
	}

	client := app.New(cfg)
	client.Run()
}
