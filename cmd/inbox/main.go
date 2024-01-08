package main

import (
	"fmt"
	"github.com/FreeZmaR/go-project-layout/config"
	"github.com/FreeZmaR/go-project-layout/config/build"
	"github.com/FreeZmaR/go-project-layout/config/types"
	"github.com/FreeZmaR/go-project-layout/internal/app/httpsrv"
	"go.uber.org/fx"
	"log/slog"
	"os"
)

func main() {
	slog.Info("Starting inbox service...")
	slog.Info(fmt.Sprintf("Build info: pid: %d  %s", os.Geteuid(), build.Info()))

	cfg, err := config.LoadInbox()
	if err != nil {
		slog.Error("Failed to load config", slog.String("err", err.Error()))

		os.Exit(1)
	}

	slog.Info("Config loaded")

	app := httpsrv.NewInboxV1(
		fx.Provide(
			func() *types.HTTPServer { return cfg.Server },
			func() *types.Postgres { return cfg.Postgres },
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
