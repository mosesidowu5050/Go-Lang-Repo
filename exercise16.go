package main

import "fmt"

func main() {
	x := 1
	total := 0

	for x <= 10 {
		total += x
		x++
	}
	fmt.Println("Total is", total)
}
