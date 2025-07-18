package main

import "fmt"

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {

	x := 5
	fmt.Printf("Factorial of 5 is: %d", factorial(uint(x)))
}
