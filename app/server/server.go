package server

import (
	"context"
	"github.com/tasuke/go-onion/config"
	"github.com/tasuke/go-onion/presentation/settings"
	"github.com/tasuke/go-onion/server/route"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(ctx context.Context, conf *config.Config) {
	api := settings.NewGinEngine()
	route.InitRoute(api)

	address := conf.Server.Address + ":" + conf.Server.Port
	log.Printf("Starting server on %s...\n", address)
	srv := &http.Server{
		Addr:              address,
		Handler:           api,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       10 * time.Minute,
		WriteTimeout:      10 * time.Minute,
	}
	go func() {
		// srv.Shutdownが呼ばれるとhttp.ErrServerClosedを返すのでこれは無視する
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		os.Exit(1)
	}
}
