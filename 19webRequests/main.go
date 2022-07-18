package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("WEB REQUESTS IN GO")

	response, err := http.Get(url)
	checkNilError(err)
	fmt.Printf("Response is of type %T\n", response)
	defer response.Body.Close() //It is callers responsibility to close connection

	databytes, err := ioutil.ReadAll(response.Body)
	checkNilError(err)
	fmt.Println(string(databytes))
}

func checkNilError(err error) {
	if err != nil {
		panic(err) //shut down execution and show error
	}
}
