package main

import (
	"reflect"
	"testing"
)

func TestAddAndGet(t *testing.T) {
	m := NewStringIntMap()
	m.Add("a", 1)
	m.Add("b", 2)

	val, ok := m.Get("a")
	if !ok || val != 1 {
		t.Errorf("ожидалось Get('a') = 1, получено %d, ok=%v", val, ok)
	}

	val, ok = m.Get("b")
	if !ok || val != 2 {
		t.Errorf("ожидалось Get('b') = 2, получено %d, ok=%v", val, ok)
	}
}

func TestRemove(t *testing.T) {
	m := NewStringIntMap()
	m.Add("x", 100)
	m.Remove("x")

	_, ok := m.Get("x")
	if ok {
		t.Errorf("элемент 'x' должен быть удален")
	}
}

func TestExists(t *testing.T) {
	m := NewStringIntMap()
	m.Add("key1", 42)

	if !m.Exists("key1") {
		t.Errorf("ожидалось, что ключ 'key1' существует")
	}

	if m.Exists("missing") {
		t.Errorf("ожидалось, что ключ 'missing' не существует")
	}
}

func TestCopy(t *testing.T) {
	m := NewStringIntMap()
	m.Add("a", 10)
	m.Add("b", 20)

	copyMap := m.Copy()
	if !reflect.DeepEqual(copyMap, m.data) {
		t.Errorf("копия не совпадает с оригиналом: %v vs %v", copyMap, m.data)
	}

	// проверяем независимость копии
	copyMap["a"] = 999
	if m.data["a"] == 999 {
		t.Errorf("копия должна быть независимой от оригинала")
	}
}
