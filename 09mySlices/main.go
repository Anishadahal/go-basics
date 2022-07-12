package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices in GO")

	var fruitList = []string{"tomato", "peach"}
	fmt.Printf("This is of type %T\n", fruitList)

	fruitList = append(fruitList, "mango", "apple", "banana")
	fmt.Println(fruitList)

	var fruitList1 = append(fruitList[1:])//starts from 1st index, not 0th
	fmt.Println("[1:]",fruitList1)

	var fruitList2 = append(fruitList[0:2])
	fmt.Println("[0:2]",fruitList2) //doesnot count 2
 
	var fruitList3 = append(fruitList[:3])//starts from 1st index, not 0th
	fmt.Println("[:3]",fruitList3)


	highScore := make([]int, 4)
	highScore[0] = 234
	highScore[1] = 945
	highScore[2] = 465
	highScore[3] = 867
	// highScore[4] = 777

	highScore = append(highScore, 777)

	fmt.Println(highScore)
	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

	//how to remove a value from slices based on index
	var courses = []string{"reactjs", "javascript","swift", "python", "ruby", "go"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]... )//deleting index 2
	fmt.Println(courses)


}
