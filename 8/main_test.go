package main

import (
	"testing"
	"time"
)

func TestSemaphoreWaitGroup_Basic(t *testing.T) {
	wg := NewSemaphoreWaitGroup()

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(50 * time.Millisecond)
		}()
	}

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		// ok
	case <-time.After(500 * time.Millisecond):
		t.Fatal("таймаут ожидания Wait()")
	}
}

func TestSemaphoreWaitGroup_DoneWithoutAdd(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("ожидалась паника при отрицательном счётчике")
		}
	}()

	wg := NewSemaphoreWaitGroup()
	wg.Done() // вызов без Add должен вызвать панику
}

func TestSemaphoreWaitGroup_WaitAfterZero(t *testing.T) {
	wg := NewSemaphoreWaitGroup()

	wg.Add(1)
	go func() {
		time.Sleep(100 * time.Millisecond)
		wg.Done()
	}()

	start := time.Now()
	wg.Wait()
	elapsed := time.Since(start)

	if elapsed < 100*time.Millisecond {
		t.Error("Wait завершился слишком рано")
	}
}
