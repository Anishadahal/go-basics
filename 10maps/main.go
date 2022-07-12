package main

import "fmt"

func main() {
	fmt.Println("MAPS")
	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println(languages)
	fmt.Println("js : ", languages["JS"])

	delete(languages, "RB")
	fmt.Println(languages)

	//loops are interesting in golang
	for key, value := range languages {
		fmt.Println(key, value)
		fmt.Printf("For key %v value is %v", key, value)
	}
	for _, value := range languages {
		fmt.Println(value)
	}
}
