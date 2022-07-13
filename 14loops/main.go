package main

import "fmt"

func main() {
	fmt.Println("Loops")

	days := []string{"Sun", "Tue", "Wed", "Fri", "Sat"}

	fmt.Println(days)

	//for loop

	for d := 0; d < len(days); d++ {
		// fmt.Println(days[d])
	}

	//like for each
	// for i := range days {
	// fmt.Println(days[i])
	// }

	//next way
	// for index, day := range days {
	// 	fmt.Printf("Index is %v and value is %v\n", index, day)
	// }

	//break, continue, goto
	rogueValue := 1

	//like while loop
	for rogueValue < 10 {

		if rogueValue == 2 {
			goto lco
		}
		if rogueValue == 5 {
			// break

			rogueValue++
			continue
		}
		fmt.Println("Value is", rogueValue)
		rogueValue++
	}

lco:
	fmt.Println("Jumping at ...goto")
}
