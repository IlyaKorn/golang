package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	currentOperation := chooseUserOperation()
	userNumbers := scanUserOperations()

	fmt.Println("RESULT: ", calculateNumbers(currentOperation, userNumbers))
}

func chooseUserOperation() string {
	var operation string

	fmt.Println("Please, choose your operation: SUM|AVG|MED")
	fmt.Scan(&operation)
	_, err := isOperationExist(operation)

	if err != nil {
		fmt.Println(err)
		return chooseUserOperation()
	}

	return operation
}

func isOperationExist(operation string) (bool, error) {
	if operation == "SUM" || operation == "AVG" || operation == "MED" {
		return true, nil
	}

	return false, errors.New(operation + " - is not a valid operation")
}

func scanUserOperations() []float64 {
	var operations string
	fmt.Println("Please, input your operations separated by commas (100, 200, 300): ")
	fmt.Scan(&operations)

	var formattedOperations = createSliceInt(operations)
	var operationsToFloatNumbers = make([]float64, 0, 10)

	for _, value := range formattedOperations {
		num, err := strconv.ParseFloat(strings.TrimSpace(value), 64)
		if err != nil {
			fmt.Println("Incorrect number", err)
			continue
		}
		operationsToFloatNumbers = append(operationsToFloatNumbers, num)
	}

	return operationsToFloatNumbers
}

func createSliceInt(operations string) []string {
	return strings.Split(operations, ",")
}

func calculateNumbers(operation string, numbers []float64) float64 {
	var result float64

	switch operation {
	case "SUM":
		{
			for _, value := range numbers {
				result += value
			}
		}
	case "AVG":
		{
			for _, value := range numbers {
				result += value
			}

			result = result / float64(len(numbers))
		}
	default:
		{
			sort.Float64s(numbers)
			if len(numbers)%2 == 0 {
				// 1 is a number that based in the middle of the slice for calculating median
				result = ((numbers[len(numbers)/2]) + (numbers[len(numbers)/2-1])) / 2
			} else {
				result = numbers[len(numbers)/2]
			}
		}
	}

	return result
}
