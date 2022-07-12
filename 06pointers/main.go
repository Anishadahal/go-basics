package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointers")

	var ptr *int
	fmt.Println("Value of ptr is", ptr)

	//referencing using &
	myNumber := 23
	var ptr2 *int = &myNumber
	fmt.Println("Value of ptr2 address is", ptr2)
	fmt.Println("Value of ptr2 value is", *ptr2)
}
