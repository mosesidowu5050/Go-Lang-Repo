package main

import "fmt"

func main() {

	fmt.Printf("N \t\tN2 \t\tN3 \t\tN4")
	fmt.Println()

	for counter := 1; counter <= 5; counter++ {
		fmt.Printf("%d\t\t%d\t\t%d\t\t%d\n", counter, counter*counter,
			counter*counter*counter, counter*counter*counter*counter)

	}
}
