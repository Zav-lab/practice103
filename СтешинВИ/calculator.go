package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите выражение: ")
	expression, _ := reader.ReadString('\n')

	result := calculate(expression)
	fmt.Println("Результат:", result)
}

func calculate(expression string) float64 {
	expression = strings.ReplaceAll(expression, " ", "")

	operands := strings.FieldsFunc(expression, func(r rune) bool {
		return r == '+' || r == '-' || r == '*' || r == '/'
	})
	operators := strings.FieldsFunc(expression, func(r rune) bool {
		return r == '+' || r == '-' || r == '*' || r == '/'
	})

	numbers := make([]float64, len(operands))
	for i, operand := range operands {
		numbers[i], _ = strconv.ParseFloat(operand, 64)
	}

	for i, operator := range operators {
		switch operator {
		case "*":
			numbers[i+1] *= numbers[i]
		case "/":
			numbers[i+1] = numbers[i] / numbers[i+1]
		}
	}

	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		switch operators[i-1] {
		case "+":
			result += numbers[i]
		case "-":
			result -= numbers[i]
		}
	}

	return result
}