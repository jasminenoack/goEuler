package main

import "fmt"

// LargestPrimeFactor find the largest prime factor of a number
func LargestPrimeFactor(number int) (int, error) {
	i := 2
	if number < i {
		err := fmt.Errorf("invalid input")
		return 0, err
	}

	for i < number {
		if number%i == 0 {
			number /= i
		} else {
			i++
		}
	}

	return i, nil
}
