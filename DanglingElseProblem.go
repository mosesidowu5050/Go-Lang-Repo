package main

import "fmt"

func main() {

	x := 1
	y := 2

	if x > 5 {
		if y > 5 {
			fmt.Println("x and y are > 5")
		}
	} else {
		fmt.Println("x and y are <= 5")
	}
}
