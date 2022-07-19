package main

import (
	"fmt"
	"net/http"
	"sync"
	// "time"
)

var wg sync.WaitGroup //pointer
var signals = []string{"test"}
var mut sync.Mutex //pointer

func main() {
	fmt.Println("GO routines")
	// go greeter("Hello") //creating go routine
	// greeter("World")

	websiteList := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://facebook.com",
		"https://github.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Second)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()

	result, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPs in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for the %s\n", result.StatusCode, endpoint)
	}
}
