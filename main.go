package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var expression string = ""
	fmt.Scan(&expression)
	fmt.Println(expression)
	fmt.Println(performCalculation(expression))
}

func performCalculation(input string) float64 {

	var operands []string
	var result float64
	if strings.Contains(input, "+") {
		operands = strings.Split(input, "+")
		x, _ := strconv.ParseFloat(operands[0], 64)
		y, _ := strconv.ParseFloat(operands[1], 64)
		result = x + y
	} else if strings.Contains(input, "-") {
		operands = strings.Split(input, "-")
		x, _ := strconv.ParseFloat(operands[0], 64)
		y, _ := strconv.ParseFloat(operands[1], 64)
		result = x - y
	} else if strings.Contains(input, "*") {
		operands = strings.Split(input, "*")
		x, _ := strconv.ParseFloat(operands[0], 64)
		y, _ := strconv.ParseFloat(operands[1], 64)
		result = x * y
	} else if strings.Contains(input, "/") {
		operands = strings.Split(input, "/")
		x, _ := strconv.ParseFloat(operands[0], 64)
		y, _ := strconv.ParseFloat(operands[1], 64)
		result = x / y
	}

	return result
}
