package main

import (
	"testing"
)

func TestLargestPrimeFactor(t *testing.T) {
	_, err := LargestPrimeFactor(0)
	if err == nil {
		t.Errorf("0 did not return error.")
	}

	_, err = LargestPrimeFactor(1)
	if err == nil {
		t.Errorf("1 did not return error")
	}

	result, _ := LargestPrimeFactor(2)
	if result != 2 {
		t.Errorf("%d, does not equal %d", result, 2)
	}

	result, _ = LargestPrimeFactor(3)
	if result != 3 {
		t.Errorf("%d, does not equal %d", result, 3)
	}

	result, _ = LargestPrimeFactor(4)
	if result != 2 {
		t.Errorf("%d, does not equal %d", result, 2)
	}

	result, _ = LargestPrimeFactor(5)
	if result != 5 {
		t.Errorf("%d, does not equal %d", result, 5)
	}

	result, _ = LargestPrimeFactor(6)
	if result != 3 {
		t.Errorf("%d, does not equal %d", result, 3)
	}

	result, _ = LargestPrimeFactor(7)
	if result != 7 {
		t.Errorf("%d, does not equal %d", result, 7)
	}

	result, _ = LargestPrimeFactor(8)
	if result != 2 {
		t.Errorf("%d, does not equal %d", result, 2)
	}

	result, _ = LargestPrimeFactor(9)
	if result != 3 {
		t.Errorf("%d, does not equal %d", result, 3)
	}

	result, _ = LargestPrimeFactor(10)
	if result != 5 {
		t.Errorf("%d, does not equal %d", result, 5)
	}
}

func BenchmarkLargestPrimeFactorFor10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LargestPrimeFactor(10)
	}
}

func BenchmarkLargestPrimeFactorFor100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LargestPrimeFactor(100)
	}
}

func BenchmarkLargestPrimeFactorFor600851475143(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LargestPrimeFactor(600851475143)
	}
}
