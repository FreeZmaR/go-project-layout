package httpsrv

import (
	"context"
	"errors"
	"github.com/FreeZmaR/go-project-layout/internal/lib/fxutils"
	"log/slog"
	"net"
	"net/http"
)

type App struct {
	srv      *http.Server
	runner   *fxutils.Runner
	serverCH chan error
}

func NewApp(srv *http.Server, runner *fxutils.Runner) *App {
	return &App{srv: srv, runner: runner, serverCH: make(chan error, 1)}
}

func (app *App) Run(_ context.Context) error {
	ln, err := app.listen()
	if err != nil {
		return err
	}

	go app.serve(ln)
	app.runner.StartTracking()

	return nil
}

func (app *App) Stop(ctx context.Context) error {
	defer app.runner.StopTracking()

	err := app.srv.Shutdown(ctx)
	if err != nil {
		return err
	}

	return <-app.serverCH
}

func (app *App) listen() (net.Listener, error) {
	ln, err := net.Listen("tcp", app.srv.Addr)
	if err != nil {
		return nil, err
	}

	slog.Info("App listen on " + app.srv.Addr)

	return ln, nil
}

func (app *App) serve(ln net.Listener) {
	err := app.srv.Serve(ln)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		app.serverCH <- err
	}

	app.serverCH <- nil
}
