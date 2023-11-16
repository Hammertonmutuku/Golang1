package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to  Get request")
	//PerformGetReuest()
	//PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetReuest() {

	const myUrl = "http://localhost:8000/get"

	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}

	

	defer response.Body.Close()

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println(byteCount)
	fmt.Println(responseString.String())

	// fmt.Println(string(content))
}

func PerformPostJsonRequest() {
	const myUrl = "http://localhost:8000/post";

	requestBody := strings.NewReader(`
	{
		"name": "John",
		"age":30,
		"city":"London"

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
	const myUrl = "http://localhost:8000/postform";

	//formdata

	data := url.Values{}

	data.Add("firstName", "Hammerton")
	data.Add("lastName", "Mutuku")
	data.Add("email", "hammertonmutuku@gmail.com")

	response, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}
