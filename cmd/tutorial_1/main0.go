package main

import (
	"errors"
	"fmt"
)

func main0() {
	printMe()
	var intNum int
	fmt.Println(intNum)

	var floatNum float32
	fmt.Println(floatNum)

	var number1 int = 10
	var number2 int = 2
	var result, remainder, err = intDivision(number1, number2)
	if err != nil {
		fmt.Printf(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the integer division is %v", result)
	} else {
		fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
	}
}
func printMe() {
	fmt.Println("hello world!")

}
func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err

}
