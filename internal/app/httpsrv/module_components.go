package httpsrv

import (
	"context"
	"crypto/tls"
	"errors"
	"github.com/FreeZmaR/go-service-structure/template/config/build"
	"github.com/FreeZmaR/go-service-structure/template/config/types"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"log/slog"
	"net/http"
	"time"
)

func InvokeAppLifeCycle(lc fx.Lifecycle, app *App) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return app.Run(ctx)
		},
		OnStop: func(ctx context.Context) error {
			return app.Stop(ctx)
		},
	})
}

func ProvideHTTPServer(cfg *types.HTTPServer, router *mux.Router) (*http.Server, error) {
	if build.IsProductionMode() && nil == cfg.TLS {
		return nil, errors.New("http-server: tls config is required")
	}

	srv := &http.Server{
		Addr:         cfg.Addr(),
		ReadTimeout:  time.Duration(cfg.ReadTimeoutSec) * time.Second,
		WriteTimeout: time.Duration(cfg.WriteTimeoutSec) * time.Second,
		IdleTimeout:  time.Duration(cfg.IdleTimeoutSec) * time.Second,
		Handler:      router,
	}

	if build.IsDevelopMode() {
		return srv, nil
	}

	if nil == cfg.TLS {
		return nil, errors.New("http-server: tls config is required")
	}

	srv.TLSConfig = &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	if cfg.TLS.MinVersion != nil {
		srv.TLSConfig.MinVersion = *cfg.TLS.MinVersion
	}
	if cfg.TLS.MaxVersion != nil {
		srv.TLSConfig.MaxVersion = *cfg.TLS.MaxVersion
	}
	if cfg.TLS.Curves != nil {
		srv.TLSConfig.CurvePreferences = *cfg.TLS.Curves
	}
	if cfg.TLS.Ciphers != nil {
		srv.TLSConfig.CipherSuites = *cfg.TLS.Ciphers
	}

	return srv, nil
}

func ProvideMuxRouter() *mux.Router {
	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("not found handler")

		w.WriteHeader(http.StatusNotFound)
	})

	return router
}
