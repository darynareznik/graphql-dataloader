package main

import (
	"context"
	"graphql-dataloader/config"
	"graphql-dataloader/graph/server"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	setupGracefulShutdown(cancel)

	cfg := config.Read()
	zerolog.SetGlobalLevel(cfg.GetLogLevel())

	if err := run(ctx, cfg); err != nil {
		cancel()
		log.Fatal().Err(err).Msg("shutting down service with error")
	}
}

func run(ctx context.Context, cfg config.Config) error {
	svr := server.CreateAndRun(ctx, cfg)
	defer closeWithTimeout(ctx, svr.Close, 5*time.Second)

	<-ctx.Done()
	return nil
}

func closeWithTimeout(ctx context.Context, c func(context.Context), d time.Duration) {
	ctx, cancel := context.WithTimeout(ctx, d)
	defer cancel()
	c(ctx)
}

func setupGracefulShutdown(stop func()) {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalChannel
		log.Error().Msg("got Interrupt signal")
		stop()
	}()
}
