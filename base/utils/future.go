package utils

import (
	"context"
	"things/base/errors"
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
	f := &Future{}

	if timeout > 0 {
		f.ctx, f.cancel = context.WithTimeout(context.TODO(), timeout)
		f.err = errors.Default("FutureTimeout", "future timed out after %s.", timeout.String())
	} else {
		f.ctx, f.cancel = context.WithCancel(context.TODO())
	}

	go f.Execute()
	return f
}

func (f *Future) SetResult(value any) {
	f.result = value
	f.err = nil
	f.cancel()
}

func (f *Future) SetError(err error) {
	f.result = nil
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

func UnwrapFuture[T any](future *Future) (v T, err error) {
	result, err := future.Wait()
	if err != nil {
		return
	}
	v, ok := result.(T)
	if !ok {
		err = errors.Invalid("UnwrapFuture", "Received unexpected result type")
	}
	return
}
