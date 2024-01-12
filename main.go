package main

import (
	"calculator/calculator/calc"
	"fmt"
)

func main() {
	sum := calc.Add(10, 5)
	fmt.Println("sum", sum)

	difference := calc.Subtract(10, 5)
	fmt.Println("difference", difference)

	product := calc.Multiply(10, 5)
	fmt.Println("product", product)

	quotient, err := calc.Divide(10, 0)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println("quotient", quotient)

}
