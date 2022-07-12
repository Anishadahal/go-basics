package main

import "fmt"

func main() {
	fmt.Println("Go Lang Arrays")

	var fruitList [5]string
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[4] = "Peach"
	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println(vegList)
	fmt.Println(len(vegList))
}
