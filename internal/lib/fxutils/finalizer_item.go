package fxutils

type FinalizerItem interface {
	Close() error
}

type finalizerItem struct {
	closeFN FinalizerItermFN
}

type FinalizerItermFN func() error

func NewFinalizerItem(closeFN FinalizerItermFN) FinalizerItem {
	return finalizerItem{
		closeFN: closeFN,
	}
}

func (f finalizerItem) Close() error {
	return f.closeFN()
}
