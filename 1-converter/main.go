package main

import (
	"fmt"
	"strconv"
	"strings"
)

const USD = "USD"
const EUR = "EUR"
const RUB = "RUB"

func getSourceCurrency() string {
	var input string

	for {
		fmt.Print("Выберите исходную валюту USD|EUR|RUB: ")
		fmt.Scan(&input)
		result := strings.ToUpper(input)
		if result == USD || result == EUR || result == RUB {
			return result
		}
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
	}
}

func getTargetCurrency(firstTargetCurrency string, secondTargetCurrency string) string {
	var input string

	for {
		fmt.Printf("Выберите целевую валюту %s|%s: ", firstTargetCurrency, secondTargetCurrency)
		fmt.Scan(&input)
		result := strings.ToUpper(input)
		if result == firstTargetCurrency || result == secondTargetCurrency {
			return result
		}
	}
}

func getValidTargetCurrencies(sourceCurrency string) (string, string) {
	switch {
	case sourceCurrency == USD:
		return EUR, RUB
	case sourceCurrency == EUR:
		return USD, RUB
	default:
		return USD, EUR
	}
}

func main() {
	sourceCurrency := getSourceCurrency()
	amount := getAmount()
	firstTargetCurrency, secondTargetCurrency := getValidTargetCurrencies(sourceCurrency)
	targetCurrency := getTargetCurrency(firstTargetCurrency, secondTargetCurrency)
	
    result := calculate(amount, sourceCurrency, targetCurrency)
	fmt.Printf("Вы получите %.2f %s\n", result, targetCurrency)
}

func calculate(amount float64, sourceCurrency string, targetCurrency string) float64 {
	const fromUstToEur = 0.96
	const fromUsdToRub = 88.45
	const fromEurToRub = (1 / fromUstToEur) * fromUsdToRub

	switch {
	case sourceCurrency == USD && targetCurrency == EUR:
		return amount * fromUstToEur
	case sourceCurrency == USD && targetCurrency == RUB:
		return amount * fromUsdToRub
	case sourceCurrency == EUR && targetCurrency == USD:
		return amount * (1 / fromUstToEur)
	case sourceCurrency == EUR && targetCurrency == RUB:
		return amount * fromEurToRub
	case sourceCurrency == RUB && targetCurrency == EUR:
		return amount * (1 / fromEurToRub)
	default:
		return amount * (1 / fromUsdToRub)
	}
}
