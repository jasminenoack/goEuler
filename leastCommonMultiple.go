package main

func isFactor(largeNumber, smallNumber int) bool {
	return largeNumber%smallNumber == 0
}

func leastCommonMultiple(numbers []int) int {
	highestNumber := numbers[len(numbers)-1]
	currentNumber := highestNumber
	for !divisibleByAll(currentNumber, numbers) {
		currentNumber += highestNumber
	}
	return currentNumber
}

func divisibleByAll(number int, numbers []int) bool {
	for _, testNumber := range numbers {
		if number%testNumber != 0 {
			return false
		}
	}
	return true
}

// LeastCommonMultiple returns the least common multiple of all numbers in range of the int
func LeastCommonMultiple(max int) int {
	numbers := []int{}
	for i := 1; i <= max; i++ {
		numbers = append(numbers, i)
	}

	for i := max; i > 0; i-- {
		newNumbers := []int{}
		for _, number := range numbers {
			if !(number != i && isFactor(i, number)) {
				newNumbers = append(newNumbers, number)
			}
		}
		numbers = newNumbers
	}

	return leastCommonMultiple(numbers)
}
