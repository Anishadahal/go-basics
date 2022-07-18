package main

import "fmt"

func main() {
	defer fmt.Println("Defer in GO")
	defer fmt.Println("......")
	fmt.Println("Lang")
	defer fmt.Println("\n uage") //LIFO

	myDefer()
}


//lang, myDefer
func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print("d", i, ", ")
		fmt.Print(i, ", ")
	}
}
