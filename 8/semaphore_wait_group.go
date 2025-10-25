package main

import (
	"sync"
)

type SemaphoreWaitGroup struct {
	ch chan struct{}
	mu sync.Mutex
	n  int
}

func NewSemaphoreWaitGroup() *SemaphoreWaitGroup {
	return &SemaphoreWaitGroup{ch: make(chan struct{})}
}

func (wg *SemaphoreWaitGroup) Add(delta int) {
	wg.mu.Lock()
	defer wg.mu.Unlock()

	prev := wg.n
	wg.n += delta

	if wg.n < 0 {
		panic("SemaphoreWaitGroup: отрицательное количество горутин")
	}
	if prev == 0 && wg.n > 0 {
		wg.ch = make(chan struct{})
	}
	if wg.n == 0 {
		close(wg.ch)
	}
}

func (wg *SemaphoreWaitGroup) Done() {
	wg.Add(-1)
}

func (wg *SemaphoreWaitGroup) Wait() {
	wg.mu.Lock()
	ch := wg.ch
	wg.mu.Unlock()
	<-ch
}
