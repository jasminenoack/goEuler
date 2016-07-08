package main

import (
	"math"
	"runtime"
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

// LargestPalindromeProduct2 with go routines
// uses go routines, actually takes longer probably due to switching routines.
func LargestPalindromeProduct2(digits int) int {
	runtime.GOMAXPROCS(10)
	maxStart := 0.0
	minStart := 0.0
	ch := make(chan int, 10000)
	largestPalindrome := 0
	for i := 0; i < digits; i++ {
		maxStart += 9.0 * math.Pow(10.0, float64(i))
		if i == digits-1 {
			minStart = 1.0 * math.Pow(10.0, float64(i))
		}
	}
	for i := maxStart; i >= minStart; i-- {
		go determineLargestPalindromeRoutine(int(i), ch)
	}
	for j := maxStart; j >= minStart; j-- {
		number := <-ch
		if number > largestPalindrome {
			largestPalindrome = number
		}
	}
	close(ch)
	return largestPalindrome
}

func determineLargestPalindromeRoutine(number int, ch chan<- int) {
	for i := number; i > 0; i-- {
		currentProduct := number * i
		isPal := isPalindrome(strconv.Itoa(currentProduct))
		if isPal {
			ch <- currentProduct
			return
		}
	}
	ch <- 0
	return
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
