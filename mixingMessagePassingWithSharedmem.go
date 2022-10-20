package main

import (
	"context"
	"errors"
	"time"
)

var (
	// ErrCancelled means wait was cancelled by context
	ErrCancelled = errors.New("cancelled by context")
)

type responseFuture interface {
	Start()
	Wait(ctx context.Context) error
	GetResponse() interface{}
	GetError() error
}

type ResponseFunc func() (interface{}, error)

type responseFutureImpl struct {
	f        ResponseFunc
	ch       chan int
	response interface{}
	err      error
}

func newResponseFuture(f ResponseFunc) responseFuture {
	return &responseFutureImpl{
		f:  f,
		ch: make(chan int),
	}
}

// Start starts a future.
func (f *responseFutureImpl) Start() {
	go func() {
		resp, err := f.f()
		f.response = resp
		f.err = err
		// intentionally added to trigger a time out
		time.Sleep(time.Second * 10)
		f.ch <- 1
	}()
}

// Wait waits for future.
func (f *responseFutureImpl) Wait(ctx context.Context) error {
	select {
	case <-f.ch:
		return nil
	case <-ctx.Done():
		// This write data races with the read in GetError() and write in Start()
		f.err = ErrCancelled
		return ErrCancelled
	}
}

// GetResponse returns future response
func (f *responseFutureImpl) GetResponse() interface{} {
	// This read data races with the write in Wait() and Start()
	return f.response
}

// GetError returns error from the future
func (f *responseFutureImpl) GetError() error {
	// This read data races with the write in Wait() and Start()
	return f.err
}

func runMeLater() (interface{}, error) {
	return "hi", nil
}

func main() {
	r := newResponseFuture(runMeLater)
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*1000))
	r.Start()
	r.Wait(ctx)
}
