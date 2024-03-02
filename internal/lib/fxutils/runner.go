package fxutils

import "sync"

type Runner struct {
	wg        *sync.WaitGroup
	waitCH    chan struct{}
	isWaiting bool
}

func NewRunner() *Runner {
	return &Runner{
		wg:        &sync.WaitGroup{},
		waitCH:    make(chan struct{}, 1),
		isWaiting: false,
	}
}

func (r *Runner) StartTracking() {
	r.wg.Add(1)
}

func (r *Runner) StopTracking() {
	r.wg.Done()
}

func (r *Runner) wait() <-chan struct{} {
	if !r.isWaiting {
		r.isWaiting = true

		go r.waitTracking()
	}

	return r.waitCH
}

func (r *Runner) waitTracking() {
	r.wg.Wait()
	r.waitCH <- struct{}{}
}
