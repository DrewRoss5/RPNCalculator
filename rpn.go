package main

import (
	"errors"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

// performs a specified operation with two floats and returns the result. Returns an error if the operator is unrecognized
func ApplyOperation(x, y float64, operator string) float64 {
	var result float64
	switch operator {
	case "+":
		result = x + y
	case "-":
		result = x - y
	case "*":
		result = x * y
	case "/":
		result = x / y
	case "//":
		result = math.Floor(x / y)
	case "^":
		result = math.Pow(x, y)
	case "%":
		result = float64(int64(x) % int64(y))
	}
	return (math.Round(result*1000) / 1000)
}

func Calculate(operators []string, operands stack) (float64, error) {
	var x, y, result float64
	// parse each token
	for _, operator := range operators {
		if len(operands) < 2 {
			return 0, errors.New("incomplete expression (missing number)")
		}
		// get the top two values off of the stack
		x, operands = Pop(operands)
		y, operands = Pop(operands)
		// append the result to the stack
		result = ApplyOperation(x, y, operator)
		operands = append(operands, result)
	}
	//  not enough operations were specified
	if len(operands) != 1 {
		return 0, errors.New("incomplete expression (missing operator)")
	}
	return result, nil
}

// spilts an equation string into tokens and validates each token
func SplitTokens(equation string) ([]string, stack, error) {
	valid_operators := []string{"+", "-", "*", "/", "//", "%", "^"}
	tokens := strings.Split(equation, " ")
	// validate each token (each token must be either a number or an operator)
	var err error
	var num float64
	var operands stack
	var operators []string
	for _, token := range tokens {
		num, err = strconv.ParseFloat(token, 64)
		if err == nil {
			operands = append(operands, num)
			continue
		} else {
			if !slices.Contains(valid_operators, token) {
				return nil, nil, fmt.Errorf("unrecognized token: %v", token)
			}
			operators = append(operators, token)
		}
	}

	return operators, operands, nil
}
