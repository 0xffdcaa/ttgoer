package main

import (
	"os"
	"os/signal"
	"syscall"

	"ttgoer/bot"
	"ttgoer/log"
	"ttgoer/tiktok"
)

func main() {
	log.S().Info("starting")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	tiktok.Init()
	bot.Start()

	<-quit
	tiktok.GracefulShutdown()
	bot.Stop()
	log.S().Info("stopping")
}
