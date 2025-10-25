// // go:build unit

package main

import (
	"strings"
	"testing"
)

func TestGetType(t *testing.T) {
	tests := []struct {
		value    any
		expected string
	}{
		{42, "int"},
		{3.14, "float64"},
		{"Golang", "string"},
		{true, "bool"},
		{complex64(1 + 2i), "complex64"},
	}

	for _, tt := range tests {
		got := GetType(tt.value)
		if got != tt.expected {
			t.Errorf("GetType(%v) = %s, ожидалось %s", tt.value, got, tt.expected)
		}
	}
}

func TestCombineToString(t *testing.T) {
	tests := []struct {
		input    []any
		expected string
	}{
		{[]any{1, "abc", true}, "1abctrue"},
		{[]any{"Go", 2024, false}, "Go2024false"},
	}

	for _, tt := range tests {
		got := CombineToString(tt.input...)
		if got != tt.expected {
			t.Errorf("CombineToString(%v) = %s, ожидалось %s", tt.input, got, tt.expected)
		}
	}
}

func TestToRunes(t *testing.T) {
	s := "Go語"
	runes := ToRunes(s)
	expected := []rune{'G', 'o', '語'}

	if len(runes) != len(expected) {
		t.Fatalf("ToRunes(%s) вернул %d рун, ожидалось %d", s, len(runes), len(expected))
	}

	for i, r := range expected {
		if runes[i] != r {
			t.Errorf("Руна %d: ожидалось %q, получено %q", i, r, runes[i])
		}
	}
}

func TestHashWithSaltSHA256(t *testing.T) {
	input := []rune("test")
	salt := "go-2024"

	hash1 := HashWithSaltSHA256(input, salt)
	hash2 := HashWithSaltSHA256(input, salt)
	if hash1 != hash2 {
		t.Error("Хэш должен быть детерминированным при одинаковых входных данных")
	}

	hash3 := HashWithSaltSHA256(input, "other-salt")
	if hash1 == hash3 {
		t.Error("Разные соли должны давать разные хэши")
	}

	if len(hash1) != 64 {
		t.Errorf("Длина SHA256 хэша должна быть 64 символа, получено %d", len(hash1))
	}

	if !strings.ContainsAny(hash1, "abcdef0123456789") {
		t.Error("Хэш должен быть в hex-формате")
	}
}
