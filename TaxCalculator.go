package main

import "fmt"

func main() {

	taxRateFifteenPercent := 0.15
	taxRateTwentyPercent := 0.20

	for {
		fmt.Print("Enter citizen's name or enter ('exist' to quit): ")
		var name string
		fmt.Scanf("%s", &name)

		if name == "exist" {
			fmt.Println("Exiting the program.")
			break
		}

		fmt.Print("Enter citizen's earnings: ")
		var earnings float64
		fmt.Scanf("%f", &earnings)

		var totalTax float64

		if earnings <= 30000 {
			totalTax = earnings * taxRateFifteenPercent
		} else {
			amountAbove := earnings - 30000
			totalTax = (30000 * taxRateFifteenPercent) + (amountAbove * taxRateTwentyPercent)
		}

		fmt.Printf("Total Tax for %s is $%.2f on earnings of $%.2f.\n\n", name, totalTax, earnings)
		fmt.Println()
	}
}
