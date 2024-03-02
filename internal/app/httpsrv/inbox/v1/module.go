package v1

import (
	"github.com/FreeZmaR/go-project-layout/internal/app/httpsrv/inbox/v1/servprovider"
	"go.uber.org/fx"
)

const moduleName = "inbox-v1"

func NewModule() fx.Option {
	return fx.Module(
		moduleName,
		servprovider.ProvideServices(),
		servprovider.ProvideRepositories(),
		servprovider.ProvideUseCases(),
		servprovider.ProvideRouter(),
	)
}
