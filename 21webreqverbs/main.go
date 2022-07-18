package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("WEB REQUEST VERB")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:8000/get"

	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("status code", response.StatusCode)
	fmt.Println("content length ", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println(string(content))
	fmt.Println(byteCount)
	fmt.Println(responseString.String())
}

func PerformPostJsonRequest() {
	const myUrl = "http://localhost:8000/post"

	//fake json payload
	requestBody := strings.NewReader(`
		{
			"courseName" : "let's go with golang",
			"price" : 0,
			"platform" : "learnCodeOnline.in"
		}
	`)
	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myUrl = "http://localhost:8000/post"

	//fake form data
	data := url.Values{}
	data.Add("firstname", "Anisha")
	data.Add("lastname", "Dahal")
	data.Add("email", "aninsha@gmail.com")

	response, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}
