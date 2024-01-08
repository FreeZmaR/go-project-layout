package httpsrv

import (
	"context"
	"errors"
	"log/slog"
	"net"
	"net/http"
)

type App struct {
	srv      *http.Server
	serverCH chan error
}

func NewApp(srv *http.Server) *App {
	return &App{srv: srv, serverCH: make(chan error, 1)}
}

func (app *App) Run(_ context.Context) error {
	ln, err := app.listen()
	if err != nil {
		return err
	}

	go app.serve(ln)

	return nil
}

func (app *App) Stop(ctx context.Context) error {
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
