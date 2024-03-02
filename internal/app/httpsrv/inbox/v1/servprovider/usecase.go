package servprovider

import (
	"github.com/FreeZmaR/go-project-layout/internal/repository"
	"github.com/FreeZmaR/go-project-layout/internal/usecase"
	"go.uber.org/fx"
)

func ProvideUseCases() fx.Option {
	return fx.Provide(
		provideInboxUC,
	)
}

func provideInboxUC(inboxRP repository.Inbox) usecase.Inbox {
	return usecase.NewInbox(inboxRP)
}
