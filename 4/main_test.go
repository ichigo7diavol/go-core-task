package main

import (
	"reflect"
	"testing"
)

func TestDifference_CommonCase(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	expected := []string{"apple", "cherry", "43", "lead", "gno1"}
	result := Difference(slice1, slice2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ожидалось %v, получено %v", expected, result)
	}
}

func TestDifference_EmptySecond(t *testing.T) {
	slice1 := []string{"a", "b", "c"}
	slice2 := []string{}

	expected := []string{"a", "b", "c"}
	result := Difference(slice1, slice2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ожидалось %v, получено %v", expected, result)
	}
}

func TestDifference_AllOverlap(t *testing.T) {
	slice1 := []string{"x", "y"}
	slice2 := []string{"x", "y", "z"}

	expected := []string{}
	result := Difference(slice1, slice2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ожидалось пусто, получено %v", result)
	}
}

func TestDifference_EmptyFirst(t *testing.T) {
	slice1 := []string{}
	slice2 := []string{"a", "b"}

	expected := []string{}
	result := Difference(slice1, slice2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ожидалось %v, получено %v", expected, result)
	}
}
