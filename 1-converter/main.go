package main

import (
	"errors"
	"fmt"
)

func main() {
	currencyFrom := getCurrencyFrom()
	currencyTo := getCurrencyTo(currencyFrom)
	amount := getAmount()

	fmt.Println("Сумма к получению:", calculate(currencyFrom, currencyTo, amount), currencyTo)

}

func getCurrencyFrom() string {
	var currencyFrom string
	fmt.Print("Введите валюту, которую хотите конвертировать (EUR|USD|RUB): ")
	fmt.Scan(&currencyFrom)
	_, err := isCurrencyInputValid(currencyFrom)
	if err != nil {
		fmt.Println("Введена некорректная валюта, попробуйте еще раз", err)
		return getCurrencyFrom()
	}

	return currencyFrom
}

func getCurrencyTo(currencyFrom string) string {
	var currencyTo string
	fmt.Println(getTitleForCurrencyFrom(currencyFrom))
	fmt.Scan(&currencyTo)
	_, err := isCurrencyInputValid(currencyTo)
	if err != nil || currencyFrom == currencyTo {
		fmt.Println("Введена некорректная валюта, попробуйте еще раз")
		return getCurrencyTo(currencyFrom)
	}

	return currencyTo
}

func getAmount() float64 {
	var result float64
	fmt.Println("Введите сумму, которую хотите конвертировать: ")
	fmt.Scan(&result)
	_, err := isAmountValid(result)
	if err != nil {
		fmt.Println("Введена некорректная сумма, попробуйте еще раз ввести число", err)
		return getAmount()
	}
	return result
}

func getTitleForCurrencyFrom(currencyFrom string) string {
	title := "Введите валюту в которую хотите конвертировать "
	switch currencyFrom {
	case "EUR":
		{
			return title + "USD|RUB: "
		}
	case "USD":
		{
			return title + "EUR|RUB: "
		}
	case "RUB":
		{
			return title + "EUR|USD: "
		}
	default:
		return title + "EUR|USD|RUB"
	}
}

func isCurrencyInputValid(currency string) (bool, error) {
	if currency == "EUR" || currency == "USD" || currency == "RUB" {
		return true, nil
	} else {
		return false, errors.New("invalid currency")
	}
}

func isAmountValid(amount float64) (bool, error) {
	if amount <= 0 {
		return false, errors.New("invalid amount")
	}
	return true, nil
}

func calculate(currencyFrom, currencyTo string, amount float64) string {
	const USDToEUR = 0.86
	const USDToRUB = 80.69
	const EURToRUB = 100

	var result float64

	switch {
	case currencyFrom == "EUR" && currencyTo == "USD":
		{
			result = amount / USDToEUR
		}

	case currencyFrom == "USD" && currencyTo == "EUR":
		{
			result = amount * USDToEUR
		}

	case currencyFrom == "EUR" && currencyTo == "RUB":
		{
			result = amount * EURToRUB
		}

	case currencyFrom == "RUB" && currencyTo == "EUR":
		{
			result = amount / EURToRUB
		}

	case currencyFrom == "USD" && currencyTo == "RUB":
		{
			result = amount * USDToRUB
		}

	case currencyFrom == "RUB" && currencyTo == "USD":
		{
			result = amount / USDToRUB
		}
	}

	formattedResult := fmt.Sprintf("%.2f", result)

	return formattedResult
}
