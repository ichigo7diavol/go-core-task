package main

import (
	"reflect"
	"sort"
	"testing"
)

func makeIntChan(vals ...int) <-chan int {
	ch := make(chan int, len(vals))
	for _, v := range vals {
		ch <- v
	}
	close(ch)
	return ch
}

func TestMerge_Basic(t *testing.T) {
	ch1 := makeIntChan(1, 2)
	ch2 := makeIntChan(3, 4)
	ch3 := makeIntChan(5, 6)

	out := Merge(ch1, ch2, ch3)

	var result []int
	for v := range out {
		result = append(result, v)
	}

	sort.Ints(result)
	expected := []int{1, 2, 3, 4, 5, 6}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ожидалось %v, получено %v", expected, result)
	}
}

func TestMerge_EmptyChannels(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	close(ch1)
	close(ch2)

	out := Merge(ch1, ch2)

	var result []int
	for v := range out {
		result = append(result, v)
	}

	if len(result) != 0 {
		t.Errorf("ожидалось пусто, получено %v", result)
	}
}

func TestMerge_SingleChannel(t *testing.T) {
	ch := makeIntChan(7, 8, 9)

	out := Merge(ch)

	var result []int
	for v := range out {
		result = append(result, v)
	}

	expected := []int{7, 8, 9}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ожидалось %v, получено %v", expected, result)
	}
}
