package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/koliader/tellmi-gateway/internal/app/api"
	"github.com/koliader/tellmi-gateway/internal/config"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}
	if config.Environment == "dev" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
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
