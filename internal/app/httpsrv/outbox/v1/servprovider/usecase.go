package servprovider

import (
	"github.com/FreeZmaR/go-project-layout/internal/repository"
	"github.com/FreeZmaR/go-project-layout/internal/usecase"
	"go.uber.org/fx"
)

func ProvideUseCases() fx.Option {
	return fx.Provide(
		provideOutboxUC,
	)
}

func provideOutboxUC(repo repository.Outbox) usecase.Outbox {
	return usecase.NewOutbox(repo)
}
