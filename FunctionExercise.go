package main

import "fmt"

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func f() (int, int) {
	return 5, 6
}

func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}

type Circle struct {
	x, y, z float64
}

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))

	x, y := f()
	fmt.Println(y, x)

	add := func(x, y float64) int {
		return int(x + y)
	}
	fmt.Println(add(5.0, 1.7))

	defer second()
	first()

	//panic("PANIC")
	//str := recover()
	//fmt.Println(str)

	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")

}
