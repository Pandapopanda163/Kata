package main

import "fmt"

func main() {
	var x, y int
	var z string
	fmt.Println("введите операцию :")
	fmt.Scan(&x, &z, &y)
	if z == "+" {
		sum := x + y
		fmt.Println(sum)
	} else if z == "-" {
		sum := x - y
		fmt.Println(sum)
	} else if z == "*" {
		sum := x * y
		fmt.Println(sum)

	} else if z == "/" {
		sum := x / y
		fmt.Println(sum)
	}
	fmt.Println("всё")
}
