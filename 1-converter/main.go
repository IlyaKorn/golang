package main

import "fmt"

func main() {
	const USDToEURO = 0.86
	const USDToRUB = 80.69

	EUROToRUB := 1 / USDToEURO * USDToRUB

	fmt.Print(EUROToRUB)
}

func getUserInput() string {
	var input string
	fmt.Print("Enter your input: ")
	fmt.Scan(&input)
	return input
}

func calculate(sum int, currencyFrom string, currencyTo string) float64 {
	return 1.0
}
