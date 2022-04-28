package main

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"rapture-haven/virtue-of-ignorance"
	"rapture-haven/virtue-of-ignorance/internal/types/config"
)

// функция инициализации сервера через cobra
func init() {
	command := &cobra.Command{
		Use:   "serve",
		Short: "serve virtue streaming server",
		Long:  `serve virtue streaming server`,
		Run:   virtue.Service.ServeCommand,
	}

	configs := []config.Config{
		virtue.Service.Server,
		virtue.Service.WebRTC,
		virtue.Service.Remote,
		virtue.Service.Broadcast,
		virtue.Service.WebSocket,
	}

	cobra.OnInitialize(func() {
		for _, cfg := range configs {
			cfg.Set()
		}
		virtue.Service.Preflight()
	})

	for _, cfg := range configs {
		if err := cfg.Init(command); err != nil {
			log.Panic().Err(err).Msg("unable to run serve command")
		}
	}

	root.AddCommand(command)
}
