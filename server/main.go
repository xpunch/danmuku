package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/xpunch/danmuku/config"
	"github.com/xpunch/danmuku/service"
)

func main() {
	cfg := config.NewConfig()
	if err := cfg.Load(); err != nil {
		log.Fatalf("[ERR]Config: %v", err)
	}
	srv := service.NewService(cfg.Address)
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		if err := srv.Run(); err != nil {
			log.Fatalf("[ERR]Service: %v", err)
		}
	}()
	<-exit
}
