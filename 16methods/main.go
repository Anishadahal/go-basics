package main

import "fmt"

func main() {
	fmt.Println("METHODS")

	anisha := User{"Anisha", "anisha@gmail.com", true, 16}
	fmt.Println(anisha)
	anisha.GetStatus()
	anisha.NewMail()
	fmt.Println(anisha)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active? \n", u.Status)
}

func (u User) NewMail() {
	u.Email = "test.go.dev"
	fmt.Printf("Email of %v is changed to %v \n", u.Name, u.Email)
}
