package main

import (
	"context"

	"github.com/voron4ikhin/go_cheatsheet_account/internal/app"
	"github.com/voron4ikhin/go_cheatsheet_account/internal/config"
	"github.com/voron4ikhin/go_cheatsheet_account/internal/logger"
)

func main() {
	ctx := context.Background()

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	// Initialization logger with server name
	l := logger.New(cfg)

	application := app.New(&l, cfg)
	if err := application.Run(ctx); err != nil {
		l.Fatal().Err(err).Msg("error")
	}
}
