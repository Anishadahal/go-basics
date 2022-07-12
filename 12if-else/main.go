package main

import "fmt"

func main() {
	fmt.Println("IF ELSE")

	loginCount := 10

	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Something Else"
	} else {
		result = "Watch out"
	}
	fmt.Println(result)

	//odd-veven
	if loginCount%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	//WEIRD
	if num := 3; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is not less than 10")
	}


	//comma ok syntax
	// if err != nil{

	// }
}
