package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var equation string
	var result float64
	var err error
	var operators []string
	var operands stack
	reader := bufio.NewReader(os.Stdin)
	for true {
		// read the user's input (and strip off the \n suffix)
		fmt.Print("RPNCalc > ")
		equation, _ = reader.ReadString('\n')
		equation = equation[:len(equation)-1]
		if equation == "exit" {
			break
		}
		// display the results of the operation
		operators, operands, err = SplitTokens(equation)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		result, err = Calculate(operators, operands)
		if err == nil {
			fmt.Printf("Result: %v\n", result)
		} else {
			fmt.Printf("Error: %v\n", err)
		}

	}
}
