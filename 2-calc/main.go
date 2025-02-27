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

func getOperation() string {
	var input string
	for {
		fmt.Print("Выберите операцию AVG|SUM|MED: ")
		fmt.Scan(&input)
		result := strings.ToUpper(input)
		if result == AVG || result == SUM || result == MED {
			return result
		}
	}
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
			err = fmt.Errorf("\033[33mНевозможно преобразовать в строку последовательность '%s', попробуйте еще раз.\033[0m", trimmedPart)
			break
		}
		numbers = append(numbers, number)
	}
	return numbers, err
}

func calculate(operation string, numbers []float64) float64 {
	switch operation {
	case MED:
		return calculateMed(numbers)
	case AVG:
		return calculateAvg(numbers)
	default:
		return calculateSum(numbers)
	}
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
