package main

import "fmt"

func main() {

	var totalMilesDriven int
	var totalGallonsUsed int

	for {
		fmt.Print("Enter the number of miles driven (-1 to quit): ")
		var milesDriven int
		fmt.Scanf("%d", &milesDriven)

		if milesDriven == -1 {
			break
		}

		fmt.Print("Enter the gallons of gas used: ")
		var gallonsUsed int
		fmt.Scanf("%d", &gallonsUsed)

		totalMilesDriven += milesDriven
		totalGallonsUsed += gallonsUsed

		tripMilesPerGallon := float64(milesDriven) / float64(gallonsUsed)
		fmt.Printf("Miles per gallon for this trip: %.2f\n\n", tripMilesPerGallon)
	}

	totalMilesPerGallon := float64(totalMilesDriven) / float64(totalGallonsUsed)
	fmt.Printf("Total miles per gallon for all trips: %.2f\n", totalMilesPerGallon)
}
