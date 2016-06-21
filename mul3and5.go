package main

// Mul3And5 Euler project 1.
func Mul3And5(max int) int {
	total := 0
	for i := 1; i < max; i++ {
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}
	return total
}
