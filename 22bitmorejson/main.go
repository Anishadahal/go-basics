package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` //aliases
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("MOre JSON")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "learnCodeInline.in", "abc123", []string{"React", "web-dev"}},
		{"JS Bootcamp", 99, "learnCodeInline.in", "abc1234", []string{"JS", "web-dev"}},
		{"NextJS Bootcamp", 199, "learnCodeInline.in", "ac123", nil},
	}

	//package this data as JSON data
	// finalJSON, err := json.Marshal(lcoCourses)
	finalJSON, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJSON)

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{
			"coursename": "JS Bootcamp",
			"Price": 99,
			"website": "learnCodeInline.in",
			"tags": [
				"JS",
				"web-dev"
			]
		}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)

		fmt.Printf("%#v\n", lcoCourse) //value of interfacess => %#v
	} else {
		fmt.Println("JSON was not valid")
	}

	//some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for key, value := range myOnlineData {
		fmt.Printf("Key %v : Value %v , Type %T\n", key, value, value)
	}
}
