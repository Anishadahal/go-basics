package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("FILES")
	content := "This needs to go in a file. Again"

	file, err := os.Create("./myGoFile.txt")
	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)
	fmt.Println(length)

	defer file.Close()

	readFile("myGoFile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilError(err)
	fmt.Println("Text data in the file is : ", string(databyte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err) //shut down execution and show error
	}
}
