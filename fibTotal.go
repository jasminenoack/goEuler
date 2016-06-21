package main

// FibEvenTotal euler 2
func FibEvenTotal(max int) int {
	total := 0
	previousFib := 0
	nextFib := 1
	for nextFib < max {
		if nextFib%2 == 0 {
			total += nextFib
		}
		previousFib, nextFib = nextFib, previousFib+nextFib
	}
	return total
}
