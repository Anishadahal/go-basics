package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time study of GO lang")
	presentTime := time.Now()
	fmt.Println(presentTime)

	//change time format
	fmt.Println(presentTime.Format("01-02-2006"))
	fmt.Println(presentTime.Format("02-01-2006 Monday  15:04:05"))

	//created date
	createdDate := time.Date(2020, time.April, 8, 12, 10, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 15:04 Monday"))
}
