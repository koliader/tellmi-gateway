package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/koliader/tellmi-gateway/internal/app/api"
	"github.com/koliader/tellmi-gateway/internal/config"
	"github.com/koliader/tellmi-sdk/health"
	"github.com/koliader/tellmi-sdk/otel"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}
	if config.Environment == "dev" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).Hook(otel.NewZerologHook())
	} else {
		log.Logger = log.Logger.Hook(otel.NewZerologHook())
	}
	zerolog.DefaultContextLogger = &log.Logger

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	otelSDK, err := otel.Init(ctx, otel.Config{ServiceName: "gateway-service", Insecure: true})
	if err != nil {
		log.Fatal().Err(err).Msg("cannot init otel")
	}
	defer func() {
		if err := otelSDK.Shutdown(context.Background()); err != nil {
			log.Error().Err(err).Msg("otel shutdown")
		}
	}()

	if config.HealthAddress != "" {
		healthServer := health.NewServer(config.HealthAddress).
			WithHandler("/metrics", otel.MetricsHandler())
		go func() {
			log.Info().Msgf("start health server at %s", config.HealthAddress)
			if err := healthServer.Start(); err != nil {
				log.Error().Err(err).Msg("health server stopped")
			}
		}()
	}

	server, err := api.NewServer(config)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server")
	}
}
