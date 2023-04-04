package utils

import (
	"context"
	"time"
)

// Future stores result or error that may return in the future
type Future struct {
	result    any
	err       error
	ctx       context.Context
	cancel    context.CancelFunc
	onResults []func(v any)
	onErrors  []func(err error)
}

func NewFuture(timeout time.Duration) *Future {
	var ctx context.Context
	var cancel context.CancelFunc
	if timeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, timeout)
	} else {
		ctx, cancel = context.WithCancel(ctx)
	}
	f := &Future{
		ctx:    ctx,
		cancel: cancel,
	}
	go f.Execute()
	return f
}

func (f *Future) SetResult(value any) {
	f.result = value
	f.cancel()
}

func (f *Future) SetError(err error) {
	f.err = err
	f.cancel()
}

func (f *Future) Wait() (any, error) {
	<-f.ctx.Done()
	return f.result, f.err
}

func (f *Future) Execute() {
	<-f.ctx.Done()
	if f.err != nil {
		for _, handler := range f.onErrors {
			handler(f.err)
		}
	} else {
		for _, handler := range f.onResults {
			handler(f.result)
		}
	}
}

func (f *Future) OnResult(onResult func(v any)) {
	f.onResults = append(f.onResults, onResult)
}

func (f *Future) OnError(onError func(err error)) {
	f.onErrors = append(f.onErrors, onError)
}
