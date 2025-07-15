package main

import "fmt"

func main() {

	fmt.Print("Enter number between 1 and 2: ")

	for {
		var number int
		fmt.Scanf("%d", &number)

		if number < 1 || number > 2 {
			fmt.Print("Invalid, Enter Number between 1 and 2: ")
			continue
		} else {
			fmt.Print("Existing...")
			break
		}
	}

}
