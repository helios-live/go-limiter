package limiter // include go.ideatocode.tech/limiter

import (
	"errors"
	"sync"
)

// New returns a new instance of limiter
func New(max int) *Limiter {
	return &Limiter{max: &max}
}

// Limiter performs connections / thread limiting
type Limiter struct {
	max     *int
	current int
	m       sync.Mutex
}

// Add adds a connection
func (sl *Limiter) Add() bool {
	sl.m.Lock()
	defer sl.m.Unlock()

	// always return true if connections are not counted
	if sl.max == nil {
		return true
	}
	if sl.current >= *sl.max {
		return false
	}

	sl.current++

	return true
}

// Done removes a connection from the pool
func (sl *Limiter) Done() error {
	sl.m.Lock()
	defer sl.m.Unlock()
	sl.current--
	if sl.current < 0 {
		sl.current = 0
		return errors.New("limiter: done would go subzero")
	}
	return nil
}

// Current returns the current count
func (sl *Limiter) Current() int {

	sl.m.Lock()
	defer sl.m.Unlock()
	return sl.current
}

// SetMax updates the max
func (sl *Limiter) SetMax(max *int) {
	sl.m.Lock()
	defer sl.m.Unlock()
	sl.max = max
}
