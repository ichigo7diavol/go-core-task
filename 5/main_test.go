package main

import (
	"reflect"
	"testing"
)

func TestIntersect_CommonCase(t *testing.T) {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	expected := []int{3, 64}
	ok, result := Intersect(a, b)

	if !ok {
		t.Errorf("ожидалось, что пересечение есть")
	}

	if !reflect.DeepEqual(result, expected) && !reflect.DeepEqual(result, []int{64, 3}) {
		t.Errorf("ожидалось %v, получено %v", expected, result)
	}
}

func TestIntersect_NoIntersection(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}

	ok, result := Intersect(a, b)
	if ok {
		t.Errorf("ожидалось отсутствие пересечения")
	}
	if len(result) != 0 {
		t.Errorf("ожидался пустой срез, получено %v", result)
	}
}

func TestIntersect_EmptyInput(t *testing.T) {
	ok, result := Intersect([]int{}, []int{1, 2, 3})
	if ok || len(result) != 0 {
		t.Errorf("ожидалось false и пустой срез, получено %v, %v", ok, result)
	}

	ok, result = Intersect([]int{1, 2, 3}, []int{})
	if ok || len(result) != 0 {
		t.Errorf("ожидалось false и пустой срез, получено %v, %v", ok, result)
	}
}

func TestIntersect_Duplicates(t *testing.T) {
	a := []int{1, 1, 2, 3, 3}
	b := []int{3, 1}

	expected := []int{1, 3}
	ok, result := Intersect(a, b)

	if !ok {
		t.Errorf("ожидалось наличие пересечения")
	}
	if !reflect.DeepEqual(result, expected) && !reflect.DeepEqual(result, []int{3, 1}) {
		t.Errorf("ожидалось %v, получено %v", expected, result)
	}
}
