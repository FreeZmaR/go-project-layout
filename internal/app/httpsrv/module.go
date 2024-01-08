package httpsrv

import (
	inboxV1 "github.com/FreeZmaR/go-service-structure/template/internal/app/httpsrv/inbox/v1"
	outboxV1 "github.com/FreeZmaR/go-service-structure/template/internal/app/httpsrv/outbox/v1"
	"go.uber.org/fx"
)

const moduleName = "HTTP-Server"

func NewInboxV1(provider fx.Option) *fx.App {
	return newModuleApp(
		provider,
		fx.Options(
			inboxV1.NewModule(),
		),
		fx.Provide(
			NewApp,
			ProvideHTTPServer,
			ProvideMuxRouter,
		),
		fx.Invoke(
			InvokeAppLifeCycle,
		),
	)
}

func NewOutboxV1(provider fx.Option) *fx.App {
	return newModuleApp(
		provider,
		fx.Options(
			outboxV1.NewModule(),
		),
		fx.Provide(
			NewApp,
			ProvideHTTPServer,
			ProvideMuxRouter,
		),
		fx.Invoke(
			InvokeAppLifeCycle,
		),
	)
}

func newModuleApp(options ...fx.Option) *fx.App {
	return fx.New(
		fx.Module(
			moduleName,
			options...,
		),
		fx.NopLogger,
	)
}
