// euler.go
package main

import "fmt"

func main() {
	fmt.Println("methods to solve the euler problems")
	fmt.Println("")

	euler1 := Mul3And5(1000)
	fmt.Printf("Euler 1: %d", euler1)
	fmt.Println("")

	euler2 := FibEvenTotal(4000000)
	fmt.Printf("Euler 2: %d", euler2)
	fmt.Println("")

	euler3, _ := LargestPrimeFactor(600851475143)
	fmt.Printf("Euler 3: %d", euler3)
	fmt.Println("")

	euler4 := LargestPalindromeProduct(3)
	fmt.Printf("Euler 4: %g", euler4)
	fmt.Println("")

	euler5 := LeastCommonMultiple(20)
	fmt.Printf("Euler 5: %d", euler5)
	fmt.Println("")
}
