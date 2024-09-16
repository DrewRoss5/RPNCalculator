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
		result, err = EvaluateRPN(equation)
		if err == nil {
			fmt.Printf("Result: %v\n", result)
		} else {
			fmt.Printf("Result: %v\n", err)
		}

	}
}
