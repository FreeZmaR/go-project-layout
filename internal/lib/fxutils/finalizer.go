package fxutils

import "log/slog"

type Finalizer struct {
	items []FinalizerItem
}

func NewFinalizer() *Finalizer {
	return &Finalizer{}
}

func (f *Finalizer) Append(item FinalizerItem) {
	f.items = append(f.items, item)
}

func (f *Finalizer) Close() {
	for _, item := range f.items {
		err := item.Close()
		if err != nil {
			slog.Error("error close finalizer", slog.String("err", err.Error()))
		}
	}
}
