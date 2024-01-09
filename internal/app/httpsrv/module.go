package httpsrv

import (
	inboxV1 "github.com/FreeZmaR/go-project-layout/internal/app/httpsrv/inbox/v1"
	outboxV1 "github.com/FreeZmaR/go-project-layout/internal/app/httpsrv/outbox/v1"
	"go.uber.org/fx"
)

const moduleName = "HTTP-Server"

func NewAppInboxV1(provider fx.Option) *fx.App {
	return newModuleApp(
		NewModuleInboxV1(provider),
	)
}

func NewAppOutboxV1(provider fx.Option) *fx.App {
	return newModuleApp(
		NewModuleOutboxV1(provider),
	)
}

func NewModuleInboxV1(provider fx.Option) fx.Option {
	return fx.Module(
		moduleName,
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

func NewModuleOutboxV1(provider fx.Option) fx.Option {
	return fx.Module(
		moduleName,
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

func newModuleApp(module fx.Option) *fx.App {
	return fx.New(
		module,
		fx.NopLogger,
	)
}
