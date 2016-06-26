package main

import (
	"testing"
)

func TestLargestPalindromeProduct(t *testing.T) {
	result := LargestPalindromeProduct(2)
	if result != 9009 {
		t.Errorf("%g, does not equal %d", result, 9009)
	}

	result = LargestPalindromeProduct(1)
	if result != 9 {
		t.Errorf("%g, does not equal %d", result, 9)
	}
}

func BenchmarkLargestPalindromeProduct1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LargestPalindromeProduct(1)
	}
}

func BenchmarkLargestPalindromeProduct2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LargestPalindromeProduct(2)
	}
}

func BenchmarkLargestPalindromeProduct3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LargestPalindromeProduct(3)
	}
}
