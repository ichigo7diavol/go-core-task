package main

import (
	"testing"
	"time"
)

func TestRandomGenerator_Basic(t *testing.T) {
	ch := RandomGenerator()

	select {
	case n := <-ch:
		if n < 0 || n >= 1000 {
			t.Errorf("ожидалось число в диапазоне [0, 1000), получено %d", n)
		}
	case <-time.After(time.Second):
		t.Error("таймаут: генератор не отправил значение в канал")
	}
}

func TestRandomGenerator_MultipleValues(t *testing.T) {
	ch := RandomGenerator()

	values := make(map[int]bool)
	for i := 0; i < 5; i++ {
		select {
		case n := <-ch:
			values[n] = true
		case <-time.After(time.Second):
			t.Fatal("генератор не отправил значения вовремя")
		}
	}

	if len(values) == 1 {
		t.Error("все числа одинаковые — возможно, генератор не случайный или не пересеивается")
	}
}
