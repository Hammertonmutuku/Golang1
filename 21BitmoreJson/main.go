package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to Json video")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{
			"ReactJs Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"},
		},
		{
			"MERN Bootcamp", 199, "LearnCodeOnline.in", "bcd123", []string{"full-stack", "js"},
		},
		{
			"Angular Bootcamp", 299, "LearnCodeOnline.in", "hit123", nil,
		},
	}

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJs Bootcamp",
		"Price": 299,
		"Platform": "LearnCodeOnline.in",
		"tags": ["web-dev","js"]
}
	`)
	var lsoCourse  course 

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json is valid")
		json.Unmarshal(jsonDataFromWeb, &lsoCourse)
		fmt.Println("")
		fmt.Printf("%#v\n", lsoCourse)
	} else {
		fmt.Println("Invalid JSON")
	}

	// some cases where you just want to add data to key value

	var myOnlineData map[string] interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is : %T\n", k, v, v)
	}
}
