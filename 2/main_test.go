package main

import (
	"reflect"
	"testing"
)

func TestSliceExample(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	expected := []int{2, 4, 6}
	result := sliceExample(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ожидалось %v, получено %v", expected, result)
	}
}

func TestAddElements(t *testing.T) {
	input := []int{1, 2, 3}
	expected := []int{1, 2, 3, 10}
	result := addElements(input, 10)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ожидалось %v, получено %v", expected, result)
	}
}

func TestCopySlice(t *testing.T) {
	input := []int{5, 6, 7}
	copy := copySlice(input)
	input[0] = 100

	if reflect.DeepEqual(input, copy) {
		t.Errorf("копия изменилась вместе с оригиналом")
	}
}

func TestRemoveElement(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := []int{1, 2, 4, 5}
	result := removeElement(input, 2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ожидалось %v, получено %v", expected, result)
	}
}

func TestRemoveElement_OutOfBounds(t *testing.T) {
	input := []int{1, 2, 3}
	result := removeElement(input, 10)
	if !reflect.DeepEqual(result, input) {
		t.Errorf("при неверном индексе результат должен быть тем же")
	}
}
