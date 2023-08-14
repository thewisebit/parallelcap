package parallelcap

import (
	"context"
	"math"
	"testing"
	"time"
)

func TestNewParallelCap(t *testing.T) {
	p := NewParallelCap(5)
	if cap(p.current) != 5 {
		t.Errorf("Expected capacity of 5, got %d", cap(p.current))
	}
}

func TestNewParallelCapZero(t *testing.T) {
	p := NewParallelCap(0)
	if cap(p.current) != math.MaxInt32 {
		t.Errorf("Expected maximum capacity, got %d", cap(p.current))
	}
}

func TestAddWithContext(t *testing.T) {
	p := NewParallelCap(2)
	err := p.AddWithContext(context.Background())
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	err = p.AddWithContext(context.Background())
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	err = p.AddWithContext(ctx)
	if err != context.DeadlineExceeded {
		t.Errorf("Expected context deadline exceeded error, got %v", err)
	}
}

func TestDone(t *testing.T) {
	p := NewParallelCap(2)
	p.Add()
	p.Add()
	p.Done()
	p.Done()
	if len(p.current) != 0 {
		t.Errorf("Expected current to be empty, got %d", len(p.current))
	}
}

func TestWait(t *testing.T) {
	p := NewParallelCap(2)
	p.Add()
	p.Add()

	go func() {
		time.Sleep(10 * time.Millisecond)
		p.Done()
		p.Done()
	}()

	p.Wait()
	if len(p.current) != 0 {
		t.Errorf("Expected current to be empty, got %d", len(p.current))
	}
}
