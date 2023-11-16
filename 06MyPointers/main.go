package main

import "fmt"

func main() {
	var ptr1 *int
	fmt.Println("Value of pointer is ", ptr1)

	myNumber := 23

	var ptr = &myNumber
	fmt.Println("Value of pointer is ", ptr)
	fmt.Println("Value of pointer is ", *ptr)

}
