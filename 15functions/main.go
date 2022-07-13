package main

import "fmt"

func main() {
	fmt.Println("FUNCTIONS")
	greeter()
	result := adder(12, 1)
	fmt.Println("Result is : ", result)
	result1, message := proAdder(12, 1)
	fmt.Println("Result is : ", result1, " ", message)
	result2, _ := proAdder(12, 1, 3)
	fmt.Println("Result is : ", result2)
	result3, _ := proAdder(12, 1, 3, 3, 3, 2, 3, 3, 8, 3)
	fmt.Println("Result is : ", result3)
}

func greeter() {
	fmt.Println("Namaste")
}

func adder(a int, b int) int { //function signatures
	return a + b
}

func proAdder(values ...int) (int, string) { //values = slice
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hi there"
}
