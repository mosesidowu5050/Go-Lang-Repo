package main

import "fmt"

func main() {
	itemPrices := []float64{239.99, 129.75, 99.95, 350.89}
	basePay := 200.00
	commissionRate := 0.09
	totalSales := 0.0

	for {
		fmt.Println("Enter item number (1–4), or -1 to quit:")
		var itemNumber int
		fmt.Scanf("%d", &itemNumber)

		if itemNumber == -1 {
			fmt.Println("Exiting input...")
			break
		}

		if itemNumber < 1 || itemNumber > len(itemPrices) {
			fmt.Printf("Invalid item number. Please enter a number between 1 and %d\n", len(itemPrices))
			continue
		}

		fmt.Printf("You selected item %d with price $%.2f\n", itemNumber, itemPrices[itemNumber-1])

		fmt.Print("Enter the quantity sold: ")
		var quantitySold int
		fmt.Scanf("%d", &quantitySold)

		itemTotalPrice := float64(quantitySold) * itemPrices[itemNumber-1]
		totalSales += itemTotalPrice

		fmt.Printf("Subtotal for this item: $%.2f\n\n", itemTotalPrice)
	}

	commission := commissionRate * totalSales
	totalPay := basePay + commission

	fmt.Printf("Total Pay: $%.2f\n", totalPay)
}
