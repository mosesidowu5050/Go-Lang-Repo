package main

import "fmt"

func main() {

	counter := 0
	largestNumber := 0

	for {
		fmt.Printf("Enter integer %d: ", counter+1)
		var number int
		fmt.Scanf("%d", &number)
		counter++

		if counter == 10 {
			fmt.Println("You have entered 10 numbers. Exiting input...")
			break
		}

		if number > largestNumber {
			largestNumber = number
		}

	}
	fmt.Printf("The largest number entered is: %d\n", largestNumber)

}
