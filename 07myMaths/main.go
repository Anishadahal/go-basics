package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	// "math/rand"
	// "time"
)

func main() {
	fmt.Println("MATHS in GO")

	// var myNum1 int = 2
	// var myNum2 float64 = 4

	// fmt.Println("THe sum is", myNum1+int(myNum2))//BAD

	//Random Number - math
	// rand.Seed(time.Now().UnixMicro())
	// fmt.Println(rand.Intn(5) + 1)

	//Random Number - crypto
	myRandomNumber, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNumber)

}
