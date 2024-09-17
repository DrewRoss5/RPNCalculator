package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var equation string
	var result float64
	var err error
	var tokens []string
	reader := bufio.NewReader(os.Stdin)
	for {
		// read the user's input (and strip off the \n suffix)
		fmt.Print("RPNCalc > ")
		equation, _ = reader.ReadString('\n')
		equation = equation[:len(equation)-1]
		if equation == "exit" {
			break
		}
		// display the results of the operation
		tokens = strings.Split(equation, " ")
		result, err = Calculate(tokens)
		if err == nil {
			fmt.Printf("Result: %v\n", result)
		} else {
			fmt.Printf("Error: %v\n", err)
		}

	}
}
