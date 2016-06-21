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

func TestLargestPalindromeProduct(t *testing.T) {
	result := LargestPalindromeProduct(2)
	if result != 9009 {
		t.Errorf("%d, does not equal %d", result, 9009)
	}

	result = LargestPalindromeProduct(1)
	if result != 9 {
		t.Errorf("%d, does not equal %d", result, 9)
	}
}
