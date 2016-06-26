package main

import (
	"testing"
)

func TestMult3or5(t *testing.T) {
	result := Mul3And5(5)
	if result != 3 {
		t.Errorf("%d, does not equal %d", result, 3)
	}

	result = Mul3And5(6)
	if result != 8 {
		t.Errorf("%d, does not equal %d", result, 8)
	}

	result = Mul3And5(7)
	if result != 14 {
		t.Errorf("%d, does not equal %d", result, 14)
	}

	result = Mul3And5(10)
	if result != 23 {
		t.Errorf("%d, does not equal %d", result, 23)
	}

	result = Mul3And5(11)
	if result != 33 {
		t.Errorf("%d, does not equal %d", result, 33)
	}
}

func BenchmarkMult3or5for10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Mul3And5(10)
	}
}

func BenchmarkMult3or5for100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Mul3And5(100)
	}
}

func BenchmarkMult3or5for1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Mul3And5(1000)
	}
}
