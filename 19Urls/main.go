package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentif=dhdhdhdh"

func main() {
	fmt.Println("Welcome to handling Urls in golang")

	fmt.Println(myUrl)

	result, _ := url.Parse(myUrl)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	qParams := result.Query()

	fmt.Printf("The results of query are : %T\n", qParams)

	fmt.Println(qParams["coursename"])
	fmt.Println(qParams["paymentif"])

	for _, val := range qParams {
		fmt.Println("Param is : ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=hitesh",

	}
    
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
