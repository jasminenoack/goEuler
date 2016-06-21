package main

import (
	"math"
	"strconv"
)

// LargestPalindromeProduct finding largest palindrome
func LargestPalindromeProduct(digits int) float64 {
	// largestPalindrome := 0
	maxStart := 0.0
	minStart := 0.0
	largestPalindrome := 0.0

	for i := 0; i < digits; i++ {
		maxStart += 9.0 * math.Pow(10.0, float64(i))
		if i == digits-1 {
			minStart = 1.0 * math.Pow(10.0, float64(i))
		}
	}

	for i := maxStart; i >= minStart; i-- {
		for j := i; j >= minStart; j-- {
			currentProduct := i * j
			if currentProduct < largestPalindrome {
				break
			}
			currentProductString := strconv.FormatFloat(currentProduct, 'f', -1, 64)
			isPal := isPalindrome(currentProductString)
			if isPal {
				largestPalindrome = currentProduct
			}
		}
	}
	return largestPalindrome
}

func isPalindrome(str string) bool {
	stringMiddle := len(str) / 2
	for i := 0; i < stringMiddle; i++ {
		lastIndex := len(str) - 1 - i
		if str[i] != str[lastIndex] {
			return false
		}
	}
	return true
}
