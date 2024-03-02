package fxutils

import (
	"log/slog"
)

type FinalizerItem interface {
	Close() error
}

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
			slog.Error("Error close item by finalizer: ", err.Error())
		}
	}
}
