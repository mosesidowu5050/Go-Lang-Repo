package main

import "fmt"

func main() {

	fmt.Print("Enter first number to compare: ")
	var numberOne int
	fmt.Scanf("%d", &numberOne)

	fmt.Print("Enter second number to compare: ")
	var numberTwo int
	fmt.Scanf("%d", &numberTwo)

	if numberOne == numberTwo {
		fmt.Print(0)
	} else if numberOne > numberTwo {
		fmt.Print(1)
	} else if numberTwo > numberOne {
		fmt.Print(-1)
	}
}
