package httpsrv

import (
	inboxV1 "github.com/FreeZmaR/go-project-layout/internal/app/httpsrv/inbox/v1"
	outboxV1 "github.com/FreeZmaR/go-project-layout/internal/app/httpsrv/outbox/v1"
	"github.com/FreeZmaR/go-project-layout/internal/lib/fxutils"
	"go.uber.org/fx"
)

const moduleName = "HTTP-Server"

func NewAppInboxV1(provider fx.Option) *fxutils.App {
	return fxutils.NewApp(
		NewModuleInboxV1(provider),
	)
}

func NewAppOutboxV1(provider fx.Option) *fxutils.App {
	return fxutils.NewApp(
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

		ProvideProvider(),
		ProvideInvoke(),
	)
}

func NewModuleOutboxV1(provider fx.Option) fx.Option {
	return fx.Module(
		moduleName,
		provider,
		fx.Options(
			outboxV1.NewModule(),
		),

		ProvideProvider(),
		ProvideInvoke(),
	)
}
