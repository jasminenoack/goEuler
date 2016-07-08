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

func TestLargestPalindromeProduct2(t *testing.T) {
	result := LargestPalindromeProduct2(2)
	if result != 9009 {
		t.Errorf("%d, does not equal %d", result, 9009)
	}

	result = LargestPalindromeProduct2(1)
	if result != 9 {
		t.Errorf("%d, does not equal %d", result, 9)
	}
}

func BenchmarkLargestPalindromeProduct1_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LargestPalindromeProduct(1)
	}
}

func BenchmarkLargestPalindromeProduct1_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LargestPalindromeProduct(2)
	}
}

func BenchmarkLargestPalindromeProduct1_3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LargestPalindromeProduct(3)
	}
}

func BenchmarkLargestPalindromeProduct2_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LargestPalindromeProduct2(1)
	}
}

func BenchmarkLargestPalindromeProduct2_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LargestPalindromeProduct2(2)
	}
}

func BenchmarkLargestPalindromeProduct2_3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LargestPalindromeProduct2(3)
	}
}
