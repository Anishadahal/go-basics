package main

import "fmt"

const LoginToken string = "lkafdjs123123123" //First letter capital = public variable

func main() {
	fmt.Println("Variables")
	var username string = "Anisha"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username) //T for type

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("This is of type: %T \n", isLoggedIn)

	var smallVal int = 256
	fmt.Println(smallVal)
	fmt.Printf("This is of type: %T \n", smallVal)

	var smallFloat float64 = 256.23423456789900876543234567890123456789
	fmt.Println(smallFloat)
	fmt.Printf("This is of type: %T \n", smallFloat)

	//default values and aliases

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("This is of type: %T \n", anotherVariable)

	var anotherVariableString string
	fmt.Println(anotherVariableString)
	fmt.Printf("This is of type: %T \n", anotherVariableString)

	//Implicit Type
	var name = "Anisha Dahal"
	fmt.Println(name)
	fmt.Printf("This is of type: %T \n", name)

	var age = 12
	fmt.Println(age)
	fmt.Printf("This is of type: %T \n", age)

	//No var style
	numberOfUser := 123 //Walrus operator allowed inside method only
	fmt.Println(numberOfUser)
	fmt.Printf("This is of type: %T \n", numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("This is of type: %T \n", LoginToken)


}
