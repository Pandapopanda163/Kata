package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		if input[0] == "exit" {
			break
		}
		if len(input) == 3 {
			switch {
			case isValidRomanChar(input[0]) && isValidRomanChar(input[2]):
				num1 := romanValue(input[0])
				num2 := romanValue(input[2])
				value := calculator(num1, num2, input[1])
				fmt.Println(valuesRoman(value))
				continue
			case !(isValidRomanChar(input[0]) || isValidRomanChar(input[2])):
				num1, _ := strconv.Atoi(input[0])
				num2, _ := strconv.Atoi(input[2])
				value := calculator(num1, num2, input[1])
				fmt.Println(value)
				continue
			default:
				panic("Работать с арабскими и римскими цифрами вместе, пока не могу")
			}
		} else {
			panic("Ввод с пробелами!")
		}

	}
}

func calculator(num1 int, num2 int, sign string) int {
	if !(0 < num1 && num1 < 11) || !(0 < num2 && num2 < 11) {
		panic("Числа должны быть в диапозоне от 1 до 10 включительно!")
	}
	switch sign {
	case "+":
		//fmt.Println(num1 + num2)
		return num1 + num2
	case "-":
		//fmt.Println(num1 - num2)
		return num1 - num2
	case "*":
		//fmt.Println(num1 * num2)
		return num1 * num2
	case "/":
		return num1 / num2
	default:
		panic("Неверная операция!")
	}
}

func isValidRomanChar(chars string) bool {
	for _, char := range chars {
		if !strings.Contains("IVX", string(char)) {
			return false
		}
	}
	return true
}

func romanValue(chars string) int {
	result := 0
	if chars == "IV" {
		return 4
	}
	if chars == "IX" {
		return 9
	}
	for _, char := range chars {
		switch char {
		case 'I':
			result += 1
		case 'V':
			result += 5
		case 'X':
			result += 10
		}
	}
	return result
}

func valuesRoman(n int) string {
	result := ""
	if n < 1 {
		panic("Число может быть только положительным")
	}
	for _, numeral := range []struct {
		value  int
		symbol string
	}{
		{100, "C"},
		{50, "L"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	} {
		for n >= numeral.value {
			result += numeral.symbol
			n -= numeral.value
		}
	}
	return result
}
