package main

import "fmt"

func main() {

	fmt.Println("If ELse ")

	loginCount := 23
    var result string

	if loginCount < 10 {
		result = "Regular User"
	} else {
		result = "Something else"
	}

	fmt.Println(result)

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	}else {
		fmt.Println("Num is Not less than 10")
	}

	// if err != nil {

	// }
}
