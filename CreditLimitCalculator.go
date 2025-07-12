package main

import "fmt"

func main() {

	var accountNumber int
	var beginningBalance float64
	var totalCharges float64
	var totalCredits float64
	var creditLimit float64

	for {
		fmt.Print("Enter account number or press (-1 to quit): ")
		fmt.Scanf("%d", &accountNumber)

		if accountNumber == -1 {
			fmt.Println("Exiting the program.")
			break
		}

		fmt.Print("Enter your starting balance: ")
		fmt.Scanf("%f", &beginningBalance)

		fmt.Print("Enter the total charges: ")
		fmt.Scanf("%f", &totalCharges)

		fmt.Print("Enter the total credits: ")
		fmt.Scanf("%f", &totalCredits)

		fmt.Print("Enter your credit limit: ")
		fmt.Scanf("%f", &creditLimit)

		currentBalance := beginningBalance + totalCharges - totalCredits
		fmt.Printf("Balance for this Account Number %d is: %.2f", accountNumber, currentBalance)
		fmt.Println()

		if currentBalance > creditLimit {
			fmt.Printf("Credit limit exceeded for account number %d\n. "+
				"Current balance: %.2f, Credit limit: %.2f\n", accountNumber, currentBalance, creditLimit)
		}
	}
}
