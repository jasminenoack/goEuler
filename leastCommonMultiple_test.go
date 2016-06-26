package main

import (
	"testing"
)

func TestLeastCommonMultiples(t *testing.T) {
	result := LeastCommonMultiple(1)
	if result != 1 {
		t.Errorf("%d, does not equal %d", result, 1)
	}

	result = LeastCommonMultiple(2)
	if result != 2 {
		t.Errorf("%d, does not equal %d", result, 2)
	}

	result = LeastCommonMultiple(3)
	if result != 6 {
		t.Errorf("%d, does not equal %d", result, 6)
	}

	result = LeastCommonMultiple(4)
	if result != 12 {
		t.Errorf("%d, does not equal %d", result, 12)
	}

	result = LeastCommonMultiple(10)
	if result != 2520 {
		t.Errorf("%d, does not equal %d", result, 2520)
	}
}

func BenchmarkLeastCommonMultiple5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LeastCommonMultiple(5)
	}
}

func BenchmarkLeastCommonMultiple10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LeastCommonMultiple(10)
	}
}

func BenchmarkLeastCommonMultiple100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LeastCommonMultiple(20)
	}
}
