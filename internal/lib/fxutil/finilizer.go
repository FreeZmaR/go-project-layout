package fxutil

import "log/slog"

type Resource interface {
	Close() error
}

func Finalizer(resources ...Resource) {
	for _, res := range resources {
		if err := res.Close(); err != nil {
			slog.Error("failed to close resource", err)
		}
	}
}
