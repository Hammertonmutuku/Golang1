package main

import "fmt"

func main() {
	fmt.Println("maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["GoLang"] = "Golang" // key and value are separated by a colon (:), the language name is JS, then we
	languages["Go"] = "Golang"
	
	fmt.Println("List of all: ", languages)
	fmt.Println("Js of all: ", languages["JS"])

	delete(languages,"JS")
	fmt.Println("List of all: ", languages)

	//Looping in golang are interesting

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	} 
}