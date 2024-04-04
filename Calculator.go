package main

import "fmt"

func main() {
	var num_1 float64
	var num_2 float64
	var opr string
	fmt.Println("введите число :")
	fmt.Scan(&num_1)
	fmt.Println("введите число :")
	fmt.Scan(&num_2)
	fmt.Println("введите операцию :")
	fmt.Scan(&opr)
	if opr == "+" {
		sum := num_1 + num_2
		fmt.Println(sum)
	} else if opr == "-" {
		sum := num_1 - num_2
		fmt.Println(sum)
	} else if opr == "*" {
		sum := num_1 * num_2
		fmt.Println(sum)

	} else if opr == "/" {
		sum := num_1 / num_2
		fmt.Println(sum)
	}
	fmt.Println("всё")
}
