package parallelcap

import (
	"context"
	"math"
	"sync"
)

type ParallelCap struct {
	cap     int
	current chan interface{}
	wg      sync.WaitGroup
}

// NewParallelCap creates a new ParallelCap with the given capacity.
func NewParallelCap(cap int) *ParallelCap {
	if cap <= 0 {
		cap = math.MaxInt32
	}

	return &ParallelCap{
		cap:     cap,
		current: make(chan interface{}, cap),
		wg:      sync.WaitGroup{},
	}
}

// Add adds one to the ParallelCap.
func (p *ParallelCap) Add() {
	_ = p.AddWithContext(context.Background())
}

// AddWithContext adds one to the ParallelCap.
func (p *ParallelCap) AddWithContext(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case p.current <- struct{}{}:
		p.wg.Add(1)
		return nil
	}
}

// Done removes one from the ParallelCap.
func (p *ParallelCap) Done() {
	<-p.current
	p.wg.Done()
}

// Wait blocks until the ParallelCap is empty.
func (p *ParallelCap) Wait() {
	p.wg.Wait()
}
