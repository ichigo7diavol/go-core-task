package main

import (
	"reflect"
	"testing"
)

func TestCubePipeline(t *testing.T) {
	in := make(chan uint8)
	out := make(chan float64)

	go CubePipeline(in, out)

	go func() {
		for _, v := range []uint8{1, 2, 3, 4} {
			in <- v
		}
		close(in)
	}()

	var results []float64
	for v := range out {
		results = append(results, v)
	}

	expected := []float64{1, 8, 27, 64}

	if !reflect.DeepEqual(results, expected) {
		t.Errorf("ожидалось %v, получено %v", expected, results)
	}
}

func TestCubePipelineEmpty(t *testing.T) {
	in := make(chan uint8)
	out := make(chan float64)

	go CubePipeline(in, out)
	close(in)

	var results []float64
	for v := range out {
		results = append(results, v)
	}

	if len(results) != 0 {
		t.Errorf("ожидался пустой срез, получено %v", results)
	}
}
