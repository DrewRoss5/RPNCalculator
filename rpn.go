package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// performs a specified operation with two floats and returns the result. Returns an error if the operator is unrecognized
func ApplyOperation(x, y float64, operator string) (float64, error) {
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
	default:
		return 0, fmt.Errorf("unrecognized operator: %v", operator)
	}
	return (math.Round(result*1000) / 1000), nil
}

func Calculate(equation string) (float64, error) {
	var operands stack
	var result float64
	var err error
	var x, y, new_val float64
	// seperate the equation into tokens
	tokens := strings.Split(equation, " ")
	// parse each token
	for _, token := range tokens {
		new_val, err = strconv.ParseFloat(token, 64)
		if err == nil {
			operands = append(operands, new_val)
		} else {
			if len(operands) < 2 {
				return 0, errors.New("incomplete expression (missing number)")
			}
			// get the top two values off of the stack
			x, operands = Pop(operands)
			y, operands = Pop(operands)
			// run the operation and check if it returned an error
			result, err = ApplyOperation(x, y, token)
			if err != nil {
				return 0, err
			}
			// append the result to the stack
			operands = append(operands, result)
		}
	}
	//  not enough operations were specified
	if len(operands) != 1 {
		return 0, errors.New("incomplete expression (missing operator)")
	}
	return result, nil
}
