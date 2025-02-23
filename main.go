package main

import "fmt"

func getInput() string {
    var result string
    fmt.Println("Введите что-нибудь:")
    fmt.Scan(&result)
    return result
}

func main()  {
    fromUstToEur := 0.96
    fromUsdToRub := 88.45
    fromEurToRub := (1 / fromUstToEur) * fromUsdToRub
    fmt.Println(fromEurToRub)
    inputText := getInput()
    fmt.Println(inputText)
}

func calculate(amount float64, sourceCurrency string, targetCurrency string) float64 {
    return 0
} 
