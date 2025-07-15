package main

import "fmt"

func main() {

	sumNumbers := 0

	fmt.Print("Enter a specified number: ")
	var specifiedNumber int
	fmt.Scanf("%d", &specifiedNumber)

	for {
		fmt.Print("Enter integer: ")
		var number int
		fmt.Scanf("%d", &number)
		sumNumbers += number

		if sumNumbers == specifiedNumber {
			fmt.Printf("Sum of those integers equals the value that was input in the beginning. Sum Is: %d.\n", sumNumbers)
			break
		} else if sumNumbers > specifiedNumber {
			fmt.Printf("Sum of those integers greater than value that was input in the beginning. Sum Is: %d.\n", sumNumbers)
			break
		}
	}
}
