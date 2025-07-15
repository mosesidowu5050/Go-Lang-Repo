package main

import "fmt"

func main() {

	firstLargest := 0
	secondLargest := 0

	for count := 0; count < 10; count++ {
		fmt.Printf("Enter number %d: ", count+1)
		var number int
		fmt.Scanf("%d", &number)

		if count == 10 {
			fmt.Println("You have entered 10 numbers. Exiting input...")
			break
		}

		if number > firstLargest {
			secondLargest = firstLargest
			firstLargest = number
		}

		if number > secondLargest && number != firstLargest {
			secondLargest = number
		}
	}

	fmt.Printf("First Largest number is %d: ", firstLargest)
	fmt.Println()
	fmt.Printf("Second Largest number is %d: ", secondLargest)
}
