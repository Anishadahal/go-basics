package main

import "fmt"

func main() {
	fmt.Println("Struct")

	//no inheritance in go
	//no super/parent

	anisha := User{"Anisha", "anisha@go.dev", true, 16}

	fmt.Println(anisha)
	fmt.Printf("The details are: %+v\n", anisha)
	fmt.Printf("Name is %v and email is %v", anisha.Name, anisha.Email)

}

type User struct { //class => capital
	Name   string //fields => public = capital
	Email  string
	Status bool
	Age    int
}
