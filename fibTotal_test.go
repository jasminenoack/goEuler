package main

import (
	"testing"
)

func TestFibTotal(t *testing.T) {
	result := FibEvenTotal(2)
	if result != 0 {
		t.Errorf("%d, does not equal %d", result, 0)
	}

	result = FibEvenTotal(3)
	if result != 2 {
		t.Errorf("%d, does not equal %d", result, 2)
	}

	result = FibEvenTotal(14)
	if result != 10 {
		t.Errorf("%d, does not equal %d", result, 10)
	}

	result = FibEvenTotal(100)
	if result != 44 {
		t.Errorf("%d, does not equal %d", result, 44)
	}
}

func BenchmarkFibTotalFor10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibEvenTotal(10)
	}
}

func BenchmarkFibTotalFor100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibEvenTotal(100)
	}
}

func BenchmarkFibTotalFor4000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibEvenTotal(4000000)
	}
}
