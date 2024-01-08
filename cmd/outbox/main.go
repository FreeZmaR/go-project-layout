package main

import (
	"fmt"
	"github.com/FreeZmaR/go-service-structure/template/config"
	"github.com/FreeZmaR/go-service-structure/template/config/build"
	"github.com/FreeZmaR/go-service-structure/template/config/types"
	"github.com/FreeZmaR/go-service-structure/template/internal/app/httpsrv"
	"go.uber.org/fx"
	"log/slog"
	"os"
)

func main() {
	slog.Info("Starting outbox service...")
	slog.Info(fmt.Sprintf("Build info: pid: %d  %s", os.Geteuid(), build.Info()))

	cfg, err := config.LoadOutbox()
	if err != nil {
		slog.Error("Failed to load config", slog.String("err", err.Error()))

		os.Exit(1)
	}

	slog.Info("Config loaded")

	app := httpsrv.NewOutboxV1(
		fx.Provide(
			func() *types.HTTPServer { return cfg.Server },
			func() *types.Postgres { return cfg.Postgres },
			func() *types.Redis { return cfg.Redis },
		),
	)

	if err = app.Err(); err != nil {
		slog.Error("Failed to create app", slog.String("err", err.Error()))

		os.Exit(1)
	}

	slog.Info("App created")

	app.Run()

	slog.Info("App stopped")
}
