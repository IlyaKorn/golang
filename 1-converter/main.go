package main

import (
	"errors"
	"fmt"
)

type CurrencyExchange = map[string]float64
type CurrenciesCalculation = map[string]map[string]func(float64) float64

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
	text := map[string]string{
		"EUR": "USD|RUB: ",
		"USD": "EUR|RUB: ",
		"RUB": "EUR|USD: ",
	}

	return title + text[currencyFrom]
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
	currenciesList := CurrencyExchange{
		"USDToEUR": 0.86, "USDToRUB": 80.69, "EURToRUB": 100,
	}

	currencyCalculation := CurrenciesCalculation{
		"EUR": {
			"USD": func(amount float64) float64 {
				return amount * currenciesList["USDToEUR"]
			},
			"RUB": func(amount float64) float64 {
				return amount * currenciesList["EURToRUB"]
			},
		},
		"USD": {
			"EUR": func(amount float64) float64 {
				return amount / currenciesList["USDToEUR"]
			},
			"RUB": func(amount float64) float64 {
				return amount * currenciesList["USDToRUB"]
			},
		},
		"RUB": {
			"EUR": func(amount float64) float64 {
				return amount / currenciesList["EURToRUB"]
			},
			"USD": func(amount float64) float64 {
				return amount / currenciesList["USDToRUB"]
			},
		},
	}

	result := fmt.Sprintf("%.2f", currencyCalculation[currencyFrom][currencyTo](amount))

	return result
}
