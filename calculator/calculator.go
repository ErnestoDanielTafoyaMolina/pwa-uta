package main

import (
	"fmt"
)

func handlePlus(num1, num2 int) int {
	return num1 + num2
}
func handleMinus(num1, num2 int) int {
	return num1 - num2
}
func handleTimes(num1, num2 int) int {
	return num1 * num2
}
func handleDivition(num1, num2 int) (int, string) {
	if num2 == 0 {
		return 0, "it is not allowed to divide by zero"
	}
	return num1 / num2, ""
}

func main() {

	var number1 int
	var number2 int
	var result int

	var operator string

	fmt.Print("\nType the first number: ")
	fmt.Scan(&number1)

	fmt.Print("\nType the operator: ")
	fmt.Scan(&operator)

	fmt.Print("\nType the second number: ")
	fmt.Scan(&number2)

	switch operator {
	case "+":
		result = handlePlus(number1, number2)
		fmt.Printf("\n%d %s %d = %d", number1, operator, number2, result)
	case "-":
		result = handleMinus(number1, number2)
		fmt.Printf("\n%d %s %d = %d", number1, operator, number2, result)
	case "*":
		result = handleTimes(number1, number2)
		fmt.Printf("\n%d %s %d = %d", number1, operator, number2, result)
	case "/":
		result, errorMessage := handleDivition(number1, number2)

		//divition has another step to complete
		if errorMessage != "" {
			fmt.Println("Error")
		} else {
			fmt.Printf("\n%d %s %d = %d", number1, operator, number2, result)
		}
	default:
		fmt.Println("invalid operator")
	}

}
