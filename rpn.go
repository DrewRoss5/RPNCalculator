package main

import (
	"errors"
	"fmt"
	"math"
	"slices"
	"strconv"
)

type math_op func(float64, float64) float64

func Add(x, y float64) float64    { return x + y }
func Sub(x, y float64) float64    { return x - y }
func Multi(x, y float64) float64  { return x * y }
func Div(x, y float64) float64    { return x / y }
func DivInt(x, y float64) float64 { return math.Floor(x / y) }
func Mod(x, y float64) float64    { return float64(int(x) % int(y)) }

// performs a specified operation with two floats and returns the result. Returns an error if the operator is unrecognized
func ApplyOperation(operands stack, operator string) (float64, stack, error) {
	valid_tokens := []string{"+", "-", "*", "/", "//", "%", "^"}
	// invalid token handler
	if !slices.Contains(valid_tokens, operator) {
		return -1, nil, fmt.Errorf("Invlaid token: %v", operator)
	}
	// too few values handler
	if len(operands) < 2 {
		return -1, nil, fmt.Errorf("invalid equation")
	}
	var x, y float64
	x, operands = Pop(operands)
	y, operands = Pop(operands)
	func_map := map[string]math_op{"+": Add, "-": Sub, "*": Multi, "/": Div, "//": DivInt, "%": Mod, "^": math.Pow}
	return (math.Round(func_map[operator](x, y)*1000) / 1000), operands, nil
}

func Calculate(tokens []string) (float64, error) {
	var err error
	var operands stack
	var operand, result float64
	// parse each token
	for _, token := range tokens {
		operand, err = strconv.ParseFloat(token, 64)
		if err == nil {
			operands = append(operands, operand)
		} else {
			result, operands, err = ApplyOperation(operands, token)
			if err != nil {
				return -1, err
			}
			operands = append(operands, result)
		}
	}
	//  not enough operations were specified
	if len(operands) != 1 {
		return 0, errors.New("incomplete expression (missing operator)")
	}
	return result, nil
}
