package main

import (
	"fmt"
	"strconv"
	"strings"
)

var currencyRates = map[string]map[string]float64{
	"USD": {"EUR": 0.95, "RUB": 91.07},
	"EUR": {"USD": 1.05, "RUB": 95.61},
	"RUB": {"USD": 0.011, "EUR": 0.01},
}

func getStringFromMapKeys[T any](data map[string]T) string {
	keys := make([]string, 0, len(data))
	for key := range data {
		keys = append(keys, key)
	}

	result := strings.Join(keys, "|")
	return result
}
func getSourceCurrency() string {
	var input string
	currenciesString := getStringFromMapKeys(currencyRates)
	for {
		inputMessage := "Выберите исходную валюту " + currenciesString + ":"
		fmt.Print(inputMessage)
		fmt.Scan(&input)
		currency := strings.ToUpper(input)
		if _, ok := currencyRates[currency]; ok {
			return currency
		}
		printRed("Введено некорректное наименование исходной валюты, попробуйте еще раз")
	}
}

func getAmount() float64 {
	var input string

	for {
		fmt.Print("Введите количество валюты: ")
		fmt.Scan(&input)
		result, err := strconv.ParseFloat(input, 64)
		if err == nil {
			return result
		}
		printRed("Введено некорректное значение количества валюты, попробуйте еще раз")
	}
}

func printRed(s string) {
	fmt.Printf("\033[31m%s\033[0m\n\n", s)
}

func getTargetCurrency(sourceCurrency string) string {
	var input string
	currenciesString := getStringFromMapKeys(currencyRates[sourceCurrency])
	for {
		inputMessage := "Выберите целевую валюту " + currenciesString + ":"
		fmt.Print(inputMessage)
		fmt.Scan(&input)
		currency := strings.ToUpper(input)
		if _, ok := currencyRates[sourceCurrency][currency]; ok {
			return currency
		}
		printRed("Введено некорректное наименование целевой валюты, попробуйте еще раз")
	}
}

func main() {
	sourceCurrency := getSourceCurrency()
	amount := getAmount()
	targetCurrency := getTargetCurrency(sourceCurrency)

	result := calculate(amount, sourceCurrency, targetCurrency)
	fmt.Printf("Вы получите %.2f %s\n", result, targetCurrency)
}

func calculate(amount float64, sourceCurrency string, targetCurrency string) float64 {
	rate := currencyRates[sourceCurrency][targetCurrency]
	return amount * rate
}
