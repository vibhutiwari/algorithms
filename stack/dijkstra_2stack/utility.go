package main

import (
	"strconv"
)

func isDigit(char string) bool {
	_, err := strconv.Atoi(char)
	return err == nil
}

func isOpenParentheses(char string) bool {
	switch char {
	case "(", "{", "[":
		return true
	}
	return false
}

func isCloseParentheses(char string) bool {
	switch char {
	case ")", "}", "]":
		return true
	}
	return false
}

func isOperator(char string) bool {
	switch char {
	case "+", "-", "*", "/", "%":
		return true
	}
	return false
}

func getPrecedence(char string) int {
	switch char {
	case "*", "/", "%":
		return 2
	case "+", "-":
		return 1
	}
	return 0
}

func computeOperation(operand1Str, operand2Str, operator string) (string, error) {
	operand1, err := strconv.Atoi(operand1Str)
	if err != nil {
		return "", err
	}

	operand2, err := strconv.Atoi(operand2Str)
	if err != nil {
		return "", err
	}

	var result int
	switch operator {
	case "+":
		result = operand1 + operand2
	case "-":
		result = operand1 - operand2
	case "*":
		result = operand1 * operand2
	case "/":
		result = operand1 / operand2
	case "%":
		result = operand1 % operand2
	}

	return strconv.Itoa(result), nil
}
