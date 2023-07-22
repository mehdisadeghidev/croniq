package cmd

import (
	"github.com/mehdisadeghidev/croniq/config"
	"github.com/mehdisadeghidev/croniq/services/api"
	"github.com/mehdisadeghidev/croniq/services/cron"
	"github.com/mehdisadeghidev/croniq/services/queue"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	config.Setup()

	_ = os.Setenv("TZ", "UTC")
}

func Setup() {
	sig := make(chan os.Signal, 10000)
	signal.Notify(sig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	api.Setup()
	queue.Setup()
	cron.Setup()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	_ = api.Shutdown()
	_ = queue.Shutdown()
}
