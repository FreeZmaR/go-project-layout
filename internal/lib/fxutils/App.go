package fxutils

import (
	"context"
	"errors"
	"go.uber.org/fx"
	"log/slog"
	"strings"
	"time"
)

type App struct {
	instance  *fx.App
	finalizer *Finalizer
	runner    *Runner
	errStack  []string
}

func NewApp(modules fx.Option) *App {
	f := NewFinalizer()
	r := NewRunner()

	instance := fx.New(
		modules,
		fx.Provide(
			func() *Finalizer { return f },
			func() *Runner { return r },
		),
		fx.NopLogger,
	)

	return &App{
		instance:  instance,
		finalizer: f,
		runner:    r,
	}
}

func (a *App) Run() {
	if err := a.start(); err != nil {
		slog.Error("Error start app: ", err.Error())

		return
	}

	defer a.finalizer.Close()

	a.listener()

	if err := a.stop(); err != nil {
		slog.Error("Error stop app: ", err.Error())
	}
}

func (a *App) start() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return a.instance.Start(ctx)
}

func (a *App) stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return a.instance.Stop(ctx)
}

func (a *App) listener() {
	select {
	case sig := <-a.instance.Done():
		slog.Info("App received signal: ", sig.String())

		return
	case <-a.runner.wait():
		return
	}
}

func (a *App) Err() error {
	return a.getLastErr()
}

func (a *App) ErrStack() []string {
	a.makeErrStack()

	return a.errStack
}

func (a *App) getLastErr() error {
	a.makeErrStack()

	if len(a.errStack) == 0 {
		return nil
	}

	return errors.New(a.errStack[len(a.errStack)-1])
}

func (a *App) makeErrStack() {
	if nil == a.instance.Err() || len(a.errStack) > 0 {
		return
	}

	stack := strings.Split(a.instance.Err().Error(), "): ")

	for i := 0; i < len(stack); i++ {
		if i == len(stack)-1 {
			break
		}

		if stack[i][len(stack[i])-1] != ')' {
			stack[i] += ")"
		}
	}

	a.errStack = stack
}
