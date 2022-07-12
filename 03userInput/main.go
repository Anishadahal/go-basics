package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	//Create reader
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza.")

	//comma ok syntax || error ok syntax
	input, _ := reader.ReadString('\n') //read as long as \n
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is, %T", input)
}
