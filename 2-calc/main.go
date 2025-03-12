package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const AVG = "AVG"
const SUM = "SUM"
const MED = "MED"

var (
    colorYellow = "\033[33m"
    colorReset = "\033[0m"
)

var operationsMap = map[string]func([]float64) float64{
    "AVG": calculateAvg,
    "SUM": calculateSum,
    "MED": calculateMed,
}

func getOperation() string {
	var input string
    operationsString := getStringFromMapKeys(operationsMap)
    inputMessage := "Выберите операцию " + operationsString + ": "
	for {
		fmt.Print(inputMessage)
		fmt.Scan(&input)
		result := strings.ToUpper(input)
		if _, ok := operationsMap[result]; ok {
            return result
		}
	}
}

func getStringFromMapKeys[T any](data map[string]T) string {
	keys := make([]string, 0, len(data))
	for key := range data {
		keys = append(keys, key)
	}

	result := strings.Join(keys, "|")
	return result
}

func getNumbers() []float64 {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Введите числа через запятую (допустимы только числа и запятые): ")
		scanner.Scan()
		input := scanner.Text()
		numbers, err := convertStringToFloatArray(input)
		if err == nil {
			return numbers
		}
		fmt.Println(err)
	}
}

func convertStringToFloatArray(str string) ([]float64, error) {
	var err error

	inputParts := strings.Split(str, ",")
	numbers := make([]float64, 0, len(inputParts))
	for _, part := range inputParts {
		trimmedPart := strings.TrimSpace(part)
		number, parseErr := strconv.ParseFloat(trimmedPart, 64)
		if parseErr != nil {
			err = fmt.Errorf("%sНевозможно преобразовать в строку последовательность '%s', попробуйте еще раз.%s", colorYellow, trimmedPart, colorReset)
			break
		}
		numbers = append(numbers, number)
	}
	return numbers, err
}

func calculate(operation string, numbers []float64) float64 {
	fn := operationsMap[operation]
    return fn(numbers)
}

func calculateMed(numbers []float64) float64 {
	sliceLen := len(numbers)
	if sliceLen%2 == 0 {
		return (numbers[sliceLen/2-1] + numbers[sliceLen/2]) / 2
	} else {
		return numbers[((sliceLen - 1) / 2)]
	}
}

func calculateAvg(numbers []float64) float64 {
	var count float64
	var sum float64
	for _, value := range numbers {
		count += 1
		sum += value
	}
	return sum / count
}

func calculateSum(numbers []float64) float64 {
	var sum float64
	for _, value := range numbers {
		sum += value
	}
	return sum
}

func main() {
	operation := getOperation()
	numbers := getNumbers()
    result := calculate(operation, numbers)
    fmt.Printf("Результат операции %s: %.2f\n", operation, result)
}
